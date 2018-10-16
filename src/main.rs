extern crate clap;
use clap::{App, Arg, SubCommand};
#[macro_use]
extern crate pest;
#[macro_use]
extern crate pest_derive;
extern crate inkwell;

use std::fs::File;

mod ast;
mod parser;
mod visit;

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
        let ast_tree = parser::parse_elz_program(file_name);
        let mut visitor = visit::Visitor::new();
        let module = visitor.visit_program(ast_tree);
        let mut output = File::create("test.bc").unwrap();
        module.write_bitcode_to_file(&output, true, true);
    }
}
