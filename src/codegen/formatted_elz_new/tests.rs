use super::formatted_elz;

#[test]
fn simple_variable() {
    let program = parse_code("x:int=1;");
    assert_eq!(program, "x: int = 1;\n");
}

#[test]
fn simple_function() {
    let program = parse_code("add(x:int,y:int):int=x+y;");
    assert_eq!(program, "add(x: int, y: int): int = x + y;\n");
}

#[test]
fn simple_function_block() {
    let program = parse_code("add(x:int,y:int):int{return x+y;}");
    assert_eq!(
        program,
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
        program,
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
        program,
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
        program,
        "class Foo {}
"
    );
}

#[test]
fn class_super_type() {
    let program = parse_code("class Foo<:Bar{}");
    assert_eq!(
        program,
        "class Foo <: Bar {}
"
    );
}

#[test]
fn class_multiple_super_types() {
    let program = parse_code("class Foo<:Bar,Tool{}");
    assert_eq!(
        program,
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
        program,
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
        program,
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
        program,
        "trait Foo {}
"
    );
}

#[test]
fn simple_comment() {
    let program = parse_code(
        "//this is comment line
trait Foo {}
",
    );
    assert_eq!(
        program,
        "//this is comment line
trait Foo {}
"
    )
}

#[test]
fn multi_blank() {
    let program = parse_code("
    class                 Car                               {name               :   string          ;new(name:string):Car;}
    ");
    assert_eq!(
        program,
        "class Car {
  name: string;
  new(name: string): Car;
}
"
    );
}

#[test]
fn multi_newline() {
    let program = parse_code(
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
        program,
        "class Car {
  name: string;
  new(name: string): Car;
}
"
    );
}

#[test]
fn test() {
    let program = parse_code(
        "//this is comment line
trait Foo {}
    class Car{name:string       ;//comment
//this is comment line;zxcupqwyeirotyhrg


    :  :     new(name:string):Car;   bar(i: int):void  ;  fn foo()   {b=555*10 ; a=1+b;}}class CarFoo{::bar(): void {return;}}",
    );
    assert_eq!(
        program,
        "//this is comment line
trait Foo {}
class Car {
  name: string;
  //comment
  //this is comment line;zxcupqwyeirotyhrg
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
fn parse_code(code: &str) -> String {
    let s = formatted_elz(code.to_string());
    s
}
