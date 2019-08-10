use super::*;
use crate::ast::Type::Defined;

#[test]
fn test_parse_namepsace() {
    let mut parser = Parser::new(
        "\
namespace math {
  let add = (x: int, y: int): int => x + y
}
",
    );

    let namespace = parser.parse_namespace().unwrap();
    assert_eq!(
        namespace,
        Top::Namespace(
            "math".to_string(),
            vec![Top::Binding(
                "add".to_string(),
                Type::None,
                Expr::Lambda(Lambda::new(
                    Defined("int".to_string()),
                    vec![
                        Parameter(Defined("int".to_string()), "x".to_string()),
                        Parameter(Defined("int".to_string()), "y".to_string()),
                    ],
                    Some(Box::new(Expr::Binary(
                        Box::new(Expr::Identifier("x".to_string())),
                        Box::new(Expr::Identifier("y".to_string())),
                        Operator::Plus,
                    ))),
                )),
            )]
        ),
    );
}

#[test]
fn test_statement() {
    let mut parser = Parser::new(
        "\
let foo: int = 1;
",
    );

    let statement = parser.parse_statement().unwrap();
    assert_eq!(
        statement,
        Statement::Let {
            name: "foo".to_string(),
            typ: Type::Defined("int".to_string()),
            expr: Expr::Int(1),
        }
    )
}

#[test]
fn test_parse_binding() {
    let mut parser = Parser::new(
        "\
let add = (x: int, y: int): int => x + y
",
    );

    let binding = parser.parse_binding().unwrap();
    assert_eq!(
        binding,
        Top::Binding(
            "add".to_string(),
            Type::None,
            Expr::Lambda(Lambda::new(
                Defined("int".to_string()),
                vec![
                    Parameter(Defined("int".to_string()), "x".to_string()),
                    Parameter(Defined("int".to_string()), "y".to_string()),
                ],
                Some(Box::new(Expr::Binary(
                    Box::new(Expr::Identifier("x".to_string())),
                    Box::new(Expr::Identifier("y".to_string())),
                    Operator::Plus,
                ))),
            )),
        )
    )
}

#[test]
fn test_parse_structure_type_define() {
    let mut parser = Parser::new(
        "\
type Car (
  name: string,
  price: int
)
",
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
",
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
                    params: vec![Parameter(Type::Unsure("a".to_string()), "a".to_string())],
                },
                SubType {
                    tag: "Nothing".to_string(),
                    params: vec![],
                },
            ],
        )
    );
}
