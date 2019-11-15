use crate::lexer::Location;

#[derive(Debug)]
pub enum SemanticError {
    NameRedefined(String, Location),
}

impl SemanticError {
    pub fn name_redefined<T: ToString>(name: T, location: Location) -> SemanticError {
        SemanticError::NameRedefined(name.to_string(), location)
    }
}

impl std::fmt::Display for SemanticError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        use SemanticError::*;
        match self {
            NameRedefined(name, loc) => write!(f, "{:?} name: {} be redefined", loc, name),
        }
    }
}
impl std::error::Error for SemanticError {
    fn description(&self) -> &str {
        use SemanticError::*;
        match self {
            NameRedefined(_, _) => "name redefined",
        }
    }
}
