#[derive(Clone, Debug, PartialEq)]
pub enum Type {
    Unit,
    I64,
    F64,
    String,
    Bool,
    TypeVar(TypeVar),
    Lambda(Vec<Type>, Box<Type>),
}

impl std::fmt::Display for Type {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        use Type::*;
        let string = match self {
            Unit => "()",
            I64 => "i64",
            F64 => "f64",
            String => "string",
            Bool => "bool",
            TypeVar(_) => "'type_var",
            Lambda(_, _) => "fn",
        };
        write!(f, "{}", string)
    }
}

#[derive(Clone, Debug, Eq, PartialEq, Hash)]
pub struct TypeVar(pub u64);
