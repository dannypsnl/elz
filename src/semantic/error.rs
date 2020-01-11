use super::types::Type;
use crate::lexer::Location;
use thiserror::Error;

pub type Result<T> = std::result::Result<T, SemanticError>;

#[derive(Debug, Error)]
#[error("{location} {err}")]
pub struct SemanticError {
    location: Location,
    err: SemanticErrorVariant,
}

#[derive(Debug, Error)]
enum SemanticErrorVariant {
    #[error("name: `{}` be redefined", .0)]
    NameRedefined(String),
    #[error("type mismatched, expected: `{}` but got: `{}`", .0, .1)]
    TypeMismatched(Type, Type),
    #[error("no variable named: `{}`", .0)]
    NoVariableNamed(String),
    #[error("no type named: `{}`",  .0)]
    NoTypeNamed(String),
    #[error("call on non-function type: `{}`", .0)]
    CallOnNonFunctionType(Type),
    #[error("following fields must be inited but haven't: {}", ShowFieldsList(.0.to_vec()))]
    FieldsMissingInit(Vec<String>),
    #[error("cannot use class construction on a non-class type: {}", .0)]
    CannotConstructNonClassType(Type),
    #[error("cannot use class construction out of class scope")]
    CannotUseClassConstructionOutOfClass(),
    #[error("can only as subtype of a trait type, but got: {}", .got_type)]
    CanOnlyAsSubtypeOfATrait { got_type: Type },
}

impl SemanticError {
    fn new(location: &Location, err: SemanticErrorVariant) -> SemanticError {
        SemanticError {
            location: location.clone(),
            err,
        }
    }
    pub fn name_redefined<T: ToString>(location: &Location, name: T) -> SemanticError {
        SemanticError::new(
            location,
            SemanticErrorVariant::NameRedefined(name.to_string()),
        )
    }
    pub fn type_mismatched(location: &Location, expected: &Type, actual: &Type) -> SemanticError {
        SemanticError::new(
            location,
            SemanticErrorVariant::TypeMismatched(expected.clone(), actual.clone()),
        )
    }
    pub fn no_variable(location: &Location, name: &str) -> SemanticError {
        SemanticError::new(
            location,
            SemanticErrorVariant::NoVariableNamed(name.to_string()),
        )
    }
    pub fn no_type(location: &Location, name: &str) -> SemanticError {
        SemanticError::new(
            location,
            SemanticErrorVariant::NoTypeNamed(name.to_string()),
        )
    }
    pub fn call_on_non_function_type(location: &Location, typ: Type) -> SemanticError {
        SemanticError::new(location, SemanticErrorVariant::CallOnNonFunctionType(typ))
    }
    pub fn fields_missing_init(location: &Location, fields: Vec<String>) -> SemanticError {
        SemanticError::new(location, SemanticErrorVariant::FieldsMissingInit(fields))
    }
    pub fn cannot_construct_non_class_type(location: &Location, typ: Type) -> SemanticError {
        SemanticError::new(
            location,
            SemanticErrorVariant::CannotConstructNonClassType(typ),
        )
    }
    pub fn cannot_use_class_construction_out_of_class(location: &Location) -> SemanticError {
        SemanticError::new(
            location,
            SemanticErrorVariant::CannotUseClassConstructionOutOfClass(),
        )
    }
    pub fn can_only_as_subtype_of_a_trait(location: &Location, got_type: &Type) -> SemanticError {
        SemanticError::new(
            location,
            SemanticErrorVariant::CanOnlyAsSubtypeOfATrait {
                got_type: got_type.clone(),
            },
        )
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
