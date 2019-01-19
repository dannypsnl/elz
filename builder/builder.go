package builder

import (
	"fmt"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/internal/collection/stack"
	"github.com/elz-lang/elz/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Builder struct {
	*parser.BaseElzListener

	bindings []*ast.Func

	exprStack *stack.Stack
	debug     bool
}

func (b *Builder) ExitProg(c *parser.ProgContext) {
	if b.debug {
		fmt.Printf("%#v\n", c)
	}
}

func New() *Builder {
	return &Builder{
		debug:     false,
		exprStack: stack.New(),
		bindings:  make([]*ast.Func, 0),
	}
}

func (b *Builder) BuildFromCode(code string) {
	input := antlr.NewInputStream(code)
	b.build(input)
}

func (b *Builder) BuildFromFile(filename string) error {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return err
	}
	b.build(input)
	return nil
}

func (b *Builder) build(input antlr.CharStream) {
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(b, tree)
}
