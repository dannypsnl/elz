use super::*;

pub(crate) trait FormattedElz {
    fn formatted_elz(&self) -> String;
}

pub(crate) struct FormatTopAstList(pub(crate) Vec<TopAst>);
impl FormattedElz for FormatTopAstList {
    fn formatted_elz(&self) -> String {
        let mut s = String::new();
        for ast in &self.0 {
            match ast {
                TopAst::Variable(v) => s.push_str(v.formatted_elz().as_str()),
                TopAst::Function(f) => s.push_str(f.formatted_elz().as_str()),
                TopAst::Class(c) => s.push_str(c.formatted_elz().as_str()),
            }
            s.push_str("\n");
        }
        s
    }
}

impl FormattedElz for Variable {
    fn formatted_elz(&self) -> String {
        let mut s = String::new();
        s.push_str(self.name.as_str());
        s.push_str(": ");
        s.push_str(self.typ.formatted_elz().as_str());
        s.push_str(" = ");
        s.push_str(self.expr.formatted_elz().as_str());
        s.push_str(";");
        s
    }
}

impl FormattedElz for Function {
    fn formatted_elz(&self) -> String {
        let mut s = String::new();
        s.push_str(self.name.as_str());
        s.push_str("(");
        for (index, param) in self.parameters.iter().enumerate() {
            s.push_str(param.1.as_str());
            s.push_str(": ");
            s.push_str(param.0.formatted_elz().as_str());
            if index < self.parameters.len() - 1 {
                s.push_str(", ");
            }
        }
        s.push_str("): ");
        s.push_str(self.ret_typ.formatted_elz().as_str());
        match &self.body {
            None => (),
            Some(b) => {
                s.push_str(b.formatted_elz().as_str());
            }
        }
        s
    }
}

impl FormattedElz for Class {
    fn formatted_elz(&self) -> String {
        unimplemented!()
    }
}

impl FormattedElz for Body {
    fn formatted_elz(&self) -> String {
        use Body::*;
        let mut s = String::new();
        match self {
            Block(block) => {
                s.push_str(block.formatted_elz().as_str());
            }
            Expr(expr) => {
                s.push_str(" = ");
                s.push_str(expr.formatted_elz().as_str());
                s.push_str(";");
            }
        }
        s
    }
}

impl FormattedElz for Block {
    fn formatted_elz(&self) -> String {
        unimplemented!()
    }
}

impl FormattedElz for Expr {
    fn formatted_elz(&self) -> String {
        use ExprVariant::*;
        match &self.value {
            Binary(le, re, op) => format!(
                "{} {} {}",
                le.formatted_elz(),
                op.formatted_elz(),
                re.formatted_elz()
            ),
            F64(f) => format!("{}", f),
            Int(i) => format!("{}", i),
            Bool(b) => format!("{}", b),
            String(s) => format!("\"{}\"", s),
            List(elems) => {
                let mut s = "".to_string();
                s.push_str("[");
                concat_with_separator(&mut s, elems, ", ");
                s.push_str("]");
                s
            }
            FuncCall(func, args) => {
                let mut s = "".to_string();
                s.push_str(func.formatted_elz().as_str());
                concat_with_separator(&mut s, args, ", ");
                s
            }
            DotAccess(var, field_name) => format!("{}.{}", var.formatted_elz(), field_name),
            Identifier(name) => name.clone(),
            ClassConstruction(name, fields_init) => unimplemented!(),
        }
    }
}

impl FormattedElz for Argument {
    fn formatted_elz(&self) -> String {
        unimplemented!()
    }
}

impl FormattedElz for Operator {
    fn formatted_elz(&self) -> String {
        use Operator::*;
        match self {
            Plus => format!("+"),
        }
    }
}

impl FormattedElz for ParsedType {
    fn formatted_elz(&self) -> String {
        use ParsedType::*;
        let mut s = String::new();
        match self {
            TypeName(name) => {
                s.push_str(name.as_str());
            }
            GenericType(name, type_parameters) => {
                s.push_str(name.as_str());
                s.push_str("[");
                concat_with_separator(&mut s, type_parameters, ", ");
                s.push_str("]");
            }
        }
        s
    }
}

// helpers
fn concat_with_separator<T: FormattedElz>(
    s: &mut String,
    vector: &Vec<T>,
    separator: &'static str,
) {
    for (index, elem) in vector.iter().enumerate() {
        s.push_str(elem.formatted_elz().as_str());
        if index < vector.len() - 1 {
            s.push_str(separator);
        }
    }
}
