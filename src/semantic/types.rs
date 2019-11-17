use super::error::Result;
use super::error::SemanticError;
use crate::ast::*;
use crate::ast::{Function, ParsedType};
use crate::lexer::Location;
use std::collections::HashMap;

pub struct TypeEnv {
    parent: Option<*const TypeEnv>,
    type_env: HashMap<String, TypeInfo>,
}

impl TypeEnv {
    pub(crate) fn type_of_expr(&self, expr: Expr) -> Result<Type> {
        use ExprVariant::*;
        let location = expr.location;
        match expr.value {
            Binary(l, r, op) => {
                let left_type = self.type_of_expr(*l)?;
                let right_type = self.type_of_expr(*r)?;
                match (left_type, right_type, op) {
                    (Type::Int, Type::Int, Operator::Plus) => Ok(Type::Int),
                    _ => panic!("unsupported operator"),
                }
            }
            F64(_) => Ok(Type::F64),
            Int(_) => Ok(Type::Int),
            Bool(_) => Ok(Type::Bool),
            String(_) => Ok(Type::String),
            List(es) => {
                if es.len() < 1 {
                    // FIXME: when there has no element in list, e.g. `[]`. We should provide FreeType for it.
                    // so type should be: `List[T]`, and `T` is unknown that can be any type unify with it later.
                    unreachable!()
                }
                let expr_type: Type = self.type_of_expr(es[0].clone())?;
                for e in es {
                    if expr_type != self.type_of_expr(e.clone())? {
                        return Err(SemanticError::type_mismatched(
                            e.location,
                            expr_type,
                            self.type_of_expr(e)?,
                        ));
                    }
                }
                Ok(Type::generic_type("List", vec![expr_type]))
            }
            FuncCall(f, args) => {
                let location = f.location;
                let f_type = self.type_of_expr(*f)?;
                match f_type {
                    Type::FunctionType(params, ret_typ) => {
                        for (p, arg) in params.iter().zip(args.iter()) {
                            self.unify(
                                arg.location,
                                p.clone(),
                                self.type_of_expr(arg.expr.clone())?,
                            )?;
                        }
                        Ok(*ret_typ)
                    }
                    _ => Err(SemanticError::call_on_non_function_type(location, f_type)),
                }
            }
            Identifier(id) => {
                let type_info = self.get_variable(location, id)?;
                Ok(type_info.typ)
            }
        }
    }

    pub(crate) fn unify(&self, location: Location, expected: Type, actual: Type) -> Result<()> {
        if expected == actual {
            Ok(())
        } else {
            Err(SemanticError::type_mismatched(location, expected, actual))
        }
    }
}

impl TypeEnv {
    pub fn new() -> TypeEnv {
        TypeEnv {
            parent: None,
            type_env: HashMap::new(),
        }
    }
    pub fn with_parent(parent: &TypeEnv) -> TypeEnv {
        TypeEnv {
            parent: Some(parent),
            type_env: HashMap::new(),
        }
    }

    pub(crate) fn add_variable(
        &mut self,
        location: Location,
        key: String,
        typ: Type,
    ) -> Result<()> {
        if self.type_env.contains_key(&key) {
            Err(SemanticError::name_redefined(location, key))
        } else {
            self.type_env.insert(key, TypeInfo::new(location, typ));
            Ok(())
        }
    }
    pub(crate) fn get_variable(&self, location: Location, k: String) -> Result<TypeInfo> {
        let result = self.type_env.get(&k);
        match result {
            Some(t) => Ok(t.clone()),
            None => match self.parent {
                Some(env) => unsafe { env.as_ref() }.unwrap().get_variable(location, k),
                None => Err(SemanticError::no_variable(location, k)),
            },
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub struct TypeInfo {
    pub location: Location,
    pub typ: Type,
}

impl TypeInfo {
    fn new(location: Location, typ: Type) -> TypeInfo {
        TypeInfo { location, typ }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum Type {
    Void,
    Int,
    Bool,
    F64,
    String,
    // name[generic_types]
    GenericType(String, Vec<Type>),
    FunctionType(Vec<Type>, Box<Type>),
    UnknownType(String),
}

impl Type {
    fn generic_type<T: ToString>(name: T, generics: Vec<Type>) -> Type {
        Type::GenericType(name.to_string(), generics)
    }
}

impl Type {
    pub fn from(typ: ParsedType) -> Type {
        use Type::*;
        match typ.name().as_str() {
            "int" => Int,
            "void" => Void,
            "f64" => F64,
            "bool" => Bool,
            "string" => String,
            "List" => Type::generic_type(
                "List",
                typ.generics()
                    .iter()
                    .map(|parsed_type| Type::from(parsed_type.clone()))
                    .collect(),
            ),
            _ => UnknownType(typ.name()),
        }
    }

    pub fn new_function(f: Function) -> Type {
        let param_types = f
            .parameters
            .into_iter()
            .map(|param| Type::from(param.0))
            .collect();
        Type::FunctionType(param_types, Type::from(f.ret_typ).into())
    }
}

impl std::fmt::Display for Type {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        use Type::*;
        match self {
            Void => write!(f, "void"),
            Int => write!(f, "int"),
            F64 => write!(f, "f64"),
            Bool => write!(f, "bool"),
            String => write!(f, "string"),
            GenericType(name, generics) => {
                write!(f, "{}[", name)?;
                for g in generics {
                    write!(f, "{} ,", g)?;
                }
                write!(f, "]")
            }
            // FIXME: print format: `(int, int): int` not `<function>`
            FunctionType(_params, _ret) => write!(f, "<function>"),
            UnknownType(s) => write!(f, "{}", s.as_str()),
        }
    }
}
