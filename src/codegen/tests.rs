use super::*;
use llvm::LLVMValue;

#[test]
fn test_codegen_main() {
    let code = "\
        main(): void {}
        x: int = 1;";
    let module = gen_code(code);
    assert_eq!(
        module.llvm_represent(),
        "@x = global i64 1
define void @main() {
  ret void
}
"
    );
}

#[test]
fn test_return_value() {
    let code = "foo(): int = 1;";
    let module = gen_code(code);
    assert_eq!(
        module.llvm_represent(),
        "define i64 @foo() {
  ret i64 1
}
"
    );
}

#[test]
fn test_function_declaration_with_parameter() {
    let code = "add(x: int, y: int): int;";
    let module = gen_code(code);
    assert_eq!(
        module.llvm_represent(),
        "declare i64 @add(i64 %x, i64 %y)\n"
    )
}

#[test]
fn test_function_define_with_parameter() {
    let code = "const(x: int): int = 1;";
    let module = gen_code(code);
    assert_eq!(
        module.llvm_represent(),
        "define i64 @const(i64 %x) {
  ret i64 1
}
"
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
        module.llvm_represent(),
        "define void @main() {
  call void @foo(i64 1)
  ret void
}
define void @foo(i64 %x) {
  ret void
}
"
    )
}

#[test]
fn binary_expr() {
    let code = "
    foo(): int = 1 + 2;
    ";
    let module = gen_code(code);
    assert_eq!(
        module.llvm_represent(),
        "define i64 @foo() {
  %1 = add i64 1, 2
  ret i64 %1
}
"
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
        module.llvm_represent(),
        "%Foo = type {i64}
declare %Foo* @\"Foo::new\"()
define void @\"Foo::bar\"(%Foo* %self) {
  ret void
}
"
    )
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
        module.llvm_represent(),
        "define void @foo() {
  br i1 true, label %label2, label %label3
label2:
  ret void
label3:
  br i1 true, label %label4, label %label5
label4:
  br label %label1
label5:
  br label %label1
label1:
  ret void
}
"
    )
}

// helpers, must put tests before this line
fn gen_code(code: &'static str) -> ir::Module {
    let mut program = crate::parser::Parser::parse_program("", code)
        .map_err(|err| {
            panic!("{}", err);
        })
        .unwrap();
    let mut prelude = crate::parser::parse_prelude();
    prelude.append(&mut program);
    let code_generator = CodeGenerator::new();
    code_generator.generate_module(&prelude)
}
