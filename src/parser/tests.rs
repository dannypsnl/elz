use super::*;
use crate::ast::Function;

#[test]
fn test_parse_function_with_block_body() {
    let code = "\
    main(): void {}
    ";

    let mut parser = Parser::new(code);

    let func = parser.parse_function().unwrap();
    assert_eq!(
        func,
        Function::new(
            (1, 0),
            "main",
            vec![],
            ParsedType::new("void"),
            Body::Block(Block::new())
        )
    )
}

#[test]
fn test_parse_function_with_expression_body() {
    let code = "\
    add(x: int, y: int): int = x + y
    ";

    let mut parser = Parser::new(code);

    let func = parser.parse_function().unwrap();
    assert_eq!(
        func,
        Function::new(
            (1, 0),
            "add",
            vec![
                Parameter::new("x", ParsedType::new("int")),
                Parameter::new("y", ParsedType::new("int")),
            ],
            ParsedType::new("int"),
            Body::Expr(Expr::binary(
                (1, 27),
                Expr::identifier((1, 27), "x"),
                Expr::identifier((2, 0), "y"),
                Operator::Plus
            ))
        )
    )
}

#[test]
fn test_parse_variable_define() {
    let code = "\
    x: int = 1
    ";

    let mut parser = Parser::new(code);

    let var = parser.parse_variable().unwrap();
    assert_eq!(
        var,
        Variable::new((1, 0), "x", ParsedType::new("int"), Expr::int((2, 0), 1))
    )
}
