use super::*;

#[test]
fn test_codegen_main() {
    let code = "\
        main(): void {}
        x: int = 1;";
    let module = gen_code(code);
    assert_eq!(
        module.llvm_represent(),
        "\
         @x = global i64 1\n\
         define void @main() {\n  ret void\n\
         }\n"
    );
}

#[test]
fn test_return_value() {
    let code = "foo(): int = 1;";
    let module = gen_code(code);
    assert_eq!(
        module.llvm_represent(),
        "\
         define i64 @foo() {\n  ret i64 1\n\
         }\n"
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

// helpers, must put tests before this line
fn gen_code(code: &'static str) -> Module {
    let program = crate::parser::Parser::parse_program("", code).unwrap();
    let backend = Backend::LLVM;
    backend.generate_module(&program)
}
