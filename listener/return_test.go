package listener

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/elz-lang/elz/parser"
)

func TestReturnNotInFunction(t *testing.T) {
	source := `
	return 10
	`

	input := antlr.NewInputStream(source)
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	listener := /*listener.*/ New()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if listener.context.Reporter.ErrMsgs[0] != "return statement must in function" {
		t.Error(listener.context.Reporter.ErrMsgs)
	}
}
