use crate::lexer::{Location, TkType, Token};
use thiserror::Error;

pub type Result<T> = std::result::Result<T, ParseError>;

#[derive(Debug, Error)]
#[error("{location} {err}")]
pub struct ParseError {
    location: Location,
    err: ParseErrorVariant,
}

#[derive(Debug, Error)]
pub enum ParseErrorVariant {
    #[error("expected one of {} but got {}", ShowTkTypeList(.0.to_vec()), .1.tk_type())]
    NotExpectedToken(Vec<TkType>, Token),
    #[error("meet eof when parsing")]
    EOF,
}

impl ParseError {
    pub fn not_expected_token(expected: Vec<TkType>, actual: Token) -> ParseError {
        use ParseErrorVariant::*;
        ParseError {
            location: actual.location(),
            err: NotExpectedToken(expected, actual),
        }
    }
    pub fn eof(location: &Location) -> ParseError {
        ParseError {
            location: location.clone(),
            err: ParseErrorVariant::EOF,
        }
    }

    pub fn location(&self) -> Location {
        self.location.clone()
    }
    pub fn message(&self) -> String {
        use ParseErrorVariant::*;
        match self.err {
            NotExpectedToken(..) => "not expected token",
            EOF => "eof",
        }
        .to_string()
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
