use clap::{App, Arg, SubCommand};
use elz::parser::Parser;
use elz::semantic::SemanticChecker;
use std::fs;

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

    if let Some(compile) = matches.subcommand_matches("compile") {
        let files: Vec<_> = compile.values_of("INPUT").unwrap().collect();

        // FIXME: for now to make code simple we only handle the first input file.
        let code = match fs::read_to_string(files[0]) {
            Ok(code) => code,
            Err(e) => {
                println!("{}", e);
                return;
            }
        };
        let program = match Parser::parse_program(files[0], &code) {
            Ok(p) => p,
            Err(e) => {
                println!("{}", e);
                return;
            }
        };
        // check program
        let mut semantic_checker = SemanticChecker::new();
        match semantic_checker.check_program(program) {
            Err(e) => println!("{}", e),
            _ => {
                // TODO: generate binary(or do nothing for library)
            }
        }
    }
}
