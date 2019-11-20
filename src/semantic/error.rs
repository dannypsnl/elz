use super::types::Type;
use crate::lexer::Location;

pub type Result<T> = std::result::Result<T, SemanticError>;

#[derive(Debug)]
pub struct SemanticError {
    location: Location,
    err: SemanticErrorVariant,
}

#[derive(Debug)]
pub enum SemanticErrorVariant {
    NameRedefined(String),
    TypeMismatched(Type, Type),
    NoVariableNamed(String),
    CallOnNonFunctionType(Type),
}

impl SemanticError {
    pub fn name_redefined<T: ToString>(location: Location, name: T) -> SemanticError {
        SemanticError {
            location,
            err: SemanticErrorVariant::NameRedefined(name.to_string()),
        }
    }
    pub fn type_mismatched(location: Location, expected: Type, actual: Type) -> SemanticError {
        SemanticError {
            location,
            err: SemanticErrorVariant::TypeMismatched(expected, actual),
        }
    }
    pub fn no_variable(location: Location, name: String) -> SemanticError {
        SemanticError {
            location,
            err: SemanticErrorVariant::NoVariableNamed(name),
        }
    }
    pub fn call_on_non_function_type(location: Location, typ: Type) -> SemanticError {
        SemanticError {
            location,
            err: SemanticErrorVariant::CallOnNonFunctionType(typ),
        }
    }
}

impl std::fmt::Display for SemanticError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "{} ", self.location)?;
        use SemanticErrorVariant::*;
        match &self.err {
            NameRedefined(name) => write!(f, "name: {} be redefined", name),
            TypeMismatched(expected, actual) => {
                write!(f, "expected: {} but got: {}", expected, actual)
            }
            NoVariableNamed(name) => write!(f, "no variable named: {}", name),
            CallOnNonFunctionType(typ) => write!(f, "call on non-function type: {}", typ),
        }
    }
}
impl std::error::Error for SemanticError {
    fn description(&self) -> &str {
        use SemanticErrorVariant::*;
        match self.err {
            NameRedefined(_) => "name redefined",
            TypeMismatched(_, _) => "type mismatched",
            NoVariableNamed(_) => "no variable named",
            CallOnNonFunctionType(_) => "call on non-function type",
        }
    }
}
