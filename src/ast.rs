use super::lexer::{TkType, Token};
use crate::lexer::Location;

#[derive(Clone, Debug, PartialEq)]
pub enum TopAst {
    Function(Function),
    Variable(Variable),
}

impl TopAst {
    pub fn name(&self) -> String {
        use TopAst::*;
        match self {
            Function(f) => f.name.clone(),
            Variable(v) => v.name.clone(),
        }
    }
    pub fn location(&self) -> Location {
        use TopAst::*;
        match self {
            Function(f) => f.loc,
            Variable(v) => v.loc,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum ParsedType {
    /// Defined: `int`, `f64`, `string`
    Defined(String),
}

impl ParsedType {
    pub fn new<T: ToString>(name: T) -> ParsedType {
        ParsedType::Defined(name.to_string())
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct Variable {
    loc: Location,
    name: String,
    pub typ: ParsedType,
    pub expr: Expr,
}

impl Variable {
    pub fn new<T: ToString>(loc: Location, name: T, typ: ParsedType, expr: Expr) -> Variable {
        Variable {
            loc,
            name: name.to_string(),
            typ,
            expr,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct Function {
    loc: Location,
    name: String,
    pub parameters: Vec<Parameter>,
    pub ret_typ: ParsedType,
    pub body: Body,
}

impl Function {
    pub fn new<T: ToString>(
        loc: Location,
        name: T,
        parameters: Vec<Parameter>,
        ret_typ: ParsedType,
        body: Body,
    ) -> Function {
        Function {
            loc,
            name: name.to_string(),
            parameters,
            ret_typ,
            body,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum Body {
    Block(Block),
    Expr(Expr),
}

#[derive(Clone, Debug, PartialEq)]
pub struct Block {
    pub statements: Vec<Statement>,
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
pub struct Parameter(pub ParsedType, pub String);

impl Parameter {
    pub fn new<T: ToString>(name: T, typ: ParsedType) -> Parameter {
        Parameter(typ, name.to_string())
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum Statement {
    /// Return:
    ///
    /// `return 1;`
    Return(Expr),
}

#[derive(Clone, Debug, PartialEq)]
pub enum Expr {
    /// `x + y`
    Binary(Box<Expr>, Box<Expr>, Operator),
    /// `1.345`
    F64(f64),
    /// `1`
    Int(i64),
    /// `"str"`
    String(String),
    /// `a(b)`
    FuncCall(Box<Expr>, Vec<Argument>),
    /// `n`
    Identifier(String),
}

impl Expr {
    pub fn binary(l: Expr, r: Expr, op: Operator) -> Expr {
        Expr::Binary(l.into(), r.into(), op)
    }
    pub fn identifier<T: ToString>(id: T) -> Expr {
        Expr::Identifier(id.to_string())
    }
}

/// Argument:
///
/// `assert(n, equal_to: 1)`
#[derive(Clone, Debug, PartialEq)]
pub struct Argument {
    pub name: Option<String>,
    pub expr: Expr,
}

impl Argument {
    pub fn new(name: Option<String>, expr: Expr) -> Argument {
        Argument { name, expr }
    }
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
