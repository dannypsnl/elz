use super::*;
use crate::lexer::TkType::EOF;
use llvm::LLVMValue;

#[test]
fn test_codegen_main() {
    let code = "main(): void {}";
    let module = gen_code(code);
    assert_eq!(
        module.functions.get("@main").unwrap().llvm_represent(),
        "define void @main() {
  ret void
}"
    );
}

#[test]
fn global_variable() {
    let code = "x: int = 1;";
    let module = gen_code(code);
    assert_eq!(module.variables[0].llvm_represent(), "@x = global i64 1");
}

#[test]
fn test_return_value() {
    let code = "foo(): int = 1;";
    let module = gen_code(code);
    assert_eq!(
        module.functions.get("@foo").unwrap().llvm_represent(),
        "define i64 @foo() {
  ret i64 1
}"
    );
}

#[test]
fn test_function_declaration_with_parameter() {
    let code = "add(x: int, y: int): int;";
    let module = gen_code(code);
    assert_eq!(
        module.functions.get("@add").unwrap().llvm_represent(),
        "declare i64 @add(i64 %x, i64 %y)"
    )
}

#[test]
fn test_function_define_with_parameter() {
    let code = "const(x: int): int = 1;";
    let module = gen_code(code);
    assert_eq!(
        module.functions.get("@const").unwrap().llvm_represent(),
        "define i64 @const(i64 %x) {
  ret i64 1
}"
    )
}

#[test]
fn test_function_call() {
    let code = "
    main(): void {
      foo(1);
    }
    foo(x: int): void {}";
    let module = gen_code(code);
    assert_eq!(
        module.functions.get("@main").unwrap().llvm_represent(),
        "define void @main() {
  call void @foo(i64 1)
  ret void
}"
    )
}

#[test]
fn binary_expr() {
    let code = "
    foo(): int = 1 + 2;
    ";
    let module = gen_code(code);
    assert_eq!(
        module.functions.get("@foo").unwrap().llvm_represent(),
        "define i64 @foo() {
  %1 = add i64 1, 2
  ret i64 %1
}"
    )
}

#[test]
fn test_class_define() {
    let code = "
    class Foo {
      x: int;
      ::new(): Foo;
      bar(): void {}
    }";
    let module = gen_code(code);
    assert_eq!(
        module.types.get("Foo").unwrap().llvm_def(),
        "%Foo = type { i64 }"
    );
    assert_eq!(
        module
            .functions
            .get("@\"Foo::new\"")
            .unwrap()
            .llvm_represent(),
        "declare %Foo* @\"Foo::new\"()"
    );
    assert_eq!(
        module
            .functions
            .get("@\"Foo::bar\"")
            .unwrap()
            .llvm_represent(),
        "define void @\"Foo::bar\"(%Foo* %self) {
  ret void
}"
    );
}

#[test]
fn llvm_if_else() {
    let code = "
    foo(): void {
      if true {
        return;
      } else if true {
      } else {
      }
    }
    ";
    let module = gen_code(code);
    assert_eq!(
        module.functions.get("@foo").unwrap().llvm_represent(),
        "define void @foo() {
  br i1 true, label %1, label %2
; <label>:1:
  ret void
; <label>:2:
  br i1 true, label %3, label %4
; <label>:3:
  br label %5
; <label>:4:
  br label %5
; <label>:5:
  ret void
}"
    )
}

// helpers, must put tests before this line
fn gen_code(code: &'static str) -> ir::Module {
    let mut parser = crate::parser::Parser::new("", code);
    let mut program = parser
        .parse_top_list(EOF)
        .map_err(|err| {
            panic!("{}", err);
        })
        .unwrap();
    let mut prelude = crate::parser::parse_prelude();
    prelude.top_list.append(&mut program);
    let code_generator = CodeGenerator::new();
    code_generator.generate_module(&prelude.top_list)
}
