package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/elz-lang/elz/parser"
	"os"
)

func main() {
	fmt.Println("vim-go")
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(NewElzListener(), tree)
}
