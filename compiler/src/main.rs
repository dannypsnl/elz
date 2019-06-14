mod ast;
mod lexer;
mod parser;

use parser::Parser;

fn main() {
    let mut parser = Parser::new("add(x: int, y: int): int { return x + y; }".to_string());
    let program = parser.parse_program().unwrap();
    println!("{:?}", program);
}
