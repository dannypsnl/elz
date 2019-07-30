use super::mir;
use super::mir::MIR;
use inkwell;

pub struct Generator {
    program: MIR,
    context: inkwell::context::Context,
    module: inkwell::module::Module,
}

impl Generator {
    pub fn new(program: MIR) -> Generator {
        let context = inkwell::context::Context::create();
        let module = context.create_module("");
        Generator {
            program,
            context,
            module,
        }
    }

    pub fn generate(&self) {
        for f in &self.program.functions {
            self.generate_function(f)
        }
    }

    fn generate_function(&self, f: &mir::Function) {
        let i64 = self.context.i64_type();
        let ft = i64.fn_type(&[], false);
        self.module.add_function(f.name.as_str(), ft, None);
    }
}
