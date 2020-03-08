use crate::ast;
use crate::ast::*;
use std::cell::RefCell;
use std::collections::HashMap;
use std::fmt::Formatter;
use std::rc::Rc;

pub struct Module {
    // helpers
    pub(crate) known_functions: HashMap<String, Type>,
    pub(crate) known_variables: HashMap<String, Type>,
    known_types: HashMap<String, TypeDefinition>,
    // output parts
    pub(crate) functions: Vec<Function>,
    pub(crate) variables: Vec<Variable>,
    pub(crate) types: Vec<TypeDefinition>,
}

impl Module {
    pub(crate) fn new() -> Module {
        Module {
            known_functions: HashMap::new(),
            known_variables: HashMap::new(),
            known_types: HashMap::new(),
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
        let typ = TypeDefinition {
            name: type_name.clone(),
            fields: fields
                .iter()
                .filter(|&member| match member {
                    ClassMember::Field(_) => true,
                    _ => false,
                })
                .map(|member| match member {
                    ClassMember::Field(field) => Field {
                        name: field.name.clone(),
                        typ: Type::from_ast(&field.typ),
                    },
                    _ => unreachable!(),
                })
                .collect(),
        };
        self.known_types.insert(type_name.clone(), typ.clone());
        self.types.push(typ);
    }
    fn lookup_type(&self, type_name: &String) -> &TypeDefinition {
        self.known_types.get(type_name).unwrap()
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct TypeDefinition {
    pub(crate) name: String,
    pub(crate) fields: Vec<Field>,
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct Field {
    pub(crate) name: String,
    pub(crate) typ: Type,
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
    GEP {
        id: Rc<RefCell<ID>>,
        load_from: Expr,
        indices: Vec<u64>,
    },
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
    Alloca {
        id: Rc<RefCell<ID>>,
        typ: Type,
    },
    Load {
        id: Rc<RefCell<ID>>,
        load_from: Expr,
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
            Load { id, .. }
            | Alloca { id, .. }
            | GEP { id, .. }
            | FunctionCall { id, .. }
            | BinaryOperation { id, .. } => id.borrow_mut().set_id(value),
            _ => false,
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) enum LocalVariable {
    Name { typ: Type, name: String },
}

impl LocalVariable {
    fn from_name(name: String, typ: Type) -> LocalVariable {
        LocalVariable::Name { typ, name }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) struct Body {
    pub(crate) instructions: Vec<Instruction>,
    // local variables(including parameters)
    variables: HashMap<String, LocalVariable>,
}

impl Body {
    fn from_ast(b: &ast::Body, module: &mut Module, parameters: &Vec<Parameter>) -> Body {
        let mut variables = HashMap::new();

        for p in parameters {
            // FIXME: type from duplicate in ir::Function, share information
            let local_var = LocalVariable::from_name(p.name.clone(), Type::from_ast(&p.typ));
            variables.insert(p.name.clone(), local_var);
        }

        let mut body = Body {
            instructions: vec![],
            variables,
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

    fn lookup_variable(&self, name: &String) -> Option<&LocalVariable> {
        self.variables.get(name)
    }

    pub(crate) fn generate_instructions(&mut self, stmts: &Vec<Statement>, module: &mut Module) {
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
                Variable(v) => {
                    self.expr_from_ast(&v.expr, module);
                }
            }
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
    pub(crate) fn from_ast(
        f: &ast::Function,
        class_name: Option<String>,
        module: &mut Module,
    ) -> Function {
        let body = match &f.body {
            Some(b) => Some(Body::from_ast(b, module, &f.parameters)),
            None => None,
        };
        let function_name = match class_name {
            None => f.name.clone(),
            Some(class_name) => format!("\"{}::{}\"", class_name, f.name),
        };
        Function::new(
            function_name,
            &f.parameters,
            Type::from_ast(&f.ret_typ),
            body,
        )
    }
    fn new(
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
    pub(crate) name: GlobalName,
    pub(crate) expr: Expr,
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) enum GlobalName {
    ID(Rc<RefCell<ID>>),
    String(String),
}

impl Variable {
    pub(crate) fn new(name: String, expr: Expr) -> Variable {
        Variable {
            name: GlobalName::String(format!("@{}", name)),
            expr,
        }
    }
    pub(crate) fn from_id(id: Rc<RefCell<ID>>, expr: Expr) -> Variable {
        Variable {
            name: GlobalName::ID(id),
            expr,
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) enum Type {
    Void,
    Int(usize),
    Float(usize),
    Pointer(Rc<Type>),
    Array { len: usize, element_type: Rc<Type> },
    UserDefined(String),
    Named(String),
}

impl Type {
    pub(crate) fn from_ast(t: &ast::ParsedType) -> Type {
        use Type::*;
        match t.name().as_str() {
            "void" => Void,
            "int" => Int(64),
            "f64" => Float(64),
            "bool" => Int(1),
            "_c_string" => Pointer(Int(8).into()),
            name => UserDefined(name.to_string()),
        }
    }

    pub(crate) fn element_type(&self) -> Rc<Type> {
        use Type::*;
        match self {
            UserDefined(name) => Named(name.clone()).into(),
            Pointer(element_type) | Array { element_type, .. } => element_type.clone(),
            _ => unreachable!("`{:?}` don't have element type", self),
        }
    }
}

impl Body {
    fn expr_from_ast(&mut self, expr: &ast::Expr, module: &mut Module) -> Expr {
        use ast::ExprVariant::*;
        match &expr.value {
            String(string_literal) => {
                let str_literal_id = ID::new();
                module.push_variable(Variable::from_id(
                    str_literal_id.clone(),
                    Expr::CString(string_literal.clone()),
                ));
                let str_load_id = ID::new();
                let array_type = Type::Array {
                    len: string_literal.len(),
                    element_type: Type::Int(8).into(),
                };
                let inst = Instruction::GEP {
                    id: str_load_id.clone(),
                    load_from: Expr::global_id(Type::Pointer(array_type.into()), str_literal_id),
                    indices: vec![0, 0],
                };
                self.instructions.push(inst);
                let ptr_to_str = Expr::local_id(Type::Pointer(Type::Int(8).into()), str_load_id);
                let id = ID::new();
                let ret_type: Type = Type::UserDefined("string".to_string()).into();
                let inst = Instruction::FunctionCall {
                    id: id.clone(),
                    func_name: format!("@\"string::new\""),
                    ret_type: ret_type.clone().into(),
                    args_expr: vec![ptr_to_str],
                };
                self.instructions.push(inst);
                Expr::local_id(ret_type.clone(), id)
            }
            ClassConstruction(class_name, field_inits) => {
                // TODO:
                //  1. gep fields
                //  2. store value to fields
                let alloca_id = ID::new();
                let type_name = Type::UserDefined(class_name.clone());
                let inst = Instruction::Alloca {
                    id: alloca_id.clone(),
                    typ: type_name.clone(),
                };
                self.instructions.push(inst);
                let id = ID::new();
                let type_name = Type::UserDefined(class_name.clone());
                let inst = Instruction::Load {
                    id: id.clone(),
                    load_from: Expr::local_id(type_name.clone(), alloca_id),
                };
                self.instructions.push(inst);
                Expr::local_id(type_name, id)
            }
            MemberAccess(from, access) => {
                let v = self.expr_from_ast(from, module);
                let type_definition = match &v.type_() {
                    Type::UserDefined(type_name) => module.lookup_type(type_name),
                    _ => unreachable!(
                        "access member on non-class type which unlikely happen: from `{:?}`",
                        from
                    ),
                };
                let mut result_type = v.type_();
                let mut i = 0; // get index of field
                for field in &type_definition.fields {
                    if &field.name == access {
                        result_type = field.typ.clone();
                        break;
                    } else {
                        i += 1;
                    }
                }
                let gep_id = ID::new();
                let inst = Instruction::GEP {
                    id: gep_id.clone(),
                    load_from: v.clone(),
                    indices: vec![0, i],
                };
                self.instructions.push(inst);
                let id = ID::new();
                let inst = Instruction::Load {
                    id: id.clone(),
                    load_from: Expr::local_id(result_type.clone(), gep_id),
                };
                self.instructions.push(inst);
                Expr::local_id(result_type, id)
            }
            Binary(lhs, rhs, op) => {
                let id = ID::new();
                let lhs = self.expr_from_ast(lhs, module);
                let rhs = self.expr_from_ast(rhs, module);
                let result_typ = lhs.type_();
                let op_name = match op {
                    Operator::Plus => "add",
                }
                .to_string();
                let inst = Instruction::BinaryOperation {
                    id: id.clone(),
                    op_name,
                    lhs,
                    rhs,
                };
                self.instructions.push(inst);
                Expr::local_id(result_typ, id)
            }
            FuncCall(f, args) => {
                let id = self.expr_from_ast(f, module);
                let name = match id {
                    Expr::Identifier(_, name) => name,
                    e => unreachable!("call on a non-function expression: {:#?}", e),
                };
                match module.known_functions.get(&name).cloned() {
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
                        Expr::local_id(ret_type.clone(), id)
                    },
                    None => unreachable!("no function named: `{}` which unlikely happened, semantic module must have a bug there!", name),
                }
            }
            Identifier(name) => match self.lookup_variable(name) {
                Some(local_var) => match local_var {
                    LocalVariable::Name { name, typ } => {
                        Expr::Identifier(typ.clone(), name.clone())
                    }
                },
                None => {
                    let ret_type = module.known_functions.get(name).expect(format!("no variable named: `{}` which unlikely happened, semantic module must have a bug there!", name).as_str());
                    Expr::Identifier(ret_type.clone(), name.clone())
                }
            },
            _ => Expr::from_ast(expr),
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) enum Expr {
    I64(i64),
    F64(f64),
    Bool(bool),
    CString(String),
    Identifier(Type, String),
    LocalIdentifier(Type, Rc<RefCell<ID>>),
    GlobalIdentifier(Type, Rc<RefCell<ID>>),
}

impl Expr {
    pub(crate) fn from_ast(a: &ast::Expr) -> Expr {
        use ExprVariant::*;
        match &a.value {
            F64(f) => Expr::F64(*f),
            Int(i) => Expr::I64(*i),
            Bool(b) => Expr::Bool(*b),
            String(s) => Expr::CString(s.clone()),
            expr => unimplemented!("codegen: expr {:#?}", expr),
        }
    }
    pub(crate) fn type_(&self) -> Type {
        match self {
            Expr::I64(..) => Type::Int(64),
            Expr::F64(..) => Type::Float(64),
            Expr::Bool(..) => Type::Int(1),
            Expr::CString(s) => Type::Array {
                len: s.len(),
                element_type: Type::Int(8).into(),
            },
            Expr::Identifier(typ, ..) => typ.clone(),
            Expr::LocalIdentifier(typ, ..) => typ.clone(),
            Expr::GlobalIdentifier(typ, ..) => typ.clone(),
        }
    }

    fn local_id(typ: Type, id: Rc<RefCell<ID>>) -> Expr {
        Expr::LocalIdentifier(typ, id)
    }
    fn global_id(typ: Type, id: Rc<RefCell<ID>>) -> Expr {
        Expr::GlobalIdentifier(typ, id)
    }
}
