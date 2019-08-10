use super::super::ast::*;
use std::collections::HashMap;

pub fn flat_package(prefix_name: impl Into<String>, program: &Vec<Top>) -> HashMap<String, &Top> {
    let mut m = HashMap::new();
    let prefix_name = prefix_name.into();
    for top_elem in program {
        match top_elem {
            Top::Binding(name, _, _) => {
                let s = if prefix_name != "".to_string() {
                    let mut s = prefix_name.clone();
                    s.push_str("::");
                    s.push_str(&name);
                    s
                } else {
                    name.clone()
                };
                m.insert(s, top_elem);
            }
            Top::Namespace(_name, _elements) => {
                let m2 = flat_package(_name, _elements);
                m = m.into_iter().chain(m2).collect();
            }
            _ => unimplemented!(),
        }
    }
    m
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_flatten() {
        let program = vec![
            Top::Binding(
                "a".to_string(),
                Type::None,
                Expr::Identifier("foo::bar".to_string()),
            ),
            Top::Namespace(
                "foo".to_string(),
                vec![Top::Binding("bar".to_string(), Type::None, Expr::Int(1))],
            ),
        ];
        let map = flat_package("", &program);
        assert_eq!(
            map["foo::bar"],
            &Top::Binding("bar".to_string(), Type::None, Expr::Int(1))
        );
        assert_eq!(
            map["a"],
            &Top::Binding(
                "a".to_string(),
                Type::None,
                Expr::Identifier("foo::bar".to_string()),
            ),
        );
    }
}
