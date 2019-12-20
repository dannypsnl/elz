use crate::ast::*;

pub mod ir;
pub mod llvm;

pub struct CodeGenerator {}

// Functionality
impl CodeGenerator {
    pub fn generate_module(&self, asts: &Vec<TopAst>) -> ir::Module {
        let mut module = ir::Module::new();
        for top in asts {
            match top {
                TopAst::Function(f) => {
                    module.remember_function(f);
                }
                TopAst::Variable(v) => {
                    module.remember_variable(v);
                }
                TopAst::Class(_) => unimplemented!(),
            }
        }
        for top in asts {
            match top {
                TopAst::Function(f) => {
                    let body = match &f.body {
                        Some(b) => Some(ir::Body::from_ast(b, &module)),
                        None => None,
                    };
                    let func = ir::Function::new(
                        f.name.clone(),
                        &f.parameters,
                        ir::Type::from_ast(&f.ret_typ),
                        body,
                    );
                    module.push_function(func);
                }
                TopAst::Variable(v) => {
                    let var =
                        ir::Variable::new(v.name.clone(), ir::Expr::from_ast(&v.expr, &module));
                    module.push_variable(var);
                }
                TopAst::Class(_) => unimplemented!(),
            }
        }
        module
    }
}

// Constructor
impl CodeGenerator {
    pub fn new() -> CodeGenerator {
        CodeGenerator {}
    }
}

#[cfg(test)]
mod tests;
