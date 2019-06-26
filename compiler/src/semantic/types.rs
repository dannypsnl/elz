use super::super::ast;
use super::Context;

#[derive(Clone, Debug, PartialEq)]
pub enum Type {
    I64,
    F64,
    String,
    Bool,
    TypeVar(TypeVar),
    Lambda(Vec<Type>, Box<Type>),
}

impl Type {
    pub fn from_ast_type(c: &mut Context, t: ast::Type) -> Result<Type, String> {
        match t {
            ast::Type::Defined(name) => c.get(&name),
            ast::Type::Unsure(name) => {
                let id = match c.type_var_id.get(&name) {
                    Some(&id) => id,
                    None => {
                        c.type_var_id.insert(name, c.count);
                        c.count += 1;
                        c.count - 1
                    }
                };
                Ok(Type::TypeVar(TypeVar(id)))
            }
        }
    }
}

#[derive(Clone, Debug, Eq, PartialEq, Hash)]
pub struct TypeVar(pub u64);