use crate::ast::*;
use crate::lexer::Location;

mod error;
mod tag;
mod type_checker;

use error::{Result, SemanticError};
use tag::SemanticTag;
use type_checker::{Type, TypeEnv};

pub struct SemanticChecker {
    type_env: TypeEnv,
}

impl SemanticChecker {
    pub fn new() -> SemanticChecker {
        SemanticChecker {
            type_env: TypeEnv::new(),
        }
    }
}

impl SemanticChecker {
    pub fn check_program(&mut self, modules: &Vec<Module>) -> Result<()> {
        for m in modules {
            self.check_module(m)?;
        }
        Ok(())
    }

    fn check_module(&mut self, module: &Module) -> Result<()> {
        self.prepare_program(&module.top_list)?;
        self.check_top_list(&module.top_list)
    }

    fn prepare_program(&mut self, ast: &Vec<TopAst>) -> Result<()> {
        for top in ast {
            use TopAst::*;
            match &top {
                Class(c) => {
                    let typ = self.type_env.new_class(c)?;
                    self.type_env.add_type(&c.location, &c.name, typ)?;
                }
                _ => (),
            }
        }
        for top in ast {
            use TopAst::*;
            match &top {
                Class(c) => {
                    for member in &c.members {
                        match member {
                            ClassMember::StaticMethod(static_method) => {
                                self.type_env.add_variable(
                                    &static_method.location,
                                    &format!("{}::{}", c.name, static_method.name),
                                    self.type_env.new_function_type(static_method)?,
                                )?;
                            }
                            _ => (),
                        }
                    }
                }
                _ => (),
            }
        }
        for top in ast {
            use TopAst::*;
            match &top {
                Variable(v) => {
                    self.type_env
                        .add_variable(&v.location, &v.name, self.type_env.from(&v.typ)?)?
                }
                Function(f) => {
                    self.type_env.add_variable(
                        &f.location,
                        &f.name,
                        self.type_env.new_function_type(f)?,
                    )?;
                }
                _ => (),
            }
        }
        Ok(())
    }

    fn check_top_list(&mut self, ast: &Vec<TopAst>) -> Result<()> {
        for top in ast {
            use TopAst::*;
            match &top {
                Variable(v) => {
                    let typ = self.type_env.type_of_expr(&v.expr)?;
                    // show where error happened
                    // we are unifying <expr> and <type>, so <expr> location is better than
                    // variable define statement location
                    self.type_env
                        .unify(&v.expr.location, &self.type_env.from(&v.typ)?, &typ)?
                }
                Function(f) => self.check_function_body(&f.location, &f, &self.type_env)?,
                Class(c) => {
                    let mut class_type_env = TypeEnv::with_parent(&self.type_env);
                    for member in &c.members {
                        match member {
                            ClassMember::Field(f) => {
                                let typ = class_type_env.from(&f.typ)?;
                                class_type_env.add_variable(&f.location, &f.name, typ)?;
                            }
                            _ => (),
                        }
                    }
                    class_type_env.in_class_scope = true;
                    for member in &c.members {
                        match member {
                            ClassMember::StaticMethod(static_method) => {
                                self.check_function_body(
                                    &static_method.location,
                                    &static_method,
                                    &class_type_env,
                                )?;
                            }
                            ClassMember::Method(method) => {
                                self.check_function_body(
                                    &method.location,
                                    &method,
                                    &class_type_env,
                                )?;
                            }
                            _ => (),
                        }
                    }
                }
                Trait(_) => unimplemented!(),
                Import(_) => unimplemented!(),
            }
        }
        Ok(())
    }

    fn check_function_body(&self, location: &Location, f: &Function, env: &TypeEnv) -> Result<()> {
        let return_type = self.type_env.from(&f.ret_typ)?;
        let mut type_env = TypeEnv::with_parent(env);
        for Parameter { name, typ } in &f.parameters {
            type_env.add_variable(location, name, type_env.from(typ)?)?;
        }
        match &f.body {
            Some(Body::Expr(e)) => {
                let e_type = type_env.type_of_expr(e)?;
                type_env.unify(location, &return_type, &e_type)
            }
            Some(Body::Block(b)) => self.check_block(&type_env, b, &return_type),
            None => {
                if f.tag.is_extern() {
                    // extern function declaration don't have body need to check
                    // e.g.
                    // ```
                    // foo(): void;
                    // ```
                    Ok(())
                } else {
                    Err(SemanticError::non_extern_function_must_have_body(
                        location,
                        f.name.as_str(),
                    ))
                }
            }
        }
    }

    fn check_block(&self, type_env: &TypeEnv, b: &Block, return_type: &Type) -> Result<()> {
        let mut type_env = TypeEnv::with_parent(type_env);
        let location = &b.location;
        if b.statements.len() == 0 {
            if type_env
                .unify(
                    location,
                    return_type,
                    &type_env.lookup_type(location, "void")?.typ,
                )
                .is_err()
            {
                return Err(SemanticError::dead_code_after_return_statement(location));
            }
        } else {
            for (i, stmt) in b.statements.iter().enumerate() {
                use StatementVariant::*;
                let location = &stmt.location;
                match &stmt.value {
                    Return(e) => {
                        let typ = match e {
                            Some(e) => type_env.type_of_expr(e)?,
                            None => type_env.lookup_type(location, "void")?.typ,
                        };
                        if i != b.statements.len() - 1 {
                            return Err(SemanticError::dead_code_after_return_statement(location));
                        }
                        type_env.unify(location, return_type, &typ)?;
                    }
                    Variable(v) => {
                        let var_def_typ = type_env.from(&v.typ)?;
                        let var_typ = type_env.type_of_expr(&v.expr)?;
                        type_env.unify(location, &var_def_typ, &var_typ)?;
                        type_env.add_variable(location, &v.name, var_def_typ)?;
                        if i == b.statements.len() - 1 {
                            type_env.unify(
                                location,
                                return_type,
                                &type_env.lookup_type(location, "void")?.typ,
                            )?;
                        }
                    }
                    Expression(func_call) => {
                        let func_call_ret_typ = type_env.type_of_expr(func_call)?;
                        type_env.unify(
                            location,
                            &type_env.lookup_type(location, "void")?.typ,
                            &func_call_ret_typ,
                        )?;
                        if i == b.statements.len() - 1 {
                            type_env.unify(
                                location,
                                return_type,
                                &type_env.lookup_type(location, "void")?.typ,
                            )?;
                        }
                    }
                    IfBlock {
                        clauses,
                        else_block,
                    } => {
                        for (condition, then_block) in clauses {
                            let cond_type = type_env.type_of_expr(condition)?;
                            type_env.unify(
                                location,
                                &type_env.lookup_type(location, "bool")?.typ,
                                &cond_type,
                            )?;
                            self.check_block(&type_env, then_block, return_type)?;
                        }
                        self.check_block(&type_env, else_block, return_type)?;
                    }
                }
            }
        }
        Ok(())
    }
}

// Must put code before tests module
#[cfg(test)]
mod tests;
