use clap::{App, Arg, SubCommand};
use elz::parser::Parser;
use elz::semantic::check_program;
use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
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
        let code = fs::read_to_string(files[0])?;
        let program = Parser::parse_program(code)?;
        // check program
        check_program(program)?;
    }
    Ok(())
}
