use super::*;

#[test]
fn test_parse_import() {
    let mut parser = Parser::new(
        "\
import foo::bar
"
        .to_string(),
    );

    let import = parser.parse_import().unwrap();
    assert_eq!(
        import,
        Top::Import(vec!["foo".to_string(), "bar".to_string()])
    );
}

#[test]
fn test_global_variable() {
    let mut parser = Parser::new(
        "\
x: int = 1
"
        .to_string(),
    );

    let import = parser.parse_global_variable().unwrap();
    assert_eq!(
        import,
        Top::GlobalVariable(
            Some(Type::Defined("int".to_string())),
            "x".to_string(),
            Expr::Int(1)
        )
    );
}

#[test]
fn test_global_variable_without_type() {
    let mut parser = Parser::new(
        "\
x = 1
"
        .to_string(),
    );

    let import = parser.parse_global_variable().unwrap();
    assert_eq!(
        import,
        Top::GlobalVariable(None, "x".to_string(), Expr::Int(1))
    );
}

#[test]
fn test_parse_function() {
    let mut parser = Parser::new(
        "\
add(x: int, y: int): int {
  return x + y;
}
"
        .to_string(),
    );

    let bind = parser.parse_function().unwrap();
    assert_eq!(
        bind,
        Top::FuncDefine(
            Type::Defined("int".to_string()),
            "add".to_string(),
            vec![
                Parameter(Type::Defined("int".to_string()), "x".to_string()),
                Parameter(Type::Defined("int".to_string()), "y".to_string())
            ],
            Block::from(vec![Statement::Return(Expr::Binary(
                Box::new(Expr::Identifier("x".to_string())),
                Box::new(Expr::Identifier("y".to_string())),
                Operator::Plus
            ))])
        )
    );
}

#[test]
fn test_parse_structure_type_define() {
    let mut parser = Parser::new(
        "\
type Car (
  name: string,
  price: int
)
"
        .to_string(),
    );

    let type_define = parser.parse_type_define().unwrap();
    assert_eq!(
        type_define,
        Top::StructureTypeDefine(
            "Car".to_string(),
            vec![],
            vec![
                Parameter(Type::Defined("string".to_string()), "name".to_string()),
                Parameter(Type::Defined("int".to_string()), "price".to_string()),
            ],
        )
    );
}

#[test]
fn test_parse_tagged_union_type_define() {
    let mut parser = Parser::new(
        "\
type Option 'a (
  Just(a: 'a)
  | Nothing
)
"
        .to_string(),
    );

    let type_define = parser.parse_type_define().unwrap();
    assert_eq!(
        type_define,
        Top::TaggedUnionTypeDefine(
            "Option".to_string(),
            vec![Type::Unsure("a".to_string())],
            vec![
                SubType {
                    tag: "Just".to_string(),
                    params: vec![Parameter(Type::Unsure("a".to_string()), "a".to_string()),]
                },
                SubType {
                    tag: "Nothing".to_string(),
                    params: vec![]
                },
            ],
        )
    );
}
