use crate::ast;
use crate::ast::*;

pub struct CodeGenerator {
    backend: Backend,
}

// Functionality
impl CodeGenerator {
    pub fn generate_module(&self, asts: &Vec<TopAst>) {
        let module = self.backend.generate_module(asts);
        println!("{}", module.llvm_represent());
    }
}

// Constructor
impl CodeGenerator {
    pub fn with_backend(backend: Backend) -> CodeGenerator {
        CodeGenerator { backend }
    }
}

pub enum Backend {
    LLVM,
}

impl Backend {
    fn generate_module(&self, asts: &Vec<TopAst>) -> Module {
        let mut module = Module::new();
        for top in asts {
            match top {
                TopAst::Function(f) => {
                    let func = Function::new(top.name(), Type::from_ast(&f.ret_typ));
                    module.push_function(func);
                }
                TopAst::Variable(v) => {
                    let var =
                        Variable::new(top.name(), Type::from_ast(&v.typ), Expr::from_ast(&v.expr));
                    module.push_variable(var);
                }
            }
        }
        module
    }
}

trait LLVMValue {
    fn llvm_represent(&self) -> String;
}

struct Module {
    functions: Vec<Function>,
    variables: Vec<Variable>,
}

impl Module {
    fn new() -> Module {
        Module {
            functions: vec![],
            variables: vec![],
        }
    }
    fn push_function(&mut self, f: Function) {
        self.functions.push(f);
    }
    fn push_variable(&mut self, v: Variable) {
        self.variables.push(v);
    }
}

impl LLVMValue for Module {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        for v in &self.variables {
            s.push_str(v.llvm_represent().as_ref());
            s.push_str("\n");
        }
        for f in &self.functions {
            s.push_str(f.llvm_represent().as_ref());
            s.push_str("\n");
        }
        s
    }
}

struct Function {
    name: String,
    ret_typ: Type,
}

impl Function {
    fn new(name: String, ret_typ: Type) -> Function {
        Function { name, ret_typ }
    }
}

impl LLVMValue for Function {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        s.push_str("define ");
        if self.name == "main".to_string() {
            s.push_str("i32 ");
        } else {
            s.push_str(self.ret_typ.llvm_represent().as_str());
            s.push_str(" ");
        }
        // function name need @, e.g. @main
        s.push_str("@");
        s.push_str(self.name.as_str());
        // TODO: parameters
        s.push_str("()");
        // TODO: body
        s.push_str(" {}");
        s
    }
}

struct Variable {
    name: String,
    typ: Type,
    expr: Expr,
}

impl Variable {
    fn new(name: String, typ: Type, expr: Expr) -> Variable {
        Variable { name, typ, expr }
    }
}

impl LLVMValue for Variable {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        s.push_str("@");
        s.push_str(self.name.as_str());
        s.push_str(" = ");
        s.push_str("global ");
        s.push_str(self.typ.llvm_represent().as_str());
        s.push_str(" ");
        s.push_str(self.expr.llvm_represent().as_str());
        s
    }
}

enum Type {
    Int(usize),
    Float(usize),
}

impl Type {
    fn from_ast(t: &ast::ParsedType) -> Type {
        use Type::*;
        match t.name().as_str() {
            "int" => Int(64),
            "void" => Int(32),
            "f64" => Float(64),
            "bool" => Int(1),
            _ => unimplemented!(),
        }
    }
}

impl LLVMValue for Type {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        match self {
            Type::Float(n) => s.push_str(format!("f{}", n).as_str()),
            Type::Int(n) => s.push_str(format!("i{}", n).as_str()),
        }
        s
    }
}

enum Expr {
    I64(i64),
    F64(f64),
    Bool(bool),
    String(String),
    Unknown,
}

impl Expr {
    fn from_ast(a: &ast::Expr) -> Expr {
        use ExprVariant::*;
        match &a.value {
            F64(f) => Expr::F64(*f),
            Int(i) => Expr::I64(*i),
            Bool(b) => Expr::Bool(*b),
            String(s) => Expr::String(s.clone()),
            _ => Expr::Unknown,
        }
    }
}

impl LLVMValue for Expr {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        match self {
            Expr::F64(f) => s.push_str(format!("{}", f).as_str()),
            Expr::I64(i) => s.push_str(format!("{}", i).as_str()),
            Expr::Bool(b) => s.push_str(format!("{}", b).as_str()),
            Expr::String(s_l) => s.push_str(s_l.as_str()),
            _ => s.push_str("<unknown>"),
        }
        s
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_() {
        let code = "\
        main(): void {}
        x: int = 1;";
        let module = gen_code(code);
        assert_eq!(
            module.llvm_represent(),
            "\
             @x = global i64 1\n\
             define i32 @main() {}\n"
        );
    }

    // helpers, must put tests before this line
    fn gen_code(code: &'static str) -> Module {
        let program = crate::parser::Parser::parse_program("", code).unwrap();
        let backend = Backend::LLVM;
        backend.generate_module(&program)
    }
}
