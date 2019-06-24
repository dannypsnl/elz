use super::lexer::{TkType, Token};

#[derive(Clone, Debug, PartialEq)]
pub enum Type {
    /// Defined: int, f64, string
    Defined(String),
    /// Unsure: 'a, 'element, 'key, 'value
    Unsure(String),
}

#[derive(Clone, Debug, PartialEq)]
pub struct Block {
    statements: Vec<Statement>,
}

impl Block {
    pub fn new() -> Block {
        Block::from(vec![])
    }
    pub fn from(statements: Vec<Statement>) -> Block {
        Block { statements }
    }
    pub fn append(&mut self, stmt: Statement) {
        self.statements.push(stmt);
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct Parameter(pub Type, pub String);

#[derive(Debug, PartialEq)]
pub enum Top {
    Binding(String, Type, Expr),
    /// StructureTypeDefine: Name, Unsure Types, Fields
    StructureTypeDefine(String, Vec<Type>, Vec<Parameter>),
    /// TaggedUnionTypeDefine: Name, Unsure Types, Subtypes
    TaggedUnionTypeDefine(String, Vec<Type>, Vec<SubType>),
}

#[derive(Clone, Debug, PartialEq)]
pub struct Lambda {
    pub  return_type: Type,
    pub  parameters: Vec<Parameter>,
    body: Option<Block>,
}

impl Lambda {
    pub fn new(return_type: Type, parameters: Vec<Parameter>, body: Option<Block>) -> Lambda {
        Lambda {
            return_type,
            parameters,
            body,
        }
    }
}

#[derive(Debug, PartialEq)]
pub struct SubType {
    pub tag: String,
    pub params: Vec<Parameter>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum Statement {
    /// Return:
    /// ```ignore
    /// return <Expr>
    /// ```
    Return(Expr),
}

#[derive(Clone, Debug, PartialEq)]
pub enum Expr {
    Lambda(Lambda),
    Binary(Box<Expr>, Box<Expr>, Operator),
    F64(f64),
    Int(i64),
    String(String),
    FuncCall(Box<Expr>, Vec<Argument>),
    Identifier(String),
}

#[derive(Clone, Debug, PartialEq)]
pub struct Argument {
    pub name: String,
    pub expr: Expr,
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
