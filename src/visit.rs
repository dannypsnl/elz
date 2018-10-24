use super::ast::*;
use inkwell::basic_block::BasicBlock;
use inkwell::builder::Builder;
use inkwell::context::Context;
use inkwell::module::Module;
use inkwell::types::{BasicType, BasicTypeEnum, FunctionType};
use inkwell::values::{BasicValue, BasicValueEnum, FunctionValue, PointerValue};
use inkwell::AddressSpace;

use std::collections::HashMap;

trait Scope {
    fn add_bind(&mut self, name: String, value: PointerValue);
    fn get_bind(&self, name: &String) -> Option<&PointerValue>;
}

struct FunctionContext {
    value: FunctionValue,
    block: Block,
}

impl FunctionContext {
    fn from(value: FunctionValue) -> FunctionContext {
        FunctionContext {
            value: value,
            block: Block::new(),
        }
    }

    fn return_type(&self) -> BasicTypeEnum {
        self.value.get_return_type()
    }
}

impl Scope for FunctionContext {
    fn add_bind(&mut self, name: String, value: PointerValue) {
        self.block.add_bind(name, value);
    }
    fn get_bind(&self, name: &String) -> Option<&PointerValue> {
        self.block.get_bind(name)
    }
}

/// Block is holding a closure in compiling, it stores binding & helps others component can access them
#[derive(Debug)]
struct Block {
    binds: HashMap<String, PointerValue>,
    // TODO: children: Vec<Block>,
}

impl Block {
    fn new() -> Block {
        Block {
            binds: HashMap::new(),
        }
    }
}
impl Scope for Block {
    fn add_bind(&mut self, name: String, value: PointerValue) {
        if self.binds.contains_key(&name) {
            panic!("bind: {} already exists!", name)
        }
        self.binds.insert(name, value);
    }
    fn get_bind(&self, name: &String) -> Option<&PointerValue> {
        self.binds.get(name)
    }
}

pub struct Visitor {
    context: Context,
    builder: Builder,
    module: Module,
    global_bind: Block,
    functions: HashMap<String, FunctionContext>,
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
            global_bind: Block::new(),
            functions: HashMap::new(),
        }
    }
    pub fn visit_program(&mut self, ast_tree: Vec<Top>) -> Module {
        for ast in ast_tree {
            match ast {
                Top::Import(chain, block) => {
                    println!("chain: {:?}, block: {:?}", chain, block);
                }
                // FIXME: use exported make sure the global value should add into shared var list or not
                Top::GlobalBind(_exported, name, expr) => {
                    self.visit_global_bind(_exported, name, expr);
                }
                Top::FnDefine(Method(return_t, name, params, statements)) => {
                    self.visit_function(return_t, name, params, statements);
                }
                _ => println!("Not implement yet"),
            }
        }
        self.module.clone()
    }
    fn visit_global_bind(&mut self, _exported: bool, name: String, expr: Expr) {
        let (expr_result, elz_type) = self.visit_expr(None, expr);
        let global_value =
            self.module
                .add_global(elz_type, Some(AddressSpace::Const), name.as_str());
        global_value.set_initializer(&expr_result);
        self.global_bind
            .add_bind(name, global_value.as_pointer_value());
    }

    fn get_param_type_list(&self, params: &Vec<Parameter>) -> Vec<BasicTypeEnum> {
        if params.len() != 0 {
            let mut params = params.clone();
            // last type must be defined!
            let mut param_t = params.pop().unwrap().1.unwrap();
            let mut param_type_list = vec![self.convert(param_t.clone())];
            while let Some(param) = params.pop() {
                if let Some(t) = param.1 {
                    param_t = t;
                }
                param_type_list.push(self.convert(param_t.clone()));
            }
            param_type_list
        } else {
            vec![]
        }
    }
    fn get_fn_type(
        &self,
        name: String,
        return_t: Option<Type>,
        param_type_list: Vec<BasicTypeEnum>,
    ) -> FunctionType {
        if name == "main" {
            if return_t != None {
                panic!("fn main can't return any type");
            }
            if param_type_list.len() != 0 {
                panic!("fn main can't have any parameters");
            }
            self.context.i32_type().fn_type(&[], false)
        } else {
            if let Some(t) = return_t {
                self.convert(t.clone())
                    .fn_type(param_type_list.as_slice(), false)
            } else {
                self.context
                    .void_type()
                    .fn_type(param_type_list.as_slice(), false)
            }
        }
    }
    fn set_params_name(&self, function: FunctionValue, params: &Vec<Parameter>) {
        for (i, param) in params.iter().enumerate() {
            if let Some(p) = function.get_nth_param(i as u32) {
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
    }
    fn visit_function(
        &mut self,
        return_t: Option<Type>,
        name: String,
        params: Vec<Parameter>,
        statements: Vec<Statement>,
    ) {
        let param_type_list = self.get_param_type_list(&params);
        let fn_type = if name == "main" {
            if return_t != None {
                panic!("fn main can't return any type");
            }
            if params.len() != 0 {
                panic!("fn main can't have any parameters");
            }
            self.context.i32_type().fn_type(&[], false)
        } else {
            self.get_fn_type(name.clone(), return_t, param_type_list)
        };
        let new_fn = self.module.add_function(name.as_str(), fn_type, None);
        self.set_params_name(new_fn, &params);
        let mut context = FunctionContext::from(new_fn);
        let basic_block = self.context.append_basic_block(&new_fn, "entry");
        self.visit_statements(&mut context, statements, &basic_block);
        if name == "main" {
            let end_of_fn = new_fn
                .get_last_basic_block()
                .expect("missing basic block in function");
            self.builder.position_at_end(&end_of_fn);
            self.builder
                .build_return(Some(&self.context.i32_type().const_int(0 as u64, true)));
        }
        let end_of_fn = new_fn
            .get_last_basic_block()
            .expect("missing basic block in function");
        self.builder.position_at_end(&end_of_fn);
        self.builder
            .build_return(Some(&self.context.i32_type().const_int(1 as u64, true)));
        self.functions.insert(name, context);
    }
    fn visit_statements(
        &mut self,
        context: &mut FunctionContext,
        statements: Vec<Statement>,
        basic_block: &BasicBlock,
    ) {
        self.builder.position_at_end(basic_block);
        for stmt in statements {
            match stmt {
                Statement::LetDefine(_mutable, name, typ, expr) => {
                    let (v, type_enum) = self.visit_expr(Some(context), expr);
                    if let Some(typ) = typ {
                        let t = self.convert(typ);
                        if t != type_enum {
                            panic!(
                                "defined type is {:?}, not matching expression type: {:?}",
                                t, type_enum
                            );
                        }
                    }
                    let pv = self.builder.build_alloca(type_enum, name.as_str());
                    self.builder.build_store(pv, v);
                    context.add_bind(name, pv);
                }
                Statement::Return(e) => {
                    let (v, t) = self.visit_expr(Some(context), e);
                    if context.return_type() != t {
                        panic!(
                            "expected return {:?}, but is {:?}",
                            context.return_type(),
                            t
                        );
                    }
                    self.builder.build_return(Some(&v));
                }
                Statement::Assign(leftValue, rightValue) => {
                    let (leftV, leftT) = self.visit_expr(Some(context), leftValue);
                    let (rightV, rightT) = self.visit_expr(Some(context), rightValue);
                }
                stmt => panic!("Not implement AST: {:?} yet", stmt),
            }
        }
    }

    fn visit_expr(&mut self, scope: Option<&Scope>, expr: Expr) -> (BasicValueEnum, BasicTypeEnum) {
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
            Expr::Ident(name) => {
                let pv = if let Some(scope) = scope {
                    match scope.get_bind(&name) {
                        Some(v) => v,
                        None => match self.global_bind.get_bind(&name) {
                            Some(v) => v,
                            None => panic!("No value named {}!", name),
                        },
                    }
                } else {
                    match self.global_bind.get_bind(&name) {
                        Some(v) => v,
                        None => panic!("No value named {}!", name),
                    }
                };
                let v = self.builder.build_load(*pv, "");
                (v, v.get_type())
            }
            // Expr::AccessChain(identChain, subExprs) => {
            //     for ident in identChain {}
            //     ()
            // }
            e => panic!("not implement expression {:?} yet", e),
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

#[cfg(test)]
mod tests {
    use super::*;
    use inkwell::targets::{InitializationConfig, Target};
    use inkwell::OptimizationLevel;

    fn test_module(module: Module, fn_name: &str, expected: u64) {
        Target::initialize_native(&InitializationConfig::default())
            .expect("Failed to initialize native target");

        let ee = module
            .create_jit_execution_engine(OptimizationLevel::None)
            .expect("failed at create JIT execution engine");
        let test_fn = ee
            .get_function_value(fn_name)
            .expect("failed at get function value");
        let result = unsafe { ee.run_function(&test_fn, &vec![]) };
        let r = result.as_int(true);
        assert_eq!(r, expected);
    }

    #[test]
    fn a_function_return_an_interger() {
        let mut visitor = Visitor::new();
        visitor.visit_function(
            Some(Type("i64".to_string(), vec![])),
            "foo".to_string(),
            vec![],
            vec![Statement::Return(Expr::Integer(1))],
        );
        test_module(visitor.module, "foo", 1);
    }

    #[test]
    fn function_return_global_value() {
        let mut visitor = Visitor::new();
        visitor.visit_global_bind(false, "g".to_string(), Expr::Integer(10));
        visitor.visit_function(
            Some(Type("i64".to_string(), vec![])),
            "foo".to_string(),
            vec![],
            vec![Statement::Return(Expr::Ident("g".to_string()))],
        );
        test_module(visitor.module, "foo", 10);
    }

    #[test]
    fn function_return_local_value() {
        let mut visitor = Visitor::new();
        visitor.visit_function(
            Some(Type("i64".to_string(), vec![])),
            "foo".to_string(),
            vec![],
            vec![
                Statement::LetDefine(false, "a".to_string(), None, Expr::Integer(15)),
                Statement::Return(Expr::Ident("a".to_string())),
            ],
        );
        test_module(visitor.module, "foo", 15);
    }
}
