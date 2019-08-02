use super::super::ast;
use super::Context;
use super::Result;

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

impl Type {
    pub fn from_ast_type(c: &mut Context, t: &ast::Type) -> Result<Type> {
        match t {
            ast::Type::Unit => Ok(Type::Unit),
            ast::Type::Defined(name) => c.get_type(&name),
            ast::Type::Unsure(name) => {
                let id = match c.type_var_id.get(name) {
                    Some(&id) => id,
                    None => {
                        c.type_var_id.insert(name.clone(), c.count);
                        c.count += 1;
                        c.count - 1
                    }
                };
                Ok(Type::TypeVar(TypeVar(id)))
            }
            ast::Type::None => {
                c.count += 1;
                let id = c.count - 1;
                Ok(Type::TypeVar(TypeVar(id)))
            }
        }
    }
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
