use super::ast::*;
use std::collections::HashMap;

mod types;
#[cfg(test)]
mod tests;

use types::Type;
use types::TypeVar;

pub struct Context {
    type_environment: HashMap<String, Type>
}


pub fn infer_expr<'start_infer>(c: &mut Context, expr: Expr, substitution: &'start_infer mut Substitution) -> Result<(Type, &'start_infer mut Substitution), String> {
    match expr {
        Expr::Int(_) => Ok((Type::I64, substitution)),
        Expr::F64(_) => Ok((Type::F64, substitution)),
        Expr::String(_) => Ok((Type::String, substitution)),
        Expr::Identifier(access_chain) => {
            match c.type_environment.get(&access_chain) {
                Some(t) => Ok((t.clone(), substitution)),
                None => Err("not found".to_string())
            }
        }
        Expr::Binary(left_e, right_e, op) => {
            let (t1, substitution) = infer_expr(c, *left_e, substitution)?;
            let (t2, substitution) = infer_expr(c, *right_e, substitution)?;
            match op {
                Operator::Plus => {
                    let substitution = unify((Type::I64, t1), substitution)?;
                    let substitution = unify((Type::I64, t2), substitution)?;
                    Ok((Type::I64, substitution))
                }
            }
        }
        _ => unimplemented!()
    }
}


pub fn unify(target: (Type, Type), sub: &mut Substitution) -> Result<&mut Substitution, String> {
    if target.0 == target.1 {
        Ok(sub)
    } else {
        Err("unimplemented".to_string())
    }
}


pub struct Substitution(HashMap<TypeVar, Type>);

impl Substitution {}