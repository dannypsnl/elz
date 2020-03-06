use clap::{App, Arg, SubCommand};
use elz::cmd;

fn main() {
    let matches = App::new("elz")
        .author("Danny Lin <dannypsnl@gmail.com>")
        .subcommand(
            SubCommand::with_name(cmd::compile::CMD_NAME)
                .about("compile input file")
                .arg(
                    Arg::with_name("INPUT")
                        .help("input file to compile")
                        .required(true)
                        .min_values(1),
                ),
        )
        .subcommand(
            SubCommand::with_name(cmd::fmt::CMD_NAME)
                .about("format all files matched *.elz under the directory")
                .arg(
                    Arg::with_name("INPUT")
                        .help("input directories or files to format")
                        .required(true)
                        .min_values(1),
                ),
        )
        .get_matches();

    if let Some(compile_args) = matches.subcommand_matches(cmd::compile::CMD_NAME) {
        let files: Vec<_> = compile_args.values_of("INPUT").unwrap().collect();
        match cmd::compile::compile(files) {
            Ok(..) => (),
            Err(..) => println!("compile failed"),
        }
    } else if let Some(compile_args) = matches.subcommand_matches(cmd::fmt::CMD_NAME) {
        let files: Vec<_> = compile_args.values_of("INPUT").unwrap().collect();
        match cmd::fmt::format(files) {
            Ok(..) => (),
            Err(..) => println!("format failed"),
        }
    }
}
