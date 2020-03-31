use std::ffi::OsStr;
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
        // FIXME:
        //   wait #247
    }
    Ok(())
}
