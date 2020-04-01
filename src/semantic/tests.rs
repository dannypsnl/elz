use super::*;
use crate::lexer::TkType;
use crate::parser::{parse_prelude, Parser};

#[test]
fn test_redefine_variable_would_get_error() {
    let code = "
    x: int = 1;
    x: int = 2;
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_function_and_variable_use_the_same_space() {
    let code = "
    x: int = 1;
    x(): void {}
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn binary_expression() -> Result<()> {
    let code = "add(x: int, y: int): int = x + y;";
    check_code(code)
}

#[test]
fn heterogeneous_list() {
    let code = "x: List[int] = [1, \"s\"];";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn call_on_non_function_type() {
    let code = "
    x: int = 1;
    main(): void {
      x();
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_type_mismatched() {
    let code = "
    x: int = \"str\";
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn call_non_void_function_as_statement() {
    let code = "
    main(): void {
      foo();
    }
    foo(): int = 1;
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn call_void_function_as_statement() -> Result<()> {
    let code = "
    main(): void {
      foo();
    }
    foo(): void {}
    ";
    check_code(code)
}

#[test]
fn test_check_function_call() -> Result<()> {
    let code = "
    x(a: int): int = a;
    y: int = x(2);
    ";
    check_code(code)
}

#[test]
fn test_unify_list_type() -> Result<()> {
    let code = "
    x: List[int] = [1, 2, 3];
    ";
    check_code(code)
}

#[test]
fn test_unify_free_var() -> Result<()> {
    let code = "
    x: List[int] = [];
    ";
    check_code(code)
}

#[test]
fn test_check_return_nothing() -> Result<()> {
    let code = "
    x(): void {
      return;
    }
    ";
    check_code(code)
}

#[test]
fn test_check_local_variable_define() -> Result<()> {
    let code = "
    x(): int {
      y: int = 1;
      return y;
    }
    ";
    check_code(code)
}

#[test]
fn static_method_check() -> Result<()> {
    let code = "
    class Foo {
      ::new(): Foo = Foo {};
    }
    ";
    check_code(code)
}

#[test]
fn method_check() -> Result<()> {
    let code = "
    class Foo {
      bar(): int = 1;
    }
    ";
    check_code(code)
}

#[test]
fn test_all_class_field_must_init() {
    let code = "
    class Foo {
      bar: int;
      ::new(): Foo = Foo {};
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_cannot_use_class_construction_on_non_class_type() {
    let code = "
    class Foo {
      ::new(): Foo = int {};
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_cannot_use_class_construction_out_of_class() {
    let code = "
    class Foo {}
    x: Foo = Foo {};
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_global_should_be_able_to_use_class_static_method() -> Result<()> {
    let code = "
    main(): void {
      foo: Foo = Foo::new();
    }

    class Foo {
      ::new(): Foo = Foo {};
    }
    ";
    check_code(code)
}

#[test]
fn test_static_method_should_not_available_in_class_scope() {
    let code = "
    class Foo {
      ::new2(): void {new1();}
      ::new1(): void {}
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn method() -> Result<()> {
    let code = "
    main(): void {
      foo: Foo = Foo::new();
      foo.bar();
    }

    class Foo {
      ::new(): Foo = Foo {};
      bar(): void {}
    }
    ";
    check_code(code)
}

#[test]
fn method_should_be_able_to_call_with_instance() -> Result<()> {
    let code = "
    class Foo {
      ::new(): Foo {
        f: Foo = Foo {};
        f.bar();
        return f;
      }
      bar(): void {}
    }
    ";
    check_code(code)
}

#[test]
fn non_extern_function_cannot_without_body() {
    let code = "
    foo(): void;
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn only_trait_can_be_super_type() {
    let code = "
    class Bar {}
    class Foo <: Bar {}
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn if_else_block_must_return_same_type_as_return_type() {
    let code = "
    foo(): int {
      if true {
        return 1;
      } else {
        return false;
      }
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn condition_must_be_a_bool() {
    let code = "
    foo(): void {
      if 1 {}
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn missing_return_statement_is_invalid() {
    let code = "
    foo(): int {
      if true {
        return 1;
      } else {
      }
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn dead_code_after_return_statement_is_invalid() {
    let code = "
    foo(): void {
      return;
      x: int = 1;
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn init_expression_must_has_same_type_as_define() {
    let code = "
    class Foo {
      flag: bool = 1;
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn redefine_field_is_invalid() {
    let code = "
    class Foo {
      a: int;
      a: int;
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn no_type_name() {
    let code = "x: Foo = Foo {};";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn no_member_name() {
    let code = "
    class Foo {
      ::new(): Foo = Foo {};
    }
    main(): void {
      foo: Foo = Foo::new();
      foo.a();
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

// helpers, must put tests before this line
fn check_code(code: &'static str) -> Result<()> {
    let mut parser = Parser::new("", code);
    let mut code = parser
        .parse_top_list(TkType::EOF)
        .map_err(|err| {
            println!("{}", err);
            err
        })
        .unwrap();

    code.push(TopAst::Import(Import {
        location: Location::none(),
        import_path: "prelude".to_string(),
        imported_component: vec![
            "int".to_string(),
            "void".to_string(),
            "f64".to_string(),
            "bool".to_string(),
            "string".to_string(),
            "List".to_string(),
            "println".to_string(),
        ],
    }));

    let prelude = parse_prelude();
    let mut checker = SemanticChecker::new();
    checker
        .check_program(&vec![
            prelude,
            Module {
                name: "test".to_string(),
                top_list: code,
            },
        ])
        .map_err(|err| {
            // map origin error and report at here
            println!("{}", err);
            err
        })
}
