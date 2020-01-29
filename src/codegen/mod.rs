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
            use TopAstVariant::*;
            match &top.ast {
                Function(f) => {
                    module.remember_function(f);
                }
                Variable(v) => {
                    module.remember_variable(v);
                }
                Class(_) => {}
                Trait(_) => unimplemented!(),
            }
        }
        for top in asts {
            use TopAstVariant::*;
            match  &top.ast {
                Function(f) => {
                    // FIXME: provide a tag, e.g.
                    // ```
                    // @Codegen(Omit)
                    // println(content: string): void;
                    // ```
                    if f.name.as_str() == "println" {
                        continue;
                    }
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
                Variable(v) => {
                    let var =
                        ir::Variable::new(v.name.clone(), ir::Expr::from_ast(&v.expr, &module));
                    module.push_variable(var);
                }
                Class(c) => {
                    match c.name.as_str() {
                        // FIXME: provide a tag, e.g.
                        // ```
                        // @Codegen(Omit)
                        // class int {}
                        // ```
                        "void" => continue,
                        "int" => continue,
                        "f64" => continue,
                        "bool" => continue,
                        "string" => continue,
                        "List" => continue,
                        _ => {}
                    }
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
                Trait(_) => unimplemented!(),
            }
        }
        module
    }
}

#[cfg(test)]
mod tests;
