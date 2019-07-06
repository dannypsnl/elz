use super::super::ast;
use super::super::ast::{Expr, Lambda, Operator, Statement};
use super::types::{Type, TypeVar};
use super::{check_program, infer_expr, Context, Substitution};

#[test]
fn test_program() {
    use super::super::parser;

    let mut parser = parser::Parser::new(
        "\
let a = 1

let add_one = (x: int): int => a + x
"
        .to_string(),
    );

    let program = parser.parse_program().unwrap();

    check_program(&program).unwrap();
}

#[test]
fn test_infer_block() {
    use ast::Block;
    use ast::Type::*;
    let block = Block {
        statements: vec![
            Statement::Let {
                name: "foo".to_string(),
                typ: Defined("int".to_string()),
                expr: Expr::Int(1),
            },
            Statement::Return(Expr::F64(1.8)),
        ],
    };

    let sub = &mut Substitution::new();
    let (typ, sub) = infer_expr(&mut Context::new(), &Expr::Block(block), sub).unwrap();
    assert_eq!(sub.get(&typ), Type::F64,)
}

#[test]
fn test_infer_binary() {
    use Expr::Int;
    let expr = Expr::Binary(Box::new(Int(1)), Box::new(Int(2)), Operator::Plus);
    let mut sub = Substitution::new();
    let t = infer_expr(&mut Context::new(), &expr, &mut sub).unwrap();
    assert_eq!(t.0, Type::I64);
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
    let (return_type, _) = infer_expr(&mut Context::new(), &expr, &mut sub).unwrap();
    assert_eq!(
        return_type,
        Type::Lambda(
            vec![Type::TypeVar(TypeVar(0)), Type::TypeVar(TypeVar(1))],
            Box::new(Type::TypeVar(TypeVar(0))),
        )
    );

    let func_call = Expr::FuncCall(
        Box::new(expr),
        vec![
            ast::Argument {
                name: "".to_string(),
                expr: Expr::Int(1),
            },
            ast::Argument {
                name: "".to_string(),
                expr: Expr::Int(2),
            },
        ],
    );
    let (return_type, sub) = infer_expr(&mut Context::new(), &func_call, &mut sub).unwrap();
    // substitution the return type should be integer
    assert_eq!(Type::I64, sub.get(&return_type));
}
