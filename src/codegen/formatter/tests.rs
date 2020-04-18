use super::*;

#[test]
fn simple_variable() {
    let formatted_code = format_elz("x:int=1;".to_string());
    assert_eq!(formatted_code, "x: int = 1;\n");
}

#[test]
fn simple_function() {
    let formatted_code = format_elz("add(x:int,y:int):int=x+y;".to_string());
    assert_eq!(formatted_code, "add(x: int, y: int): int = x + y;\n");
}

#[test]
fn simple_function_block() {
    let formatted_code = format_elz("add(x:int,y:int):int{return x+y;}".to_string());
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
    let formatted_code = format_elz("foo():void{x:int=1;}".to_string());
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
    let formatted_code = format_elz(
        "foo():void{//comment
  x(a:1,2);}
"
        .to_string(),
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
    let formatted_code = format_elz("class Foo{}".to_string());
    assert_eq!(
        formatted_code,
        "class Foo {}
"
    );
}

#[test]
fn class_super_type() {
    let formatted_code = format_elz("class Foo<:Bar{}".to_string());
    assert_eq!(
        formatted_code,
        "class Foo <: Bar {}
"
    );
}

#[test]
fn class_multiple_super_types() {
    let formatted_code = format_elz("class Foo<:Bar,Tool{}".to_string());
    assert_eq!(
        formatted_code,
        "class Foo <: Bar, Tool {}
"
    );
}

#[test]
fn class_members() {
    let formatted_code = format_elz(
        "class Foo{
x:int;
::new() :Foo =Foo{x:1};
bar(): void{}
}"
        .to_string(),
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
    let formatted_code = format_elz(
        "class Foo{
::bar(): void{return;}
}"
        .to_string(),
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
    let formatted_code = format_elz("trait Foo{}".to_string());
    assert_eq!(
        formatted_code,
        "trait Foo {}
"
    );
}

#[test]
fn simple_comment() {
    let formatted_code = format_elz(
        "//this is comment line
trait Foo {}
"
        .to_string(),
    );
    assert_eq!(
        formatted_code,
        "// this is comment line
trait Foo {}
"
    );
}

#[test]
fn multi_blank() {
    let formatted_code = format_elz("
    class                 Car                               {name               :   string          ;new(name:string):Car;}
    ".to_string());
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
    let formatted_code = format_elz(
        "
    class Car
    {
        name:string
;
new(name:string):Car
;
}
"
        .to_string(),
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
    let formatted_code = format_elz(
        "module main
//this is comment line
trait Foo {}
    class Car{name:string       ;//comment line one
//comment line two;
    :  :     new(name:string):Car;   bar(i: int):void  ;  fn foo()   {b=555*10 ; a=1+b;}}class CarFoo{::bar(): void {return;}}".to_string(),
    );
    assert_eq!(
        formatted_code,
        "module main
// this is comment line
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

#[test]
fn simple_module() {
    let formatted_code = "module main

main(): void {}
";
    assert_eq!(
        formatted_code,
        "module main

main(): void {}
"
    );
}
