use super::ast;
use super::ast::*;
use std::collections::HashMap;

#[cfg(test)]
mod tests;
mod types;
mod error;

use types::Type;
use types::TypeVar;
use error::{
    Result,
    CheckError,
};

pub struct Context<'current, 'parent> {
    parent: Option<&'current Context<'parent, 'parent>>,
    type_environment: HashMap<String, Type>,
    type_var_id: HashMap<String, u64>,
    count: u64,
}

impl<'current, 'parent> Context<'current, 'parent> {
    /// new returns a new context for inference expression
    pub fn new() -> Context<'current, 'static> {
        Context {
            parent: None,
            type_environment: HashMap::new(),
            type_var_id: HashMap::new(),
            count: 0,
        }
    }
    pub fn with_parent(ctx: &'parent Context<'_, 'parent>) -> Context<'current, 'parent> {
        Context {
            parent: Some(ctx),
            type_environment: HashMap::new(),
            type_var_id: HashMap::new(),
            count: 0,
        }
    }

    fn get(&self, key: &String) -> Result<Type> {
        match self.type_environment.get(key) {
            Some(t) => Ok(t.clone()),
            None => if let Some(p_ctx) = self.parent {
                p_ctx.get(key)
            } else {
                Err(CheckError::not_found(key.clone()))
            },
        }
    }
}

#[allow(mutable_borrow_reservation_conflict)]
pub fn check_program(program: Vec<ast::Top>) -> Result<()> {
    use ast::Type;

    let mut ctx = Context::new();

    ctx.type_environment.insert("int".to_string(), types::Type::I64);
    for top_elem in program {
        match top_elem {
            Top::Binding(name, typ, expr) => {
                match typ {
                    Type::Unsure(_) => {
                        let mut c = Context::with_parent(&ctx);
                        let mut sub = Substitution::new();
                        ctx.type_environment.insert(name, infer_expr(&mut c, expr, &mut sub)?.0)
                    }
                    Type::Defined(_) => {
                        let mut c = Context::with_parent(&ctx);
                        ctx.type_environment.insert(name, types::Type::from_ast_type(&mut c, typ)?)
                    }
                }
            }
            _ => unimplemented!()
        };
    }
    Ok(())
}

pub fn infer_expr<'start_infer>(
    c: &mut Context,
    expr: Expr,
    substitution: &'start_infer mut Substitution,
) -> Result<(Type, &'start_infer mut Substitution)> {
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
                        return Err(CheckError::MismatchedArguments);
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
) -> Result<&'start_infer mut Substitution> {
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

pub fn occurs(x: &TypeVar, t: &Type) -> Result<()> {
    match t {
        Type::TypeVar(id) => {
            if id.0 == x.0 {
                Err(CheckError::CyclicType)
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