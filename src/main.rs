use clap::{App, Arg, SubCommand};
use elz;
use elz::mir;
use elz::mir::MIRError;
use elz::parser::Parser;
use elz::semantic;
use quick_protobuf::serialize_into_vec;
use std::ffi::CString;
use std::fs;
use std::os::raw::c_char;

#[link(name = "code_generate")]
extern "C" {
    fn generate(s: *const c_char);
}

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
                        .index(1),
                ),
        )
        .get_matches();

    if let Some(compile) = matches.subcommand_matches("compile") {
        let compile_file = compile.value_of("INPUT").unwrap();
        let code = fs::read_to_string(compile_file).expect("failed at read content of input file");

        let mut parser = Parser::new(code);
        // run parser
        let program = parser.parse_program()?;
        // type inference and check
        semantic::check_program(&program)?;
        // generate MIR
        let mir = mir::generate_mir_program(&program);
        match mir {
            Err(err) => {
                if err == MIRError::NoMain {
                    // if no main then we only check the program
                    Ok(())
                } else {
                    Err(Box::new(err))
                }
            }
            Ok(mir_program) => {
                let out_vec = serialize_into_vec(&mir_program).expect("Cannot write MIR!");
                let c_str = CString::new(out_vec)?;
                unsafe {
                    generate(c_str.as_ptr());
                }
                Ok(())
            }
        }
    } else {
        Ok(())
    }
}
