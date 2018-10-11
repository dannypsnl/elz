use std::fs::File;
use std::io::prelude::*;

use pest::Parser;

use super::ast::*;

#[derive(Parser)]
#[grammar = "grammar.pest"]
pub struct ElzParser;

use pest::iterators::Pair;

fn parse_method(method: Pair<Rule>) -> Method {
    let mut pairs = method.into_inner();
    let name = pairs.next().unwrap();
    let mut params = vec![];
    while let Some(p) = pairs.next() {
        if p.as_rule() != Rule::parameter {
            break;
        }
        let mut pairs = p.into_inner();
        let p_name = pairs.next().unwrap();
        let mut p_type = None;
        if let Some(typ) = pairs.next() {
            p_type = parse_elz_type(typ);
        }
        params.push(Parameter(p_name.as_str().to_string(), p_type));
    }
    Method(name.as_str().to_string(), params)
}
fn parse_function_define(fn_def: Pair<Rule>) -> Top {
    let mut pairs = fn_def.into_inner();
    let method = pairs.next().unwrap();
    Top::FnDefine(parse_method(method))
}

fn parse_elz_type(elz_type: Pair<Rule>) -> Option<Type> {
    let mut pairs = elz_type.into_inner();
    if let Some(type_name) = pairs.next() {
        let mut templates = vec![];
        while let Some(typ) = pairs.next() {
            match parse_elz_type(typ) {
                Some(t) => templates.push(t),
                None => (),
            }
        }
        Some(Type(type_name.as_str().to_string(), templates))
    } else {
        None
    }
}
fn parse_type_field(rule: Pair<Rule>) -> TypeField {
    let mut pairs = rule.into_inner();
    let field_name = pairs.next().unwrap();
    let field_type = parse_elz_type(pairs.next().unwrap()).unwrap();
    TypeField(field_name.as_str().to_string(), field_type)
}
fn parse_type_define(rule: Pair<Rule>) -> Top {
    let mut pairs = rule.into_inner();
    let type_name = pairs.next().unwrap();
    let mut templates = vec![];
    let mut fields = vec![];
    while let Some(r) = pairs.next() {
        if r.as_rule() != Rule::ident {
            fields.push(parse_type_field(r));
            break;
        }
        templates.push(r.as_str().to_string());
    }
    for field in pairs {
        fields.push(parse_type_field(field));
    }
    Top::TypeDefine(type_name.as_str().to_string(), templates, fields)
}

fn parse_import_stmt(rule: Pair<Rule>) -> Top {
    let pairs = rule.into_inner();
    let mut chain = vec![];
    let mut block = vec![];
    for pair in pairs {
        match pair.as_rule() {
            Rule::ident => chain.push(pair.as_str().to_string()),
            Rule::import_block => {
                let mut pairs = pair.into_inner();
                for pair in pairs {
                    match pair.as_rule() {
                        Rule::ident => block.push(pair.as_str().to_string()),
                        _ => panic!("import block expect ident only"),
                    }
                }
            }
            _ => panic!("import statement expect ident & import block only"),
        }
    }
    Top::Import(chain, block)
}

fn parse_expr(rule: Pair<Rule>) -> Expr {
    match rule.as_rule() {
        Rule::number => Expr::Number(rule.as_str().to_string().parse::<f64>().unwrap()),
        _ => panic!("unknown"),
    }
}
fn parse_global_binding(rule: Pair<Rule>) -> Top {
    let mut pairs = rule.into_inner();
    let export = pairs.next().unwrap();
    let mut name = export.clone();
    let mut exported = false;
    if export.as_rule() == Rule::symbol_export {
        exported = true;
        name = pairs.next().unwrap();
    }
    let expr = pairs.next().unwrap();
    Top::GlobalBind(
        exported,
        name.as_str().to_string(),
        parse_expr(expr.clone()),
    )
}

pub fn parse_elz_program(file_name: &str) {
    let mut f = File::open(file_name).expect("file not found");
    let mut program_content = String::new();
    f.read_to_string(&mut program_content)
        .expect("failed at read file");

    println!("start compiling");
    let program = ElzParser::parse(Rule::elz_program, program_content.as_str())
        .expect("unsuccesful compile")
        .next()
        .unwrap();
    for rule in program.into_inner() {
        match rule.as_rule() {
            Rule::import_stmt => {
                let ast = parse_import_stmt(rule);
                println!("ast: {:?}", ast);
            }
            Rule::global_binding => {
                let ast = parse_global_binding(rule);
                println!("ast: {:?}", ast);
            }
            Rule::type_define => {
                let ast = parse_type_define(rule);
                println!("ast: {:?}", ast);
            }
            Rule::function_define => {
                let ast = parse_function_define(rule);
                println!("ast: {:?}", ast);
            }
            Rule::EOI => {
                println!("end of compiling");
            }
            _ => {
                println!("unhandled rule");
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::collections::HashMap;

    #[test]
    fn test_import_stmt() {
        let test_cases: HashMap<&str, Top> = vec![
            ("import lib", Top::Import(vec!["lib".to_string()], vec![])),
            (
                "import lib::sub",
                Top::Import(vec!["lib".to_string(), "sub".to_string()], vec![]),
            ),
            (
                "import lib::sub::sub",
                Top::Import(
                    vec!["lib".to_string(), "sub".to_string(), "sub".to_string()],
                    vec![],
                ),
            ),
            (
                "import lib::sub::{block0, block1}",
                Top::Import(
                    vec!["lib".to_string(), "sub".to_string()],
                    vec!["block0".to_string(), "block1".to_string()],
                ),
            ),
        ].into_iter()
        .collect();
        for (input, ast) in test_cases {
            let r = ElzParser::parse(Rule::import_stmt, input)
                .unwrap()
                .next()
                .unwrap();
            assert_eq!(ast, parse_import_stmt(r));
        }
    }
    #[test]
    fn test_global_binding() {
        let test_cases: HashMap<&str, Top> = vec![
            (
                "_ab_c1 =1",
                Top::GlobalBind(false, "_ab_c1".to_string(), Expr::Number(1.0)),
            ),
            (
                "+a= 3.1415926",
                Top::GlobalBind(true, "a".to_string(), Expr::Number(3.1415926)),
            ),
        ].into_iter()
        .collect();
        for (input, ast) in test_cases {
            let r = ElzParser::parse(Rule::global_binding, input)
                .unwrap()
                .next()
                .unwrap();
            assert_eq!(ast, parse_global_binding(r));
        }
    }
    #[test]
    fn test_function_define() {
        let test_cases: HashMap<&str, Top> = vec![
            (
                "fn test() {}",
                Top::FnDefine(Method("test".to_string(), vec![])),
            ),
            (
                "fn add(l, r: i32) {}",
                Top::FnDefine(Method(
                    "add".to_string(),
                    vec![
                        Parameter("l".to_string(), None),
                        Parameter("r".to_string(), Some(Type("i32".to_string(), vec![]))),
                    ],
                )),
            ),
        ].into_iter()
        .collect();
        for (input, ast) in test_cases {
            let r = ElzParser::parse(Rule::function_define, input)
                .unwrap()
                .next()
                .unwrap();
            assert_eq!(ast, parse_function_define(r));
        }
        parses_to! {
            parser: ElzParser,
            input: "fn test(a, b: i32) {}",
            rule: Rule::function_define,
            tokens: [
                function_define(0, 21, [
                    method(3, 21, [
                        ident(3, 7),
                        parameter(8, 9, [ident(8, 9)]),
                        parameter(11, 17, [ident(11, 12), elz_type(14, 17, [ident(14, 17)])])
                    ])
                ])
            ]
        }
    }
    #[test]
    fn test_type_define() {
        let test_cases: HashMap<&str, Top> = vec![
            (
                "type A ()",
                Top::TypeDefine("A".to_string(), vec![], vec![]),
            ),
            (
                "type List<Elem> ()",
                Top::TypeDefine("List".to_string(), vec!["Elem".to_string()], vec![]),
            ),
            (
                "type Node<Elem> ( next: Node<Elem>, elem: Elem )",
                Top::TypeDefine(
                    "Node".to_string(),
                    vec!["Elem".to_string()],
                    vec![
                        TypeField(
                            "next".to_string(),
                            Type("Node".to_string(), vec![Type("Elem".to_string(), vec![])]),
                        ),
                        TypeField("elem".to_string(), Type("Elem".to_string(), vec![])),
                    ],
                ),
            ),
        ].into_iter()
        .collect();
        for (input, ast) in test_cases {
            let r = ElzParser::parse(Rule::type_define, input)
                .unwrap()
                .next()
                .unwrap();
            assert_eq!(ast, parse_type_define(r));
        }
    }
    #[test]
    fn parse_impl_block() {
        parses_to!(
            parser: ElzParser,
            input: "impl Type {\n  method1() {}\n}",
            rule: Rule::impl_block,
            tokens: [
                impl_block(0, 28, [
                    ident(5, 9),
                    method(14, 26, [ident(14, 21)])
                ])
            ]
        )
    }
}
