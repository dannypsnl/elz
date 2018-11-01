#[derive(Clone, PartialEq, Debug)]
pub enum Expr {
    Integer(i64),
    Number(f64),
    Ident(String),
    // ident chain, sub expression
    AccessChain(Vec<Expr>, Option<Box<Expr>>),
}
#[derive(Clone, PartialEq, Debug)]
pub enum Top {
    // export, name, expression
    GlobalBind(bool, String, Expr),
    // import lib::sub::{block0, block1, block2}
    // chain, block
    Import(Vec<String>, Vec<String>),
    // name, template types, type fields
    TypeDefine(String, Vec<String>, Vec<TypeField>),
    // function proto
    FnDefine(Method),
}
#[derive(Clone, PartialEq, Debug)]
pub struct Type(pub String, pub Vec<Type>);
#[derive(Clone, PartialEq, Debug)]
pub struct TypeField(pub String, pub Type);

#[derive(Clone, PartialEq, Debug)]
pub struct Method(
    pub Option<Type>,
    pub String,
    pub Vec<Parameter>,
    pub Vec<Statement>,
);
#[derive(Clone, PartialEq, Debug)]
pub struct Parameter(pub String, pub Option<Type>);
#[derive(Clone, PartialEq, Debug)]
pub enum Statement {
    // mutable, name, type, expression
    LetDefine(bool, String, Option<Type>, Expr),
    Return(Expr),
    // left-value, expression
    Assign(Expr, Expr),
    AccessChain,
}
