use super::error::Result;
use super::error::SemanticError;
use crate::ast::{Function, ParsedType};
use crate::lexer::Location;
use std::collections::HashMap;

pub struct TypeEnv {
    parent: Option<*const TypeEnv>,
    type_env: HashMap<String, TypeInfo>,
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
    F64,
    String,
    FunctionType(Vec<Type>, Box<Type>),
    UnknownType(String),
}

impl Type {
    pub fn from(typ: ParsedType) -> Type {
        use Type::*;
        match typ.name().as_str() {
            "int" => Int,
            "void" => Void,
            "f64" => F64,
            "string" => String,
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
        let r = match self {
            Void => "void",
            Int => "int",
            F64 => "f64",
            String => "string",
            // FIXME: print format: `(int, int): int` not `<function>`
            FunctionType(_params, _ret) => "<function>",
            UnknownType(s) => s.as_str(),
        };
        write!(f, "{}", r)
    }
}
