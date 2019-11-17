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
    pub fn parse_all(&mut self, end_token_type: TkType) -> Result<Vec<TopAst>> {
        let mut program = vec![];
        while self.peek(0)?.tk_type() != &end_token_type {
            let tok = self.peek(0)?;
            match tok.tk_type() {
                TkType::Ident => {
                    // found `<identifier> :`
                    if self.predict(vec![TkType::Ident, TkType::Colon]).is_ok() {
                        let v = self.parse_variable()?;
                        program.push(TopAst::Variable(v));
                    } else {
                        // else we just seems it as a function to parse
                        let f = self.parse_function()?;
                        program.push(TopAst::Function(f));
                    }
                }
                _ => unimplemented!(),
            }
        }
        Ok(program)
    }
    /// parse_variable:
    ///
    /// handle `x: int = 1;`
    pub fn parse_variable(&mut self) -> Result<Variable> {
        let loc = self.peek(0)?.location();
        // x: int = 1;
        let var_name = self.parse_identifier()?;
        // : int = 1;
        self.predict_and_consume(vec![TkType::Colon])?;
        // int = 1;
        let typ = self.parse_type()?;
        // = 1;
        self.predict_and_consume(vec![TkType::Assign])?;
        let expr = self.parse_expression(None, None)?;
        self.predict_and_consume(vec![TkType::Semicolon])?;
        Ok(Variable::new(loc, var_name, typ, expr))
    }
    /// parse_function:
    ///
    /// handle `main(): void {}` or `add(x: int, y: int): int = x + y;`
    pub fn parse_function(&mut self) -> Result<Function> {
        let loc = self.peek(0)?.location();
        // main(): void {}
        let fn_name = self.parse_identifier()?;
        // (): void {}
        let tok = self.peek(0)?;
        if tok.tk_type() == &TkType::LParen {
            let params = self.parse_parameters()?;
            // : void {}
            self.predict_and_consume(vec![TkType::Colon])?;
            // void {}
            let ret_typ = self.parse_type()?;
            // {}
            let body = self.parse_body()?;
            // now parsing done
            Ok(Function::new(loc, fn_name, params, ret_typ, body))
        } else {
            Err(ParseError::not_expected_token(vec![TkType::LParen], tok))
        }
    }
    /// parse_parameters:
    ///
    /// ()
    /// (x: int, y: int)
    fn parse_parameters(&mut self) -> Result<Vec<Parameter>> {
        self.predict_and_consume(vec![TkType::LParen])?;
        let mut params = vec![];
        while self.peek(0)?.tk_type() != &TkType::RParen {
            self.predict(vec![TkType::Ident, TkType::Colon])?;
            let param_name = self.take()?.value();
            self.consume()?;
            let typ = self.parse_type()?;
            params.push(Parameter(typ, param_name));
            let tok = self.peek(0)?;
            match tok.tk_type() {
                TkType::Comma => self.consume()?,
                TkType::RParen => (),
                _ => {
                    return Err(ParseError::not_expected_token(
                        vec![TkType::Comma, TkType::RParen],
                        tok,
                    ));
                }
            }
        }
        self.predict_and_consume(vec![TkType::RParen])?;
        Ok(params)
    }
    fn parse_body(&mut self) -> Result<Body> {
        let tok = self.peek(0)?;
        match tok.tk_type() {
            TkType::LBrace => Ok(Body::Block(self.parse_block()?)),
            TkType::Assign => {
                self.predict_and_consume(vec![TkType::Assign])?;
                let e = self.parse_expression(None, None)?;
                self.predict_and_consume(vec![TkType::Semicolon])?;
                Ok(Body::Expr(e))
            }
            _ => Err(ParseError::not_expected_token(
                vec![TkType::LBrace, TkType::Assign],
                tok,
            )),
        }
    }
    /// parse_identifier:
    ///
    /// foo::bar
    pub fn parse_identifier(&mut self) -> Result<String> {
        let mut chain = vec![];
        self.predict(vec![TkType::Ident])?;
        chain.push(self.take()?.value());
        while self.peek(0)?.tk_type() == &TkType::Accessor {
            self.predict_and_consume(vec![TkType::Accessor])?;
            self.predict(vec![TkType::Ident])?;
            chain.push(self.take()?.value());
        }
        Ok(chain.join("::"))
    }

    /// parse_type:
    ///
    /// `<identifier>`
    /// | `<identifier> [ <generic_type_list> ]`
    pub fn parse_type(&mut self) -> Result<ParsedType> {
        // ensure is <identifier>
        self.predict(vec![TkType::Ident])?;
        let type_name = self.parse_identifier()?;
        if self.predict(vec![TkType::LBracket]).is_ok() {
            let mut list = vec![];
            self.predict_and_consume(vec![TkType::LBracket])?;
            while self.peek(0)?.tk_type() != &TkType::RBracket {
                let typ = self.parse_type()?;
                list.push(typ);
                if self.predict(vec![TkType::Comma]).is_err() {
                    break;
                } else {
                    self.predict_and_consume(vec![TkType::Comma])?;
                }
            }
            self.predict_and_consume(vec![TkType::RBracket])?;
            Ok(ParsedType::generic_type(type_name, list))
        } else {
            Ok(ParsedType::type_name(type_name))
        }
    }
}

// for block
impl Parser {
    /// parse_block:
    ///
    /// {
    ///   <statement>*
    /// }
    pub fn parse_block(&mut self) -> Result<Block> {
        self.predict_and_consume(vec![TkType::LBrace])?;
        let mut block = Block::new();
        while self.peek(0)?.tk_type() != &TkType::RBrace {
            let stmt = self.parse_statement()?;
            block.append(stmt);
        }
        self.predict_and_consume(vec![TkType::RBrace])?;
        Ok(block)
    }
    pub fn parse_statement(&mut self) -> Result<Statement> {
        let tok = self.peek(0)?;
        let stmt = match tok.tk_type() {
            // `x: int = 1;`
            TkType::Ident => {
                let name = self.parse_identifier()?;
                self.predict_and_consume(vec![TkType::Colon])?;
                let typ = self.parse_type()?;
                self.predict_and_consume(vec![TkType::Assign])?;
                let expr = self.parse_expression(None, None)?;
                Ok(Statement::variable(tok.location(), name, typ, expr))
            }
            // `return 1;`
            TkType::Return => {
                self.consume()?;
                let expr = if self.peek(0)?.tk_type() == &TkType::Semicolon {
                    None
                } else {
                    Some(self.parse_expression(None, None)?)
                };
                Ok(Statement::return_stmt(tok.location(), expr))
            }
            _ => unimplemented!(),
        };
        self.predict_and_consume(vec![TkType::Semicolon])?;
        stmt
    }
}

// for expression
impl Parser {
    /// parse_expression:
    ///
    /// 1 + 2
    pub fn parse_expression(
        &mut self,
        left_hand_side: Option<Expr>,
        previous_primary: Option<u64>,
    ) -> Result<Expr> {
        let mut lhs = left_hand_side.unwrap_or(self.parse_primary()?);
        let mut lookahead = self.peek(0)?;
        while precedence(lookahead.clone()) >= previous_primary.unwrap_or(1) {
            let operator = lookahead.clone();
            self.consume()?;
            let mut rhs = self.parse_primary()?;
            lookahead = self.peek(0)?;
            while precedence(lookahead.clone()) > precedence(operator.clone())
                || (is_right_associative(lookahead.clone())
                    && (precedence(lookahead.clone()) == precedence(operator.clone())))
            {
                rhs =
                    self.parse_expression(Some(lhs.clone()), Some(precedence(lookahead.clone())))?;
                lookahead = self.peek(0)?;
            }
            lhs = Expr::binary(lhs.location, lhs, rhs, Operator::from_token(operator));
        }
        Ok(lhs)
    }
    /// parse_primary:
    ///
    /// foo()
    pub fn parse_primary(&mut self) -> Result<Expr> {
        let unary = self.parse_unary()?;
        match self.peek(0)?.tk_type() {
            TkType::LParen => self.parse_function_call(unary),
            _ => Ok(unary),
        }
    }
    /// parse_unary:
    ///
    /// <integer>
    /// | <float64>
    /// | <string_literal>
    /// | <identifier>
    /// | <bool>
    /// | <list>
    pub fn parse_unary(&mut self) -> Result<Expr> {
        let tok = self.peek(0)?;
        Ok(match tok.tk_type() {
            // FIXME: lexer should emit int & float token directly
            TkType::Integer => {
                let num = self.take()?.value();
                if num.parse::<i64>().is_ok() {
                    Expr::int(tok.location(), num.parse::<i64>().unwrap())
                } else if num.parse::<f64>().is_ok() {
                    Expr::f64(tok.location(), num.parse::<f64>().unwrap())
                } else {
                    panic!(
                        "lexing bug causes a number token can't be convert to number: {:?}",
                        num
                    )
                }
            }
            TkType::Ident => Expr::identifier(tok.location(), self.parse_identifier()?),
            TkType::True => {
                self.take()?;
                Expr::bool(tok.location(), true)
            }
            TkType::False => {
                self.take()?;
                Expr::bool(tok.location(), false)
            }
            TkType::String => Expr::string(tok.location(), self.take()?.value()),
            TkType::LBracket => {
                let list = self.parse_list()?;
                Expr::list(tok.location(), list)
            }
            _ => panic!("unimplemented primary for {:?}", self.peek(0)?),
        })
    }
    pub fn parse_function_call(&mut self, func: Expr) -> Result<Expr> {
        self.predict_and_consume(vec![TkType::LParen])?;

        let mut args = vec![];
        while self.peek(0)?.tk_type() != &TkType::RParen {
            let identifier = if self.predict(vec![TkType::Ident, TkType::Colon]).is_ok() {
                let identifier = self.take()?.value();
                self.predict_and_consume(vec![TkType::Colon])?;
                identifier
            } else {
                "".to_string()
            };
            let expr = self.parse_expression(None, None)?;
            args.push(Argument::new(expr.location, Some(identifier), expr));
            if self.predict(vec![TkType::Comma]).is_err() {
                break;
            } else {
                self.predict_and_consume(vec![TkType::Comma])?;
            }
        }
        self.predict_and_consume(vec![TkType::RParen])?;

        Ok(Expr::func_call(func.location, func, args))
    }
    pub fn parse_list(&mut self) -> Result<Vec<Expr>> {
        let mut list = vec![];
        self.predict_and_consume(vec![TkType::LBracket])?;
        while self.peek(0)?.tk_type() != &TkType::RBracket {
            let expr = self.parse_expression(None, None)?;
            list.push(expr);
            if self.predict(vec![TkType::Comma]).is_err() {
                break;
            } else {
                self.predict_and_consume(vec![TkType::Comma])?;
            }
        }
        self.predict_and_consume(vec![TkType::RBracket])?;
        Ok(list)
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

/// This block puts fundamental helpers
impl Parser {
    pub fn parse_program<T: Into<String>>(code: T) -> Result<Vec<TopAst>> {
        let mut parser = Parser::new(code);
        parser.parse_all(TkType::EOF)
    }
    /// new create Parser from code
    pub fn new<T: Into<String>>(code: T) -> Parser {
        let tokens = lexer::lex(code);
        Parser { tokens, offset: 0 }
    }
    /// peek get the token by (current position + n)
    pub fn peek(&self, n: usize) -> Result<Token> {
        self.get_token(self.offset + n)
    }
    /// consume take the token but don't use it
    pub fn consume(&mut self) -> Result<()> {
        self.take()?;
        Ok(())
    }
    /// take increment current token position
    pub fn take(&mut self) -> Result<Token> {
        self.offset += 1;
        self.get_token(self.offset - 1)
    }
    fn get_token(&self, n: usize) -> Result<Token> {
        if self.tokens.len() <= n {
            Err(ParseError::EOF)
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
                return Err(ParseError::not_expected_token(wants, tk));
            }
        }
        Ok(())
    }
}
