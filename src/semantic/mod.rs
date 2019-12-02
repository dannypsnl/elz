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
                Variable(v) => self.type_env.add_variable(
                    top.location(),
                    top.name(),
                    Type::from(v.typ.clone()),
                )?,
                Function(f) => {
                    let typ = Type::new_function(f.clone());
                    self.type_env
                        .add_variable(top.location(), top.name(), typ)?;
                }
                Class(_) => unimplemented!(),
            }
        }
        for top in ast {
            use TopAst::*;
            let location = top.location();
            match top {
                Variable(v) => {
                    let location = v.expr.location.clone();
                    let typ = self.type_env.type_of_expr(v.expr.clone())?;
                    // show where error happened
                    // we are unifying <expr> and <type>, so <expr> location is better than
                    // variable define statement location
                    self.type_env
                        .unify(location, Type::from(v.typ.clone()), typ)?
                }
                Function(f) => self.check_function_body(location, f.clone())?,
                Class(_) => unimplemented!(),
            }
        }
        Ok(())
    }

    fn check_function_body(&self, location: Location, f: Function) -> Result<()> {
        let return_type = Type::from(f.ret_typ);
        let mut type_env = TypeEnv::with_parent(&self.type_env);
        for Parameter(p_type, p_name) in f.parameters {
            type_env.add_variable(location.clone(), p_name, Type::from(p_type))?;
        }
        match f.body {
            Some(Body::Expr(e)) => {
                let e_type = type_env.type_of_expr(e)?;
                type_env.unify(location, return_type.clone(), e_type)
            }
            Some(Body::Block(b)) => {
                for stmt in b.statements {
                    use StatementVariant::*;
                    match stmt.value {
                        Return(e) => {
                            let typ = match e {
                                Some(e) => type_env.type_of_expr(e)?,
                                None => Type::Void,
                            };
                            type_env.unify(stmt.location, return_type.clone(), typ)?
                        }
                        Variable(v) => {
                            let var_def_typ = Type::from(v.typ);
                            let var_typ = type_env.type_of_expr(v.expr)?;
                            type_env.unify(stmt.location.clone(), var_def_typ.clone(), var_typ)?;
                            type_env.add_variable(stmt.location, v.name, var_def_typ)?
                        }
                        FunctionCall(func_call) => {
                            let func_call_ret_typ = type_env.type_of_expr(func_call)?;
                            type_env.unify(stmt.location.clone(), Type::Void, func_call_ret_typ)?;
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
