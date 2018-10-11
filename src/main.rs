extern crate clap;
use clap::{App, Arg, SubCommand};
#[macro_use]
extern crate pest;
#[macro_use]
extern crate pest_derive;

mod ast;
mod parser;

fn main() {
    let matches = App::new("elz")
        .version("0.1.0")
        .author("Danny Lin. <dannypsnl@gmail.com>")
        .about("elz compiler")
        .subcommand(
            SubCommand::with_name("compile")
                .about("compile single elz program")
                .arg(
                    Arg::with_name("INPUT")
                        .help("your elz program")
                        .required(true)
                        .index(1),
                ),
        ).get_matches();

    if let Some(matches) = matches.subcommand_matches("compile") {
        let file_name = matches.value_of("INPUT").unwrap(); // missing at is fine, just panic.

        // FIXME: return AST tree here
        parser::parse_elz_program(file_name);
    }
}
