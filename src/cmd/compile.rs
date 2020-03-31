use crate::ast::Module;
use crate::diagnostic::Reporter;
use crate::parser::{parse_prelude, Parser};
use crate::semantic::SemanticChecker;

pub const CMD_NAME: &'static str = "compile";

pub fn compile(files: Vec<&str>) -> Result<(), Box<dyn std::error::Error>> {
    let mut reporter = Reporter::new();
    // FIXME: comment out code generator for now to focus on semantic checking
    // let program =
    check(&mut reporter, files)?;
    // let code_generator = CodeGenerator::new();
    // let module = code_generator.generate_module(&program);
    // println!("{}", module.llvm_represent());
    Ok(())
}

fn check(
    reporter: &mut Reporter,
    files: Vec<&str>,
) -> Result<Vec<Module>, Box<dyn std::error::Error>> {
    // FIXME: for now to make code simple we only handle the first input file.
    let code = std::fs::read_to_string(files[0])?;
    let mut file_reporter = reporter.for_file(files[0], &code);
    let module = match Parser::parse_program(files[0], &code) {
        Ok(p) => p,
        Err(err) => {
            file_reporter.add_diagnostic(err.location(), format!("{}", err), err.message());
            file_reporter.report(reporter);
            return Err(err.into());
        }
    };
    let prelude = parse_prelude();
    let program = vec![module, prelude];
    // check program
    let mut semantic_checker = SemanticChecker::new();
    match semantic_checker.check_program(&program) {
        Ok(..) => Ok(program),
        Err(err) => {
            file_reporter.add_diagnostic(err.location(), format!("{}", err), err.message());
            file_reporter.report(reporter);
            Err(err.into())
        }
    }
}
