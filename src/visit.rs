use super::ast::*;
use inkwell::builder::Builder;
use inkwell::context::Context;
use inkwell::module::Module;
use inkwell::AddressSpace;

pub struct Visitor {
    context: Context,
    builder: Builder,
    module: Module,
}

impl Visitor {
    pub fn new() -> Visitor {
        let context = Context::create();
        let builder = context.create_builder();
        let module = context.create_module("default");
        Visitor {
            context: context,
            builder: builder,
            module: module,
        }
    }
    pub fn visit_program(&mut self, ast_tree: Vec<Top>) {
        for ast in ast_tree {
            match ast {
                Top::Import(chain, block) => {
                    println!("chain: {:?}, block: {:?}", chain, block);
                }
                Top::GlobalBind(exported, name, e) => {
                    let i32_t = self.context.i32_type();
                    let global_value =
                        self.module
                            .add_global(i32_t, Some(AddressSpace::Const), name.as_str());
                    global_value.set_initializer(&i32_t.const_int(10, true));
                    println!("global bind: {} = {:?}, exported: {}", name, e, exported);
                    println!("get global IR: {:?}", self.module.get_global(name.as_str()));
                }
                _ => println!("Not implement yet"),
            }
        }
    }
}
