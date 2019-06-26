use super::{
    infer_expr,
    Substitution,
    Context,
    types::{
        TypeVar,
        Type,
    },
};
use super::super::ast;
use super::super::ast::{
    Expr,
    Operator,
    Lambda,
};

#[test]
fn test_infer_binary() {
    use Expr::Int;
    let expr = Expr::Binary(
        Box::new(Int(1)),
        Box::new(Int(2)),
        Operator::Plus,
    );
    let mut sub = Substitution::new();
    let t = infer_expr(&mut Context::new(), expr, &mut sub).unwrap();
    assert_eq!(
        t.0,
        Type::I64
    );
}

#[test]
fn test_infer_lambda() {
    let expr = Expr::Lambda(Lambda::new(
        ast::Type::Unsure("a".to_string()),
        vec![
            ast::Parameter(ast::Type::Unsure("a".to_string()), "x".to_string()),
            ast::Parameter(ast::Type::Unsure("b".to_string()), "y".to_string()),
        ],
        Some(Box::new(Expr::Binary(
            Box::new(Expr::Identifier("x".to_string())),
            Box::new(Expr::Identifier("y".to_string())),
            Operator::Plus,
        ))),
    ));
    let mut sub = Substitution::new();
    let (return_type, _) = infer_expr(&mut Context::new(), expr, &mut sub).unwrap();
    assert_eq!(
        return_type,
        Type::Lambda(
            vec![
                Type::TypeVar(TypeVar(0)),
                Type::TypeVar(TypeVar(1))
            ],
            Box::new(Type::TypeVar(TypeVar(0))),
        )
    );
}