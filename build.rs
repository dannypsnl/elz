fn main() {
    println!("cargo:rustc-link-lib=static=code_generate");
    println!("cargo:rustc-link-search=native=.");
}
