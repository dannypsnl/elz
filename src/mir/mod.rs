use super::ast;
use super::ast::{Expr, Top};
use std::collections::HashMap;

pub enum Type {
    /// use as placeholder, and won't do any checking
    /// since semantic should ensure all mismatched aren't
    /// existed.
    Unsure,
    Void,
    I64,
    F64,
    Bool,
    String,
}

pub struct Context {
    parent: Option<*const Context>,
    binding_map: HashMap<String, Bind>,
}

pub struct Bind(String, Expr);

pub fn generate_mir_program(program: &Vec<ast::Top>) {
    let mut ctx = Context {
        parent: None,
        binding_map: HashMap::new(),
    };
    for top_elem in program {
        match top_elem {
            Top::Binding(name, _, expr) => {
                ctx.binding_map
                    .insert(name.clone(), Bind(name.clone(), expr.clone()));
            }
            _ => unimplemented!(),
        };
    }

    let main_binding = ctx.binding_map.get("main").unwrap();
    let Bind(name, expr) = main_binding;
    if let Expr::Lambda(lambda) = expr {
        if let Some(expr) = lambda.body.clone() {
            if let Expr::Block(block) = expr.as_ref() {
                let mut stmts = vec![];
                for stmt in &block.statements {
                    stmts.push(generate_stmt_mir(stmt));
                }
                let f = Function {
                    name: name.clone(),
                    block: stmts,
                };
            } else {
                panic!("main lambda must be a block");
            }
        } else {
            panic!("main lambda can't have an empty body");
        }
    } else {
        panic!("main is the entry point of executable and must be a lambda");
    }
}

pub fn generate_stmt_mir(stmt: &ast::Statement) -> Statement {
    use ast::Statement::*;
    match stmt {
        Return(e) => Statement::Return,
        _ => unimplemented!(),
    }
}

pub struct Function {
    name: String,
    block: Vec<Statement>,
}

pub enum Statement {
    Return,
}
