use crate::ast::TopAst;
use std::collections::HashMap;

mod error;
use error::SemanticError;

pub type Result<T> = std::result::Result<T, SemanticError>;

pub fn check_program(program: Vec<TopAst>) -> Result<()> {
    let _ = Program::from_ast(program)?;
    Ok(())
}

#[derive(Clone, Debug, PartialEq)]
pub struct Program {
    variables: HashMap<String, TopAst>,
}

impl Program {
    fn new() -> Program {
        Program {
            variables: HashMap::new(),
        }
    }
    fn from_ast(program: Vec<TopAst>) -> Result<Program> {
        let mut p = Program::new();
        for ast in program {
            p.add_variable(ast)?
        }
        Ok(p)
    }

    fn add_variable(&mut self, v: TopAst) -> Result<()> {
        if self.variables.contains_key(&v.name()) {
            Err(SemanticError::name_redefined(v.name(), v.location()))
        } else {
            self.variables.insert(v.name(), v);
            Ok(())
        }
    }
}
