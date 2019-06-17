use super::*;
use crate::lexer::TkType::Accessor;

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
        Top::Import(AccessChain::from("foo::bar"))
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
            Some(Type::Defined(AccessChain::from("int"))),
            "x".to_string(),
            Expr::Int(1),
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
fn test_parse_function_declare() {
    let mut parser = Parser::new(
        "\
add(x: int, y: int): int;
"
            .to_string(),
    );

    let bind = parser.parse_function().unwrap();
    assert_eq!(
        bind,
        Func::new(
            Type::Defined(AccessChain::from("int")),
            "add".to_string(),
            vec![
                Parameter(Type::Defined(AccessChain::from("int")), "x".to_string()),
                Parameter(Type::Defined(AccessChain::from("int")), "y".to_string())
            ],
            None,
        ),
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
        Func::new(
            Type::Defined(AccessChain::from("int")),
            "add".to_string(),
            vec![
                Parameter(Type::Defined(AccessChain::from("int")), "x".to_string()),
                Parameter(Type::Defined(AccessChain::from("int")), "y".to_string())
            ],
            Some(Block::from(vec![Statement::Return(Expr::Binary(
                Box::new(Expr::Identifier(AccessChain::from("x"))),
                Box::new(Expr::Identifier(AccessChain::from("y"))),
                Operator::Plus,
            ))])),
        ),
    );
}

#[test]
fn test_parse_contract() {
    let mut parser = Parser::new(
        "\
contract Show (
  to_string(from: Self): string;
)
"
            .to_string(),
    );

    let contract = parser.parse_contract().unwrap();
    assert_eq!(
        contract,
        Top::Contract(
            "Show".to_string(),
            vec![
                Func::new(
                    Type::Defined(AccessChain::from("string")),
                    "to_string".to_string(),
                    vec![
                        Parameter(
                            Type::Defined(AccessChain::from("Self")),
                            "from".to_string(),
                        ),
                    ],
                    None,
                ),
            ],
        )
    );
}

#[test]
fn test_parse_impl_contract() {
    let mut parser = Parser::new(
        "\
impl Show for int (
  to_string(from: int): string {
    return sprintf(\"%d\", from);
  }
)
"
            .to_string(),
    );

    let impl_contract = parser.parse_impl_contract().unwrap();
    assert_eq!(
        impl_contract,
        Top::ImplContract(
            AccessChain::from("Show"),
            vec![
                Func::new(
                    Type::Defined(AccessChain::from("string")),
                    "to_string".to_string(),
                    vec![Parameter(Type::Defined(AccessChain::from("int")), "from".to_string())],
                    Some(Block::from(
                        vec![
                            Statement::Return(Expr::FuncCall(
                                Box::new(Expr::Identifier(AccessChain::from("sprintf"))),
                                vec![
                                    Expr::Argument("".to_string(), Box::new(Expr::String("\"%d\"".to_string()))),
                                    Expr::Argument("".to_string(), Box::new(Expr::Identifier(AccessChain::from("from"))))
                                ],
                            ))
                        ]
                    )),
                ),
            ],
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
