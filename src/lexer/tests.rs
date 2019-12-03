use self::TkType::*;
use super::*;

#[test]
fn test_unicode_identifier_a() {
    let code = "測試: int = 1";

    assert_eq!(
        lex("", code),
        vec![
            Token(Location::from(1, 0), Ident, "測試".to_string()),
            Token(Location::from(1, 2), Colon, ":".to_string()),
            Token(Location::from(1, 4), Ident, "int".to_string()),
            Token(Location::from(1, 8), Equal, "=".to_string()),
            Token(Location::from(1, 10), Integer, "1".to_string()),
            Token(Location::from(1, 11), EOF, "".to_string()),
        ]
    );
}

#[test]
fn get_number_tokens() {
    let ts = lex("", "10 30");
    assert_eq!(
        ts,
        vec![
            Token(Location::from(1, 0), Integer, "10".to_string()),
            Token(Location::from(1, 3), Integer, "30".to_string()),
            Token(Location::from(1, 5), EOF, "".to_string()),
        ]
    );
}

#[test]
fn get_ident_tokens() {
    let ts = lex("", " abc6");
    assert_eq!(
        ts,
        vec![
            Token(Location::from(1, 1), Ident, "abc6".to_string()),
            Token(Location::from(1, 5), EOF, "".to_string()),
        ]
    )
}

#[test]
fn get_escape_char_in_string() {
    let ts = lex("", "\"\\\"\"");
    assert_eq!(
        ts,
        vec![
            Token(Location::from(1, 0), String, "\"\\\"\"".to_string()),
            Token(Location::from(1, 4), EOF, "".to_string()),
        ]
    )
}

#[test]
fn comment_would_be_discard() {
    let ts = lex("", "//\n1");
    assert_eq!(
        ts,
        vec![
            Token(Location::from(2, 0), Integer, "1".to_string()),
            Token(Location::from(2, 1), EOF, "".to_string()),
        ]
    )
}
