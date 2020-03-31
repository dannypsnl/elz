use super::*;
use crate::lexer::Location;
use crate::lexer::TkType::EOF;
use std::collections::HashMap;

#[test]
fn parse_function_with_block_body() {
    let code = "\
    main(): void {}
    ";

    let mut parser = Parser::new("", code);

    let func = parser.parse_function(None).unwrap();
    assert_eq!(
        func,
        Function::new(
            Location::from(1, 0),
            None,
            "main",
            vec![],
            ParsedType::type_name("void"),
            Body::Block(Block::new(Location::from(1, 13)))
        )
    )
}

#[test]
fn parse_function_with_expression_body() {
    let code = "\
    add(x: int, y: int): int = x + y;
    ";

    let mut parser = Parser::new("", code);

    let func = parser.parse_function(None).unwrap();
    assert_eq!(
        func,
        Function::new(
            Location::from(1, 0),
            None,
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
fn parse_function_declaration() {
    let code = "\
    foo(): void;
    ";

    let mut parser = Parser::new("", code);

    let func = parser.parse_function(None).unwrap();
    assert_eq!(
        func,
        Function::new_declaration(
            Location::from(1, 0),
            None,
            "foo",
            vec![],
            ParsedType::type_name("void"),
        )
    )
}

#[test]
fn parse_variable_define() {
    let code = "\
    x: int = 1;
    ";

    let mut parser = Parser::new("", code);

    let var = parser.parse_variable(None).unwrap();
    assert_eq!(
        var,
        Variable::new(
            Location::from(1, 0),
            None,
            "x",
            ParsedType::type_name("int"),
            Expr::int(Location::from(1, 9), 1)
        )
    )
}

#[test]
fn parse_variable_define_with_list_value() {
    let code = "\
    x: List[int] = [1, 2, 3];
    ";

    let mut parser = Parser::new("", code);

    let var = parser.parse_variable(None).unwrap();
    assert_eq!(
        var,
        Variable::new(
            Location::from(1, 0),
            None,
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
fn parse_statement_if_block() {
    let code = "if true {} else if false {} else {}";

    let mut parser = Parser::new("", code);

    assert_eq!(
        parser.parse_statement().unwrap(),
        Statement::if_block(
            Location::from(1, 0),
            vec![
                (
                    Expr::bool(Location::from(1, 3), true),
                    Block::new(Location::from(1, 8))
                ),
                (
                    Expr::bool(Location::from(1, 19), false),
                    Block::new(Location::from(1, 25))
                )
            ],
            Block::new(Location::from(1, 33))
        )
    )
}

#[test]
fn parse_expr_string() {
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
fn parse_expr_class_construction() {
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

#[test]
fn parse_class() {
    let code = "\
                class Car {\n\
                name: string;\n\
                ::new(name: string): Car;\n\
                bar(i: int): void;\n\
                }";

    let mut parser = Parser::new("", code);
    let class = parser.parse_class(None).unwrap();
    assert_eq!(
        class,
        Class::new(
            Location::from(1, 0),
            None,
            vec![],
            "Car",
            vec![],
            vec![
                ClassMember::Field(Field::new(
                    Location::from(2, 0),
                    "name",
                    ParsedType::type_name("string"),
                    None
                )),
                ClassMember::StaticMethod(Function::new_declaration(
                    Location::from(3, 2),
                    None,
                    "new",
                    vec![Parameter::new("name", ParsedType::type_name("string"))],
                    ParsedType::type_name("Car"),
                )),
                ClassMember::Method(Function::new_declaration(
                    Location::from(4, 0),
                    None,
                    "bar",
                    vec![Parameter::new("i", ParsedType::type_name("int"))],
                    ParsedType::type_name("void"),
                )),
            ]
        )
    )
}

#[test]
fn parse_class_inherit() {
    let code = "class Foo <: Bar {}";

    let mut parser = Parser::new("", code);
    let class = parser.parse_class(None).unwrap();
    assert_eq!(
        class,
        Class::new(
            Location::from(1, 0),
            None,
            vec!["Bar".to_string()],
            "Foo",
            vec![],
            vec![],
        )
    )
}

#[test]
fn parse_class_with_type_parameters() {
    let code = "class Foo[T] {}";

    let mut parser = Parser::new("", code);
    let class = parser.parse_class(None).unwrap();
    assert_eq!(
        class,
        Class::new(
            Location::from(1, 0),
            None,
            vec![],
            "Foo",
            vec![TypeParameter::new("T", vec![])],
            vec![],
        )
    )
}

#[test]
fn module() {
    let code = "module foo.bar";

    let mut parser = Parser::new("", code);
    let module = parser.parse_module(EOF).unwrap();
    assert_eq!(
        module,
        Module {
            name: "foo.bar".to_string(),
            top_list: vec![]
        }
    )
}

#[test]
fn import_all() {
    let code = "import foo.bar";

    let mut parser = Parser::new("", code);
    let i = parser.parse_import().unwrap();
    assert_eq!(
        i,
        Import {
            location: Location::from(1, 0),
            import_path: "foo.bar".to_string(),
            imported_component: vec![]
        }
    )
}

#[test]
fn import() {
    let code = "import foo ( bar )";

    let mut parser = Parser::new("", code);
    let i = parser.parse_import().unwrap();
    assert_eq!(
        i,
        Import {
            location: Location::from(1, 0),
            import_path: "foo".to_string(),
            imported_component: vec!["bar".to_string()]
        }
    )
}

#[test]
fn parse_tag() {
    let code = "@builtin";

    let mut parser = Parser::new("", code);
    let tag = parser.parse_tag().unwrap().unwrap();
    assert_eq!(tag, Tag::new("builtin", vec![]))
}
