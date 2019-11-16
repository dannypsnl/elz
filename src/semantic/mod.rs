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
            self.type_env.add_variable(top.clone())?;
        }
        for top in ast {
            use TopAst::*;
            let location = top.location();
            match top {
                Variable(v) => self.unify(location, Type::from(v.typ), self.type_of(v.expr)?)?,
                Function(f) => {}
            }
        }
        Ok(())
    }

    fn type_of(&self, expr: Expr) -> Result<Type> {
        use ExprVariant::*;
        let location = expr.location();
        let typ = match expr.value {
            Binary(l, r, op) => {
                let left_type = self.type_of(*l)?;
                let right_type = self.type_of(*r)?;
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
