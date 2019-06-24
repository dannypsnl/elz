use super::ast::*;
use std::collections::HashMap;

mod types;

use types::Type;

pub struct Context {
    type_environment: HashMap<String, Type>
}


pub fn infer_expr(c: &mut Context, expr: Expr) -> Result<Type, String> {
    match expr {
        Expr::Int(_) => Ok(Type::I64),
        Expr::F64(_) => Ok(Type::F64),
        Expr::String(_) => Ok(Type::String),
        Expr::Identifier(access_chain) => {
            match c.type_environment.get(&access_chain) {
                Some(t) => Ok(t.clone()),
                None => Err("not found".to_string())
            }
        }
        _ => unimplemented!()
    }
}