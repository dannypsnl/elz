use elz::codegen::formatted_elz::{FormatTopAstList, FormattedElz, DEFAULT_LEVEL};
use elz::parser::Parser;
use std::ffi::OsStr;
use std::io::Write;
use std::path::Path;
use walkdir::WalkDir;

pub const CMD_NAME: &'static str = "fmt";

pub fn format(paths: Vec<&str>) -> Result<(), Box<dyn std::error::Error>> {
    for path_str in paths {
        let path = Path::new(path_str);
        path.extension().map_or_else(
            || handle_dir(path),
            |extension| {
                if extension == "elz" {
                    handle_entry(path)
                } else {
                    handle_dir(path)
                }
            },
        )?;
    }
    Ok(())
}

fn handle_dir(dir: &Path) -> Result<(), Box<dyn std::error::Error>> {
    for entry in WalkDir::new(dir).into_iter().filter_map(|e| e.ok()) {
        handle_entry(entry.path())?;
    }
    Ok(())
}

fn handle_entry(file_path: &Path) -> Result<(), Box<dyn std::error::Error>> {
    if file_path.extension() == Some(OsStr::new("elz")) {
        let code = std::fs::read_to_string(file_path)?;
        let program = Parser::parse_program(file_path.to_string_lossy().to_string(), code)?;
        let result = FormatTopAstList(program).formatted_elz(DEFAULT_LEVEL);
        let mut origin_file = std::fs::File::create(file_path)?;
        origin_file.write_all(result.as_bytes())?;
        origin_file.write_all("\n".as_bytes())?;
    }
    Ok(())
}
