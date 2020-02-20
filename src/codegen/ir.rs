use crate::ast;
use crate::ast::*;
use std::collections::HashMap;
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

/// Label represents a location which can be the target of jump instructions
///
/// LLVM example: `; <label>:9:`, here `id` is `9`
#[derive(Debug, Clone, PartialEq)]
pub(crate) struct Label {
    pub(crate) id: u64,
}

impl Label {
    pub(crate) fn new(id: u64) -> Label {
        Label { id }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) enum Instruction {
    Return(Option<Expr>),
    TempVariable(u64, Expr),
    Label(Rc<Label>),
    Branch {
        cond: Expr,
        if_true: Rc<Label>,
        if_false: Rc<Label>,
    },
    Goto(Rc<Label>),
}

impl Instruction {
    pub(crate) fn is_terminator(&self) -> bool {
        use Instruction::*;
        match self {
            Return(..) | Branch { .. } | Goto(..) => true,
            _ => false,
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct Body {
    // helper part
    pub(crate) counter: u64,
    // value part
    pub(crate) statements: Vec<Instruction>,
}

impl Body {
    pub(crate) fn from_ast(b: &ast::Body, module: &Module) -> Body {
        let mut body = Body {
            counter: 0,
            statements: vec![],
        };
        match b {
            ast::Body::Expr(e) => {
                body.statements = vec![Instruction::Return(Some(Expr::from_ast(e, module)))];
            }
            ast::Body::Block(b) => body.generate_instructions(&b.statements, module),
        };
        body
    }
    pub(crate) fn generate_instructions(&mut self, stmts: &Vec<Statement>, module: &Module) {
        for stmt in stmts {
            use ast::StatementVariant::*;
            match &stmt.value {
                Return(e) => {
                    let inst = match e {
                        None => Instruction::Return(None),
                        Some(ex) => Instruction::Return(Some(Expr::from_ast(ex, module))),
                    };
                    self.statements.push(inst)
                }
                Expression(expr) => {
                    let instruction = Instruction::TempVariable(
                        self.get_and_update_count(),
                        Expr::from_ast(expr, module),
                    );
                    self.statements.push(instruction)
                }
                IfBlock {
                    clauses,
                    else_block,
                } => {
                    let leave_label = Rc::new(Label::new(self.get_and_update_count()));
                    for (cond, then_block) in clauses {
                        let if_then_label = Rc::new(Label::new(self.get_and_update_count()));
                        let else_then_label = Rc::new(Label::new(self.get_and_update_count()));
                        let inst = Instruction::Branch {
                            cond: Expr::from_ast(cond, module),
                            if_true: Rc::clone(&if_then_label),
                            if_false: Rc::clone(&else_then_label),
                        };
                        self.statements.push(inst);
                        // if then
                        self.statements
                            .push(Instruction::Label(Rc::clone(&if_then_label)));
                        self.generate_instructions(&then_block.statements, module);
                        self.goto(&leave_label);
                        // else then
                        self.statements
                            .push(Instruction::Label(Rc::clone(&else_then_label)));
                    }
                    self.generate_instructions(&else_block.statements, module);
                    self.goto(&leave_label);
                    self.statements
                        .push(Instruction::Label(Rc::clone(&leave_label)));
                }
                st => unimplemented!("for statement: {:#?}", st),
            }
        }
    }
    fn get_and_update_count(&mut self) -> u64 {
        let tmp = self.counter;
        self.counter += 1;
        tmp
    }
    fn goto(&mut self, label: &Rc<Label>) {
        if let Some(inst) = self.statements.last() {
            if inst.is_terminator() {
                return;
            }
        }
        self.statements.push(Instruction::Goto(Rc::clone(label)));
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
    Identifier(String),
    // name, return type, argument expression
    FunctionCall(String, Box<Type>, Vec<Expr>),
}

impl Expr {
    pub(crate) fn from_ast(a: &ast::Expr, module: &Module) -> Expr {
        use ExprVariant::*;
        match &a.value {
            F64(f) => Expr::F64(*f),
            Int(i) => Expr::I64(*i),
            Bool(b) => Expr::Bool(*b),
            String(s) => Expr::String(s.clone()),
            Identifier(name) => Expr::Identifier(name.clone()),
            FuncCall(f, args) => {
                let id = Expr::from_ast(f, module);
                let name = match id {
                    Expr::Identifier(name) => name,
                    e => unreachable!("call on a non-function expression: {:#?}", e),
                };
                match module.known_functions.get(&name) {
                    Some(ret_type) => {
                        let args_expr: Vec<Expr> = args
                            .iter()
                            .map(|arg| Expr::from_ast(&arg.expr, module))
                            .collect();
                        Expr::FunctionCall(
                            format!("@{}", name),
                            ret_type.to_owned().into(),
                            args_expr,
                        )
                    },
                    None => unreachable!("no function named: {} which unlikely happened, semantic module must has bug there!", name),
                }
            }
            expr => unimplemented!("codegen: expr {:#?}", expr),
        }
    }
}

impl Expr {
    pub(crate) fn return_void(&self) -> bool {
        match self {
            Expr::FunctionCall(_, ret_type, _) => {
                if ret_type == &Box::new(Type::Void) {
                    true
                } else {
                    false
                }
            }
            _ => false,
        }
    }
}
