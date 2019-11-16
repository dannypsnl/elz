use super::*;
use crate::parser::Parser;
use std::error::Error;

#[test]
fn test_redefine_variable_would_get_error() {
    let code = "\
    x: int = 1
    x: int = 2
    ";
    let err = check_code(code).unwrap_err();
    assert_eq!(err.description(), "name redefined");
}

#[test]
fn test_function_and_variable_use_the_same_space() {
    let code = "\
    x: int = 1
    x(): void {}
    ";
    let err = check_code(code).unwrap_err();
    assert_eq!(err.description(), "name redefined");
}

#[test]
fn test_type_mismatched() {
    let code = "\
    x: int = \"str\"
    ";
    let err = check_code(code).unwrap_err();
    assert_eq!(err.description(), "type mismatched");
}

// helpers, must put tests before this line
fn check_code<T: ToString>(code: T) -> Result<()> {
    let program = Parser::parse_program(code.to_string()).unwrap();
    let mut checker = SemanticChecker::new();
    checker.check_program(program)
}
