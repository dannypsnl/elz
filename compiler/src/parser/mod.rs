use super::ast::*;
use super::lexer;
use super::lexer::{TkType, Token};

mod error;
#[cfg(test)]
mod tests;

use error::ParseError;
use error::Result;

/// Parser is a parsing helper
pub struct Parser {
    tokens: Vec<Token>,
    offset: usize,
}

impl Parser {
    pub fn parse_program(&mut self) -> Result<Vec<Top>> {
        let mut program = vec![];
        match self.peek(0)?.tk_type() {
            TkType::Ident => {
                let f = self.parse_function()?;
                program.push(f);
            }
            _ => {
                panic!("unsupport yet");
            }
        }
        Ok(program)
    }
    /// parse_function:
    /// ```
    /// add(x: int, y: int): int {
    ///   return x + y;
    /// }
    /// ```
    pub fn parse_function(&mut self) -> Result<Top> {
        self.predict(vec![TkType::Ident])?;
        let func_name = self.take()?;
        let params = self.parse_parameters()?;
        self.predict_and_consume(vec![TkType::Colon])?;
        let return_type = self.parse_type()?;
        let block = self.parse_block()?;
        Ok(Top::Func(return_type, func_name.value(), params, block))
    }
    /// parse_block:
    /// ```
    /// {
    ///   <statement>*
    /// }
    /// ```
    pub fn parse_block(&mut self) -> Result<Block> {
        self.predict_and_consume(vec![TkType::LBrace])?;
        let mut block = Block::new();
        while self.peek(0)?.tk_type() != &TkType::RBrace {
            let stmt = self.parse_statement()?;
            block.append(stmt);
        }
        Ok(block)
    }
    /// parse_statement:
    /// ```
    /// return 1;
    /// ```
    pub fn parse_statement(&mut self) -> Result<Statement> {
        let stmt = match self.peek(0)?.tk_type() {
            TkType::Return => {
                self.consume()?;
                Ok(Statement::Return(self.parse_expression(None, 1)?))
            }
            _ => Err(ParseError::new(format!("unimplement"))),
        };
        self.predict_and_consume(vec![TkType::Semicolon])?;
        stmt
    }
    /// parse_expression:
    /// ```
    /// 1 + 2
    /// ```
    pub fn parse_expression(
        &mut self,
        left_hand_side: Option<Expr>,
        previous_primary: u64,
    ) -> Result<Expr> {
        let mut lhs = left_hand_side.unwrap_or(self.parse_primary()?);
        let mut lookahead = self.peek(0)?;
        while precedence(lookahead.clone()) >= previous_primary {
            let operator = lookahead.clone();
            self.consume()?;
            let mut rhs = self.parse_primary()?;
            lookahead = self.peek(0)?;
            while precedence(lookahead.clone()) > precedence(operator.clone())
                || (is_right_associative(lookahead.clone())
                    && (precedence(lookahead.clone()) == precedence(operator.clone())))
            {
                rhs = self.parse_expression(Some(lhs.clone()), precedence(lookahead.clone()))?;
                lookahead = self.peek(0)?;
            }
            lhs = Expr::Binary(Box::new(lhs), Box::new(rhs), Operator::from_token(operator));
        }
        Ok(lhs)
    }
    /// parse_primary:
    /// ```
    /// 1, x
    /// ```
    pub fn parse_primary(&mut self) -> Result<Expr> {
        match self.peek(0)?.tk_type() {
            TkType::Num => {
                let num = self.take()?.value();
                if num.parse::<i64>().is_ok() {
                    Ok(Expr::Int(num.parse::<i64>().unwrap()))
                } else if num.parse::<f64>().is_ok() {
                    Ok(Expr::F64(num.parse::<f64>().unwrap()))
                } else {
                    panic!(
                        "lexing error, number token can't be convert to number: {:?}",
                        num
                    )
                }
            }
            TkType::Ident => Ok(Expr::Identifier(self.take()?.value())),
            _ => Err(ParseError::new(format!(
                "unimplement primary for {:?}",
                self.peek(0)?
            ))),
        }
    }
    /// parse_parameters:
    /// ```
    /// (x: int, y: int)
    /// ```
    pub fn parse_parameters(&mut self) -> Result<Vec<Parameter>> {
        self.predict_and_consume(vec![TkType::LParen])?;
        let mut params = vec![];
        loop {
            self.predict(vec![TkType::Ident, TkType::Colon])?;
            let param_name = self.take()?.value();
            self.consume()?;
            let typ = self.parse_type()?;
            params.push(Parameter(typ, param_name));
            if self.predict(vec![TkType::Comma]).is_ok() {
                self.consume()?;
            } else if self.predict(vec![TkType::RParen]).is_ok() {
                self.consume()?;
                return Ok(params);
            } else {
                return Err(ParseError::new(format!(
                    "expected `,` or `)` but got unexpected: {:?} while parsing parameters",
                    self.peek(0)?,
                )));
            }
        }
    }
    /// parse_type:
    /// ```
    /// int
    /// ```
    pub fn parse_type(&mut self) -> Result<Type> {
        self.predict(vec![TkType::Ident])?;
        let typ = self.take()?.value();
        Ok(Type::Normal(typ))
    }
}

fn is_right_associative(_op: Token) -> bool {
    false
}
fn precedence(op: Token) -> u64 {
    match op.tk_type() {
        TkType::Plus => 2,
        _ => 0,
    }
}

/// This block puts helpers
impl Parser {
    /// new create Parser from code
    pub fn new(code: String) -> Parser {
        let tokens = lexer::lex(code);
        Parser {
            tokens: tokens,
            offset: 0,
        }
    }
    /// peek get the token by current add n
    pub fn peek(&self, n: usize) -> Result<Token> {
        self.get_token(self.offset + n)
    }
    /// consume take the token but don't use it
    pub fn consume(&mut self) -> Result<()> {
        self.take()?;
        Ok(())
    }
    /// take add current token position
    pub fn take(&mut self) -> Result<Token> {
        self.offset += 1;
        self.get_token(self.offset - 1)
    }
    fn get_token(&self, n: usize) -> Result<Token> {
        if self.tokens.len() <= n {
            Err(ParseError::new("eof".to_string()))
        } else {
            Ok(self.tokens[n].clone())
        }
    }
    fn matched(&self, token_type: &TkType, expected_type: &TkType) -> bool {
        *token_type == *expected_type
    }
    pub fn predict_and_consume(&mut self, wants: Vec<TkType>) -> Result<()> {
        let len = wants.len();
        self.predict(wants)?;
        for _ in 1..=len {
            self.consume()?;
        }
        Ok(())
    }
    pub fn predict(&self, wants: Vec<TkType>) -> Result<()> {
        for (i, v) in wants.iter().enumerate() {
            let tk = self.peek(i)?;
            if !self.matched(tk.tk_type(), v) {
                return Err(ParseError::new(format!(
                    "expected: {:?} but got {:?} at {:?}",
                    v,
                    tk.tk_type(),
                    tk.location(),
                )));
            }
        }
        Ok(())
    }
}
