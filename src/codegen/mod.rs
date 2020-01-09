use crate::ast::*;

pub mod formatted_elz;
pub mod ir;
pub mod llvm;

pub struct CodeGenerator {}

impl CodeGenerator {
    pub fn new() -> CodeGenerator {
        CodeGenerator {}
    }

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
                TopAst::Class(_) => {}
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
                TopAst::Class(c) => {
                    module.push_type(&c.name, &c.members);

                    for member in &c.members {
                        match member {
                            ClassMember::StaticMethod(static_method) => {
                                let body = match &static_method.body {
                                    Some(b) => Some(ir::Body::from_ast(b, &module)),
                                    None => None,
                                };
                                let func = ir::Function::new(
                                    format!("\"{}::{}\"", c.name, static_method.name),
                                    &static_method.parameters,
                                    ir::Type::from_ast(&static_method.ret_typ),
                                    body,
                                );
                                module.push_function(func);
                            }
                            ClassMember::Method(method) => {
                                let body = match &method.body {
                                    Some(b) => Some(ir::Body::from_ast(b, &module)),
                                    None => None,
                                };
                                let func = ir::Function::new(
                                    format!("\"{}::{}\"", c.name, method.name),
                                    &method.parameters,
                                    ir::Type::from_ast(&method.ret_typ),
                                    body,
                                );
                                module.push_function(func);
                            }
                            _ => (),
                        }
                    }
                }
            }
        }
        module
    }
}

#[cfg(test)]
mod tests;
