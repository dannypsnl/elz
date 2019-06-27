use super::ast;
use super::ast::*;
use std::collections::HashMap;

#[cfg(test)]
mod tests;
mod types;
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
            None => Err("not found".to_string()),
        }
    }
}

pub fn infer_expr<'start_infer>(
    c: &mut Context,
    expr: Expr,
    substitution: &'start_infer mut Substitution,
) -> Result<(Type, &'start_infer mut Substitution), String> {
    match expr {
        Expr::Int(_) => Ok((Type::I64, substitution)),
        Expr::F64(_) => Ok((Type::F64, substitution)),
        Expr::String(_) => Ok((Type::String, substitution)),
        Expr::Identifier(access_chain) => Ok((c.get(&access_chain)?, substitution)),
        Expr::Binary(left_e, right_e, op) => {
            let (t1, substitution) = infer_expr(c, *left_e, substitution)?;
            let (t2, substitution) = infer_expr(c, *right_e, substitution)?;
            match op {
                Operator::Plus => {
                    let substitution = unify((&Type::I64, &t1), substitution)?;
                    let substitution = unify((&Type::I64, &t2), substitution)?;
                    Ok((Type::I64, substitution))
                }
            }
        }
        Expr::Lambda(lambda) => {
            let return_type = Type::from_ast_type(c, lambda.return_type)?;
            let mut pts = vec![];
            for param in lambda.parameters {
                let param_type = Type::from_ast_type(c, param.0)?;
                c.type_environment.insert(param.1, param_type.clone());
                pts.push(param_type);
            }
            let (expression_type, substitution) = match lambda.body {
                Some(expr) => infer_expr(c, *expr, substitution)?,
                None => (
                    Type::from_ast_type(c, ast::Type::Unsure("a".to_string()))?,
                    substitution,
                ),
            };
            let substitution = unify((&return_type, &expression_type), substitution)?;
            Ok((Type::Lambda(pts, Box::new(return_type)), substitution))
        }
        Expr::FuncCall(func, args) => {
            let (typ, substitution) = infer_expr(c, *func, substitution)?;
            match typ {
                Type::Lambda(params, return_type) => {
                    if params.len() != args.len() {
                        return Err("mismatched arguments".to_string());
                    }
                    for i in 0..params.len() {
                        let (arg_type, substitution) =
                            infer_expr(c, args[i].expr.clone(), substitution)?;
                        unify((&arg_type, &params[i]), substitution)?;
                    }
                    Ok((*return_type, substitution))
                }
                _ => unimplemented!(),
            }
        }
        _ => unimplemented!(),
    }
}

pub fn unify<'start_infer>(
    target: (&Type, &Type),
    sub: &'start_infer mut Substitution,
) -> Result<&'start_infer mut Substitution, String> {
    let (x, t) = target;
    let (sub_x, sub_t) = (sub.get(x), sub.get(t));
    if sub_x == sub_t {
        Ok(sub)
    } else {
        match target {
            (Type::TypeVar(x), t) | (t, Type::TypeVar(x)) => {
                occurs(x, t)?;
                sub.0.insert(x.clone(), t.clone());
                Ok(sub)
            }
            _ => unimplemented!(),
        }
    }
}

pub fn occurs(x: &TypeVar, t: &Type) -> Result<(), String> {
    match t {
        Type::TypeVar(id) => {
            if id.0 == x.0 {
                Err("cyclic type detected".to_string())
            } else {
                Ok(())
            }
        }
        _ => Ok(()),
    }
}


pub struct Substitution(HashMap<TypeVar, Type>);

impl Substitution {
    pub fn new() -> Substitution {
        Substitution(HashMap::new())
    }

    fn get(&self, t: &Type) -> Type {
        match t {
            Type::TypeVar(t) => match self.0.get(&t) {
                Some(t) => t.clone(),
                None => Type::TypeVar(t.clone()),
            },
            _ => t.clone(),
        }
    }
}