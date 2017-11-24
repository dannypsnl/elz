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

topStat: fnDefine // fn foo( params ) { stats }
    | varDefine // let (mut) var: type = expr
    | typeDefine // type newType ( props )
    | implBlock // impl type { methods }
    | traitDefine // trait DB { methods }
    | importStat
    ;
importStat: 'import' ID;

statList: stat+;
stat: varDefine
    | loopStat
    | returnStat
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
typeList: typePass (',' typePass)*;

methodList: method+;
method:
    exportor? ID '(' paramList? ')' ('->' typePass)? '{'
        statList?
    '}'
    ;
implBlock:
    'impl' ID (':' typeList)? '{'
        methodList?
    '}'
    ;
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
tmethod: exportor? ID '(' typeList? ')' ('->' typePass)?;
traitDefine:
    'trait' exportor ID '{'
        tmethodList?
    '}'
    ;

expr: '(' expr ')' // FIXME: Waiting for implement
    | expr op='^' expr // FIXME: We had not support translate it.
    | expr op=('*'|'/') expr
    | expr op=('+'|'-') expr
    | factor
    ;
factor: exprStat // Important, exprStat have match & functionCall yet!
    | NUM
    | ID
    | STRING
    ;
