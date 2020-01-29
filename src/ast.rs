use super::lexer::{TkType, Token};
use crate::lexer::Location;
use std::collections::HashMap;

#[derive(Clone, Debug, PartialEq)]
pub struct Tag {
    pub name: String,
}

#[derive(Clone, Debug, PartialEq)]
pub struct TopAst {
    pub tag: Option<Tag>,
    pub ast: TopAstVariant,
}


impl TopAst {
    pub fn new( tag: Option<Tag>, ast: TopAstVariant) -> TopAst {
        TopAst {
            tag,
            ast,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum TopAstVariant {
    Function(Function),
    Variable(Variable),
    Class(Class),
    Trait(Trait),
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
pub struct Trait {
    pub location: Location,
    pub with_traits: Vec<String>,
    pub name: String,
    pub type_parameters: Vec<TypeParameter>,
    pub members: Vec<TraitMember>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum TraitMember {
    Field(Field),
    Method(Function),
}

impl Trait {
    pub fn new<T: ToString>(
        location: Location,
        with_traits: Vec<String>,
        name: T,
        type_parameters: Vec<TypeParameter>,
        members: Vec<TraitMember>,
    ) -> Trait {
        Trait {
            location,
            with_traits,
            name: name.to_string(),
            type_parameters,
            members,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct Class {
    pub location: Location,
    pub parents: Vec<String>,
    pub name: String,
    pub type_parameters: Vec<TypeParameter>,
    pub members: Vec<ClassMember>,
}

#[derive(Clone, Debug, PartialEq)]
pub enum ClassMember {
    Field(Field),
    Method(Function),
    StaticMethod(Function),
}

impl Class {
    pub fn new<T: ToString>(
        location: Location,
        parents: Vec<String>,
        name: T,
        type_parameters: Vec<TypeParameter>,
        members: Vec<ClassMember>,
    ) -> Class {
        Class {
            location,
            parents,
            name: name.to_string(),
            type_parameters,
            members,
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
pub struct Parameter {
    pub name: String,
    pub typ: ParsedType,
}

impl Parameter {
    pub fn new<T: ToString>(name: T, typ: ParsedType) -> Parameter {
        Parameter {
            name: name.to_string(),
            typ,
        }
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
    pub fn if_block(
        location: Location,
        clauses: Vec<(Expr, Block)>,
        else_block: Option<Block>,
    ) -> Statement {
        Statement {
            location,
            value: StatementVariant::IfBlock {
                clauses,
                else_block,
            },
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum StatementVariant {
    /// `return 1;`
    Return(Option<Expr>),
    /// `x: int = 1;`
    Variable(Variable),
    /// `println("hello");`
    /// `foo.bar();`
    Expression(Expr),
    /// `if <condition> {} else if <condition> else {}`
    IfBlock {
        clauses: Vec<(Expr, Block)>,
        else_block: Option<Block>,
    },
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
    /// We can have a class construction expression: `Foo { bar: 0 }` for definition `class Foo { bar: int; }`
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
            tok => unimplemented!("{:?} is not a operator", tok),
        }
    }
}
