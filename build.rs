fn main() {
    println!("cargo:rustc-link-lib=dylib=code_generate");
    println!("cargo:rustc-link-search=.");
}
