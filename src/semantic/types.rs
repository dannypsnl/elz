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
                            expr_type,
                            self.type_of_expr(e)?,
                        ));
                    }
                }
                Ok(Type::generic_type("List", vec![expr_type]))
            }
            FuncCall(f, args) => {
                let f_type = self.type_of_expr(f)?;
                match f_type {
                    Type::FunctionType(params, ret_typ) => {
                        for (p, arg) in params.iter().zip(args.iter()) {
                            let typ = self.type_of_expr(&arg.expr)?;
                            self.unify(&arg.location, p.clone(), typ)?;
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
                    Type::ClassType(name, _, _) => {
                        let transform_name = format!("{}::{}", name, access);
                        self.type_of_expr(&Expr::identifier(location.clone(), transform_name))
                    }
                    _ => unreachable!(),
                }
            }
            Identifier(id) => {
                let type_info = self.get_variable(location, id.clone())?;
                Ok(type_info.typ)
            }
            ClassConstruction(name, field_inits) => {
                if !self.in_class_scope {
                    return Err(SemanticError::cannot_use_class_construction_out_of_class(
                        location,
                    ));
                }
                let type_info = self.get_type(location, name.clone())?;
                match &type_info.typ {
                    Type::ClassType(_, _, should_inits) => {
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

    pub(crate) fn unify(&self, location: &Location, expected: Type, actual: Type) -> Result<()> {
        use Type::*;
        match (expected.clone(), actual.clone()) {
            (Void, Void) | (Int, Int) | (Bool, Bool) | (F64, F64) | (String, String) => {
                if expected == actual {
                    Ok(())
                } else {
                    Err(SemanticError::type_mismatched(location, expected, actual))
                }
            }
            (GenericType(name, generics), GenericType(name2, generics2)) => {
                if name != name2 {
                    Err(SemanticError::type_mismatched(location, expected, actual))
                } else {
                    self.unify_type_list(location, generics, generics2)?;
                    Ok(())
                }
            }
            (ClassType(name, type_parameters, _), ClassType(name2, type_parameters2, _)) => {
                if name != name2 {
                    Err(SemanticError::type_mismatched(location, expected, actual))
                } else {
                    self.unify_type_list(location, type_parameters, type_parameters2)?;
                    Ok(())
                }
            }
            (FunctionType(ft, arg), FunctionType(ft_p, arg_p)) => {
                self.unify_type_list(location, ft, ft_p)?;
                self.unify(location, *arg, *arg_p)
            }
            (FreeVar(_), t) => self.unify(location, t, expected),
            (t, f @ FreeVar(_)) => {
                if t == f || !f.occurs(t) {
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
        expected: Vec<Type>,
        actual: Vec<Type>,
    ) -> Result<()> {
        for (t1, t2) in expected.iter().zip(actual.iter()) {
            self.unify(location, t1.clone(), t2.clone())?;
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
        let mut types = HashMap::new();
        types.insert(
            "int".to_string(),
            TypeInfo::new(&Location::none(), Type::Int),
        );
        types.insert(
            "string".to_string(),
            TypeInfo::new(&Location::none(), Type::String),
        );
        types.insert(
            "int".to_string(),
            TypeInfo::new(&Location::none(), Type::Int),
        );
        types.insert(
            "void".to_string(),
            TypeInfo::new(&Location::none(), Type::Void),
        );
        types.insert(
            "f64".to_string(),
            TypeInfo::new(&Location::none(), Type::F64),
        );
        types.insert(
            "bool".to_string(),
            TypeInfo::new(&Location::none(), Type::Bool),
        );
        types.insert(
            "string".to_string(),
            TypeInfo::new(&Location::none(), Type::String),
        );
        // FIXME: Get list type should with free variable type?
        types.insert(
            "List".to_string(),
            TypeInfo::new(
                &Location::none(),
                Type::generic_type("List", vec![Type::FreeVar(0)]),
            ),
        );
        TypeEnv {
            parent: None,
            variables: HashMap::new(),
            types,
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
        match typ.name().as_str() {
            "List" => {
                let mut type_parameters = vec![];
                for parsed_type in &typ.generics() {
                    type_parameters.push(self.from(parsed_type)?);
                }
                Ok(Type::generic_type("List", type_parameters))
            }
            _ => Ok(self.get_type(&Location::none(), typ.name())?.typ),
        }
    }
    pub fn new_function_type(&self, f: &Function) -> Result<Type> {
        let mut param_types = vec![];
        for param in &f.parameters {
            param_types.push(self.from(&param.0)?);
        }
        Ok(Type::FunctionType(
            param_types,
            self.from(&f.ret_typ)?.into(),
        ))
    }
}

impl TypeEnv {
    pub(crate) fn add_variable(
        &mut self,
        location: &Location,
        key: &String,
        typ: Type,
    ) -> Result<()> {
        if self.variables.contains_key(key) {
            Err(SemanticError::name_redefined(location, key))
        } else {
            self.variables
                .insert(key.clone(), TypeInfo::new(location, typ));
            Ok(())
        }
    }
    pub(crate) fn get_variable(&self, location: &Location, k: String) -> Result<TypeInfo> {
        let result = self.variables.get(&k);
        match result {
            Some(t) => Ok(t.clone()),
            None => match self.parent {
                Some(env) => unsafe { env.as_ref() }.unwrap().get_variable(location, k),
                None => Err(SemanticError::no_variable(location, k)),
            },
        }
    }

    pub(crate) fn add_type(&mut self, location: &Location, key: &String, typ: Type) -> Result<()> {
        if self.types.contains_key(key) {
            Err(SemanticError::name_redefined(location, key))
        } else {
            self.types.insert(key.clone(), TypeInfo::new(location, typ));
            Ok(())
        }
    }
    pub(crate) fn get_type(&self, location: &Location, k: String) -> Result<TypeInfo> {
        let result = self.types.get(&k);
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
    ClassType(String, Vec<Type>, Vec<String>),
    // name[generic_types]
    GenericType(String, Vec<Type>),
    FunctionType(Vec<Type>, Box<Type>),
    FreeVar(usize),
}

impl Type {
    fn generic_type<T: ToString>(name: T, generics: Vec<Type>) -> Type {
        Type::GenericType(name.to_string(), generics)
    }
}

impl Type {
    pub fn new_class(c: &Class) -> Type {
        let mut should_inits = vec![];
        for field in &c.fields {
            match &field.expr {
                None => {
                    should_inits.push(field.name.clone());
                }
                _ => (),
            }
        }
        Type::ClassType(c.name.clone(), vec![], should_inits)
    }

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
            GenericType(_, ts) => {
                for t in ts {
                    if self.occurs(t) {
                        return true;
                    }
                }
                false
            }
            ClassType(_, type_parameters, _) => {
                for t in type_parameters {
                    if self.occurs(t) {
                        return true;
                    }
                }
                false
            }
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
            GenericType(name, generics) => {
                write!(f, "{}", name)?;
                if generics.len() > 0 {
                    write!(f, "[")?;
                    for (i, g) in generics.iter().enumerate() {
                        if i == generics.len() - 1 {
                            write!(f, "{}", g)?;
                        } else {
                            write!(f, "{}, ", g)?;
                        }
                    }
                    write!(f, "]")?;
                }
                write!(f, "")
            }
            ClassType(name, type_parameters, _) => {
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
            // FIXME: print format: `(int, int): int` not `<function>`
            FunctionType(_params, _ret) => write!(f, "<function>"),
            FreeVar(n) => write!(f, "'{}", n),
        }
    }
}
