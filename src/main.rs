use clap::{App, Arg, SubCommand};
use elz;
use elz::mir;
use elz::parser::Parser;
use elz::semantic;
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
        let compile_files: Vec<_> = compile.values_of("INPUT").unwrap().collect();

        let program = compile_files
            .into_iter()
            .flat_map(|source_code| {
                let code = fs::read_to_string(source_code).expect(
                    format!("failed at read content of input file: {}", source_code).as_str(),
                );

                Parser::parse_program(code)
                    .expect(format!("failed at compile file: {}", source_code).as_str())
            })
            .collect();

        let remapped = semantic::helper::flat_package("", &program);

        // type inference and check
        semantic::check_program(&remapped)?;

        if let Some(_) = remapped.get("main") {
            // generate MIR when main binding exist
            let mir = mir::generate_mir_program(remapped).expect("MIR: {}");
            let g = elz::codegenerate::Generator::new(mir);
            g.generate();
            g.binary();
        }
    }
    Ok(())
}
