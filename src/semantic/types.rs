use super::error::Result;
use super::error::SemanticError;
use crate::ast::ParsedType;
use crate::ast::TopAst;
use crate::lexer::Location;
use std::collections::HashMap;

pub struct TypeEnv {
    type_env: HashMap<String, TypeInfo>,
}

impl TypeEnv {
    pub fn new() -> TypeEnv {
        TypeEnv {
            type_env: HashMap::new(),
        }
    }

    pub(crate) fn add_variable(&mut self, top: TopAst) -> Result<()> {
        if self.type_env.contains_key(&top.name()) {
            Err(SemanticError::name_redefined(top.location(), top.name()))
        } else {
            use TopAst::*;
            let location = top.location();
            let key = top.name();
            match top {
                Variable(v) => {
                    self.type_env
                        .insert(key, TypeInfo::new(location, Type::from(v.typ)));
                }
                Function(f) => {}
            }
            Ok(())
        }
    }
    pub(crate) fn get_variable(&self, location: Location, k: String) -> Result<TypeInfo> {
        let result = self.type_env.get(&k);
        match result {
            Some(t) => Ok(t.clone()),
            None => Err(SemanticError::no_variable(location, k)),
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
}

impl std::fmt::Display for Type {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        use Type::*;
        let r = match self {
            Void => "void",
            Int => "int",
            F64 => "f64",
            String => "string",
            UnknownType(s) => s.as_str(),
        };
        write!(f, "{}", r)
    }
}
