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
                program.push(
                    if self.predict(vec![TkType::Ident, TkType::LParen]).is_ok() {
                        Top::FuncDefine(self.parse_function()?)
                    } else {
                        self.parse_global_variable()?
                    },
                );
            }
            TkType::Type => {
                let t = self.parse_type_define()?;
                program.push(t);
            }
            TkType::Import => {
                let i = self.parse_import()?;
                program.push(i);
            }
            TkType::Contract => {
                let c = self.parse_contract()?;
                program.push(c);
            }
            _ => {
                panic!("unsupport yet");
            }
        }
        Ok(program)
    }
    /// parse_import:
    /// ```ignore
    /// import foo::bar
    /// import pkg::hello
    /// ```
    pub fn parse_import(&mut self) -> Result<Top> {
        self.predict_and_consume(vec![TkType::Import])?;
        Ok(Top::Import(self.parse_access_chain()?))
    }
    /// parse_access_chain:
    /// ```ignore
    /// foo::bar
    /// ```
    pub fn parse_access_chain(&mut self) -> Result<Vec<String>> {
        let mut chain = vec![];
        self.predict(vec![TkType::Ident])?;
        chain.push(self.take()?.value());
        while self.peek(0)?.tk_type() == &TkType::Accessor {
            self.predict_and_consume(vec![TkType::Accessor])?;
            self.predict(vec![TkType::Ident])?;
            chain.push(self.take()?.value());
        }
        Ok(chain)
    }
    /// parse_type_define:
    ///
    /// * tagged union type
    ///     ```ignore
    ///     type Option 'a (
    ///       Just(a: 'a)
    ///       | Nothing
    ///     )
    ///     ```
    /// * structure type
    ///     ```ignore
    ///     type Car (
    ///       name: string,
    ///       price: int,
    ///     )
    ///     ```
    pub fn parse_type_define(&mut self) -> Result<Top> {
        self.predict_and_consume(vec![TkType::Type])?;
        self.predict(vec![TkType::Ident])?;
        let type_name = self.take()?.value();
        let mut unsure_types = vec![];
        while self.peek(0)?.tk_type() != &TkType::LParen {
            unsure_types.push(self.parse_unsure_type()?);
        }
        if self
            .predict(vec![TkType::LParen, TkType::Ident, TkType::Colon])
            .is_ok()
        {
            // structure type
            let params = self.parse_parameters()?;
            Ok(Top::StructureTypeDefine(type_name, unsure_types, params))
        } else if
        // Just(a: 'a) | Nothing
        self.predict(vec![TkType::LParen, TkType::Ident, TkType::LParen]).is_ok() ||
            // Red | Blue | Green
            self.predict(vec![TkType::LParen, TkType::Ident, TkType::VerticalLine]).is_ok()
        {
            // tagged union type
            self.predict_and_consume(vec![TkType::LParen])?;
            let mut subtypes = vec![];
            subtypes.push(self.parse_tagged_subtype()?);
            while self.peek(0)?.tk_type() == &TkType::VerticalLine {
                self.predict_and_consume(vec![TkType::VerticalLine])?;
                subtypes.push(self.parse_tagged_subtype()?);
            }
            self.predict_and_consume(vec![TkType::RParen])?;
            Ok(Top::TaggedUnionTypeDefine(
                type_name,
                unsure_types,
                subtypes,
            ))
        } else {
            Err(ParseError::new(format!(
                "{} {} can't be a part of type define",
                self.peek(0)?.value(),
                self.peek(1)?.value(),
            )))
        }
    }
    /// parse_tagged_subtype:
    /// ```ignore
    /// Just(a: 'a)
    /// Nothing
    /// ```
    pub fn parse_tagged_subtype(&mut self) -> Result<SubType> {
        self.predict(vec![TkType::Ident])?;
        let tag = self.take()?.value();
        if self.predict(vec![TkType::LParen]).is_err() {
            Ok(SubType {
                tag,
                params: vec![],
            })
        } else {
            let params = self.parse_parameters()?;
            Ok(SubType { tag, params })
        }
    }
    /// parse_global_variable:
    /// ```ignore
    /// x: int = 1
    /// ```
    pub fn parse_global_variable(&mut self) -> Result<Top> {
        self.predict(vec![TkType::Ident])?;
        let variable_name = self.take()?.value();
        let typ = if self.predict(vec![TkType::Colon]).is_ok() {
            self.predict_and_consume(vec![TkType::Colon])?;
            Some(self.parse_type()?)
        } else {
            None
        };
        self.predict_and_consume(vec![TkType::Assign])?;
        Ok(Top::GlobalVariable(
            typ,
            variable_name,
            self.parse_expression(None, 1)?,
        ))
    }
    /// parse_contract:
    /// ```ignore
    /// contract Show (
    ///   to_string(from: Self): string;
    /// )
    /// ```
    pub fn parse_contract(&mut self) -> Result<Top> {
        self.predict_and_consume(vec![TkType::Contract])?;
        self.predict(vec![TkType::Ident])?;
        let contract_name = self.take()?.value();
        self.predict_and_consume(vec![TkType::LParen])?;
        let mut func_declare_set = vec![];
        while self.peek(0)?.tk_type() != &TkType::RParen {
            let func_declare = self.parse_function_declare()?;
            self.predict_and_consume(vec![TkType::Semicolon])?;
            func_declare_set.push(func_declare);
        }
        self.predict_and_consume(vec![TkType::RParen])?;
        Ok(Top::Contract(contract_name, func_declare_set))
    }
    /// parse_impl_contract
    /// ```ignore
    /// type Car (
    ///   name: string,
    ///   price: int
    /// )
    ///
    /// impl Comparable for Car (
    ///   equal(c1: Car, c2: Car): bool {
    ///     return c1.name == c2.name;
    ///   }
    /// )
    /// ```
    pub fn parse_impl_contract(&mut self) -> Result<Top> {
        self.predict_and_consume(vec![TkType::Impl])?;
        let access_chain = self.parse_access_chain()?;
        self.predict_and_consume(vec![TkType::For])?;
        self.parse_type()?;
        self.predict_and_consume(vec![TkType::LParen])?;
        let mut func_defines = vec![];
        while self.peek(0)?.tk_type() != &TkType::RParen {
            let func_define = self.parse_function()?;
            func_defines.push(func_define);
        }
        self.predict_and_consume(vec![TkType::RParen])?;
        Ok(Top::ImplContract(access_chain, func_defines))
    }
    /// parse_function:
    /// ```ignore
    /// add(x: int, y: int): int {
    ///   return x + y;
    /// }
    /// ```
    pub fn parse_function(&mut self) -> Result<Func> {
        let mut func = self.parse_function_declare()?;
        if self.peek(0)?.tk_type() == &TkType::Semicolon {
            self.predict_and_consume(vec![TkType::Semicolon])?;
        } else {
            let block = self.parse_block()?;
            func.set_body(block);
        }
        Ok(func)
    }
    /// parse_function_declare:
    /// ```ignore
    /// add(x: int, y: int): int
    /// ```
    pub fn parse_function_declare(&mut self) -> Result<Func> {
        self.predict(vec![TkType::Ident])?;
        let func_name = self.take()?.value();
        let params = self.parse_parameters()?;
        self.predict_and_consume(vec![TkType::Colon])?;
        let return_type = self.parse_type()?;
        Ok(Func::new(return_type, func_name, params, None))
    }
    /// parse_block:
    /// ```ignore
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
        self.predict_and_consume(vec![TkType::RBrace])?;
        Ok(block)
    }
    /// parse_statement:
    /// ```ignore
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
    /// ```ignore
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
    /// ```ignore
    /// foo()
    /// ```
    pub fn parse_primary(&mut self) -> Result<Expr> {
        let unary = self.parse_unary()?;
        match self.peek(0)?.tk_type() {
            TkType::LParen => {
                self.parse_argument(unary)
            }
            _ => Ok(unary)
        }
    }
    /// parse_unary:
    /// ```ignore
    /// <number>
    /// | <string_literal>
    /// | <identifier>
    /// ```
    pub fn parse_unary(&mut self) -> Result<Expr> {
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
            TkType::String => Ok(Expr::String(self.take()?.value())),
            _ => Err(ParseError::new(format!(
                "unimplemented primary for {:?}",
                self.peek(0)?
            ))),
        }
    }
    pub fn parse_argument(&mut self, func: Expr) -> Result<Expr> {
        self.predict_and_consume(vec![TkType::LParen])?;

        let mut args = vec![];
        while self.peek(0)?.tk_type() != &TkType::RParen {
            let identifier =
                if self.predict(vec![TkType::Ident, TkType::Colon]).is_ok() {
                    let identifier = self.take()?.value();
                    self.predict_and_consume(vec![TkType::Colon])?;
                    identifier
                } else { "".to_string() };
            let expr = self.parse_expression(None, 1)?;
            args.push(Expr::Argument(identifier, Box::new(expr)));
            if self.predict(vec![TkType::Comma]).is_err() {
                break;
            } else {
                self.predict_and_consume(vec![TkType::Comma])?;
            }
        }
        self.predict_and_consume(vec![TkType::RParen])?;

        Ok(Expr::FuncCall(Box::new(func), args))
    }
    /// parse_parameters:
    /// ```ignore
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
    /// ```ignore
    /// int, 'a
    /// ```
    pub fn parse_type(&mut self) -> Result<Type> {
        match self.peek(0)?.tk_type() {
            TkType::Prime => self.parse_unsure_type(),
            TkType::Ident => Ok(Type::Defined(self.take()?.value())),
            _ => Err(ParseError::new(format!(
                "expected `'` for unsure type(e.g. `'element`) or <identifier> for defined type but got {:?} while parsing type",
                self.peek(0)?,
            )))
        }
    }
    /// parse_unsure_type:
    /// ```ignore
    /// 'a, 'b
    /// ```
    pub fn parse_unsure_type(&mut self) -> Result<Type> {
        self.predict_and_consume(vec![TkType::Prime])?;
        Ok(Type::Unsure(self.take()?.value()))
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
