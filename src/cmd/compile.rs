use crate::ast::TopAst;
use crate::codegen::llvm::LLVMValue;
use crate::codegen::CodeGenerator;
use crate::diagnostic::Reporter;
use crate::parser::{parse_prelude, Parser};
use crate::semantic::SemanticChecker;

pub const CMD_NAME: &'static str = "compile";

pub fn compile(files: Vec<&str>) -> Result<(), Box<dyn std::error::Error>> {
    let mut reporter = Reporter::new();
    let program = check(&mut reporter, files)?;
    let code_generator = CodeGenerator::new();
    let module = code_generator.generate_module(&program);
    println!("{}", module.llvm_represent());
    Ok(())
}

fn check(
    reporter: &mut Reporter,
    files: Vec<&str>,
) -> Result<Vec<TopAst>, Box<dyn std::error::Error>> {
    // FIXME: for now to make code simple we only handle the first input file.
    let code = std::fs::read_to_string(files[0])?;
    let mut file_reporter = reporter.for_file(files[0], &code);
    let mut program = match Parser::parse_program(files[0], &code) {
        Ok(p) => p,
        Err(err) => {
            file_reporter.add_diagnostic(err.location(), format!("{}", err), err.message());
            file_reporter.report(reporter);
            return Err(err.into());
        }
    };
    let mut prelude = parse_prelude();
    prelude.append(&mut program);
    // check program
    let mut semantic_checker = SemanticChecker::new();
    match semantic_checker.check_program(&prelude) {
        Ok(..) => Ok(prelude),
        Err(err) => {
            file_reporter.add_diagnostic(err.location(), format!("{}", err), err.message());
            file_reporter.report(reporter);
            Err(err.into())
        }
    }
}
