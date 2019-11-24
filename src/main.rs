use clap::{App, Arg, SubCommand};
use elz::ast::TopAst;
use elz::codegen::{Backend, CodeGenerator};
use elz::parser::Parser;
use elz::semantic::SemanticChecker;

fn main() {
    let matches = App::new("elz")
        .author("Danny Lin <dannypsnl@gmail.com>")
        .subcommand(
            SubCommand::with_name("compile")
                .about("compile input file")
                .arg(
                    Arg::with_name("INPUT")
                        .help("input file to compile")
                        .required(true)
                        .min_values(1),
                ),
        )
        .get_matches();

    if let Some(compile_args) = matches.subcommand_matches("compile") {
        let files: Vec<_> = compile_args.values_of("INPUT").unwrap().collect();

        match check(files) {
            Err(e) => println!("{}", e),
            Ok(program) => {
                let code_generator = CodeGenerator::with_backend(Backend::LLVM);
                code_generator.generate_module(&program);
            }
        }
    }
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
