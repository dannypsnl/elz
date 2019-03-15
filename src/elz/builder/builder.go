package builder

import (
	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/codegen"
	"github.com/elz-lang/elz/src/elz/internal/collection/stack"
	"github.com/elz-lang/elz/src/elz/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Builder struct {
	*parser.BaseElzListener

	bindingType  map[string][]ast.Type
	bindTypeList []ast.Type
	exprStack    *stack.Stack

	astTree *codegen.Tree
}

func (b *Builder) GetTree() *codegen.Tree {
	return b.astTree
}

func (b *Builder) ExitProgram(c *parser.ProgramContext) {
}

func New() *Builder {
	return &Builder{
		bindingType:  make(map[string][]ast.Type),
		bindTypeList: make([]ast.Type, 0),
		exprStack:    stack.New(),
		astTree:      codegen.NewTree(),
	}
}

func NewFromFile(file string) (*codegen.Tree, error) {
	b := New()
	err := b.BuildFromFile(file)
	if err != nil {
		return nil, err
	}
	return b.GetTree(), nil
}

func NewFromCode(code string) *codegen.Tree {
	b := New()
	b.BuildFromCode(code)
	return b.GetTree()
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
