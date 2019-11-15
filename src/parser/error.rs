use crate::lexer::{TkType, Token};

pub type Result<T> = std::result::Result<T, ParseError>;

#[derive(Debug)]
pub enum ParseError {
    NotExpectedToken(Vec<TkType>, Token),
    EOF,
}

impl ParseError {
    pub fn not_expected_token(expected: Vec<TkType>, actual: Token) -> ParseError {
        use ParseError::*;
        NotExpectedToken(expected, actual)
    }
}

impl std::fmt::Display for ParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        use ParseError::*;
        match self {
            NotExpectedToken(expected, actual) => write!(
                f,
                "{:?} expected one of {:?} but got {}",
                actual.location(),
                expected,
                actual
            ),
            EOF => write!(f, "meet eof when parsing"),
        }
    }
}
impl std::error::Error for ParseError {
    fn description(&self) -> &str {
        use ParseError::*;
        match self {
            NotExpectedToken(_, _) => "not expected token",
            EOF => "eof",
        }
    }
}
