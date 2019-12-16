use strum_macros::Display;

#[derive(Display, Clone, Debug, PartialEq)]
pub enum TkType {
    #[strum(serialize = "<eof>")]
    EOF,
    // literal
    #[strum(serialize = "<identifier>")]
    Identifier,
    #[strum(serialize = "<integer>")]
    Integer,
    #[strum(serialize = "<string>")]
    String,
    // keyword
    #[strum(serialize = "return")]
    Return,
    #[strum(serialize = "class")]
    Class,
    #[strum(serialize = "true")]
    True,
    #[strum(serialize = "false")]
    False,
    // symbol
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
    #[strum(serialize = ".")]
    Dot,
    // ignored
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
    pub fn none() -> Location {
        Location::from(0, 0)
    }
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
            "class" => self.new_token(TkType::Class, s),
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
        Some('.') => {
            lexer.next();
            lexer.emit(TkType::Dot);
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
    lexer.emit(TkType::Identifier);
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
mod tests;
