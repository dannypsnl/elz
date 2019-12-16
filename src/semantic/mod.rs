use crate::ast::*;
mod error;
mod types;
use crate::lexer::Location;
use crate::semantic::types::TypeEnv;
use error::Result;
use types::Type;

pub struct SemanticChecker {
    type_env: TypeEnv,
}

impl SemanticChecker {
    pub fn new() -> SemanticChecker {
        SemanticChecker {
            type_env: TypeEnv::new(),
        }
    }
}

impl SemanticChecker {
    pub fn check_program(&mut self, ast: &Vec<TopAst>) -> Result<()> {
        for top in ast {
            use TopAst::*;
            match top {
                Class(c) => {
                    let typ = Type::new_class(c);
                    self.type_env.add_type(&c.location, &c.name, typ)?;
                    for static_method in &c.static_methods {
                        self.type_env.add_variable(
                            &static_method.location,
                            &format!("{}::{}", c.name, static_method.name).to_string(),
                            self.type_env.new_function_type(static_method)?,
                        )?;
                    }
                }
                _ => (),
            }
        }
        for top in ast {
            use TopAst::*;
            match top {
                Variable(v) => {
                    self.type_env
                        .add_variable(&v.location, &v.name, self.type_env.from(&v.typ)?)?
                }
                Function(f) => {
                    self.type_env.add_variable(
                        &f.location,
                        &f.name,
                        self.type_env.new_function_type(f)?,
                    )?;
                }
                _ => (),
            }
        }
        for top in ast {
            use TopAst::*;
            match top {
                Variable(v) => {
                    let typ = self.type_env.type_of_expr(&v.expr)?;
                    // show where error happened
                    // we are unifying <expr> and <type>, so <expr> location is better than
                    // variable define statement location
                    self.type_env
                        .unify(&v.expr.location, self.type_env.from(&v.typ)?, typ)?
                }
                Function(f) => self.check_function_body(&f.location, &f, &self.type_env)?,
                Class(c) => {
                    let mut class_type_env = TypeEnv::with_parent(&self.type_env);
                    for f in &c.fields {
                        let typ = class_type_env.from(&f.typ)?;
                        class_type_env.add_variable(&f.location, &f.name, typ)?;
                    }
                    class_type_env.in_class_scope = true;
                    for static_method in &c.static_methods {
                        class_type_env.add_variable(
                            &static_method.location,
                            &static_method.name,
                            self.type_env.new_function_type(static_method)?,
                        )?;
                    }
                    for static_method in &c.static_methods {
                        self.check_function_body(
                            &static_method.location,
                            &static_method,
                            &class_type_env,
                        )?;
                    }
                    for method in &c.methods {
                        self.check_function_body(&method.location, &method, &class_type_env)?;
                    }
                }
            }
        }
        Ok(())
    }

    fn check_function_body(&self, location: &Location, f: &Function, env: &TypeEnv) -> Result<()> {
        let return_type = self.type_env.from(&f.ret_typ)?;
        let mut type_env = TypeEnv::with_parent(env);
        for Parameter(p_type, p_name) in &f.parameters {
            type_env.add_variable(location, p_name, type_env.from(p_type)?)?;
        }
        match &f.body {
            Some(Body::Expr(e)) => {
                let e_type = type_env.type_of_expr(e)?;
                type_env.unify(location, return_type.clone(), e_type)
            }
            Some(Body::Block(b)) => {
                for stmt in &b.statements {
                    use StatementVariant::*;
                    match &stmt.value {
                        Return(e) => {
                            let typ = match e {
                                Some(e) => type_env.type_of_expr(e)?,
                                None => Type::Void,
                            };
                            type_env.unify(&stmt.location, return_type.clone(), typ)?
                        }
                        Variable(v) => {
                            let var_def_typ = type_env.from(&v.typ)?;
                            let var_typ = type_env.type_of_expr(&v.expr)?;
                            type_env.unify(&stmt.location, var_def_typ.clone(), var_typ)?;
                            type_env.add_variable(&stmt.location, &v.name, var_def_typ)?
                        }
                        Expression(func_call) => {
                            let func_call_ret_typ = type_env.type_of_expr(func_call)?;
                            type_env.unify(&stmt.location, Type::Void, func_call_ret_typ)?;
                        }
                    }
                }
                Ok(())
            }
            None => {
                // function declaration has no body need to check
                // e.g.
                // ```
                // foo(): void;
                // ```
                Ok(())
            }
        }
    }
}

// Must put code before tests module
#[cfg(test)]
mod tests;
