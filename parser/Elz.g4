grammar Elz;

options {
    language = Go; // target language
}

WS: [ \t\r\n]+ -> channel(HIDDEN);
COMMENT: '//' .*? '\n' -> channel(HIDDEN);

ID : StartLetter Letter*;
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

intSuffix: '\'i' ('8'|'16'|'32'|'64')
    | '\'f' ('32'|'64')
    ;
floatSuffix: '\'f' ('32'|'64');

// 1.34, .1, 0.4
FLOAT: Digit* '.' Digit+;
// 123, 1, 2, 54, 67, 98760, 0
INT: Digit+;
fragment
Digit: [0-9];

STRING: '"' .*? '"';

prog: topStatList?;

compilerNotation: '#' '[' ID ']'
    | '#' '[' ID '(' ID (',' ID)* ')' ']'
    ;

topStatList: (compilerNotation? topStat)+;

topStat: fnDefine     // fn foo( $params ) { $stat... }
    | declareFn       // fn printf(ref<i8>) -> i32
    | typeDefine      // type Writer ( prop... )
    | implBlock       // impl Stack { $method... }
    | traitDefine     // trait DB { method... }
    | importStat      // import ( Module... )
    | globalVarDef    // pi: num = $expr
    | stat            // Just for recognizing!
    ;

importMod: ID ('::' ID)*;
importStat: 'import' '(' importMod ')';

statList: stat+;
stat: localVarDef
    | loopStat // loop { stats }
    | returnStat // return expr
    | assign
    | exprStat
    ;

returnStat:
    'return' expr
    ;

loopStat:
    'loop' '{'
        statList?
    '}'
    ;

exprStat: matchRule
    | fnCall
    ;

// match i {
//   10 => { break },"
//   _ => { i = i + 1 }
// }
matchRule:
    'match' expr '{'
        expr '=>' stat
        (',' expr '=>' stat)*
        ','?
    '}'
    ;

// var = 1
assign:
    ID '=' expr
    ;

// 1, 2, 3, 4, 5
exprList: expr (',' expr)*;
// add(1, 2)
fnCall:
    ID '(' exprList? ')'
    ;

// mean a typeForm, but typeForm already be use by Go, so need an alternative name
typeForm : ID
    | ID '<' typeInfoList '>'
    | '...' typeForm
    | '[' typeForm ';' INT ']'
    ;
typeInfoList: (typeForm|expr) (',' (typeForm|expr))*;
typeList: typeForm (',' typeForm)*;

// @op
annotation: '@' ID ('(' expr ')')? ;

methodList: method+;
method:
    exportor? ID '(' paramList? ')' ('->' typeForm)? '{'
        statList?
    '}'
    ;
implBlock:
    'impl' ID (':' typeList)? '{'
        methodList?
    '}'
    ;

exportor: '+';
globalVarDef: exportor? define;

define: ID (':' typeForm)? '=' expr;
localVarDef:
    'let' mut='mut'? define (',' define)*
    ;

paramList: param (',' param)*;
paramType: ':' typeForm;
param: ID paramType?;
returnType: '->' typeForm;
fnDefine:
    // because fn also handle operator, so if we use exportor after keyword fn will cause we hard to divide ++ && + +
    exportor? 'fn' ID '(' paramList? ')' returnType? '{'
        statList?
    '}'
    ;

declareFn:
    'fn' ID '(' typeList? ')' returnType?
    ;

attrList: attr (',' attr)*;
attr: exportor? ID ':' typeForm;
typeDefine:
    'type' exportor? ID '(' attrList? ')'
    ;

tmethodList: tmethod+;
tmethod: ID '(' typeList? ')' ('->' typeForm)?;
traitDefine:
    exportor? 'trait' ID '{'
        tmethodList?
    '}'
    ;

// Explain for expr, because Antlr support the operation precedence by declared order
// So we don't have to consider that
expr: op='&' expr                      # Ref // operation prec
    | op='*' expr                      # DeRef
    | ID '[' INT ']'                   # AccessArrayElement
    | expr op='as' typeForm            # As
    | expr op='^' expr                 # Pow
    | expr op=('*'|'/') expr           # MulOrDiv
    | expr op=('+'|'-') expr           # AddOrSub
    | expr op=('<'|'>'|'<='|'>=') expr # Cmp
    | expr op='!=' expr                # NotEq
    | expr op='==' expr                # Eq
    | expr op=('&&'|'||') expr         # AndOrOr
    | expr '?' expr ':' expr           # ThreeOpCmp
    | '(' expr ')'                     # SubExpr
    | exprStat                         # StatExpr
    | '[' expr ';' INT ']'             # ArrWithLen
    | '[' exprList ']'                 # ArrWithList
    | INT intSuffix?                   # Int
    | FLOAT floatSuffix?               # Float
    | ID                               # Id
    | STRING                           # Str
    ;
