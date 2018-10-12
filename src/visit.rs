use super::ast::*;
use inkwell::builder::Builder;
use inkwell::context::Context;
use inkwell::module::Module;
use inkwell::types::{BasicType, BasicTypeEnum};
use inkwell::values::BasicValue;
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
                // FIXME: use exported make sure the global value should add into shared var list or not
                Top::GlobalBind(_exported, name, e) => {
                    let (expr_result, elz_type) = self.visit_const_expr(e);
                    let global_value =
                        self.module
                            .add_global(elz_type, Some(AddressSpace::Const), name.as_str());
                    global_value.set_initializer(expr_result.as_ref());
                    println!("get global IR: {:?}", self.module.get_global(name.as_str()));
                }
                _ => println!("Not implement yet"),
            }
        }
    }
    pub fn visit_const_expr(&mut self, expr: Expr) -> (Box<BasicValue>, BasicTypeEnum) {
        match expr {
            Expr::Integer(iv) => (
                Box::new(self.context.i64_type().const_int(iv as u64, true)),
                self.context.i64_type().as_basic_type_enum(),
            ),
            Expr::Number(fv) => (
                Box::new(self.context.f64_type().const_float(fv)),
                self.context.f64_type().as_basic_type_enum(),
            ),
        }
    }
}
