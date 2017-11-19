package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/elz-lang/elz/parser"
	"os"
)

func main() {
	// TODO: Use CharStream & a function wrap this all, here should handle other mession.
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(NewElzListener(), tree)
}
