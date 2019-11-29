use super::types::Type;
use crate::lexer::Location;
use thiserror::Error;

pub type Result<T> = std::result::Result<T, SemanticError>;

#[derive(Debug, Error)]
pub enum SemanticError {
    #[error("{} name: `{}` be redefined", .0, .1)]
    NameRedefined(Location, String),
    #[error("{} expected: `{}` but got: `{}`", .0, .1, .2)]
    TypeMismatched(Location, Type, Type),
    #[error("{} no variable named: `{}`", .0, .1)]
    NoVariableNamed(Location, String),
    #[error("{} call on non-function type: `{}`", .0, .1)]
    CallOnNonFunctionType(Location, Type),
}

impl SemanticError {
    pub fn name_redefined<T: ToString>(location: Location, name: T) -> SemanticError {
        SemanticError::NameRedefined(location, name.to_string())
    }
    pub fn type_mismatched(location: Location, expected: Type, actual: Type) -> SemanticError {
        SemanticError::TypeMismatched(location, expected, actual)
    }
    pub fn no_variable(location: Location, name: String) -> SemanticError {
        SemanticError::NoVariableNamed(location, name)
    }
    pub fn call_on_non_function_type(location: Location, typ: Type) -> SemanticError {
        SemanticError::CallOnNonFunctionType(location, typ)
    }
}
