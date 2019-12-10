use super::*;
use crate::lexer::Location;
use std::collections::{HashMap, HashSet};

#[test]
fn test_parse_function_with_block_body() {
    let code = "\
    main(): void {}
    ";

    let mut parser = Parser::new("", code);

    let func = parser.parse_function().unwrap();
    assert_eq!(
        func,
        Function::new(
            Location::from(1, 0),
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

    let mut parser = Parser::new("", code);

    let func = parser.parse_function().unwrap();
    assert_eq!(
        func,
        Function::new(
            Location::from(1, 0),
            "add",
            vec![
                Parameter::new("x", ParsedType::type_name("int")),
                Parameter::new("y", ParsedType::type_name("int")),
            ],
            ParsedType::type_name("int"),
            Body::Expr(Expr::binary(
                Location::from(1, 27),
                Expr::identifier(Location::from(1, 27), "x"),
                Expr::identifier(Location::from(1, 31), "y"),
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

    let mut parser = Parser::new("", code);

    let var = parser.parse_variable().unwrap();
    assert_eq!(
        var,
        Variable::new(
            Location::from(1, 0),
            "x",
            ParsedType::type_name("int"),
            Expr::int(Location::from(1, 9), 1)
        )
    )
}

#[test]
fn test_parse_variable_define_with_list_value() {
    let code = "\
    x: List[int] = [1, 2, 3];
    ";

    let mut parser = Parser::new("", code);

    let var = parser.parse_variable().unwrap();
    assert_eq!(
        var,
        Variable::new(
            Location::from(1, 0),
            "x",
            ParsedType::generic_type("List", vec![ParsedType::type_name("int")]),
            Expr::list(
                Location::from(1, 15),
                vec![
                    Expr::int(Location::from(1, 16), 1),
                    Expr::int(Location::from(1, 19), 2),
                    Expr::int(Location::from(1, 22), 3),
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

    let mut parser = Parser::new("", code);

    let s = parser.parse_string().unwrap();
    let location = Location::from(1, 0);
    let expected = Expr::binary(
        location.clone(),
        Expr::binary(
            location.clone(),
            Expr::string(location.clone(), "str \"\\ value "),
            Expr::identifier(location.clone(), "a"),
            Operator::Plus,
        ),
        Expr::string(location, ""),
        Operator::Plus,
    );
    assert_eq!(s, expected)
}

#[test]
fn test_parse_function_declaration() {
    let code = "\
    foo(): void;
    ";

    let mut parser = Parser::new("", code);

    let func = parser.parse_function().unwrap();
    assert_eq!(
        func,
        Function::new_declaration(
            Location::from(1, 0),
            "foo",
            vec![],
            ParsedType::type_name("void"),
        )
    )
}

#[test]
fn test_parse_class() {
    let code = "\
                class Car {\n\
                name: string;\n\
                ::new(name: string): Car;\n\
                }";

    let mut parser = Parser::new("", code);
    let class = parser.parse_class().unwrap();
    assert_eq!(
        class,
        Class::new(
            Location::from(1, 0),
            "Car",
            vec![Field::new(
                Location::from(2, 0),
                "name",
                ParsedType::type_name("string"),
                None
            )],
            vec![],
            vec![Function::new_declaration(
                Location::from(3, 2),
                "new",
                vec![Parameter::new("name", ParsedType::type_name("string"))],
                ParsedType::type_name("Car"),
            )]
        )
    )
}

#[test]
fn test_class_construction() {
    let code = "Car { name: \"\", price: 10000 }";

    let mut fields_inits = HashMap::<String, Expr>::new();
    fields_inits.insert("name".to_string(), Expr::string(Location::from(1, 12), ""));
    fields_inits.insert("price".to_string(), Expr::int(Location::from(1, 23), 10000));

    let mut parser = Parser::new("", code);
    let class_construction = parser.parse_expression(None, None).unwrap();
    assert_eq!(
        class_construction,
        Expr::class_construction(Location::from(1, 0), "Car", fields_inits)
    )
}
