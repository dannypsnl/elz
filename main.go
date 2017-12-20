package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
	"llvm.org/llvm/bindings/go/llvm"
)

func main() {
	// TODO: Use CharStream & a function wrap this all, here should handle other mession.
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	eal := NewElzListener()
	antlr.ParseTreeWalkerDefault.Walk(eal, tree)

	// Module is temprary variable
	ctx := &ast.Context{
		llvm.NewModule("main"),
		llvm.NewContext(),
		make(map[string]llvm.Value),
		llvm.NewBuilder(),
	}
	for _, ast := range eal.AstList {
		ast.Codegen(ctx)
	}
	fmt.Println(ctx.Module)
}
