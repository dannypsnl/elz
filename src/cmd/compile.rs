use crate::ast::TopAst;
use crate::codegen::llvm::LLVMValue;
use crate::codegen::CodeGenerator;
use crate::parser::Parser;
use crate::semantic::SemanticChecker;

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
    let program = Parser::parse_program(files[0], &code)?;
    // check program
    let mut semantic_checker = SemanticChecker::new();
    semantic_checker.check_program(&program)?;
    Ok(program)
}
