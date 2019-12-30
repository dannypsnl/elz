use crate::ast::format::{FormatTopAstList, FormattedElz, DEFAULT_LEVEL};
use crate::parser::Parser;
use std::io::Write;
use walkdir::{DirEntry, WalkDir};

pub const CMD_NAME: &'static str = "fmt";

pub fn format(directory: Vec<&str>) -> Result<(), Box<dyn std::error::Error>> {
    let dir = directory[0];
    for entry in WalkDir::new(dir).into_iter().filter_map(|e| e.ok()) {
        handle_entry(entry)?;
    }
    Ok(())
}

fn handle_entry(entry: DirEntry) -> Result<(), Box<dyn std::error::Error>> {
    let file_name = entry.file_name().to_string_lossy().to_string();
    if file_name.ends_with(".elz") {
        let code = std::fs::read_to_string(entry.path())?;
        let program = Parser::parse_program(entry.path().to_string_lossy().to_string(), code)?;
        let result = FormatTopAstList(program).formatted_elz(DEFAULT_LEVEL);
        let mut origin_file = std::fs::File::create(entry.path())?;
        origin_file.write_all(result.as_bytes())?;
        origin_file.write_all("\n".as_bytes())?;
    }
    Ok(())
}
