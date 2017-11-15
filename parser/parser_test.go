package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	parser := Parse("main.elz", `ident`)
	asts := parser.parseProgram()
	if (asts[0] != Error{
		"Pos 0, token 'ident' is not allow at top level\n",
	}) {
		t.Error(`error`)
	}
}
