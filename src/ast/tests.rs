use crate::ast::format::{FormatTopAstList, FormattedElz};
use crate::parser::Parser;

#[test]
fn simple_variable() {
    let program = parse_code("x:int=1;");
    assert_eq!(program.formatted_elz(), "x: int = 1;\n");
}

#[test]
fn simple_function() {
    let program = parse_code("add(x:int,y:int):int=x+y;");
    assert_eq!(
        program.formatted_elz(),
        "add(x: int, y: int): int = x + y;\n"
    );
}

// helpers, must put tests before this line
fn parse_code(code: &'static str) -> FormatTopAstList {
    FormatTopAstList(Parser::parse_program("", code).unwrap())
}
