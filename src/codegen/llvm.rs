use super::ir;

pub trait LLVMValue {
    fn llvm_represent(&self) -> String;
}

impl LLVMValue for ir::Module {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        for t in &self.types {
            s.push_str(t.llvm_represent().as_str());
            s.push_str("\n");
        }
        for v in &self.variables {
            s.push_str(v.llvm_represent().as_str());
            s.push_str("\n");
        }
        for f in &self.functions {
            s.push_str(f.llvm_represent().as_str());
            s.push_str("\n");
        }
        s
    }
}

impl LLVMValue for ir::TypeDefinition {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        s.push_str(format!("%{}", self.name).as_str());
        s.push_str(" = type {");
        for (index, field) in self.fields.iter().enumerate() {
            s.push_str(field.llvm_represent().as_str());
            if index < self.fields.len() - 1 {
                s.push_str(" ");
            }
        }
        s.push_str("}");
        s
    }
}

impl ir::Instruction {
    pub(crate) fn return_void(&self) -> bool {
        match self {
            ir::Instruction::FunctionCall { ret_type, .. } => {
                if ret_type == &Box::new(ir::Type::Void) {
                    true
                } else {
                    false
                }
            }
            _ => false,
        }
    }
}

impl LLVMValue for ir::Instruction {
    fn llvm_represent(&self) -> String {
        use ir::Instruction::*;
        match self {
            Return(e) => match e {
                None => "ret void".to_string(),
                Some(ex) => {
                    let es = ex.llvm_represent();
                    let ret_typ = ex.type_();
                    format!("ret {} {}", ret_typ.llvm_represent(), es)
                }
            },
            BinaryOperation {
                id,
                op_name,
                lhs,
                rhs,
            } => {
                let mut s = String::new();
                let ret_type = lhs.type_();
                s.push_str(
                    format!(
                        "%{} = {} {} {}, {}",
                        id.borrow(),
                        op_name,
                        ret_type.llvm_represent(),
                        lhs.llvm_represent(),
                        rhs.llvm_represent()
                    )
                    .as_str(),
                );
                s
            }
            FunctionCall {
                id,
                func_name,
                ret_type,
                args_expr,
            } => {
                let mut s = String::new();
                if !self.return_void() {
                    s.push_str(format!("%{} = ", id.borrow()).as_str());
                }
                s.push_str("call ");
                s.push_str(format!("{} ", ret_type.llvm_represent()).as_str());
                s.push_str(func_name.as_str());
                s.push_str("(");
                for (index, arg_expr) in args_expr.iter().enumerate() {
                    s.push_str(arg_expr.type_().llvm_represent().as_str());
                    s.push_str(" ");
                    s.push_str(arg_expr.llvm_represent().as_str());
                    if index < args_expr.len() - 1 {
                        s.push_str(", ");
                    }
                }
                s.push_str(")");
                s
            }
            Branch {
                cond,
                if_true,
                if_false,
            } => format!(
                "br {} {}, {}, {}",
                cond.type_().llvm_represent(),
                cond.llvm_represent(),
                if_true.llvm_represent(),
                if_false.llvm_represent(),
            ),
            Goto(block) => format!("br {}", block.llvm_represent()),
            Label(label) => format!("; <label>:{}:", label.id.borrow()),
        }
    }
}

impl LLVMValue for ir::Label {
    fn llvm_represent(&self) -> String {
        format!("label %{}", self.id.borrow())
    }
}

impl LLVMValue for ir::Body {
    fn llvm_represent(&self) -> String {
        let mut s = String::new();
        for instruction in &self.instructions {
            match instruction {
                ir::Instruction::Label(..) => {
                    s.push_str(format!("{}\n", instruction.llvm_represent()).as_str());
                }
                _ => {
                    s.push_str(format!("  {}\n", instruction.llvm_represent()).as_str());
                }
            }
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
        s.push_str(self.expr.type_().llvm_represent().as_str());
        s.push_str(" ");
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
            Type::UserDefined(name) => s.push_str(format!("%{}*", name).as_str()),
        }
        s
    }
}

impl LLVMValue for ir::Expr {
    fn llvm_represent(&self) -> String {
        use ir::Expr;
        let mut s = String::new();
        match self {
            Expr::F64(f) => s.push_str(format!("{}", f).as_str()),
            Expr::I64(i) => s.push_str(format!("{}", i).as_str()),
            Expr::Bool(b) => s.push_str(format!("{}", b).as_str()),
            Expr::String(s_l) => s.push_str(format!("{}", s_l).as_str()),
            Expr::Identifier(_, name) => s.push_str(format!("%{}", name).as_str()),
            Expr::LocalIdentifier(_, id) => s.push_str(format!("%{}", id.borrow()).as_str()),
        }
        s
    }
}
