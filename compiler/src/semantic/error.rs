pub type Result<T> = std::result::Result<T, CheckError>;

#[derive(Debug)]
pub enum CheckError {
    NotFound(String),
    MismatchedArguments,
    CyclicType,
}

use CheckError::*;

impl CheckError {
    pub fn not_found(identifier: String) -> CheckError {
        NotFound(identifier)
    }
    pub fn cyclic_type() -> CheckError {
        CyclicType
    }
    pub fn mismatch_arguments() -> CheckError {
        MismatchedArguments
    }
}

impl std::fmt::Display for CheckError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            NotFound(name) => write!(f, "no any identifier name: {}", name),
            CyclicType => write!(f, "cyclic type detected"),
            MismatchedArguments => write!(f, "mismatched arguments"),
        }
    }
}

impl std::error::Error for CheckError {
    fn description(&self) -> &str {
        match self {
            NotFound(_) => "not found",
            CyclicType => "cyclic type detected",
            MismatchedArguments => "mismatched arguments",
        }
    }
}
