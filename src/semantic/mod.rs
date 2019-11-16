use crate::ast::*;
mod error;
mod types;
use crate::lexer::Location;
use crate::semantic::types::TypeEnv;
use error::Result;
use error::SemanticError;
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
                    self.unify(location, Type::from(v.typ), self.type_of_expr(v.expr)?)?
                }
                Function(f) => self.check_function_body(location, f)?,
            }
        }
        Ok(())
    }

    fn check_function_body(&self, location: Location, f: Function) -> Result<()> {
        let return_type = Type::from(f.ret_typ);
        match f.body {
            Body::Expr(e) => self.unify(location, return_type.clone(), self.type_of_expr(e)?),
            Body::Block(b) => {
                let mut type_env = TypeEnv::with_parent(&self.type_env);
                for param in f.parameters {
                    type_env.add_variable(location, param.1, Type::from(param.0))?;
                }
                for stmt in b.statements {
                    use StatementVariant::*;
                    match stmt.value {
                        Return(e) => self.unify(stmt.location, return_type.clone(), self.type_of_expr(e)?)?,
                    }
                }
                Ok(())
            }
        }
    }

    fn type_of_expr(&self, expr: Expr) -> Result<Type> {
        use ExprVariant::*;
        let location = expr.location;
        let typ = match expr.value {
            Binary(l, r, op) => {
                let left_type = self.type_of_expr(*l)?;
                let right_type = self.type_of_expr(*r)?;
                match (left_type, right_type, op) {
                    (Type::Int, Type::Int, Operator::Plus) => Type::Int,
                    _ => panic!("unsupported operator"),
                }
            }
            F64(_) => Type::F64,
            Int(_) => Type::Int,
            String(_) => Type::String,
            FuncCall(f, args) => unimplemented!(),
            Identifier(id) => {
                let type_info = self.type_env.get_variable(location, id)?;
                type_info.typ
            }
        };
        Ok(typ)
    }

    pub fn unify(&self, location: Location, expected: Type, actual: Type) -> Result<()> {
        if expected == actual {
            Ok(())
        } else {
            Err(SemanticError::type_mismatched(location, expected, actual))
        }
    }
}
