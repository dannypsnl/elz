use super::ast::*;
use std::collections::HashMap;

mod types;
#[cfg(test)]
mod tests;

use types::Type;
use types::TypeVar;

pub struct Context {
    type_environment: HashMap<String, Type>,
    type_var_id: HashMap<String, u64>,
    count: u64,
}

impl Context {
    /// new returns a new context for inference expression
    pub fn new() -> Context {
        Context {
            type_environment: HashMap::new(),
            type_var_id: HashMap::new(),
            count: 0,
        }
    }

    fn get(&self, key: &String) -> Result<Type, String> {
        match self.type_environment.get(key) {
            Some(t) => Ok(t.clone()),
            None => Err("not found".to_string())
        }
    }
}


pub fn infer_expr<'start_infer>(c: &mut Context, expr: Expr, substitution: &'start_infer mut Substitution) -> Result<(Type, &'start_infer mut Substitution), String> {
    match expr {
        Expr::Int(_) => Ok((Type::I64, substitution)),
        Expr::F64(_) => Ok((Type::F64, substitution)),
        Expr::String(_) => Ok((Type::String, substitution)),
        Expr::Identifier(access_chain) => {
            Ok((c.get(&access_chain)?, substitution))
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
        Expr::Lambda(lambda) => {
            let return_type = Type::from_ast_type(c, lambda.return_type)?;
            let mut pts = vec![];
            for param in lambda.parameters {
                let param_type = Type::from_ast_type(c, param.0)?;
                pts.push(param_type);
            }
            Ok((
                Type::Lambda(pts, Box::new(return_type)),
                substitution
            ))
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

impl Substitution {
    pub fn new() -> Substitution {
        Substitution(
            HashMap::new()
        )
    }
}