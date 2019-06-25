use super::*;
use crate::ast::Expr::{Binary, Int};

#[test]
fn test_infer_binary() {
    let mut ctx = Context { type_environment: HashMap::new() };
    let expr = Binary(
        Box::new(Int(1)),
        Box::new(Int(2)),
        Operator::Plus,
    );
    let mut sub = Substitution(HashMap::new());
    let t = infer_expr(&mut ctx, expr, &mut sub).unwrap();
    assert_eq!(
        t.0,
        Type::I64
    );
}