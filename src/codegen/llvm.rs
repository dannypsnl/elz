use super::ir;

pub trait LLVMValue {
    fn llvm_represent(&self) -> String;
}

impl LLVMValue for ir::Module {
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

impl LLVMValue for ir::Instruction {
    fn llvm_represent(&self) -> String {
        use ir::Instruction::*;
        match self {
            Return(e) => {
                let es = match e {
                    None => "void".to_string(),
                    Some(ex) => ex.llvm_represent(),
                };
                format!("ret {}", es)
            }
            TempVariable(counter, e) => {
                if e.return_void() {
                    format!("{}", e.llvm_represent().as_str())
                } else {
                    format!("%{} = {}", counter, e.llvm_represent().as_str())
                }
            }
        }
    }
}

impl LLVMValue for ir::Body {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        for stmt in &self.statements {
            s.push_str(format!("  {}\n", stmt.llvm_represent()).as_str());
        }
        s
    }
}

impl LLVMValue for ir::Function {
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
                    ir::Type::Void => {
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

impl LLVMValue for ir::Variable {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        s.push_str(self.name.as_str());
        s.push_str(" = ");
        s.push_str("global ");
        s.push_str(self.expr.llvm_represent().as_str());
        s
    }
}

impl LLVMValue for ir::Type {
    fn llvm_represent(&self) -> String {
        use ir::Type;
        let mut s = String::new();
        match self {
            Type::Void => s.push_str("void"),
            Type::Float(n) => s.push_str(format!("f{}", n).as_str()),
            Type::Int(n) => s.push_str(format!("i{}", n).as_str()),
        }
        s
    }
}

impl LLVMValue for ir::Expr {
    fn llvm_represent(&self) -> String {
        use ir::Expr;
        let mut s = String::new();
        match self {
            Expr::F64(f) => s.push_str(format!("f64 {}", f).as_str()),
            Expr::I64(i) => s.push_str(format!("i64 {}", i).as_str()),
            Expr::Bool(b) => s.push_str(format!("i1 {}", b).as_str()),
            Expr::String(s_l) => s.push_str(format!("String* {}", s_l).as_str()),
            Expr::FunctionCall(f_name, ret_type, args) => {
                s.push_str("call ");
                s.push_str(format!("{} ", ret_type.llvm_represent()).as_str());
                s.push_str(f_name.as_str());
                s.push_str("(");
                for (index, arg_expr) in args.iter().enumerate() {
                    s.push_str(arg_expr.llvm_represent().as_str());
                    if index < args.len() - 1 {
                        s.push_str(", ");
                    }
                }
                s.push_str(")");
            }
            e => unreachable!("we shouldn't call llvm_represent on: {:#?}", e),
        }
        s
    }
}
