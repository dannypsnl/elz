package parser

import (
	"fmt"
	"github.com/elz-lang/elz/lexer"
)

type Ast interface{}

type ErrorTree struct {
	errorMsg string
}

type DefineVar struct {
	immutable  bool
	ident      string
	expression Expr
}

type Expr struct {
}

func Parse(filename, source_code string) []Ast {
	return startParse(lexer.Lex(filename, source_code), make([]Ast, 0), ' ')
}

func startParse(l *lexer.Lexer, tree []Ast, lookingFor rune) {
	for token := l.NextItem(); token.Type != lexer.ItemEOF; {
		switch t := token.Type; t {
		// First, we handle keyword[let], it's a statement call DefineVar, so we will get a Ast[DefineVar].
		case lexer.ItemKwLet:
			tree = append(tree, parseDefineVar(l))
		// If meet keyword[fn], it's a statement call DefineFn, so we will get a Ast[DefineFn]
		case lexer.ItemKwFn:
			tree = append(tree, parseDefineFn(l))
		// If we meet a error item, that mean we didn't receive this token, it could be a bug that we didn't consider. Please report.
		case lexer.ItemError:
			fmt.Println(token.Val)
		default:
			panic("Seems a bug!")
		}
	}
}

func parseDefineVar(l *lexer.Lexer) *DefineVar {
	tk := l.NextItem()
	switch t := tk.Type; t {
	case lexer.ItemKwMut:
	case lexer.ItemIdent:
	default:
		return ErrorTree{"Keyword let expected mut or identifier."}
	}
}
