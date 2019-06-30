use elz;

use std::fs;
use clap::{App, SubCommand, Arg};
use elz::parser::Parser;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let matches = App::new("elz")
        .author("Danny Lin <dannypsnl@gmail.com>")
        .subcommand(SubCommand::with_name("compile")
            .about("compile input file")
            .arg(
                Arg::with_name("INPUT")
                    .help("input file to compile")
                    .required(true)
                    .index(1)
            )
        )
        .get_matches();

    if let Some(compile) = matches.subcommand_matches("compile") {
        let compile_file = compile.value_of("INPUT").unwrap();
        let code = fs::read_to_string(compile_file).expect("failed at read content of input file");

        let mut parser =
            Parser::new(code);

        let program = parser.parse_program()?;
        println!("{:?}", program);
    }

    Ok(())
}
