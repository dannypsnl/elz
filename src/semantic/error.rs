use super::types::Type;
use crate::lexer::Location;
use thiserror::Error;

pub type Result<T> = std::result::Result<T, SemanticError>;

#[derive(Debug, Error)]
pub enum SemanticError {
    #[error("{} name: `{}` be redefined", .0, .1)]
    NameRedefined(Location, String),
    #[error("{} type mismatched, expected: `{}` but got: `{}`", .0, .1, .2)]
    TypeMismatched(Location, Type, Type),
    #[error("{} no variable named: `{}`", .0, .1)]
    NoVariableNamed(Location, String),
    #[error("{} no type named: `{}`", .0,  .1)]
    NoTypeNamed(Location, String),
    #[error("{} call on non-function type: `{}`", .0, .1)]
    CallOnNonFunctionType(Location, Type),
    #[error("{} following fields must be inited but haven't: {}", .0, ShowFieldsList(.1.to_vec()))]
    FieldsMissingInit(Location, Vec<String>),
    #[error("{} cannot use class construct on a non-class type: {}", .0, .1)]
    CannotConstructNonClassType(Location, Type),
}

impl SemanticError {
    pub fn name_redefined<T: ToString>(location: &Location, name: T) -> SemanticError {
        SemanticError::NameRedefined(location.clone(), name.to_string())
    }
    pub fn type_mismatched(location: &Location, expected: Type, actual: Type) -> SemanticError {
        SemanticError::TypeMismatched(location.clone(), expected, actual)
    }
    pub fn no_variable(location: &Location, name: String) -> SemanticError {
        SemanticError::NoVariableNamed(location.clone(), name)
    }
    pub fn no_type(location: &Location, name: String) -> SemanticError {
        SemanticError::NoTypeNamed(location.clone(), name)
    }
    pub fn call_on_non_function_type(location: &Location, typ: Type) -> SemanticError {
        SemanticError::CallOnNonFunctionType(location.clone(), typ)
    }
    pub fn fields_missing_init(location: &Location, fields: Vec<String>) -> SemanticError {
        SemanticError::FieldsMissingInit(location.clone(), fields)
    }
    pub fn cannot_construct_non_class_type(location: &Location, typ: Type) -> SemanticError {
        SemanticError::CannotConstructNonClassType(location.clone(), typ)
    }
}

struct ShowFieldsList(Vec<String>);
impl std::fmt::Display for ShowFieldsList {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        for e in &self.0 {
            write!(f, "`{}` ", e)?;
        }
        write!(f, "")
    }
}
