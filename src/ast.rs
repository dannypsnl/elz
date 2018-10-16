#[derive(Clone, PartialEq, Debug)]
pub enum Expr {
    Integer(i64),
    Number(f64),
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
pub struct Method(pub Option<Type>, pub String, pub Vec<Parameter>);
#[derive(Clone, PartialEq, Debug)]
pub struct Parameter(pub String, pub Option<Type>);
#[derive(Clone, PartialEq, Debug)]
pub enum Statement {
    LetDefine,
    LetMutDefine,
    Assign,
    AccessChain,
}
