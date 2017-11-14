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
		case lexer.ItemKwLet:
			tree = append(tree, parseDefineVar(l))
		case lexer.ItemError:
			fmt.Println(token.Val)
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
