use super::ast;
use super::ast::*;
use std::collections::HashMap;

mod error;
#[cfg(test)]
mod tests;
mod types;

use error::{CheckError, Result};
use types::Type;
use types::TypeVar;

pub struct Context {
    parent: Option<*const Context>,
    type_map: HashMap<String, Type>,
    type_environment: HashMap<String, Type>,
    type_var_id: HashMap<String, u64>,
    count: u64,
}

impl Context {
    /// new returns a new context for inference expression
    pub fn new() -> Context {
        let mut ctx = Context {
            parent: None,
            type_map: HashMap::new(),
            type_environment: HashMap::new(),
            type_var_id: HashMap::new(),
            count: 0,
        };
        ctx.add_type("int".to_string(), types::Type::I64);

        ctx
    }
    pub fn with_parent(ctx: &Context) -> Context {
        Context {
            parent: Some(ctx),
            type_map: HashMap::new(),
            type_environment: HashMap::new(),
            type_var_id: HashMap::new(),
            count: 0,
        }
    }

    fn get_type(&self, type_name: &String) -> Result<Type> {
        match self.type_map.get(type_name) {
            Some(t) => Ok(t.clone()),
            None => {
                if let Some(p_ctx) = self.parent {
                    unsafe { p_ctx.as_ref() }.unwrap().get_type(type_name)
                } else {
                    Err(CheckError::not_found(type_name.clone()))
                }
            }
        }
    }
    fn add_type(&mut self, type_name: String, typ: Type) {
        self.type_map.insert(type_name, typ);
    }

    fn get_identifier(&self, key: &String) -> Result<Type> {
        match self.type_environment.get(key) {
            Some(t) => Ok(t.clone()),
            None => {
                if let Some(p_ctx) = self.parent {
                    unsafe { p_ctx.as_ref() }.unwrap().get_identifier(key)
                } else {
                    Err(CheckError::not_found(key.clone()))
                }
            }
        }
    }
    fn add_identifier(&mut self, key: String, typ: Type) {
        self.type_environment.insert(key, typ);
    }
}

#[allow(mutable_borrow_reservation_conflict)]
pub fn check_program(program: &Vec<ast::Top>) -> Result<()> {
    use ast::Type;

    let mut ctx = Context::new();

    for top_elem in program {
        match top_elem {
            Top::Binding(name, typ, expr) => {
                let expr_type = match typ {
                    Type::None | Type::Unsure(_) => {
                        let (expr_type, _) = infer_expr(
                            &mut Context::with_parent(&ctx),
                            expr,
                            &mut Substitution::new(),
                        )?;
                        expr_type
                    }
                    Type::Defined(_) | Type::Unit => {
                        let mut c = Context::with_parent(&ctx);
                        let (expr_type, _) = infer_expr(&mut c, expr, &mut Substitution::new())?;
                        let defined_type = types::Type::from_ast_type(&mut c, typ)?;
                        if expr_type != defined_type {
                            return Err(CheckError::type_mismatched(defined_type, expr_type));
                        }
                        defined_type
                    }
                };
                ctx.add_identifier(name.clone(), expr_type);
            }
            _ => unimplemented!(),
        };
    }
    Ok(())
}

#[allow(mutable_borrow_reservation_conflict)]
pub fn infer_expr<'start_infer>(
    c: &mut Context,
    expr: &Expr,
    substitution: &'start_infer mut Substitution,
) -> Result<(Type, &'start_infer mut Substitution)> {
    match expr {
        Expr::Int(_) => Ok((Type::I64, substitution)),
        Expr::F64(_) => Ok((Type::F64, substitution)),
        Expr::String(_) => Ok((Type::String, substitution)),
        Expr::Identifier(access_chain) => Ok((c.get_identifier(&access_chain)?, substitution)),
        Expr::Binary(left_e, right_e, op) => {
            let (t1, substitution) = infer_expr(c, left_e, substitution)?;
            let (t2, substitution) = infer_expr(c, right_e, substitution)?;
            match op {
                Operator::Plus => {
                    let substitution = unify((&Type::I64, &t1), substitution)?;
                    let substitution = unify((&Type::I64, &t2), substitution)?;
                    Ok((Type::I64, substitution))
                }
            }
        }
        Expr::Block(block) => {
            let mut return_type = Type::Unit;
            let mut substitution = substitution;
            let c = &mut Context::with_parent(c);
            for stmt in &block.statements {
                match stmt {
                    Statement::Let { name, typ, expr } => {
                        let mut binding_ctx = Context::with_parent(c);
                        match typ {
                            ast::Type::None | ast::Type::Unsure(_) => {
                                let (typ, sub) = infer_expr(&mut binding_ctx, &expr, substitution)?;
                                substitution = sub;
                                c.add_identifier(name.clone(), typ);
                            }
                            ast::Type::Defined(_) | ast::Type::Unit => {
                                c.add_identifier(
                                    name.clone(),
                                    types::Type::from_ast_type(&mut binding_ctx, typ)?,
                                );
                            }
                        }
                    }
                    Statement::Return(expr) => {
                        let (rt, sub) = infer_expr(c, &expr, substitution)?;
                        return_type = rt;
                        substitution = sub;
                    }
                }
            }
            Ok((return_type, substitution))
        }
        Expr::Lambda(lambda) => {
            let return_type = Type::from_ast_type(c, &lambda.return_type)?;
            let mut new_ctx = Context::with_parent(c);
            let pts: Result<Vec<_>> = lambda
                .parameters
                .iter()
                .map(|param| {
                    let ast::Parameter(param_defined_type, param_name) = param;
                    let param_type = Type::from_ast_type(&mut new_ctx, param_defined_type);
                    if let Ok(typ) = &param_type {
                        new_ctx.add_identifier(param_name.clone(), typ.clone());
                    }
                    param_type
                })
                .collect();
            let pts = pts?;
            let (expression_type, substitution) = match &lambda.body {
                Some(expr) => infer_expr(&mut new_ctx, expr, substitution)?,
                None => (
                    Type::from_ast_type(&mut new_ctx, &ast::Type::Unsure("a".to_string()))?,
                    substitution,
                ),
            };
            let substitution = unify((&return_type, &expression_type), substitution)?;
            Ok((Type::Lambda(pts, Box::new(return_type)), substitution))
        }
        Expr::FuncCall(func, args) => {
            let (typ, substitution) = infer_expr(c, func, substitution)?;
            match typ {
                Type::Lambda(params, return_type) => {
                    if params.len() != args.len() {
                        return Err(CheckError::MismatchedArguments);
                    }
                    for i in 0..params.len() {
                        let (arg_type, substitution) = infer_expr(c, &args[i].expr, substitution)?;
                        unify((&arg_type, &params[i]), substitution)?;
                    }
                    Ok((*return_type, substitution))
                }
                _ => unimplemented!(),
            }
        }
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
            (left, right) => {
                if sub.get(left) != sub.get(right) {
                    Err(CheckError::type_mismatched(left.clone(), right.clone()))
                } else {
                    Ok(sub)
                }
            }
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
