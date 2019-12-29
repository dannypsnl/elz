use super::lexer::{TkType, Token};
use crate::lexer::Location;
use std::collections::HashMap;

#[derive(Clone, Debug, PartialEq)]
pub enum TopAst {
    Function(Function),
    Variable(Variable),
    Class(Class),
}

#[derive(Clone, Debug, PartialEq)]
pub struct TypeParameter {
    pub name: String,
    pub parent_types: Vec<ParsedType>,
}

impl TypeParameter {
    pub fn new<T: ToString>(name: T, parent_types: Vec<ParsedType>) -> TypeParameter {
        TypeParameter {
            name: name.to_string(),
            parent_types,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct Class {
    pub location: Location,
    pub parent_class_name: Option<String>,
    pub name: String,
    pub type_parameters: Vec<TypeParameter>,
    pub fields: Vec<Field>,
    pub methods: Vec<Function>,
    pub static_methods: Vec<Function>,
}

impl Class {
    pub fn new<T: ToString>(
        location: Location,
        parent_class_name: Option<String>,
        name: T,
        type_parameters: Vec<TypeParameter>,
        fields: Vec<Field>,
        methods: Vec<Function>,
        static_methods: Vec<Function>,
    ) -> Class {
        Class {
            location,
            parent_class_name,
            name: name.to_string(),
            type_parameters,
            fields,
            methods,
            static_methods,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct Field {
    pub location: Location,
    pub name: String,
    pub typ: ParsedType,
    pub expr: Option<Expr>,
}

impl Field {
    pub fn new<T: ToString>(
        location: Location,
        name: T,
        typ: ParsedType,
        expr: Option<Expr>,
    ) -> Field {
        Field {
            location,
            name: name.to_string(),
            typ,
            expr,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum ParsedType {
    TypeName(String),
    GenericType(String, Vec<ParsedType>),
}

impl ParsedType {
    pub fn type_name<T: ToString>(name: T) -> ParsedType {
        ParsedType::TypeName(name.to_string())
    }
    pub fn generic_type<T: ToString>(name: T, generics: Vec<ParsedType>) -> ParsedType {
        ParsedType::GenericType(name.to_string(), generics)
    }

    pub fn name(&self) -> String {
        match self {
            ParsedType::TypeName(name) => name.clone(),
            ParsedType::GenericType(name, _) => name.clone(),
        }
    }
    pub fn generics(&self) -> Vec<ParsedType> {
        match self {
            ParsedType::TypeName(_) => vec![],
            ParsedType::GenericType(_, lst) => lst.clone(),
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct Variable {
    pub location: Location,
    pub name: String,
    pub typ: ParsedType,
    pub expr: Expr,
}

impl Variable {
    pub fn new<T: ToString>(location: Location, name: T, typ: ParsedType, expr: Expr) -> Variable {
        Variable {
            location,
            name: name.to_string(),
            typ,
            expr,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct Function {
    pub location: Location,
    pub name: String,
    pub parameters: Vec<Parameter>,
    pub ret_typ: ParsedType,
    pub body: Option<Body>,
}

impl Function {
    pub fn new<T: ToString>(
        location: Location,
        name: T,
        parameters: Vec<Parameter>,
        ret_typ: ParsedType,
        body: Body,
    ) -> Function {
        Function {
            location,
            name: name.to_string(),
            parameters,
            ret_typ,
            body: Some(body),
        }
    }
    pub fn new_declaration<T: ToString>(
        location: Location,
        name: T,
        parameters: Vec<Parameter>,
        ret_typ: ParsedType,
    ) -> Function {
        Function {
            location,
            name: name.to_string(),
            parameters,
            ret_typ,
            body: None,
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
    pub fn return_stmt(location: Location, e: Option<Expr>) -> Statement {
        Statement {
            location,
            value: StatementVariant::Return(e),
        }
    }
    pub fn variable(location: Location, variable: Variable) -> Statement {
        Statement {
            location: location.clone(),
            value: StatementVariant::Variable(variable),
        }
    }
    pub fn expression(location: Location, expr: Expr) -> Statement {
        Statement {
            location,
            value: StatementVariant::Expression(expr),
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum StatementVariant {
    /// Return:
    ///
    /// `return 1;`
    Return(Option<Expr>),
    /// `x: int = 1;`
    Variable(Variable),
    /// `println("hello");`
    /// `foo.bar();`
    Expression(Expr),
}

#[derive(Clone, Debug, PartialEq)]
pub struct Expr {
    pub location: Location,
    pub value: ExprVariant,
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
    pub fn bool(location: Location, b: bool) -> Expr {
        Expr {
            location,
            value: ExprVariant::Bool(b),
        }
    }
    pub fn string<T: ToString>(location: Location, s: T) -> Expr {
        Expr {
            location,
            value: ExprVariant::String(s.to_string()),
        }
    }
    pub fn list(location: Location, lst: Vec<Expr>) -> Expr {
        Expr {
            location,
            value: ExprVariant::List(lst),
        }
    }
    pub fn func_call(location: Location, expr: Expr, args: Vec<Argument>) -> Expr {
        Expr {
            location,
            value: ExprVariant::FuncCall(expr.into(), args),
        }
    }
    pub fn dot_access<T: ToString>(location: Location, from: Expr, access: T) -> Expr {
        Expr {
            location,
            value: ExprVariant::DotAccess(from.into(), access.to_string()),
        }
    }
    pub fn identifier<T: ToString>(location: Location, id: T) -> Expr {
        Expr {
            location,
            value: ExprVariant::Identifier(id.to_string()),
        }
    }
    pub fn class_construction<T: ToString>(
        location: Location,
        class_name: T,
        field_inits: HashMap<String, Expr>,
    ) -> Expr {
        Expr {
            location,
            value: ExprVariant::ClassConstruction(class_name.to_string(), field_inits),
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum ExprVariant {
    /// `x + y`
    Binary(Box<Expr>, Box<Expr>, Operator),
    /// `1.345`
    F64(f64),
    /// `1`
    Int(i64),
    /// `true` or `false`
    Bool(bool),
    /// `"str"`
    String(String),
    /// `[1, 2, 3]`
    List(Vec<Expr>),
    /// `a(b)`
    FuncCall(Box<Expr>, Vec<Argument>),
    /// `foo.bar`, `foo.bar()`, `foo().bar`
    DotAccess(Box<Expr>, String),
    /// `n`
    Identifier(String),
    /// assume there has a class definition:
    /// `class Foo { bar: int; }`
    /// We can have a class construction expression
    /// `Foo { bar: 0 }`
    ClassConstruction(String, HashMap<String, Expr>),
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
