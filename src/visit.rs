use super::ast::*;
use inkwell::builder::Builder;
use inkwell::context::Context;
use inkwell::module::Module;
use inkwell::types::{BasicType, BasicTypeEnum};
use inkwell::values::{BasicValue, BasicValueEnum, FunctionValue};
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
    pub fn visit_program(&mut self, ast_tree: Vec<Top>) -> Module {
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
                    global_value.set_initializer(&expr_result);
                }
                Top::FnDefine(Method(return_t, name, params, statements)) => {
                    let param_type_list = if params.len() != 0 {
                        let mut params = params.clone();
                        // last type must be defined!
                        let mut param_t = params.pop().unwrap().1.unwrap();
                        let mut param_type_list = vec![];
                        while let Some(param) = params.pop() {
                            if let Some(t) = param.1 {
                                param_t = t;
                            }
                            param_type_list.push(self.convert(param_t.clone()));
                        }
                        param_type_list
                    } else {
                        vec![]
                    };
                    let fn_type = if let Some(t) = return_t {
                        self.convert(t).fn_type(param_type_list.as_slice(), false)
                    } else {
                        self.context
                            .void_type()
                            .fn_type(param_type_list.as_slice(), false)
                    };
                    let new_fn = self.module.add_function(name.as_str(), fn_type, None);
                    for (i, param) in params.iter().enumerate() {
                        if let Some(p) = new_fn.get_nth_param(i as u32) {
                            let name = param.0.as_str();
                            use inkwell::values::BasicValueEnum::*;
                            match p {
                                ArrayValue(t) => t.set_name(name),
                                IntValue(t) => t.set_name(name),
                                FloatValue(t) => t.set_name(name),
                                PointerValue(t) => t.set_name(name),
                                StructValue(t) => t.set_name(name),
                                VectorValue(t) => t.set_name(name),
                            };
                        }
                    }
                    for stmt in statements {
                        self.visit_statement(stmt, new_fn);
                    }
                }
                _ => println!("Not implement yet"),
            }
        }
        self.module.clone()
    }
    fn visit_statement(&mut self, stmt: Statement, new_fn: FunctionValue) {
        let basic_block = self.context.append_basic_block(&new_fn, "entry");
        self.builder.position_at_end(&basic_block);
        match stmt {
            Statement::LetDefine(_mutable, name, typ, expr) => {
                let (v, type_enum) = self.visit_const_expr(expr);
                let pv = self.builder.build_alloca(type_enum, name.as_str());
                self.builder.build_store(pv, v);
            }
            Statement::Return(e) => {
                let (v, _t) = self.visit_const_expr(e);
                self.builder.build_return(Some(&v));
            }
            stmt => panic!("Not implement AST: {:?} yet", stmt),
        }
    }

    pub fn visit_const_expr(&mut self, expr: Expr) -> (BasicValueEnum, BasicTypeEnum) {
        match expr {
            Expr::Integer(iv) => (
                self.context
                    .i64_type()
                    .const_int(iv as u64, true)
                    .as_basic_value_enum(),
                self.context.i64_type().as_basic_type_enum(),
            ),
            Expr::Number(fv) => (
                self.context
                    .f64_type()
                    .const_float(fv)
                    .as_basic_value_enum(),
                self.context.f64_type().as_basic_type_enum(),
            ),
        }
    }

    fn convert(&self, typ: Type) -> BasicTypeEnum {
        match typ {
            Type(t, v) => {
                if v.len() == 0 {
                    match t.as_str() {
                        "i8" => self.context.i8_type().as_basic_type_enum(),
                        "i16" => self.context.i16_type().as_basic_type_enum(),
                        "i32" => self.context.i32_type().as_basic_type_enum(),
                        "i64" => self.context.i64_type().as_basic_type_enum(),
                        "f32" => self.context.f32_type().as_basic_type_enum(),
                        "f64" => self.context.f64_type().as_basic_type_enum(),
                        t => panic!("unknown {}", t),
                    }
                } else {
                    self.context.i32_type().as_basic_type_enum()
                }
            }
        }
    }
}
