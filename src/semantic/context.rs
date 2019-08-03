use super::super::ast;
use super::error::{CheckError, Result};
use super::types;
use std::collections::HashMap;
use types::{Type, TypeVar};

pub struct Context {
    parent: Option<*const Context>,
    type_map: HashMap<String, Type>,
    type_environment: HashMap<String, Type>,
    type_var_id: HashMap<String, u64>,
    count: u64,
}

impl Context {
    /// new returns a new context for inference expression
    pub fn new() -> Context {
        let mut ctx = Context {
            parent: None,
            type_map: HashMap::new(),
            type_environment: HashMap::new(),
            type_var_id: HashMap::new(),
            count: 0,
        };
        ctx.add_type("int".to_string(), types::Type::I64);

        ctx
    }
    pub fn with_parent(ctx: &Context) -> Context {
        Context {
            parent: Some(ctx),
            type_map: HashMap::new(),
            type_environment: HashMap::new(),
            type_var_id: HashMap::new(),
            count: 0,
        }
    }

    fn get_type(&self, type_name: &String) -> Result<Type> {
        match self.type_map.get(type_name) {
            Some(t) => Ok(t.clone()),
            None => {
                if let Some(p_ctx) = self.parent {
                    unsafe { p_ctx.as_ref() }.unwrap().get_type(type_name)
                } else {
                    Err(CheckError::not_found(type_name.clone()))
                }
            }
        }
    }
    fn add_type(&mut self, type_name: String, typ: Type) {
        self.type_map.insert(type_name, typ);
    }

    pub(crate) fn get_identifier(&self, key: &String) -> Result<Type> {
        match self.type_environment.get(key) {
            Some(t) => Ok(t.clone()),
            None => {
                if let Some(p_ctx) = self.parent {
                    unsafe { p_ctx.as_ref() }.unwrap().get_identifier(key)
                } else {
                    Err(CheckError::not_found(key.clone()))
                }
            }
        }
    }
    pub(crate) fn add_identifier(&mut self, key: String, typ: Type) {
        self.type_environment.insert(key, typ);
    }

    pub(crate) fn from_ast_type(&mut self, t: &ast::Type) -> Result<Type> {
        match t {
            ast::Type::Unit => Ok(Type::Unit),
            ast::Type::Defined(name) => self.get_type(&name),
            ast::Type::Unsure(name) => {
                let id = match self.type_var_id.get(name) {
                    Some(&id) => id,
                    None => {
                        self.type_var_id.insert(name.clone(), self.count);
                        self.count += 1;
                        self.count - 1
                    }
                };
                Ok(Type::TypeVar(TypeVar(id)))
            }
            ast::Type::None => {
                self.count += 1;
                let id = self.count - 1;
                Ok(Type::TypeVar(TypeVar(id)))
            }
        }
    }
}
