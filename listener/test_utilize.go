package listener

import (
	"strings"
	"testing"

	"github.com/elz-lang/elz/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func hasTestTemplate(t *testing.T, source, expectedIn string) {
	res := NewParse(source)

	if !strings.Contains(res, expectedIn) {
		t.Errorf("expected has: `%s`, actual: `%s`", expectedIn, res)
	}
}

func listener(code string) *ElzListener {
	input := antlr.NewInputStream(code)
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	listener := /*listener.*/ New()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener
}
