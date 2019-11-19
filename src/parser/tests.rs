use super::*;
use crate::lexer::Location;

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
            ParsedType::type_name("void"),
            Body::Block(Block::new())
        )
    )
}

#[test]
fn test_parse_function_with_expression_body() {
    let code = "\
    add(x: int, y: int): int = x + y;
    ";

    let mut parser = Parser::new(code);

    let func = parser.parse_function().unwrap();
    assert_eq!(
        func,
        Function::new(
            (1, 0),
            "add",
            vec![
                Parameter::new("x", ParsedType::type_name("int")),
                Parameter::new("y", ParsedType::type_name("int")),
            ],
            ParsedType::type_name("int"),
            Body::Expr(Expr::binary(
                (1, 27),
                Expr::identifier((1, 27), "x"),
                Expr::identifier((1, 31), "y"),
                Operator::Plus
            ))
        )
    )
}

#[test]
fn test_parse_variable_define() {
    let code = "\
    x: int = 1;
    ";

    let mut parser = Parser::new(code);

    let var = parser.parse_variable().unwrap();
    assert_eq!(
        var,
        Variable::new(
            (1, 0),
            "x",
            ParsedType::type_name("int"),
            Expr::int((1, 9), 1)
        )
    )
}

#[test]
fn test_parse_variable_define_with_list_value() {
    let code = "\
    x: List[int] = [1, 2, 3];
    ";

    let mut parser = Parser::new(code);

    let var = parser.parse_variable().unwrap();
    assert_eq!(
        var,
        Variable::new(
            (1, 0),
            "x",
            ParsedType::generic_type("List", vec![ParsedType::type_name("int")]),
            Expr::list(
                (1, 15),
                vec![
                    Expr::int((1, 16), 1),
                    Expr::int((1, 19), 2),
                    Expr::int((1, 22), 3),
                ]
            )
        )
    )
}

#[test]
fn test_parse_string() {
    let code = "\
    \"str \\\"\\\\ value {a}\"
    ";

    let mut parser = Parser::new(code);

    let s = parser.parse_string().unwrap();
    let location: Location = (1, 0);
    let expected = Expr::binary(
        location,
        Expr::binary(
            location,
            Expr::string(location, "str \"\\ value "),
            Expr::identifier(location, "a"),
            Operator::Plus,
        ),
        Expr::string(location, ""),
        Operator::Plus,
    );
    assert_eq!(s, expected)
}
