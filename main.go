package main

import (
	"fmt"
)

func main() {
<<<<<<< HEAD
	fmt.Println("elz compiler...")
=======
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
	module := ir.NewModule()
	for _, ast := range eal.AstList {
		ast.Codegen(module)
	}
	fmt.Println(module)
>>>>>>> @{-1}
}
