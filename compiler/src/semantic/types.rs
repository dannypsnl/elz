#[derive(Clone, Debug, PartialEq)]
pub enum Type {
    I64,
    F64,
    String,
    Bool,
}

#[derive(Clone, Debug, Eq, PartialEq, Hash)]
pub struct TypeVar;