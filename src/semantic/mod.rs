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
    pub fn check_program(&mut self, ast: Vec<TopAst>) -> Result<()> {
        for top in &ast {
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
            }
        }
        for top in ast {
            use TopAst::*;
            let location = top.location();
            match top {
                Variable(v) => {
                    let typ = self.type_env.type_of_expr(v.expr)?;
                    self.type_env.unify(location, Type::from(v.typ), typ)?
                }
                Function(f) => self.check_function_body(location, f)?,
            }
        }
        Ok(())
    }

    fn check_function_body(&self, location: Location, f: Function) -> Result<()> {
        let return_type = Type::from(f.ret_typ);
        let mut type_env = TypeEnv::with_parent(&self.type_env);
        for Parameter(p_type, p_name) in f.parameters {
            type_env.add_variable(location, p_name, Type::from(p_type))?;
        }
        match f.body {
            Body::Expr(e) => {
                self.type_env
                    .unify(location, return_type.clone(), type_env.type_of_expr(e)?)
            }
            Body::Block(b) => {
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
                    }
                }
                Ok(())
            }
        }
    }
}

// Must put code before tests module
#[cfg(test)]
mod tests;
