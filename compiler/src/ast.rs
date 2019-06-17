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
    /// Import
    Import(Vec<String>),
    /// Contract: Name, FuncDefines
    Contract(String, Vec<Func>),
    /// ContractFuncDefine: contract name, function define list
    ImplContract(Vec<String>, Vec<Func>),
    /// FuncDefine
    FuncDefine(Func),
    /// GlobalVariable: Type, Name, Expr
    GlobalVariable(Option<Type>, String, Expr),
    /// StructureTypeDefine: Name, Unsure Types, Fields
    StructureTypeDefine(String, Vec<Type>, Vec<Parameter>),
    /// TaggedUnionTypeDefine: Name, Unsure Types, Subtypes
    TaggedUnionTypeDefine(String, Vec<Type>, Vec<SubType>),
}

#[derive(Debug, PartialEq)]
pub struct Func {
    pub  return_type: Type,
    pub name: String,
    pub  parameters: Vec<Parameter>,
    body: Option<Block>,
}

impl Func {
    pub fn new(return_type: Type, name: String, parameters: Vec<Parameter>, body: Option<Block>) -> Func {
        Func {
            return_type,
            name,
            parameters,
            body,
        }
    }

    pub fn set_body(&mut self, block: Block) {
        self.body = Some(block);
    }
}

#[derive(Debug, PartialEq)]
pub struct SubType {
    pub tag: String,
    pub params: Vec<Parameter>,
}

#[derive(Debug, PartialEq)]
pub enum Statement {
    /// Return:
    /// ```ignore
    /// return <Expr>
    /// ```
    Return(Expr),
}

#[derive(Clone, Debug, PartialEq)]
pub enum Expr {
    Binary(Box<Expr>, Box<Expr>, Operator),
    F64(f64),
    Int(i64),
    String(String),
    Argument(String, Box<Expr>),
    FuncCall(Box<Expr>, Vec<Expr>),
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
