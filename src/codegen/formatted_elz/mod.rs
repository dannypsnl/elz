pub fn formatted_elz(code: String) -> String {
    let mut indent = 0i32;
    let mut comment_line = false;
    let mut semicolon_symbol = false;
    let mut past_symbol = false;
    let mut s = String::from("");
    let clear_multiple_blank = code.replace("  ", "").replace("\n", "\\n");
    let code_to_char: Vec<&str> = clear_multiple_blank.split("").collect();

    let mut i = 0;
    while i < code_to_char.len() {
        if comment_line {
            if code_to_char[i] == "\\" && code_to_char[i + 1] == "n" {
                s.push_str("\n");
                s.push_str(add_indent(indent).as_str());
                comment_line = false;
                i += 1;
            } else {
                s.push_str(code_to_char[i]);
            }
        } else {
            match code_to_char[i] {
                //{ } ; need to add newline
                //add double blank
                "{" => {
                    indent += 1;
                    if code_to_char[i - 1] != " " && !past_symbol {
                        s.push_str(" ");
                    }

                    if code_to_char[i + 1] == "}" || past_symbol {
                        s.push_str("{");
                        past_symbol = false;
                    } else {
                        s.push_str("{");
                        s.push_str("\n");
                        s.push_str(add_indent(indent).as_str());
                    }
                }
                //delete double blank
                "}" => {
                    indent -= 1;
                    if code_to_char[i + 1] == ";" {
                        s.push_str("}");
                    } else if code_to_char[i - 1] == "{" {
                        s.push_str("}");
                        s.push_str("\n");
                    } else {
                        s.push_str(add_indent(indent).as_str());
                        s.push_str("}");
                        s.push_str("\n");
                    }
                }
                ";" => {
                    s.push_str(";");
                    s.push_str("\n");
                    if code_to_char[i + 1] == "\\"
                        && code_to_char[i + 2] == "n"
                        && code_to_char[i + 3] == "}"
                    {
                    } else if code_to_char[i + 1] == "}" {
                    } else {
                        s.push_str(add_indent(indent).as_str());
                    }
                    semicolon_symbol = true;
                    past_symbol = false;
                }

                //add blank in the next position and check double symbol
                ":" => {
                    if code_to_char[i + 1] == ":" {
                        s.push_str(":");
                        s.push_str(":");
                        i += 1;
                    } else {
                        s.push_str(add_blank(":", " ", code_to_char[i + 1]).as_str());
                    }
                }

                //add blank in the next position
                "," => {
                    s.push_str(add_blank(",", " ", code_to_char[i + 1]).as_str());
                }

                //add blank in the pre and next position
                "+" => {
                    s.push_str(add_blank("+", code_to_char[i - 1], code_to_char[i + 1]).as_str());
                }
                "-" => {
                    s.push_str(add_blank("-", code_to_char[i - 1], code_to_char[i + 1]).as_str());
                }
                "*" => {
                    s.push_str(add_blank("*", code_to_char[i - 1], code_to_char[i + 1]).as_str());
                }
                "/" => {
                    if code_to_char[i + 1] == "/" {
                        s.push_str("//");
                        s.push_str(" ");
                        comment_line = true;
                        i += 1;
                    } else {
                        s.push_str(
                            add_blank("/", code_to_char[i - 1], code_to_char[i + 1]).as_str(),
                        );
                    }
                }
                "=" => {
                    s.push_str(add_blank("=", code_to_char[i - 1], code_to_char[i + 1]).as_str());
                    past_symbol = true;
                }
                ">" => {
                    s.push_str(add_blank(">", code_to_char[i - 1], code_to_char[i + 1]).as_str());
                }
                "<" => {
                    if code_to_char[i + 1] == ":" {
                        s.push_str(add_blank("<", code_to_char[i - 1], " ").as_str());
                    } else {
                        s.push_str(
                            add_blank("<", code_to_char[i - 1], code_to_char[i + 1]).as_str(),
                        );
                    }
                }
                "\\" => {
                    if code_to_char[i + 1] == "n" {
                        i += 1;
                    }
                }

                _ => {
                    if semicolon_symbol {
                        if code_to_char[i] != " " {
                            s.push_str(code_to_char[i]);
                            semicolon_symbol = false;
                        }
                    } else {
                        if code_to_char[i] == " " && check_symbol_behind_space(code_to_char[i + 1])
                        {
                        } else {
                            s.push_str(code_to_char[i]);
                        }
                    }
                }
            }
        }
        i += 1;
    }
    s
}

fn add_indent(tab: i32) -> String {
    let mut count = 0i32;
    let mut s = String::from("");
    loop {
        count += 1;
        if count > tab {
            break;
        }
        s.push_str("  ");
    }
    s
}

fn add_blank(c: &str, pre: &str, next: &str) -> String {
    let mut s = String::from("");
    if pre != " " {
        s.push_str(" ");
    }
    s.push_str(c);
    if next != " " {
        s.push_str(" ");
    }
    s
}

fn check_symbol_behind_space(symbol: &str) -> bool {
    if symbol == ";" {
        return true;
    } else if symbol == ":" {
        return true;
    } else if symbol == "(" {
        return true;
    } else if symbol == "[" {
        return true;
    } else {
        return false;
    }
}

#[cfg(test)]
mod tests;
