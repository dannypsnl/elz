use super::*;

#[test]
fn test_parse_function() {
	let mut parser = Parser::new(
		"\
add(x: int, y: int): int { \
  return x + y;            \
}
"
		.to_string(),
	);

	let bind = parser.parse_function().unwrap();
	assert_eq!(
		bind,
		Top::Func(
			Type::Normal("int".to_string()),
			"add".to_string(),
			vec![
				Parameter(Type::Normal("int".to_string()), "x".to_string()),
				Parameter(Type::Normal("int".to_string()), "y".to_string())
			],
			Block::from(vec![Statement::Return(Expr::Binary(
				Box::new(Expr::Identifier("x".to_string())),
				Box::new(Expr::Identifier("y".to_string())),
				Operator::Plus
			))])
		)
	);
}
