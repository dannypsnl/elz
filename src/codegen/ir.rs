use crate::ast;
use crate::ast::*;
use std::cell::RefCell;
use std::collections::HashMap;
use std::fmt::Formatter;
use std::rc::Rc;

pub struct Module {
    pub(crate) known_functions: HashMap<String, Type>,
    pub(crate) known_variables: HashMap<String, Type>,
    pub(crate) functions: Vec<Function>,
    pub(crate) variables: Vec<Variable>,
    pub(crate) types: Vec<TypeDefinition>,
}

impl Module {
    pub(crate) fn new() -> Module {
        Module {
            known_functions: HashMap::new(),
            known_variables: HashMap::new(),
            functions: vec![],
            variables: vec![],
            types: vec![],
        }
    }
    pub(crate) fn remember_function(&mut self, f: &ast::Function) {
        let ret_type = Type::from_ast(&f.ret_typ);
        self.known_functions.insert(f.name.clone(), ret_type);
    }
    pub(crate) fn remember_variable(&mut self, v: &ast::Variable) {
        self.known_variables
            .insert(v.name.clone(), Type::from_ast(&v.typ));
    }
    pub(crate) fn push_function(&mut self, f: Function) {
        self.functions.push(f);
    }
    pub(crate) fn push_variable(&mut self, v: Variable) {
        self.variables.push(v);
    }
    pub(crate) fn push_type(&mut self, type_name: &String, fields: &Vec<ClassMember>) {
        self.types.push(TypeDefinition {
            name: type_name.clone(),
            fields: fields
                .iter()
                .filter(|&member| match member {
                    ClassMember::Field(_) => true,
                    _ => false,
                })
                .map(|member| match member {
                    ClassMember::Field(field) => Type::from_ast(&field.typ),
                    _ => unreachable!(),
                })
                .collect(),
        });
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct TypeDefinition {
    pub(crate) name: String,
    pub(crate) fields: Vec<Type>,
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct ID {
    value: u64,
}

impl ID {
    fn new() -> Rc<RefCell<ID>> {
        Rc::new(RefCell::new(ID { value: 0 }))
    }
    fn set_id(&mut self, value: u64) -> bool {
        self.value = value;
        true
    }
}

impl std::fmt::Display for ID {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.value)
    }
}

/// Label represents a location which can be the target of jump instructions
#[derive(Debug, Clone, PartialEq)]
pub(crate) struct Label {
    pub(crate) id: Rc<RefCell<ID>>,
}

impl Label {
    pub(crate) fn new(id: Rc<RefCell<ID>>) -> Rc<Label> {
        Rc::new(Label { id })
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) enum Instruction {
    Return(Option<Expr>),
    Label(Rc<Label>),
    Branch {
        cond: Expr,
        if_true: Rc<Label>,
        if_false: Rc<Label>,
    },
    Goto(Rc<Label>),
    FunctionCall {
        id: Rc<RefCell<ID>>,
        func_name: String,
        ret_type: Box<Type>,
        args_expr: Vec<Expr>,
    },
    BinaryOperation {
        id: Rc<RefCell<ID>>,
        op_name: String,
        lhs: Expr,
        rhs: Expr,
    },
}

impl Instruction {
    pub(crate) fn is_terminator(&self) -> bool {
        use Instruction::*;
        match self {
            Return(..) | Branch { .. } | Goto(..) => true,
            _ => false,
        }
    }

    fn set_id(&mut self, value: u64) -> bool {
        use Instruction::*;
        match self {
            Label(label) => label.id.borrow_mut().set_id(value),
            FunctionCall { id, .. } | BinaryOperation { id, .. } => id.borrow_mut().set_id(value),
            _ => false,
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct Body {
    pub(crate) instructions: Vec<Instruction>,
}

impl Body {
    pub(crate) fn from_ast(b: &ast::Body, module: &Module) -> Body {
        let mut body = Body {
            instructions: vec![],
        };
        match b {
            ast::Body::Expr(e) => {
                let e = body.expr_from_ast(e, module);
                body.instructions.push(Instruction::Return(Some(e)));
            }
            ast::Body::Block(b) => body.generate_instructions(&b.statements, module),
        };
        // update local identifier value
        let mut counter = 1;
        for inst in &mut body.instructions {
            if inst.set_id(counter) {
                counter += 1;
            }
        }
        body
    }
    pub(crate) fn generate_instructions(&mut self, stmts: &Vec<Statement>, module: &Module) {
        for stmt in stmts {
            use ast::StatementVariant::*;
            match &stmt.value {
                Return(e) => {
                    let inst = match e {
                        None => Instruction::Return(None),
                        Some(ex) => Instruction::Return(Some(self.expr_from_ast(ex, module))),
                    };
                    self.instructions.push(inst)
                }
                Expression(expr) => {
                    self.expr_from_ast(expr, module);
                }
                IfBlock {
                    clauses,
                    else_block,
                } => {
                    let leave_label = Label::new(ID::new());
                    for (cond, then_block) in clauses {
                        let if_then_label = Label::new(ID::new());
                        let else_then_label = Label::new(ID::new());
                        let inst = Instruction::Branch {
                            cond: self.expr_from_ast(cond, module),
                            if_true: if_then_label.clone(),
                            if_false: else_then_label.clone(),
                        };
                        self.instructions.push(inst);
                        // if then
                        self.instructions
                            .push(Instruction::Label(if_then_label.clone()));
                        self.generate_instructions(&then_block.statements, module);
                        if !self.end_with_terminator() {
                            self.goto(&leave_label);
                        }
                        // else then
                        self.instructions
                            .push(Instruction::Label(else_then_label.clone()));
                    }
                    self.generate_instructions(&else_block.statements, module);
                    if !self.end_with_terminator() {
                        self.goto(&leave_label);
                    }
                    self.instructions
                        .push(Instruction::Label(leave_label.clone()));
                }
                st => unimplemented!("for statement: {:#?}", st),
            }
        }
    }
    fn expr_from_ast(&mut self, expr: &ast::Expr, module: &Module) -> Expr {
        use ast::ExprVariant::*;
        match &expr.value {
            Binary(lhs, rhs, op) => {
                let id= ID::new();
                let lhs = self.expr_from_ast(lhs, module);
                let rhs = self.expr_from_ast(rhs, module);
                let result_typ = lhs.type_();
                let op_name = match op {
                    Operator::Plus => "add",
                }
                .to_string();
                let inst = Instruction::BinaryOperation {
                    id:  id.clone(),
                    op_name,
                    lhs,
                    rhs,
                };
                self.instructions.push(inst);
                Expr::id(result_typ, id)
            }
            FuncCall(f, args) => {
                let id = self.expr_from_ast(f, module);
                let name = match id {
                    Expr::Identifier(_, name) => name,
                    e => unreachable!("call on a non-function expression: {:#?}", e),
                };
                match module.known_functions.get(&name) {
                    Some(ret_type) => {
                        let args_expr: Vec<Expr> = args
                            .iter()
                            .map(|arg| self.expr_from_ast(&arg.expr, module))
                            .collect();
                        let id = ID::new();
                        let inst = Instruction::FunctionCall{
                            id: id.clone(),
                            func_name: format!("@{}", name),
                            ret_type: ret_type.clone().into(),
                            args_expr,
                        };
                        self.instructions.push(inst);
                        Expr::id(ret_type.clone(), id)
                    },
                    None => unreachable!("no function named: {} which unlikely happened, semantic module must have a bug there!", name),
                }
            }
            Identifier(name) => {
                match module.known_functions.get(name) {
                    Some(ret_type) => {
                        Expr::Identifier(ret_type.clone(), name.clone())
                    },
                    None => unreachable!("no variable named: {} which unlikely happened, semantic module must have a bug there!", name),
                }
            }
            _ => Expr::from_ast(expr),
        }
    }
    fn end_with_terminator(&self) -> bool {
        match self.instructions.last() {
            None => false,
            Some(inst) => inst.is_terminator(),
        }
    }
    fn goto(&mut self, label: &Rc<Label>) {
        self.instructions.push(Instruction::Goto(label.clone()));
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct Function {
    pub(crate) name: String,
    pub(crate) parameters: Vec<(String, Type)>,
    pub(crate) ret_typ: Type,
    pub(crate) body: Option<Body>,
}

impl Function {
    pub(crate) fn new(
        name: String,
        parsed_params: &Vec<Parameter>,
        ret_typ: Type,
        body: Option<Body>,
    ) -> Function {
        let parameters: Vec<(String, Type)> = parsed_params
            .iter()
            .map(|p| (p.name.clone(), Type::from_ast(&p.typ)))
            .collect();
        Function {
            // function name need @, e.g. @main
            name: format!("@{}", name),
            parameters,
            ret_typ,
            body,
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct Variable {
    pub(crate) name: String,
    pub(crate) expr: Expr,
}

impl Variable {
    pub(crate) fn new(name: String, expr: Expr) -> Variable {
        Variable {
            name: format!("@{}", name),
            expr,
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) enum Type {
    Void,
    Int(usize),
    Float(usize),
    UserDefined(String),
}

impl Type {
    pub(crate) fn from_ast(t: &ast::ParsedType) -> Type {
        use Type::*;
        match t.name().as_str() {
            "void" => Void,
            "int" => Int(64),
            "f64" => Float(64),
            "bool" => Int(1),
            name => UserDefined(name.to_string()),
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) enum Expr {
    I64(i64),
    F64(f64),
    Bool(bool),
    String(String),
    Identifier(Type, String),
    LocalIdentifier(Type, Rc<RefCell<ID>>),
}

impl Expr {
    pub(crate) fn from_ast(a: &ast::Expr) -> Expr {
        use ExprVariant::*;
        match &a.value {
            F64(f) => Expr::F64(*f),
            Int(i) => Expr::I64(*i),
            Bool(b) => Expr::Bool(*b),
            String(s) => Expr::String(s.clone()),
            expr => unimplemented!("codegen: expr {:#?}", expr),
        }
    }
    pub(crate) fn type_(&self) -> Type {
        match self {
            Expr::I64(..) => Type::Int(64),
            Expr::F64(..) => Type::Float(64),
            Expr::Bool(..) => Type::Int(1),
            Expr::String(..) => Type::UserDefined("string".to_string()),
            Expr::Identifier(typ, ..) => typ.clone(),
            Expr::LocalIdentifier(typ, ..) => typ.clone(),
        }
    }

    fn id(typ: Type, id: Rc<RefCell<ID>>) -> Expr {
        Expr::LocalIdentifier(typ, id)
    }
}
