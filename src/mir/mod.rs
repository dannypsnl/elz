use super::ast;
use super::ast::Top;
use std::collections::HashMap;

mod mir;
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

pub fn generate_mir_program(program: &Vec<ast::Top>) -> Result<mir::MIR> {
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

    let Bind(name, expr) = get_main_binding(&ctx)?;

    let mut mir = mir::MIR { functions: vec![] };
    let lambda = check_main_is_lambda(expr)?;
    check_main_return_type(lambda)?;
    let expr = ensure_main_body_is_not_empty(lambda)?;
    let block = check_main_body_is_block(expr.as_ref())?;
    let mut stmts = vec![];
    for stmt in &block.statements {
        stmts.push(generate_stmt_mir(stmt)?);
    }
    let f = mir::Function {
        name: std::borrow::Cow::from(name.clone()),
        block: stmts,
    };
    mir.functions.push(f);
    Ok(mir)
}

pub fn generate_stmt_mir(stmt: &ast::Statement) -> Result<mir::Statement> {
    use ast::Statement::*;
    let stmt = match stmt {
        Return(e) => mir::Statement {
            statement: mir::mod_Statement::OneOfstatement::return_pb(generate_expr_mir(e)?),
        },
        _ => unimplemented!(),
    };
    Ok(stmt)
}

pub fn generate_expr_mir(expr: &ast::Expr) -> Result<mir::Expr> {
    use ast::Expr::*;
    let expr = match expr {
        Int(i) => mir::Expr {
            expr: mir::mod_Expr::OneOfexpr::int(*i),
        },
        _ => unimplemented!(),
    };
    Ok(expr)
}
