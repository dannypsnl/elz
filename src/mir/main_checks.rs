use super::super::ast;
use super::Bind;
use super::Context;
use super::MIRError;
use super::Result;

pub(crate) fn check_main_body_is_block(expr: &ast::Expr) -> Result<&ast::Block> {
    match expr {
        ast::Expr::Block(block) => Ok(block),
        _ => Err(MIRError::new("main lambda must be a block")),
    }
}

pub(crate) fn ensure_main_body_is_not_empty(lambda: &ast::Lambda) -> Result<Box<ast::Expr>> {
    match lambda.body.clone() {
        Some(expr) => Ok(expr),
        _ => Err(MIRError::new("main lambda can't have an empty body")),
    }
}

pub(crate) fn check_main_return_type(lambda: &ast::Lambda) -> Result<()> {
    match lambda.return_type {
        ast::Type::Unit => Ok(()),
        _ => Err(MIRError::new("main lambda must return unit type")),
    }
}

pub(crate) fn check_main_is_lambda(expr: &ast::Expr) -> Result<&ast::Lambda> {
    match expr {
        ast::Expr::Lambda(lambda) => Ok(lambda),
        _ => Err(MIRError::new(
            "main is the entry point of executable and must be a lambda",
        )),
    }
}

pub(crate) fn get_main_binding(ctx: &Context) -> Result<&Bind> {
    match ctx.binding_map.get("main") {
        Some(bind) => Ok(bind),
        None => Err(MIRError::NoMain),
    }
}
