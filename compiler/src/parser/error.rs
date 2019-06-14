pub type Result<T> = std::result::Result<T, ParseError>;

#[derive(Debug)]
pub struct ParseError {
	descript: String,
}

impl ParseError {
	pub fn new(descript: String) -> ParseError {
		ParseError { descript: descript }
	}
}

impl std::fmt::Display for ParseError {
	fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
		write!(f, "ParseError: {}", self.descript)
	}
}
impl std::error::Error for ParseError {
	fn description(&self) -> &str {
		"parse error"
	}
}
