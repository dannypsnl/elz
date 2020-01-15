use super::error::Result;
use super::error::SemanticError;
use crate::ast::*;
use crate::ast::{Function, ParsedType};
use crate::lexer::Location;
use std::collections::HashMap;

pub struct TypeEnv {
    parent: Option<*const TypeEnv>,
    variables: HashMap<String, TypeInfo>,
    types: HashMap<String, TypeInfo>,
    free_var_count: usize,
    // flag
    pub in_class_scope: bool,
}

impl TypeEnv {
    pub(crate) fn type_of_expr(&mut self, expr: &Expr) -> Result<Type> {
        use ExprVariant::*;
        let location = &expr.location;
        match &expr.value {
            Binary(l, r, op) => {
                let left_type = self.type_of_expr(l)?;
                let right_type = self.type_of_expr(r)?;
                match (left_type, right_type, op) {
                    (Type::Int, Type::Int, Operator::Plus) => Ok(Type::Int),
                    (l, r, op) => panic!("unsupported operator, {} {:?} {}", l, op, r),
                }
            }
            F64(_) => Ok(Type::F64),
            Int(_) => Ok(Type::Int),
            Bool(_) => Ok(Type::Bool),
            String(_) => Ok(Type::String),
            List(es) => {
                let expr_type: Type = if es.len() < 1 {
                    self.free_var()
                } else {
                    self.type_of_expr(&es[0])?
                };
                for e in es {
                    if expr_type != self.type_of_expr(e)? {
                        return Err(SemanticError::type_mismatched(
                            &e.location,
                            &expr_type,
                            &self.type_of_expr(e)?,
                        ));
                    }
                }
                Ok(self.get_type(location, "List")?.typ)
            }
            FuncCall(f, args) => {
                let f_type = self.type_of_expr(f)?;
                match f_type {
                    Type::FunctionType(params, ret_typ) => {
                        for (p, arg) in params.iter().zip(args.iter()) {
                            let typ = self.type_of_expr(&arg.expr)?;
                            self.unify(&arg.location, p, &typ)?;
                        }
                        Ok(*ret_typ)
                    }
                    _ => Err(SemanticError::call_on_non_function_type(
                        &f.location,
                        f_type,
                    )),
                }
            }
            DotAccess(from, access) => {
                let typ = self.type_of_expr(from)?;
                match typ {
                    Type::ClassType(name, ..) => {
                        let transform_name = format!("{}::{}", name, access);
                        self.type_of_expr(&Expr::identifier(location.clone(), transform_name))
                    }
                    _ => unreachable!(),
                }
            }
            Identifier(id) => {
                let type_info = self.get_variable(location, id.as_str())?;
                Ok(type_info.typ)
            }
            ClassConstruction(name, field_inits) => {
                if !self.in_class_scope {
                    return Err(SemanticError::cannot_use_class_construction_out_of_class(
                        location,
                    ));
                }
                let type_info = self.get_type(location, name)?;
                match &type_info.typ {
                    Type::ClassType(.., should_inits) => {
                        let mut missing_init_fields = vec![];
                        for should_init in should_inits {
                            if !field_inits.contains_key(should_init) {
                                missing_init_fields.push(should_init.clone())
                            }
                        }
                        if !missing_init_fields.is_empty() {
                            return Err(SemanticError::fields_missing_init(
                                location,
                                missing_init_fields,
                            ));
                        }
                    }
                    rest => {
                        return Err(SemanticError::cannot_construct_non_class_type(
                            location,
                            rest.clone(),
                        ));
                    }
                }
                Ok(type_info.typ)
            }
        }
    }

    pub(crate) fn unify(&self, location: &Location, expected: &Type, actual: &Type) -> Result<()> {
        use Type::*;
        match (expected, actual) {
            (Void, Void) | (Int, Int) | (Bool, Bool) | (F64, F64) | (String, String) => {
                if expected == actual {
                    Ok(())
                } else {
                    Err(SemanticError::type_mismatched(location, expected, actual))
                }
            }
            (
                ClassType(name, _, type_parameters, _),
                ClassType(name2, parents, type_parameters2, _),
            ) => {
                if name != name2 {
                    for parent in parents {
                        if self.unify(location, expected, parent).is_ok() {
                            return Ok(());
                        }
                    }
                    Err(SemanticError::type_mismatched(location, expected, actual))
                } else {
                    self.unify_type_list(location, type_parameters, type_parameters2)?;
                    Ok(())
                }
            }
            (FunctionType(ft, arg), FunctionType(ft_p, arg_p)) => {
                self.unify_type_list(location, ft, ft_p)?;
                self.unify(location, arg, arg_p)
            }
            (FreeVar(_), t) => self.unify(location, t, expected),
            (t, f @ FreeVar(_)) => {
                if t == f || !f.occurs(t.clone()) {
                    // FIXME: here we should update substitution map, but we haven't create it
                    // pseudo code would like:
                    // self.substitution_map.insert(f, t)
                    // means f is t now
                    Ok(())
                } else {
                    Ok(())
                }
            }
            (_, _) => Err(SemanticError::type_mismatched(location, expected, actual)),
        }
    }

    fn unify_type_list(
        &self,
        location: &Location,
        expected: &Vec<Type>,
        actual: &Vec<Type>,
    ) -> Result<()> {
        for (t1, t2) in expected.iter().zip(actual.iter()) {
            self.unify(location, t1, t2)?;
        }
        Ok(())
    }

    fn free_var(&mut self) -> Type {
        let typ = Type::FreeVar(self.free_var_count);
        self.free_var_count += 1;
        typ
    }
}

impl TypeEnv {
    pub fn new() -> TypeEnv {
        TypeEnv {
            parent: None,
            variables: HashMap::new(),
            types: HashMap::new(),
            free_var_count: 1,
            in_class_scope: false,
        }
    }
    pub fn with_parent(parent: &TypeEnv) -> TypeEnv {
        let mut type_env = TypeEnv::new();
        type_env.parent = Some(parent);
        // inherit the attribute from parent
        // if parent is in class scope, this of course is in class scope
        type_env.in_class_scope = parent.in_class_scope;
        type_env
    }
    pub fn from(&self, typ: &ParsedType) -> Result<Type> {
        Ok(self.get_type(&Location::none(), typ.name().as_str())?.typ)
    }
    pub fn new_function_type(&self, f: &Function) -> Result<Type> {
        let mut param_types = vec![];
        for param in &f.parameters {
            param_types.push(self.from(&param.typ)?);
        }
        Ok(Type::FunctionType(
            param_types,
            self.from(&f.ret_typ)?.into(),
        ))
    }
    pub fn new_class(&self, c: &Class) -> Result<Type> {
        match c.name.as_str() {
            // FIXME: provide a tag, e.g.
            // ```
            // @Codegen(Omit)
            // class int {}
            // ```
            "void" => Ok(Type::Void),
            "int" => Ok(Type::Int),
            "f64" => Ok(Type::F64),
            "bool" => Ok(Type::Bool),
            "string" => Ok(Type::String),
            _ => {
                let mut should_inits = vec![];
                for member in &c.members {
                    match member {
                        ClassMember::Field(field) => match &field.expr {
                            None => {
                                should_inits.push(field.name.clone());
                            }
                            _ => (),
                        },
                        _ => (),
                    }
                }
                let mut parents = vec![];
                match &c.parent_class_name {
                    Some(p_name) => {
                        let parent_typ = self.get_type(&c.location, p_name)?;
                        match &parent_typ.typ {
                            Type::TraitType => (),
                            t => {
                                return Err(SemanticError::only_trait_can_be_super_type(
                                    &c.location,
                                    t,
                                ))
                            }
                        };
                        parents.push(parent_typ.typ);
                    }
                    None => (),
                }
                Ok(Type::ClassType(
                    c.name.clone(),
                    parents,
                    vec![],
                    should_inits,
                ))
            }
        }
    }
}

impl TypeEnv {
    pub(crate) fn add_variable(&mut self, location: &Location, key: &str, typ: Type) -> Result<()> {
        if self.variables.contains_key(key) {
            Err(SemanticError::name_redefined(location, key))
        } else {
            self.variables
                .insert(key.to_string(), TypeInfo::new(location, typ));
            Ok(())
        }
    }
    pub(crate) fn get_variable(&self, location: &Location, k: &str) -> Result<TypeInfo> {
        let result = self.variables.get(k);
        match result {
            Some(t) => Ok(t.clone()),
            None => match self.parent {
                Some(env) => unsafe { env.as_ref() }.unwrap().get_variable(location, k),
                None => Err(SemanticError::no_variable(location, k)),
            },
        }
    }

    pub(crate) fn add_type(&mut self, location: &Location, key: &str, typ: Type) -> Result<()> {
        if self.types.contains_key(key) {
            Err(SemanticError::name_redefined(location, key))
        } else {
            self.types
                .insert(key.to_string(), TypeInfo::new(location, typ));
            Ok(())
        }
    }
    pub(crate) fn get_type(&self, location: &Location, k: &str) -> Result<TypeInfo> {
        let result = self.types.get(k);
        match result {
            Some(t) => Ok(t.clone()),
            None => match self.parent {
                Some(env) => unsafe { env.as_ref() }.unwrap().get_type(location, k),
                None => Err(SemanticError::no_type(location, k)),
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
    fn new(location: &Location, typ: Type) -> TypeInfo {
        TypeInfo {
            location: location.clone(),
            typ,
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum Type {
    Void,
    Int,
    Bool,
    F64,
    String,
    // TODO: complete definition
    TraitType,
    // name, parents, type parameters, uninitialized fields
    ClassType(String, Vec<Type>, Vec<Type>, Vec<String>),
    FunctionType(Vec<Type>, Box<Type>),
    FreeVar(usize),
}

impl Type {
    fn occurs(&self, t: Type) -> bool {
        use Type::*;
        match t {
            Void | Int | Bool | F64 | String => false,
            FunctionType(t1, t2) => {
                for t in t1 {
                    if self.occurs(t) {
                        return true;
                    }
                }
                self.occurs(*t2)
            }
            ClassType(_, _, type_parameters, _) => {
                for t in type_parameters {
                    if self.occurs(t) {
                        return true;
                    }
                }
                false
            }
            TraitType => unimplemented!("trait type"),
            FreeVar(_) => self.clone() == t,
        }
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
            ClassType(name, _, type_parameters, _) => {
                write!(f, "{}", name)?;
                if type_parameters.len() > 0 {
                    write!(f, "[")?;
                    for (i, g) in type_parameters.iter().enumerate() {
                        if i == type_parameters.len() - 1 {
                            write!(f, "{}", g)?;
                        } else {
                            write!(f, "{}, ", g)?;
                        }
                    }
                    write!(f, "]")?;
                }
                write!(f, "")
            }
            TraitType => unimplemented!("trait type"),
            // FIXME: print format: `(int, int): int` not `<function>`
            FunctionType(_params, _ret) => write!(f, "<function>"),
            FreeVar(n) => write!(f, "'{}", n),
        }
    }
}
