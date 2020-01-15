use rust_embed::RustEmbed;

#[derive(RustEmbed)]
#[folder = "lib/prelude/"]
pub struct Asset;
