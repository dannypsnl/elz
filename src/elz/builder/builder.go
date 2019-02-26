package builder

import (
	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/internal/collection/stack"
	"github.com/elz-lang/elz/src/elz/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Builder struct {
	*parser.BaseElzListener

	bindTypeList []ast.Type
	exprStack    *stack.Stack

	astTree *ast.Tree
}

func (b *Builder) GetTree() *ast.Tree {
	return b.astTree
}

func (b *Builder) ExitProgram(c *parser.ProgramContext) {
}

func New() *Builder {
	return &Builder{
		bindTypeList: make([]ast.Type, 0),
		exprStack:    stack.New(),
		astTree:      ast.NewTree(),
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
	tree := p.Program()
	antlr.ParseTreeWalkerDefault.Walk(b, tree)
}
