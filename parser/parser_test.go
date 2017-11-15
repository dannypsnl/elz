package parser

import (
	_ "fmt"
	"testing"

	"github.com/elz-lang/elz/lexer"
)

func TestParser(t *testing.T) {
	parser := Parse("main.elz", `ident
	let about:num = 1
	`)
	asts := parser.parseProgram()
	if (asts[0] != Error{
		"At(0), token 'ident' is not allow at top level\n",
	}) {
		t.Error(`error`)
	}

	vd := &VarDefination{
		immutable: true,
		export:    false,
		name:      "about",
		varType:   "num",
		expression: lexer.Item{
			lexer.ItemNumber,
			23,
			"1",
		},
	}
	if asts[1].(*VarDefination).immutable != vd.immutable {
		t.Error(`error`)
	}
}
