grammar Elz;

WS: [ \t\r\n]+ -> channel(HIDDEN);
COMMENT: '//' .*? '\n' -> channel(HIDDEN);

ID : StartLetter Letter*;
fragment
StartLetter: [a-zA-Z_]
    ;
fragment
Letter: [0-9]
    | StartLetter
    ;

NUM: StartDigit Digit*;
fragment
StartDigit: [0-9.];
fragment
Digit: [0-9];

STRING: '"' .*? '"';

prog: topStatList?;

topStatList: topStat+;

topStat: fnDefine
    | varDefine
    | typeDefine
    | traitDefine
    | importStat
    ;
importStat: 'import' ID;

statList: stat+;
stat: varDefine
    | assign
    | exprStat
    ;

exprStat: matchRule
    | fnCall
    ;

matchRule:
    'match' expr '{'
        expr '=>' stat
        (',' expr '=>' stat)*
        ','?
    '}'
    ;

assign:
    ID '=' expr
    ;

exprList: expr (',' expr)*;
fnCall:
    ID '(' exprList? ')'
    ;

typePass : ID;
exportor: '+';
define: exportor? ID (':' typePass)? '=' expr;
varDefine:
    'let' mut='mut'? define (',' define)*
    ;
paramList: param (',' param)*;
param: ID ':' typePass;
fnDefine:
    'fn' exportor? ID '(' paramList? ')' ('->' typePass)? '{'
        statList?
    '}'
    ;
attrList: attr+;
attr: ID ':' typePass;
typeDefine:
    'type' exportor? ID '(' attrList ')'
    ;
tmethodList: tmethod+;
tmethod: exportor? ID '(' paramList? ')' ('->' typePass)?;
traitDefine:
    'trait' exportor ID '{'
        tmethodList?
    '}'
    ;

expr: '(' expr ')'
    | expr op='^' expr
    | expr op=('*'|'/') expr
    | expr op=('+'|'-') expr
    | factor
    ;
factor: exprStat
    | NUM
    | ID
    | STRING
    ;
