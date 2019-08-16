use super::mir;
use inkwell;
use mir::MIR;

pub struct Generator {
    program: MIR,
    context: inkwell::context::Context,
    module: inkwell::module::Module,
    builder: inkwell::builder::Builder,
}

impl Generator {
    pub fn new(program: MIR) -> Generator {
        let context = inkwell::context::Context::create();
        let module = context.create_module("");
        let builder = context.create_builder();
        Generator {
            program,
            context,
            module,
            builder,
        }
    }

    pub fn generate(&self) {
        let i64 = self.context.i64_type();
        let ft = i64.fn_type(&[], false);
        let main_fn = self.module.add_function("main", ft, None);
        let main_entry = main_fn.append_basic_block("");
        self.builder.position_at_end(&main_entry);
        self.builder.build_return(Some(&i64.const_int(0, false)));
    }

    pub fn binary(&self) {
        let f = std::fs::File::create("test.bc").expect("failed at create file");
        self.module.write_bitcode_to_file(&f, true, true);
    }
}
