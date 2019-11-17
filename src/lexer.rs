#[derive(Clone, Debug, PartialEq)]
pub enum TkType {
    EOF,
    /// e.g. a, ab, foo
    Ident,
    /// return
    Return,
    /// e.g. 1, 10, 34
    Integer,
    /// true
    True,
    /// false
    False,
    /// "string_literal"
    String,
    /// +
    Plus,
    /// ,
    Comma,
    /// =
    Assign,
    /// (
    LParen,
    /// )
    RParen,
    /// [
    LBracket,
    /// ]
    RBracket,
    /// {
    LBrace,
    /// }
    RBrace,
    /// :
    Colon,
    /// ::
    Accessor,
    /// ;
    Semicolon,
}

impl std::fmt::Display for TkType {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        use TkType::*;
        let r = match self {
            EOF => "<eof>",
            Ident => "<identifier>",
            Return => "return",
            Integer => "<integer>",
            True => "true",
            False => "false",
            String => "<string>",
            Plus => "+",
            Comma => ",",
            Assign => "=",
            LParen => "(",
            RParen => ")",
            LBracket => "[",
            RBracket => "]",
            LBrace => "{",
            RBrace => "}",
            Colon => ":",
            Accessor => "::",
            Semicolon => ";",
        };
        write!(f, "{}", r)
    }
}

pub type Location = (u32, u32);

#[derive(Clone, Debug, PartialEq)]
pub struct Token(Location, TkType, String);

impl Token {
    pub fn location(&self) -> Location {
        self.0
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
        let (line, pos) = self.location();
        write!(
            f,
            "loc: ({}, {}), type: {:?}, `{}`",
            line,
            pos,
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
    fn new<T: Into<String>>(code: T) -> Lexer {
        Lexer {
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
        Token((self.line, self.pos), token_type, value)
    }
    fn emit(&mut self, token_type: TkType) {
        let s: String = self.code[self.start..self.offset].into_iter().collect();
        let tok = match s.as_str() {
            "return" => self.new_token(TkType::Return, s),
            "true" => self.new_token(TkType::True, s),
            "false" => self.new_token(TkType::False, s),
            _ => self.new_token(token_type, s),
        };
        self.tokens.push(tok);
        self.ignore();
    }
}

fn whitespace(lexer: &mut Lexer) -> State {
    while let Some(c) = lexer.peek() {
        if c == ' ' || c == '\n' {
            if c == '\n' {
                lexer.pos = 0;
                lexer.line += 1;
            }
            lexer.next();
        } else {
            break;
        }
    }
    lexer.ignore();

    match lexer.peek() {
        Some(_c @ '0'..='9') => State::Fn(number),
        Some('=') => {
            lexer.next();
            lexer.emit(TkType::Assign);
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
        Some('(') => {
            lexer.next();
            lexer.emit(TkType::LParen);
            State::Fn(whitespace)
        }
        Some(')') => {
            lexer.next();
            lexer.emit(TkType::RParen);
            State::Fn(whitespace)
        }
        Some('[') => {
            lexer.next();
            lexer.emit(TkType::LBracket);
            State::Fn(whitespace)
        }
        Some(']') => {
            lexer.next();
            lexer.emit(TkType::RBracket);
            State::Fn(whitespace)
        }
        Some('{') => {
            lexer.next();
            lexer.emit(TkType::LBrace);
            State::Fn(whitespace)
        }
        Some('}') => {
            lexer.next();
            lexer.emit(TkType::RBrace);
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
    lexer.next();
    while let Some(c) = lexer.next() {
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

pub fn lex<T: Into<String>>(source: T) -> Vec<Token> {
    let mut lexer = Lexer::new(source);
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
            lex(code),
            vec![
                Token((1, 0), Ident, "測試".to_string()),
                Token((1, 2), Colon, ":".to_string()),
                Token((1, 4), Ident, "int".to_string()),
                Token((1, 8), Assign, "=".to_string()),
                Token((1, 10), Integer, "1".to_string()),
                Token((1, 11), EOF, "".to_string()),
            ]
        );
    }

    #[test]
    fn get_number_tokens() {
        let ts = lex("10 30");
        assert_eq!(
            ts,
            vec![
                Token((1, 0), Integer, "10".to_string()),
                Token((1, 3), Integer, "30".to_string()),
                Token((1, 5), EOF, "".to_string()),
            ]
        );
    }

    #[test]
    fn get_ident_tokens() {
        let ts = lex(" abc6");
        assert_eq!(
            ts,
            vec![
                Token((1, 1), Ident, "abc6".to_string()),
                Token((1, 5), EOF, "".to_string()),
            ]
        )
    }
}
