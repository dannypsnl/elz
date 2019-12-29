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
                        .help("input directory to format")
                        .required(true)
                        .max_values(1),
                ),
        )
        .get_matches();

    let result = if let Some(compile_args) = matches.subcommand_matches(cmd::compile::CMD_NAME) {
        let files: Vec<_> = compile_args.values_of("INPUT").unwrap().collect();
        cmd::compile::compile(files)
    } else if let Some(compile_args) = matches.subcommand_matches(cmd::fmt::CMD_NAME) {
        let files: Vec<_> = compile_args.values_of("INPUT").unwrap().collect();
        cmd::fmt::format(files)
    } else {
        Ok(())
    };
    match result {
        Ok(()) => (),
        Err(err) => println!("{}", err),
    }
}
