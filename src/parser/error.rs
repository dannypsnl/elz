use crate::lexer::{TkType, Token};
use thiserror::Error;

pub type Result<T> = std::result::Result<T, ParseError>;

#[derive(Debug, Error)]
pub enum ParseError {
    #[error("{} expected one of {} but got {}", .1.location(), ShowTkTypeList(.0.to_vec()), .1.tk_type())]
    NotExpectedToken(Vec<TkType>, Token),
    #[error("meet eof when parsing")]
    EOF,
}

impl ParseError {
    pub fn not_expected_token(expected: Vec<TkType>, actual: Token) -> ParseError {
        use ParseError::*;
        NotExpectedToken(expected, actual)
    }
}

struct ShowTkTypeList(Vec<TkType>);
impl std::fmt::Display for ShowTkTypeList {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        for e in &self.0 {
            write!(f, "`{}` ", e)?;
        }
        write!(f, "")
    }
}
