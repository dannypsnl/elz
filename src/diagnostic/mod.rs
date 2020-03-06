use crate::lexer::Location;
use codespan::Files;
use codespan_reporting::diagnostic::{Diagnostic, Label};
use codespan_reporting::term::emit;
use codespan_reporting::term::termcolor::{ColorChoice, StandardStream};

pub struct Reporter {
    files: Files<String>,
}

impl Reporter {
    pub fn new() -> Reporter {
        Reporter {
            files: Files::new(),
        }
    }

    pub(crate) fn for_file<T: Into<String>>(&mut self, file_name: T, source: T) -> FileID {
        FileID {
            value: self.files.add(file_name.into(), source.into()),
            diagnostics: vec![],
        }
    }
}

#[derive(Clone)]
pub(crate) struct FileID {
    value: codespan::FileId,
    diagnostics: Vec<Diagnostic>,
}

impl FileID {
    pub(crate) fn add_diagnostic(
        &mut self,
        location: Location,
        long_message: String,
        message: String,
    ) {
        self.diagnostics.push(Diagnostic::new_error(
            long_message,
            Label::new(self.value, location.start..location.end, message),
        ));
    }
    pub(crate) fn report(&self, reporter: &Reporter) {
        let writer = StandardStream::stderr(ColorChoice::Auto);
        let config = codespan_reporting::term::Config::default();
        for diagnostic in &self.diagnostics {
            emit(&mut writer.lock(), &config, &reporter.files, &diagnostic).unwrap();
        }
    }
}
