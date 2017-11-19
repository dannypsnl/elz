package main

import (
	_ "fmt"
	"testing"

	"github.com/elz-lang/elz/ast"
)

func TestParser(t *testing.T) {
	parser := Parse("main.elz", `ident
	let +about:num = 1
	`)
	asts := parser.parseProgram()
	if (asts[0] != ast.Error{
		"At(0), token 'ident' is not allow at top level\n",
	}) {
		t.Error(`error`)
	}

	vd := &ast.VarDefination{
		Immutable:  true,
		Export:     true,
		Name:       "about",
		VarType:    "num",
		Expression: &ast.Number{"1"},
	}
	if asts[1].(*ast.VarDefination).Immutable != vd.Immutable {
		t.Error(`error`)
	}
}
