use std::fs::File;
use std::io::prelude::*;

use pest::Parser;

#[derive(Parser)]
#[grammar = "grammar.pest"]
pub struct ElzParser;

use pest::iterators::Pair;

#[derive(Clone, PartialEq, Debug)]
enum Expr {
    Number(f64),
}
#[derive(Clone, PartialEq, Debug)]
enum Top {
    // export, name, expression
    GlobalBind(bool, String, Expr),
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
                println!("import statement");
            }
            Rule::global_binding => {
                let ast = parse_global_binding(rule);
                println!("ast: {:?}", ast);
            }
            Rule::EOI => {
                println!("end of file");
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
    fn parse_import_stmt() {
        parses_to!{
            parser: ElzParser,
            input: "import lib",
            rule: Rule::import_stmt,
            tokens: [ import_stmt(0, 10, [ ident(7, 10) ]) ]
        }
        parses_to!{
            parser: ElzParser,
            input: "import lib::one",
            rule: Rule::import_stmt,
            tokens: [ import_stmt(0, 15, [ ident(7, 10), ident(12, 15) ]) ]
        }
        parses_to!{
            parser:ElzParser,
            input: "import lib::sub::{block}",
            rule: Rule::import_stmt,
            tokens: [ import_stmt(0, 24, [
                ident(7, 10),
                ident(12, 15),
                import_block(17, 24, [ ident(18, 23) ])
            ]) ]
        }
    }
    #[test]
    fn test_global_binding() {
        parses_to! {
            parser: ElzParser,
            input: "_ab_c1 = 1",
            rule: Rule::global_binding,
            tokens: [
                global_binding(0, 10, [
                    ident(0, 6),
                    number(9, 10)
                ])
            ]
        }
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
    fn parse_function_define() {
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
    fn parse_type_define() {
        parses_to!{
            parser: ElzParser,
            input: "type Test (\n  field1: i32,\n  field2: f32\n)",
            rule: Rule::type_define,
            tokens: [
                type_define(0, 42, [
                    ident(5, 9),
                    type_field(14, 25, [ident(14, 20), elz_type(22, 25, [ident(22, 25)])]),
                    type_field(29, 40, [ident(29, 35), elz_type(37, 40, [ident(37, 40)])])
                ])
            ]
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
                    elz_type(5, 9, [ident(5, 9)]),
                    method(14, 26, [ident(14, 21)])
                ])
            ]
        )
    }
}
