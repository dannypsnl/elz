use crate::ast::*;

pub(crate) const DEFAULT_LEVEL: usize = 0;

pub(crate) trait FormattedElz {
    fn formatted_elz(&self, level: usize) -> String;
}

pub(crate) struct FormatTopAstList(pub(crate) Vec<TopAst>);
impl FormattedElz for FormatTopAstList {
    fn formatted_elz(&self, level: usize) -> String {
        let mut s = String::new();
        for ast in &self.0 {
            use TopAstVariant::*;
            match &ast.ast {
               Variable(v) => s.push_str(v.formatted_elz(level).as_str()),
               Function(f) => s.push_str(f.formatted_elz(level).as_str()),
               Class(c) => s.push_str(c.formatted_elz(level).as_str()),
               Trait(t) => s.push_str(t.formatted_elz(level).as_str()),
            }
            s.push_str("\n");
        }
        s
    }
}

impl FormattedElz for Trait {
    fn formatted_elz(&self, level: usize) -> String {
        let mut s = String::new();
        s.push_str("trait ");
        s.push_str(self.name.as_str());
        s.push_str(" ");
        for with_trait in &self.with_traits {
            s.push_str(format!("{} ", with_trait).as_str());
        }
        s.push_str("{");
        if self.members.len() > 0 {
            s.push_str("\n");
        }
        for member in &self.members {
            s.push_str("  ".repeat(level + 1).as_str());
            s.push_str(member.formatted_elz(level + 1).as_str());
            s.push_str("\n");
        }
        s.push_str("}");

        s
    }
}

impl FormattedElz for TraitMember {
    fn formatted_elz(&self, level: usize) -> String {
        use TraitMember::*;
        match self {
            Field(f) => {
                let mut s = "".to_string();
                s.push_str(f.name.as_str());
                s.push_str(": ");
                s.push_str(f.typ.formatted_elz(level).as_str());
                s.push_str(
                    f.expr
                        .as_ref()
                        .map_or_else(|| "".to_string(), |e| format!("{}", e.formatted_elz(level)))
                        .as_str(),
                );
                s.push_str(";");
                s
            }
            Method(m) => m.formatted_elz(level),
        }
    }
}

impl FormattedElz for Variable {
    fn formatted_elz(&self, level: usize) -> String {
        let mut s = String::new();
        s.push_str(self.name.as_str());
        s.push_str(": ");
        s.push_str(self.typ.formatted_elz(level).as_str());
        s.push_str(" = ");
        s.push_str(self.expr.formatted_elz(level).as_str());
        s.push_str(";");
        s
    }
}

impl FormattedElz for Function {
    fn formatted_elz(&self, level: usize) -> String {
        let mut s = String::new();
        s.push_str(self.name.as_str());
        s.push_str("(");
        concat_with_separator(&mut s, &self.parameters, ", ", level);
        s.push_str("): ");
        s.push_str(self.ret_typ.formatted_elz(level).as_str());
        match &self.body {
            None => s.push_str(";"),
            Some(b) => {
                s.push_str(" ");
                s.push_str(b.formatted_elz(level).as_str());
            }
        }
        s
    }
}

impl FormattedElz for Parameter {
    fn formatted_elz(&self, level: usize) -> String {
        match self.name.as_str() {
            "self" => "".to_string(),
            name => {
                let mut s = "".to_string();
                s.push_str(name);
                s.push_str(": ");
                s.push_str(self.typ.formatted_elz(level).as_str());
                s
            }
        }
    }
}

impl FormattedElz for String {
    fn formatted_elz(&self, _level: usize) -> String {
        self.clone()
    }
}

impl FormattedElz for Class {
    fn formatted_elz(&self, level: usize) -> String {
        let mut s = "".to_string();
        s.push_str("class ");
        s.push_str(self.name.as_str());
        s.push_str(" ");
        if self.parents.len() > 0 {
            s.push_str("<: ");
            concat_with_separator(&mut s, &self.parents, ", ", level);
            s.push_str(" ");
        }
        s.push_str("{");
        if self.members.len() > 0 {
            s.push_str("\n");
        }
        for member in &self.members {
            s.push_str("  ".repeat(level + 1).as_str());
            s.push_str(member.formatted_elz(level + 1).as_str());
            s.push_str("\n");
        }
        s.push_str("}");
        s
    }
}

impl FormattedElz for ClassMember {
    fn formatted_elz(&self, level: usize) -> String {
        use ClassMember::*;
        match self {
            Field(f) => {
                let mut s = "".to_string();
                s.push_str(f.name.as_str());
                s.push_str(": ");
                s.push_str(f.typ.formatted_elz(level).as_str());
                s.push_str(
                    f.expr
                        .as_ref()
                        .map_or_else(|| "".to_string(), |e| format!("{}", e.formatted_elz(level)))
                        .as_str(),
                );
                s.push_str(";");
                s
            }
            StaticMethod(s) => format!("::{}", s.formatted_elz(level)),
            Method(m) => m.formatted_elz(level),
        }
    }
}

impl FormattedElz for Body {
    fn formatted_elz(&self, level: usize) -> String {
        use Body::*;
        let mut s = String::new();
        match self {
            Block(block) => s.push_str(block.formatted_elz(level).as_str()),
            Expr(expr) => {
                s.push_str("= ");
                s.push_str(expr.formatted_elz(level).as_str());
                s.push_str(";");
            }
        }
        s
    }
}

impl FormattedElz for Block {
    fn formatted_elz(&self, level: usize) -> String {
        let mut s = "".to_string();
        s.push_str("{");
        if self.statements.len() > 0 {
            s.push_str("\n");
            for stmt in &self.statements {
                s.push_str("  ".repeat(level + 1).as_str());
                s.push_str(stmt.formatted_elz(level + 1).as_str());
                s.push_str("\n");
            }
            s.push_str("  ".repeat(level).as_str());
        }
        s.push_str("}");
        s
    }
}

impl FormattedElz for Statement {
    fn formatted_elz(&self, level: usize) -> String {
        use StatementVariant::*;
        match &self.value {
            Return(e) => {
                let mut s = "".to_string();
                s.push_str("return");
                s.push_str(
                    e.as_ref()
                        .map_or_else(
                            || "".to_string(),
                            |expr| format!(" {}", expr.formatted_elz(level)),
                        )
                        .as_str(),
                );
                s.push_str(";");
                s
            }
            Variable(v) => v.formatted_elz(level),
            Expression(e) => format!("{};", e.formatted_elz(level)),
            IfBlock { .. } => unimplemented!(),
        }
    }
}

impl FormattedElz for Expr {
    fn formatted_elz(&self, level: usize) -> String {
        match &self.value {
            ExprVariant::Binary(le, re, op) => format!(
                "{} {} {}",
                le.formatted_elz(level),
                op.formatted_elz(level),
                re.formatted_elz(level)
            ),
            ExprVariant::F64(f) => format!("{}", f),
            ExprVariant::Int(i) => format!("{}", i),
            ExprVariant::Bool(b) => format!("{}", b),
            ExprVariant::String(s) => format!("\"{}\"", s),
            ExprVariant::List(elems) => {
                let mut s = "".to_string();
                s.push_str("[");
                concat_with_separator(&mut s, elems, ", ", level);
                s.push_str("]");
                s
            }
            ExprVariant::FuncCall(func, args) => {
                let mut s = "".to_string();
                s.push_str(func.formatted_elz(level).as_str());
                s.push_str("(");
                concat_with_separator(&mut s, args, ", ", level);
                s.push_str(")");
                s
            }
            ExprVariant::DotAccess(var, field_name) => {
                format!("{}.{}", var.formatted_elz(level), field_name)
            }
            ExprVariant::Identifier(name) => name.clone(),
            ExprVariant::ClassConstruction(name, fields_init) => {
                let mut s = "".to_string();
                s.push_str(name);
                s.push_str("{");
                let fields: Vec<FormatFieldInit> = fields_init
                    .into_iter()
                    .map(|(field_name, field_init)| {
                        FormatFieldInit(field_name.clone(), field_init.clone())
                    })
                    .collect();
                concat_with_separator(&mut s, &fields, ", ", level);
                s.push_str("}");
                s
            }
        }
    }
}

struct FormatFieldInit(String, Expr);

impl FormattedElz for FormatFieldInit {
    fn formatted_elz(&self, level: usize) -> String {
        format!("{}: {}", self.0, self.1.formatted_elz(level))
    }
}

impl FormattedElz for Argument {
    fn formatted_elz(&self, level: usize) -> String {
        let mut s = "".to_string();
        s.push_str(
            self.name
                .as_ref()
                .map_or_else(|| "".to_string(), |name| format!("{}: ", name))
                .as_str(),
        );
        s.push_str(self.expr.formatted_elz(level).as_str());
        s
    }
}

impl FormattedElz for Operator {
    fn formatted_elz(&self, _level: usize) -> String {
        use Operator::*;
        match self {
            Plus => format!("+"),
        }
    }
}

impl FormattedElz for ParsedType {
    fn formatted_elz(&self, level: usize) -> String {
        use ParsedType::*;
        match self {
            TypeName(name) => name.clone(),
            GenericType(name, type_parameters) => {
                let mut s = String::new();
                s.push_str(name.as_str());
                s.push_str("[");
                concat_with_separator(&mut s, type_parameters, ", ", level);
                s.push_str("]");
                s
            }
        }
    }
}

// helpers
fn concat_with_separator<T: FormattedElz>(
    s: &mut String,
    vector: &Vec<T>,
    separator: &'static str,
    level: usize,
) {
    for (index, elem) in vector.iter().enumerate() {
        let elem_str = elem.formatted_elz(level);
        s.push_str(elem_str.as_str());
        if index < vector.len() - 1 {
            if elem_str != "".to_string() {
                s.push_str(separator);
            }
        }
    }
}

#[cfg(test)]
mod tests;
