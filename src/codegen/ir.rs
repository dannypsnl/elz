use crate::ast;
use crate::ast::*;
use std::collections::HashMap;

pub struct Module {
    pub(crate) known_functions: HashMap<String, Type>,
    pub(crate) known_variables: HashMap<String, Type>,
    pub(crate) functions: Vec<Function>,
    pub(crate) variables: Vec<Variable>,
}

impl Module {
    pub(crate) fn new() -> Module {
        Module {
            known_functions: HashMap::new(),
            known_variables: HashMap::new(),
            functions: vec![],
            variables: vec![],
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
}

#[derive(Debug, Clone, PartialEq)]
pub(crate) enum Instruction {
    Return(Option<Expr>),
    TempVariable(u64, Expr),
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
        self.statements = stmts
            .iter()
            .map(|stmt| {
                use ast::StatementVariant::*;
                match &stmt.value {
                    Return(e) => match e {
                        None => Instruction::Return(None),
                        Some(ex) => Instruction::Return(Some(Expr::from_ast(ex, module))),
                    },
                    Expression(expr) => {
                        let instruction =
                            Instruction::TempVariable(self.counter, Expr::from_ast(expr, module));
                        self.counter += 1;
                        instruction
                    }
                    st => unimplemented!("for statement: {:#?}", st),
                }
            })
            .collect();
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
            .map(|p| {
                let name = p.1.clone();
                let typ = Type::from_ast(&p.0);
                (name, typ)
            })
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
}

impl Type {
    pub(crate) fn from_ast(t: &ast::ParsedType) -> Type {
        use Type::*;
        match t.name().as_str() {
            "void" => Void,
            "int" => Int(64),
            "f64" => Float(64),
            "bool" => Int(1),
            _ => unimplemented!("type `{}`", t.name()),
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
