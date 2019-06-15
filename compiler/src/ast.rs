use super::lexer::{TkType, Token};

#[derive(Debug, PartialEq)]
pub enum Type {
    /// Defined: int, f64, string
    Defined(String),
    /// Unsure: 'a, 'element, 'key, 'value
    Unsure(String),
}
#[derive(Debug, PartialEq)]
pub struct Block {
    statements: Vec<Statement>,
}
impl Block {
    pub fn new() -> Block {
        Block::from(vec![])
    }
    pub fn from(stmts: Vec<Statement>) -> Block {
        Block { statements: stmts }
    }
    pub fn append(&mut self, stmt: Statement) {
        self.statements.push(stmt);
    }
}
#[derive(Debug, PartialEq)]
pub struct Parameter(pub Type, pub String);
#[derive(Debug, PartialEq)]
pub enum Top {
    /// FuncDefine: ReturnType, Name, Parameters, Block
    FuncDefine(Type, String, Vec<Parameter>, Block),
    /// TypeDefine: Name, Unsure Types, Subtypes, Fields
    TypeDefine(String, Vec<Type>, Vec<SubType>, Vec<Parameter>),
}
#[derive(Debug, PartialEq)]
pub struct SubType {
    pub tag: String,
    pub params: Vec<Parameter>,
}
#[derive(Debug, PartialEq)]
pub enum Statement {
    /// Return:
    /// ```
    /// return <Expr>
    /// ```
    Return(Expr),
}
#[derive(Clone, Debug, PartialEq)]
pub enum Expr {
    Binary(Box<Expr>, Box<Expr>, Operator),
    F64(f64),
    Int(i64),
    Identifier(String),
}
#[derive(Clone, Debug, PartialEq)]
pub enum Operator {
    Plus,
}
impl Operator {
    pub fn from_token(token: Token) -> Operator {
        match token.tk_type() {
            TkType::Plus => Operator::Plus,
            tok => panic!("{:?} is not a operator", tok),
        }
    }
}
