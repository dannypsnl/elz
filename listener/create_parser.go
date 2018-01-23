package listener

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/elz-lang/elz/parser"
)

// TODO: Put parse task at here
func NewParse(source string) string {
	input := antlr.NewInputStream(source)
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	eal := /*listener.*/ New()
	antlr.ParseTreeWalkerDefault.Walk(eal, tree)
	return fmt.Sprint(eal.Module())
}
