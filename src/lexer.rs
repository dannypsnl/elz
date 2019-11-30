use strum_macros::Display;

#[derive(Display, Clone, Debug, PartialEq)]
pub enum TkType {
    #[strum(serialize = "<eof>")]
    EOF,
    #[strum(serialize = "<identifier>")]
    Ident,
    #[strum(serialize = "return")]
    Return,
    #[strum(serialize = "<integer>")]
    Integer,
    #[strum(serialize = "true")]
    True,
    #[strum(serialize = "false")]
    False,
    #[strum(serialize = "<string>")]
    String,
    #[strum(serialize = "+")]
    Plus,
    #[strum(serialize = "-")]
    Minus,
    #[strum(serialize = "*")]
    Multiple,
    #[strum(serialize = "/")]
    Divide,
    #[strum(serialize = ",")]
    Comma,
    #[strum(serialize = "=")]
    Equal,
    #[strum(serialize = "(")]
    OpenParen,
    #[strum(serialize = ")")]
    CloseParen,
    #[strum(serialize = "[")]
    OpenBracket,
    #[strum(serialize = "]")]
    CloseBracket,
    #[strum(serialize = "{")]
    OpenBrace,
    #[strum(serialize = "}")]
    CloseBrace,
    #[strum(serialize = ":")]
    Colon,
    #[strum(serialize = "::")]
    Accessor,
    #[strum(serialize = ";")]
    Semicolon,
    #[strum(serialize = "<comment>")]
    Comment,
}

#[derive(Clone, Debug, PartialEq)]
pub struct Location {
    file_name: String,
    line: u32,
    column: u32,
}

impl Location {
    pub fn from(line: u32, column: u32) -> Location {
        Location::new("", line, column)
    }
    pub fn new<T: ToString>(file_name: T, line: u32, column: u32) -> Location {
        Location {
            file_name: file_name.to_string(),
            line,
            column,
        }
    }
}

impl std::fmt::Display for Location {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "{}:{}:{}", self.file_name, self.line, self.column)
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct Token(Location, TkType, String);

impl Token {
    pub fn location(&self) -> Location {
        self.0.clone()
    }
    pub fn tk_type(&self) -> &TkType {
        &self.1
    }
    pub fn value(&self) -> String {
        self.2.clone()
    }
}

impl std::fmt::Display for Token {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(
            f,
            "loc: {}, type: {}, `{}`",
            self.location(),
            self.tk_type(),
            self.value()
        )
    }
}

enum State {
    Fn(fn(&mut Lexer) -> State),
    EOF,
}

struct Lexer {
    file_name: String,
    code: Vec<char>,
    tokens: Vec<Token>,
    state_fn: State,
    start: usize,
    offset: usize,
    // (line, pos) represent the position for user
    pos: u32,
    line: u32,
}

impl Lexer {
    fn new<T: Into<String>>(file_name: T, code: T) -> Lexer {
        Lexer {
            file_name: file_name.into(),
            code: code.into().chars().collect(),
            tokens: vec![],
            state_fn: State::Fn(whitespace),
            start: 0,
            offset: 0,
            pos: 0,
            line: 1,
        }
    }

    fn ignore(&mut self) {
        self.pos += (self.offset - self.start) as u32;
        self.start = self.offset;
    }
    fn peek(&self) -> Option<char> {
        match self.code.get(self.offset) {
            Some(c) => Some(*c),
            None => None,
        }
    }
    fn next(&mut self) -> Option<char> {
        self.offset += 1;
        self.peek()
    }
    fn new_token(&mut self, token_type: TkType, value: String) -> Token {
        Token(
            Location::new(self.file_name.clone(), self.line, self.pos),
            token_type,
            value,
        )
    }
    fn emit(&mut self, token_type: TkType) {
        let s: String = self.code[self.start..self.offset].into_iter().collect();
        let tok = match s.as_str() {
            "return" => self.new_token(TkType::Return, s),
            "true" => self.new_token(TkType::True, s),
            "false" => self.new_token(TkType::False, s),
            _ => self.new_token(token_type.clone(), s),
        };
        match token_type {
            TkType::Comment => {}
            _ => self.tokens.push(tok),
        }
        self.ignore();
    }
}

fn whitespace(lexer: &mut Lexer) -> State {
    while let Some(c) = lexer.peek() {
        if c == ' ' || c == '\n' {
            if c == '\n' {
                lexer.next();
                lexer.start = lexer.offset;
                lexer.pos = 0;
                lexer.line += 1;
            } else {
                lexer.next();
            }
        } else {
            break;
        }
    }
    lexer.ignore();

    match lexer.peek() {
        Some(_c @ '0'..='9') => State::Fn(number),
        Some('=') => {
            lexer.next();
            lexer.emit(TkType::Equal);
            State::Fn(whitespace)
        }
        Some(',') => {
            lexer.next();
            lexer.emit(TkType::Comma);
            State::Fn(whitespace)
        }
        Some('+') => {
            lexer.next();
            lexer.emit(TkType::Plus);
            State::Fn(whitespace)
        }
        Some('-') => {
            lexer.next();
            lexer.emit(TkType::Minus);
            State::Fn(whitespace)
        }
        Some('*') => {
            lexer.next();
            lexer.emit(TkType::Multiple);
            State::Fn(whitespace)
        }
        Some('/') => {
            lexer.next();
            if lexer.peek() == Some('/') {
                // found `//`
                lexer.next();
                while let Some(c) = lexer.peek() {
                    if c == '\n' {
                        break;
                    } else {
                        lexer.next();
                    }
                }
                lexer.emit(TkType::Comment);
            } else {
                lexer.emit(TkType::Divide);
            }
            State::Fn(whitespace)
        }
        Some('(') => {
            lexer.next();
            lexer.emit(TkType::OpenParen);
            State::Fn(whitespace)
        }
        Some(')') => {
            lexer.next();
            lexer.emit(TkType::CloseParen);
            State::Fn(whitespace)
        }
        Some('[') => {
            lexer.next();
            lexer.emit(TkType::OpenBracket);
            State::Fn(whitespace)
        }
        Some(']') => {
            lexer.next();
            lexer.emit(TkType::CloseBracket);
            State::Fn(whitespace)
        }
        Some('{') => {
            lexer.next();
            lexer.emit(TkType::OpenBrace);
            State::Fn(whitespace)
        }
        Some('}') => {
            lexer.next();
            lexer.emit(TkType::CloseBrace);
            State::Fn(whitespace)
        }
        Some(':') => {
            lexer.next();
            if lexer.peek() == Some(':') {
                lexer.next();
                lexer.emit(TkType::Accessor);
            } else {
                lexer.emit(TkType::Colon);
            }
            State::Fn(whitespace)
        }
        Some(';') => {
            lexer.next();
            lexer.emit(TkType::Semicolon);
            State::Fn(whitespace)
        }
        Some('"') => State::Fn(string),
        Some(c) => {
            if in_identifier_set(c) {
                State::Fn(ident)
            } else {
                panic!("Not implemented for {} yet", c);
            }
        }
        None => State::EOF,
    }
}

fn in_identifier_set(c: char) -> bool {
    c.is_alphanumeric() || c == '_'
}

fn ident(lexer: &mut Lexer) -> State {
    while let Some(c) = lexer.next() {
        if !in_identifier_set(c) {
            break;
        }
    }
    lexer.emit(TkType::Ident);
    State::Fn(whitespace)
}

fn string(lexer: &mut Lexer) -> State {
    while let Some(c) = lexer.next() {
        if c == '\\' {
            lexer.next();
        }
        if c == '"' {
            break;
        }
    }
    lexer.next();
    lexer.emit(TkType::String);
    State::Fn(whitespace)
}

fn number(lexer: &mut Lexer) -> State {
    while let Some(c) = lexer.next() {
        if !c.is_digit(10) {
            break;
        }
    }
    lexer.emit(TkType::Integer);
    State::Fn(whitespace)
}

pub fn lex<T: Into<String>>(file_name: T, source: T) -> Vec<Token> {
    let mut lexer = Lexer::new(file_name, source);
    while let State::Fn(f) = lexer.state_fn {
        lexer.state_fn = f(&mut lexer);
    }
    lexer.emit(TkType::EOF);
    lexer.tokens
}

#[cfg(test)]
mod tests {
    use self::TkType::*;
    use super::*;

    #[test]
    fn test_unicode_identifier_a() {
        let code = "測試: int = 1";

        assert_eq!(
            lex("", code),
            vec![
                Token(Location::from(1, 0), Ident, "測試".to_string()),
                Token(Location::from(1, 2), Colon, ":".to_string()),
                Token(Location::from(1, 4), Ident, "int".to_string()),
                Token(Location::from(1, 8), Equal, "=".to_string()),
                Token(Location::from(1, 10), Integer, "1".to_string()),
                Token(Location::from(1, 11), EOF, "".to_string()),
            ]
        );
    }

    #[test]
    fn get_number_tokens() {
        let ts = lex("", "10 30");
        assert_eq!(
            ts,
            vec![
                Token(Location::from(1, 0), Integer, "10".to_string()),
                Token(Location::from(1, 3), Integer, "30".to_string()),
                Token(Location::from(1, 5), EOF, "".to_string()),
            ]
        );
    }

    #[test]
    fn get_ident_tokens() {
        let ts = lex("", " abc6");
        assert_eq!(
            ts,
            vec![
                Token(Location::from(1, 1), Ident, "abc6".to_string()),
                Token(Location::from(1, 5), EOF, "".to_string()),
            ]
        )
    }

    #[test]
    fn get_escape_char_in_string() {
        let ts = lex("", "\"\\\"\"");
        assert_eq!(
            ts,
            vec![
                Token(Location::from(1, 0), String, "\"\\\"\"".to_string()),
                Token(Location::from(1, 4), EOF, "".to_string()),
            ]
        )
    }

    #[test]
    fn comment_would_be_discard() {
        let ts = lex("", "//\n1");
        assert_eq!(
            ts,
            vec![
                Token(Location::from(2, 0), Integer, "1".to_string()),
                Token(Location::from(2, 1), EOF, "".to_string()),
            ]
        )
    }
}
