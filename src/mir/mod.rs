use super::ast;
use super::ast::Top;
use std::collections::HashMap;

mod main_checks;
use main_checks::*;

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

pub struct Context {
    parent: Option<*const Context>,
    binding_map: HashMap<String, Bind>,
}

pub struct Bind(String, ast::Expr);

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

    let mut mir = MIR {
        binary_entry: None,
        functions: vec![],
    };

    mir.binary_entry = match ctx.binding_map.get("main") {
        Some(main_bind) => {
            let Bind(name, expr) = main_bind;
            let lambda = check_main_is_lambda(expr)?;
            check_main_return_type(lambda)?;
            let expr = ensure_main_body_is_not_empty(lambda)?;
            let block = check_main_body_is_block(expr.as_ref())?;
            let statements: Result<Vec<_>> = block.clone()
                .statements
                .into_iter()
                .map(|stmt| generate_stmt_mir(&stmt))
                .collect();
            let main_fn = Function {
                name: name.clone(),
                block: statements?,
            };
            Some(main_fn)
        }
        None => None,
    };

    Ok(mir)
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

pub struct MIR {
    pub binary_entry: Option<Function>,
    pub functions: Vec<Function>,
}

pub struct Function {
    pub name: String,
    pub block: Vec<Statement>,
}

pub enum Statement {
    Return(Expr),
}

pub enum Expr {
    Int(i64),
}
