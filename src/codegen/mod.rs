use crate::ast;
use crate::ast::*;
use crate::codegen::Instruction::TempVariable;

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
                    let body = match &f.body {
                        Some(b) => Some(Body::from_ast(b)),
                        None => None,
                    };
                    let func = Function::new(
                        f.name.clone(),
                        &f.parameters,
                        Type::from_ast(&f.ret_typ),
                        body,
                    );
                    module.push_function(func);
                }
                TopAst::Variable(v) => {
                    let var = Variable::new(v.name.clone(), Expr::from_ast(&v.expr));
                    module.push_variable(var);
                }
                TopAst::Class(_) => unimplemented!(),
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

enum Instruction {
    Return(Option<Expr>),
    TempVariable(u64, Expr),
}

impl LLVMValue for Instruction {
    fn llvm_represent(&self) -> String {
        use Instruction::*;
        match self {
            Return(e) => {
                let es = match e {
                    None => "void".to_string(),
                    Some(ex) => ex.llvm_represent(),
                };
                format!("ret {}", es)
            }
            TempVariable(counter, e) => format!("%{} = {}", counter, e.llvm_represent().as_str()),
        }
    }
}

struct Body {
    counter: u64,
    statements: Vec<Instruction>,
}

impl Body {
    fn from_ast(b: &ast::Body) -> Body {
        let mut body = Body {
            counter: 0,
            statements: vec![],
        };
        match b {
            ast::Body::Expr(e) => {
                body.statements = vec![Instruction::Return(Some(Expr::from_ast(e)))];
            }
            ast::Body::Block(b) => body.generate_instructions(&b.statements),
        };
        body
    }
    fn generate_instructions(&mut self, stmts: &Vec<Statement>) {
        self.statements = stmts
            .iter()
            .map(|stmt| {
                use ast::StatementVariant::*;
                match &stmt.value {
                    Return(e) => match e {
                        None => Instruction::Return(None),
                        Some(ex) => Instruction::Return(Some(Expr::from_ast(ex))),
                    },
                    Expression(expr) => {
                        let instruction = TempVariable(self.counter, Expr::from_ast(expr));
                        self.counter += 1;
                        instruction
                    }
                    st => unimplemented!("for statement: {:#?}", st),
                }
            })
            .collect();
    }
}

impl LLVMValue for Body {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        for stmt in &self.statements {
            s.push_str(format!("  {}\n", stmt.llvm_represent()).as_str());
        }
        s
    }
}

struct Function {
    name: String,
    parameters: Vec<(String, Type)>,
    ret_typ: Type,
    body: Option<Body>,
}

impl Function {
    fn new(
        name: String,
        parsed_params: &Vec<Parameter>,
        ret_typ: Type,
        body: Option<Body>,
    ) -> Function {
        let parameters: Vec<(String, Type)> = parsed_params
            .iter()
            .map(|p| {
                let name = p.1.clone();
                let typ = Type::from_ast(&p.0);
                (name, typ)
            })
            .collect();
        Function {
            name,
            parameters,
            ret_typ,
            body,
        }
    }
}

impl LLVMValue for Function {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        let is_declaration = self.body.is_none();
        if is_declaration {
            s.push_str("declare ");
        } else {
            s.push_str("define ");
        }
        s.push_str(self.ret_typ.llvm_represent().as_str());
        s.push_str(" ");
        // function name need @, e.g. @main
        s.push_str("@");
        s.push_str(self.name.as_str());
        s.push_str("(");
        for (index, (name, typ)) in self.parameters.iter().enumerate() {
            s.push_str(typ.llvm_represent().as_str());
            s.push_str(" %");
            s.push_str(name.as_str());
            if index < self.parameters.len() - 1 {
                s.push_str(", ");
            }
        }
        s.push_str(")");
        match &self.body {
            Some(b) => {
                s.push_str(" {\n");
                s.push_str(b.llvm_represent().as_str());
                match self.ret_typ {
                    Type::Void => {
                        s.push_str("  ret void\n");
                    }
                    _ => {}
                }
                s.push_str("}");
            }
            None => (),
        };
        s
    }
}

struct Variable {
    name: String,
    expr: Expr,
}

impl Variable {
    fn new(name: String, expr: Expr) -> Variable {
        Variable { name, expr }
    }
}

impl LLVMValue for Variable {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        s.push_str("@");
        s.push_str(self.name.as_str());
        s.push_str(" = ");
        s.push_str("global ");
        s.push_str(self.expr.llvm_represent().as_str());
        s
    }
}

enum Type {
    Void,
    Int(usize),
    Float(usize),
}

impl Type {
    fn from_ast(t: &ast::ParsedType) -> Type {
        use Type::*;
        match t.name().as_str() {
            "void" => Void,
            "int" => Int(64),
            "f64" => Float(64),
            "bool" => Int(1),
            _ => unimplemented!("type `{}`", t.name()),
        }
    }
}

impl LLVMValue for Type {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        match self {
            Type::Void => s.push_str("void"),
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
            Expr::F64(f) => s.push_str(format!("f64 {}", f).as_str()),
            Expr::I64(i) => s.push_str(format!("i64 {}", i).as_str()),
            Expr::Bool(b) => s.push_str(format!("i1 {}", b).as_str()),
            Expr::String(s_l) => s.push_str(format!("String* {}", s_l).as_str()),
            _ => s.push_str("<unknown>"),
        }
        s
    }
}

#[cfg(test)]
mod tests;
