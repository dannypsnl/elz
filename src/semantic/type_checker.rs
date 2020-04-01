use super::error::Result;
use super::error::SemanticError;
use crate::ast;
use crate::ast::*;
use crate::ast::{Function, ParsedType};
use crate::lexer::Location;
use std::collections::HashMap;

pub struct TypeEnv {
    parent: Option<*const TypeEnv>,
    /// imports store information about how to lookup imported name
    pub(crate) imports: HashMap<String, String>,
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
                    (
                        Type::ClassType { name: n1, .. },
                        Type::ClassType { name: n2, .. },
                        Operator::Plus,
                    ) => {
                        if n1.as_str() == "int" && n1 == n2 {
                            Ok(self.lookup_type(location, "int")?.typ)
                        } else {
                            panic!("sadasdasda")
                        }
                    }
                    (l, r, op) => panic!("unsupported operator, {} {:?} {}", l, op, r),
                }
            }
            F64(_) => Ok(self.lookup_type(location, "f64")?.typ),
            Int(_) => Ok(self.lookup_type(location, "int")?.typ),
            Bool(_) => Ok(self.lookup_type(location, "bool")?.typ),
            String(_) => Ok(self.lookup_type(location, "string")?.typ),
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
                Ok(self.lookup_type(location, "List")?.typ)
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
            MemberAccess(from, access) => {
                let typ = self.type_of_expr(from)?;
                match typ {
                    Type::ClassType { name, members, .. } => {
                        let member = members.get_member(location, name, access)?;
                        Ok(member.typ.clone())
                    }
                    _ => unreachable!(),
                }
            }
            Identifier(id) => {
                let type_info = self.lookup_variable(location, id.as_str())?;
                Ok(type_info.typ)
            }
            ClassConstruction(name, field_inits) => {
                if !self.in_class_scope {
                    return Err(SemanticError::cannot_use_class_construction_out_of_class(
                        location,
                    ));
                }
                let type_info = self.lookup_type(location, name)?;
                match &type_info.typ {
                    Type::ClassType {
                        uninitialized_fields,
                        ..
                    } => {
                        let should_inits = uninitialized_fields;
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
            (
                ClassType {
                    name,
                    type_parameters,
                    ..
                },
                ClassType {
                    name: name2,
                    parents,
                    type_parameters: type_parameters2,
                    ..
                },
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
            imports: HashMap::new(),
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
        Ok(self
            .lookup_type(&Location::none(), typ.name().as_str())?
            .typ)
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
    pub fn new_class(&mut self, c: &Class) -> Result<Type> {
        let mut uninitialized_fields = vec![];
        let mut members = ClassMembers::new();
        for member in &c.members {
            match member {
                ast::ClassMember::Field(field) => {
                    let field_type = self.from(&field.typ)?;
                    members.add_member(
                        c.name.clone(),
                        ClassMember {
                            name: field.name.clone(),
                            location: field.location.clone(),
                            typ: field_type.clone(),
                        },
                    )?;
                    match &field.expr {
                        None => uninitialized_fields.push(field.name.clone()),
                        Some(expr) => {
                            // check expression type same as field type
                            let expr_type = self.type_of_expr(expr)?;
                            self.unify(&field.location, &field_type, &expr_type)?;
                        }
                    }
                }
                ast::ClassMember::Method(method) => {
                    members.add_member(
                        c.name.clone(),
                        ClassMember {
                            name: method.name.clone(),
                            location: method.location.clone(),
                            typ: self.new_function_type(method)?,
                        },
                    )?;
                }
                _ => (),
            }
        }
        let mut parents = vec![];
        for p_name in &c.parents {
            let parent_typ = self.lookup_type(&c.location, p_name.as_str())?;
            match &parent_typ.typ {
                Type::TraitType => parents.push(parent_typ.typ),
                t => return Err(SemanticError::only_trait_can_be_super_type(&c.location, t)),
            }
        }
        Ok(Type::ClassType {
            name: c.name.clone(),
            parents,
            type_parameters: vec![],
            uninitialized_fields,
            members,
        })
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
    pub(crate) fn lookup_variable(&self, location: &Location, k: &str) -> Result<TypeInfo> {
        let result = self.variables.get(k);
        match result {
            Some(t) => Ok(t.clone()),
            None => match self.parent {
                Some(env) => {
                    let k = match self.imports.get(k) {
                        None => k,
                        Some(v) => v,
                    };
                    unsafe { env.as_ref() }
                        .unwrap()
                        .lookup_variable(location, k)
                }
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
    pub(crate) fn lookup_type(&self, location: &Location, k: &str) -> Result<TypeInfo> {
        let result = self.types.get(k);
        match result {
            Some(t) => Ok(t.clone()),
            None => match self.parent {
                Some(env) => {
                    let k = match self.imports.get(k) {
                        None => k,
                        Some(v) => v,
                    };
                    unsafe { env.as_ref() }
                        .unwrap()
                        .lookup_variable(location, k)
                }
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
pub struct ClassMember {
    name: String,
    location: Location,
    typ: Type,
}

#[derive(Clone, Debug, PartialEq)]
pub struct ClassMembers(HashMap<String, ClassMember>);

impl ClassMembers {
    fn new() -> ClassMembers {
        ClassMembers(HashMap::new())
    }
    fn add_member(&mut self, class_name: String, member: ClassMember) -> Result<()> {
        let location = &member.location.clone();
        let member_name = member.name.clone();
        match self.0.insert(member.name.clone(), member) {
            Some(previous_field) => Err(SemanticError::redefined_member(
                location,
                member_name,
                class_name,
                previous_field.location,
            )),
            None => Ok(()),
        }
    }
    fn get_member(
        &self,
        location: &Location,
        class_name: String,
        name: &String,
    ) -> Result<&ClassMember> {
        match self.0.get(name) {
            Some(v) => Ok(v),
            None => Err(SemanticError::no_member_named(
                location,
                class_name,
                name.clone(),
            )),
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
pub enum Type {
    // TODO: complete definition
    TraitType,
    ClassType {
        name: String,
        parents: Vec<Type>,
        type_parameters: Vec<Type>,
        uninitialized_fields: Vec<String>,
        members: ClassMembers,
    },
    FunctionType(Vec<Type>, Box<Type>),
    FreeVar(usize),
}

impl Type {
    fn occurs(&self, t: Type) -> bool {
        use Type::*;
        match t {
            FunctionType(t1, t2) => {
                for t in t1 {
                    if self.occurs(t) {
                        return true;
                    }
                }
                self.occurs(*t2)
            }
            ClassType {
                name,
                type_parameters,
                ..
            } => match name.as_str() {
                "void" | "int" | "bool" | "f64" | "string" => false,
                _ => {
                    for t in type_parameters {
                        if self.occurs(t) {
                            return true;
                        }
                    }
                    false
                }
            },
            TraitType => unimplemented!("trait type"),
            FreeVar(_) => self.clone() == t,
        }
    }
}

impl std::fmt::Display for Type {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        use Type::*;
        match self {
            ClassType {
                name,
                type_parameters,
                ..
            } => {
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
