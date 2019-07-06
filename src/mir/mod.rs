use super::ast;
use super::ast::Top;
use std::collections::HashMap;

#[derive(Debug, PartialEq)]
pub enum MIRError {
    NoMain,
    Message(String),
}

impl std::fmt::Display for MIRError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "MIRError: {}", self)
    }
}

impl std::error::Error for MIRError {
    fn description(&self) -> &str {
        match self {
            MIRError::NoMain => "no main",
            MIRError::Message(s) => s.as_str(),
        }
    }
}

impl MIRError {
    fn new<T: Into<String>>(message: T) -> MIRError {
        MIRError::Message(message.into())
    }
}

type Result<T> = std::result::Result<T, MIRError>;

pub enum Type {
    /// use as placeholder, and won't do any checking
    /// since semantic should ensure all mismatched aren't
    /// existed.
    Unsure,
    Void,
    I64,
    F64,
    Bool,
    String,
}

pub struct Context {
    parent: Option<*const Context>,
    binding_map: HashMap<String, Bind>,
}

pub struct Bind(String, ast::Expr);

pub struct MIR {
    functions: Vec<Function>,
}

pub fn generate_mir_program(program: &Vec<ast::Top>) -> Result<MIR> {
    let mut ctx = Context {
        parent: None,
        binding_map: HashMap::new(),
    };
    for top_elem in program {
        match top_elem {
            Top::Binding(name, _, expr) => {
                ctx.binding_map
                    .insert(name.clone(), Bind(name.clone(), expr.clone()));
            }
            _ => unimplemented!(),
        };
    }

    let main_binding = ctx.binding_map.get("main");
    if main_binding.is_none() {
        return Err(MIRError::NoMain);
    }

    let mut mir = MIR { functions: vec![] };

    let Bind(name, expr) = main_binding.unwrap();
    if let ast::Expr::Lambda(lambda) = expr {
        if lambda.return_type == ast::Type::Unit {
            if let Some(expr) = lambda.body.clone() {
                if let ast::Expr::Block(block) = expr.as_ref() {
                    let mut stmts = vec![];
                    for stmt in &block.statements {
                        stmts.push(generate_stmt_mir(stmt)?);
                    }
                    let f = Function {
                        name: name.clone(),
                        block: stmts,
                    };
                    mir.functions.push(f);
                    Ok(mir)
                } else {
                    Err(MIRError::new("main lambda must be a block"))
                }
            } else {
                Err(MIRError::new("main lambda can't have an empty body"))
            }
        } else {
            Err(MIRError::new("main lambda must return unit type"))
        }
    } else {
        Err(MIRError::new(
            "main is the entry point of executable and must be a lambda",
        ))
    }
}

pub fn generate_stmt_mir(stmt: &ast::Statement) -> Result<Statement> {
    use ast::Statement::*;
    let stmt = match stmt {
        Return(e) => Statement::Return(generate_expr_mir(e)?),
        _ => unimplemented!(),
    };
    Ok(stmt)
}

pub fn generate_expr_mir(expr: &ast::Expr) -> Result<Expr> {
    use ast::Expr::*;
    let expr = match expr {
        Int(i) => Expr::Int(*i),
        _ => unimplemented!(),
    };
    Ok(expr)
}

pub struct Function {
    name: String,
    block: Vec<Statement>,
}

pub enum Statement {
    Return(Expr),
}

pub enum Expr {
    Int(i64),
}
