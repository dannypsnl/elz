grammar Elz;

options {
    language = Go; // target language
}

WS: [ \t\r\n]+ -> channel(HIDDEN);
COMMENT: '//' .*? '\n' -> channel(HIDDEN);

BOOLEAN: 'true' | 'false';
IDENT : StartLetter Letter*;
fragment
StartLetter: [a-zA-Z_]
    | '\u00C0'..'\u00D6'
    | '\u00D8'..'\u00F6'
    | '\u00F8'..'\u02FF'
    | '\u0370'..'\u037D'
    | '\u037F'..'\u1FFF'
    | '\u200C'..'\u200D'
    | '\u2070'..'\u218F'
    | '\u2C00'..'\u2FEF'
    | '\u3001'..'\uD7FF'
    | '\uF900'..'\uFDCF'
    | '\uFDF0'..'\uFFFD'
    ;
fragment
Letter: StartLetter
    | [0-9]
    | '\u00B7'
    | '\u0300'..'\u036F'
    | '\u203F'..'\u2040'
    ;

// 1.34, .1, 0.4
FLOAT: Digit* '.' Digit+;
// 123, 1, 2, 54, 67, 98760, 0
INT: Digit+;
fragment
Digit: [0-9];
// e.g. "hello", "say hello"
STRING: '"' .*? '"';

prog: topLevel* EOF;

topLevel: bind_type
    | binding
    ;
bind_type: IDENT '::' ;
binding:
    IDENT+ '=' pipe;

// print "hello" |> print "world"
// parseIdent "a = 1", [] |> parseBinding |> parseExpression
pipe: expr ('|>' expr)*;
funcCall: IDENT expr*;
expr:
    funcCall                     # FnCall
    | BOOLEAN                    # Boolean
    | STRING                     # String
    | FLOAT                      # Float
    | INT                        # Int
    | '(' expr ')'               # SubExpr
    | '[' expr? (',' expr)* ']'  # List
    | expr op=('*'|'/') expr     # MulDiv
    | expr op=('+'|'-') expr     # AddSub
    ;
