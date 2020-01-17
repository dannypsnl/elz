use super::{FormatTopAstList, FormattedElz, DEFAULT_LEVEL};
use crate::parser::Parser;

#[test]
fn simple_variable() {
    let program = parse_code("x:int=1;");
    assert_eq!(program.formatted_elz(DEFAULT_LEVEL), "x: int = 1;\n");
}

#[test]
fn simple_function() {
    let program = parse_code("add(x:int,y:int):int=x+y;");
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "add(x: int, y: int): int = x + y;\n"
    );
}

#[test]
fn simple_function_block() {
    let program = parse_code("add(x:int,y:int):int{return x+y;}");
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "add(x: int, y: int): int {
  return x + y;
}
"
    );
}

#[test]
fn local_variable_in_function() {
    let program = parse_code("foo():void{x:int=1;}");
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "foo(): void {
  x: int = 1;
}
"
    );
}

#[test]
fn function_call_in_function() {
    let program = parse_code("foo():void{x(a:1,2);}");
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "foo(): void {
  x(a: 1, 2);
}
"
    );
}

#[test]
fn simple_class() {
    let program = parse_code("class Foo{}");
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "class Foo {}
"
    );
}

#[test]
fn class_super_type() {
    let program = parse_code("class Foo<:Bar{}");
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "class Foo <: Bar {}
"
    );
}

#[test]
fn class_multiple_super_types() {
    let program = parse_code("class Foo<:Bar,Tool{}");
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "class Foo <: Bar, Tool {}
"
    );
}

#[test]
fn class_members() {
    let program = parse_code(
        "class Foo{
x:int;
::new() :Foo =Foo{x:1};
bar(): void{}
}",
    );
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "class Foo {
  x: int;
  ::new(): Foo = Foo{x: 1};
  bar(): void {}
}
"
    );
}

#[test]
fn nested_block() {
    let program = parse_code(
        "class Foo{
::bar(): void{return;}
}",
    );
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "class Foo {
  ::bar(): void {
    return;
  }
}
"
    );
}

#[test]
fn simple_trait() {
    let program = parse_code("trait Foo{}");
    assert_eq!(
        program.formatted_elz(DEFAULT_LEVEL),
        "trait Foo {}
"
    );
}

// helpers, must put tests before this line
fn parse_code(code: &'static str) -> FormatTopAstList {
    let program = match Parser::parse_program("", code) {
        Ok(program) => program,
        Err(err) => panic!("{}", err),
    };
    FormatTopAstList(program)
}
