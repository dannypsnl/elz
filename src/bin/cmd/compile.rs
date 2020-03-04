use elz::ast::TopAst;
use elz::codegen::llvm::LLVMValue;
use elz::codegen::CodeGenerator;
use elz::parser::{parse_prelude, Parser};
use elz::semantic::SemanticChecker;

pub const CMD_NAME: &'static str = "compile";

pub fn compile(files: Vec<&str>) -> Result<(), Box<dyn std::error::Error>> {
    let program = check(files)?;
    let code_generator = CodeGenerator::new();
    let module = code_generator.generate_module(&program);
    println!("{}", module.llvm_represent());
    Ok(())
}

fn check(files: Vec<&str>) -> Result<Vec<TopAst>, Box<dyn std::error::Error>> {
    // FIXME: for now to make code simple we only handle the first input file.
    let code = std::fs::read_to_string(files[0])?;
    let mut program = Parser::parse_program(files[0], &code)?;
    let mut prelude = parse_prelude();
    prelude.append(&mut program);
    // check program
    let mut semantic_checker = SemanticChecker::new();
    semantic_checker.check_program(&prelude)?;
    Ok(program)
}
