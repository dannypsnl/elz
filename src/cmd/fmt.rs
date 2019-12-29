use crate::parser::Parser;
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
        let _program = Parser::parse_program(entry.path().to_string_lossy().to_string(), code)?;
        // TODO: implements program format function
        println!("{}", entry.path().display());
    }
    Ok(())
}
