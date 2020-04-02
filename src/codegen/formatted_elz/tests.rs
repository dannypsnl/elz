use super::*;

#[test]
fn simple_variable() {
    let formatted_code = format_code("x:int=1;");
    assert_eq!(formatted_code, "x: int = 1;\n");
}

#[test]
fn simple_function() {
    let formatted_code = format_code("add(x:int,y:int):int=x+y;");
    assert_eq!(formatted_code, "add(x: int, y: int): int = x + y;\n");
}

#[test]
fn simple_function_block() {
    let formatted_code = format_code("add(x:int,y:int):int{return x+y;}");
    assert_eq!(
        formatted_code,
        "add(x: int, y: int): int {
  return x + y;
}
"
    );
}

#[test]
fn local_variable_in_function() {
    let formatted_code = format_code("foo():void{x:int=1;}");
    assert_eq!(
        formatted_code,
        "foo(): void {
  x: int = 1;
}
"
    );
}

#[test]
fn comment_in_function() {
    let formatted_code = format_code(
        "foo():void{//comment
  x(a:1,2);}
",
    );
    assert_eq!(
        formatted_code,
        "foo(): void {
  // comment
  x(a: 1, 2);
}
"
    );
}

#[test]
fn simple_class() {
    let formatted_code = format_code("class Foo{}");
    assert_eq!(
        formatted_code,
        "class Foo {}
"
    );
}

#[test]
fn class_super_type() {
    let formatted_code = format_code("class Foo<:Bar{}");
    assert_eq!(
        formatted_code,
        "class Foo <: Bar {}
"
    );
}

#[test]
fn class_multiple_super_types() {
    let formatted_code = format_code("class Foo<:Bar,Tool{}");
    assert_eq!(
        formatted_code,
        "class Foo <: Bar, Tool {}
"
    );
}

#[test]
fn class_members() {
    let formatted_code = format_code(
        "class Foo{
x:int;
::new() :Foo =Foo{x:1};
bar(): void{}
}",
    );
    assert_eq!(
        formatted_code,
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
    let formatted_code = format_code(
        "class Foo{
::bar(): void{return;}
}",
    );
    assert_eq!(
        formatted_code,
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
    let formatted_code = format_code("trait Foo{}");
    assert_eq!(
        formatted_code,
        "trait Foo {}
"
    );
}

#[test]
fn simple_comment() {
    let formatted_code = format_code(
        "//this is comment line
trait Foo {}
",
    );
    assert_eq!(
        formatted_code,
        "// this is comment line
trait Foo {}
"
    )
}

#[test]
fn multi_blank() {
    let formatted_code = format_code("
    class                 Car                               {name               :   string          ;new(name:string):Car;}
    ");
    assert_eq!(
        formatted_code,
        "class Car {
  name: string;
  new(name: string): Car;
}
"
    );
}

#[test]
fn multi_newline() {
    let formatted_code = format_code(
        "
    class Car
    {
        name:string
;
new(name:string):Car
;
}
",
    );
    assert_eq!(
        formatted_code,
        "class Car {
  name: string;
  new(name: string): Car;
}
"
    );
}

#[test]
fn test() {
    let formatted_code = format_code(
        "//this is comment line
trait Foo {}
    class Car{name:string       ;//comment line one
//comment line two;
    :  :     new(name:string):Car;   bar(i: int):void  ;  fn foo()   {b=555*10 ; a=1+b;}}class CarFoo{::bar(): void {return;}}",
    );
    assert_eq!(
        formatted_code,
        "// this is comment line
trait Foo {}
class Car {
  name: string;
  // comment line one
  // comment line two;
  ::new(name: string): Car;
  bar(i: int): void;
  fn foo() {
    b = 555 * 10;
    a = 1 + b;
  }
}
class CarFoo {
  ::bar(): void {
    return;
  }
}
"
    );
}

// helpers, must put tests before this line
fn format_code(code: &str) -> String {
    let s = formatted_elz(code.to_string());
    s
}
