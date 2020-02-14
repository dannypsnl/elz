use super::*;
use crate::parser::{parse_prelude, Parser};

#[test]
fn test_redefine_variable_would_get_error() {
    let code = "\
    x: int = 1;
    x: int = 2;
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_function_and_variable_use_the_same_space() {
    let code = "\
    x: int = 1;
    x(): void {}
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_type_mismatched() {
    let code = "\
    x: int = \"str\";
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn call_non_void_function_as_statement() {
    let code = "\
    main(): void {
      foo();
    }
    foo(): int = 1;
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn call_void_function_as_statement() {
    let code = "\
    main(): void {
      foo();
    }
    foo(): void {}
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_check_function_call() {
    let code = "\
    x(a: int): int = a;
    y: int = x(2);
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_unify_list_type() {
    let code = "\
    x: List[int] = [1, 2, 3];
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_unify_free_var() {
    let code = "\
    x: List[int] = [];
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_check_return_nothing() {
    let code = "\
    x(): void {
      return;
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_check_local_variable_define() {
    let code = "\
    x(): int {
      y: int = 1;
      return y;
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_static_method_check() {
    let code = "\
    class Foo {
      ::new(): Foo = Foo {};
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_method_check() {
    let code = "\
    class Foo {
      bar(): int = 1;
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_all_class_field_must_init() {
    let code = "\
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
    let code = "\
    class Foo {
      ::new(): Foo = int {};
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_cannot_use_class_construction_out_of_class() {
    let code = "\
    class Foo {}
    x: Foo = Foo {};
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_global_should_be_able_to_use_class_static_method() {
    let code = "\
    main(): void {
      foo: Foo = Foo::new();
    }

    class Foo {
      ::new(): Foo = Foo {};
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_static_method_should_not_available_in_class_scope() {
    let code = "\
    class Foo {
      ::new2(): void {new1();}
      ::new1(): void {}
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_err(), true);
}

#[test]
fn test_method_should_get_transform() {
    let code = "\
    main(): void {
      foo: Foo = Foo::new();
      foo.bar();
    }

    class Foo {
      ::new(): Foo = Foo {};
      bar(): void;
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
}

#[test]
fn test_method_should_be_able_to_call_with_instance() {
    let code = "\
    class Foo {
      ::new(): Foo {
        f: Foo = Foo {};
        f.bar();
        return f;
      }
      bar(): void;
    }
    ";
    let result = check_code(code);
    assert_eq!(result.is_ok(), true);
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

// helpers, must put tests before this line
fn check_code(code: &'static str) -> Result<()> {
    let mut program = Parser::parse_program("", code).unwrap();
    let mut prelude = parse_prelude();
    prelude.append(&mut program);
    let mut checker = SemanticChecker::new();
    checker.check_program(&prelude).map_err(|err| {
        // map origin error and report at here
        println!("{}", err);
        err
    })
}
