use super::*;
use crate::ast::Type::Defined;

#[test]
fn test_parse_binding() {
    let mut parser = Parser::new(
        "\
let add = (x: int, y: int): int => {
  return x + y;
}
"
            .to_string(),
    );

    let binding = parser.parse_binding().unwrap();
    assert_eq!(
        binding,
        Top::Binding(
            "add".to_string(),
            Expr::Lambda(Lambda::new(
                Defined(AccessChain::from("int")),
                vec![
                    Parameter(Defined(AccessChain::from("int")), "x".to_string()),
                    Parameter(Defined(AccessChain::from("int")), "y".to_string()),
                ],
                Some(Block::from(vec![
                    Statement::Return(Expr::Binary(
                        Box::new(Expr::Identifier(AccessChain::from("x"))),
                        Box::new(Expr::Identifier(AccessChain::from("y"))),
                        Operator::Plus,
                    )),
                ])),
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
                Parameter(Type::Defined(AccessChain::from("string")), "name".to_string()),
                Parameter(Type::Defined(AccessChain::from("int")), "price".to_string()),
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
