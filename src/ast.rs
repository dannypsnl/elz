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
pub struct ParsedType {
    name: String,
}

impl ParsedType {
    pub fn new<T: ToString>(name: T) -> ParsedType {
        ParsedType {
            name: name.to_string(),
        }
    }
    pub fn name(&self) -> String {
        self.name.clone()
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
pub struct Statement {
    pub location: Location,
    pub value: StatementVariant,
}

impl Statement {
    pub fn return_stmt(location: Location, e: Expr) -> Statement {
        Statement {
            location,
            value: StatementVariant::Return(e),
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum StatementVariant {
    /// Return:
    ///
    /// `return 1;`
    Return(Expr),
}

#[derive(Clone, Debug, PartialEq)]
pub struct Expr {
    pub location: Location,
    pub value: ExprVariant,
}

#[derive(Clone, Debug, PartialEq)]
pub enum ExprVariant {
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
    pub fn binary(location: Location, l: Expr, r: Expr, op: Operator) -> Expr {
        Expr {
            location,
            value: ExprVariant::Binary(l.into(), r.into(), op),
        }
    }
    pub fn f64(location: Location, f: f64) -> Expr {
        Expr {
            location,
            value: ExprVariant::F64(f),
        }
    }
    pub fn int(location: Location, i: i64) -> Expr {
        Expr {
            location,
            value: ExprVariant::Int(i),
        }
    }
    pub fn string(location: Location, s: String) -> Expr {
        Expr {
            location,
            value: ExprVariant::String(s),
        }
    }
    pub fn func_call(location: Location, expr: Expr, args: Vec<Argument>) -> Expr {
        Expr {
            location,
            value: ExprVariant::FuncCall(expr.into(), args),
        }
    }
    pub fn identifier<T: ToString>(location: Location, id: T) -> Expr {
        Expr {
            location,
            value: ExprVariant::Identifier(id.to_string()),
        }
    }
}

/// Argument:
///
/// `assert(n, equal_to: 1)`
#[derive(Clone, Debug, PartialEq)]
pub struct Argument {
    pub location: Location,
    pub name: Option<String>,
    pub expr: Expr,
}

impl Argument {
    pub fn new(location: Location, name: Option<String>, expr: Expr) -> Argument {
        Argument {
            location,
            name,
            expr,
        }
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
