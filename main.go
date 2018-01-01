package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/listener"
	"github.com/elz-lang/elz/parser"
)

func main() {
	// TODO: Use CharStream & a function wrap this all, here should handle other mession.
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	eal := listener.New()
	antlr.ParseTreeWalkerDefault.Walk(eal, tree)

	// Module is temprary variable
	ctx := ast.NewContext()
	for _, ast := range eal.AstList {
		ast.Codegen(ctx)
	}
	fmt.Println(ctx.Module)
}
