package listener

import (
	"fmt"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/collection/stack"
	"github.com/elz-lang/elz/parser"
	"github.com/elz-lang/elz/util"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"llvm.org/llvm/bindings/go/llvm"
)

// ElzListener listen signal from antlr parser then provide API let user can access the compiled result
type ElzListener struct {
	*parser.BaseElzListener

	context *ast.Context
	// AstList contain top level's ast
	AstList []ast.Ast
	// exprStack help we implement expression percedence table.
	exprStack *stack.Stack // Stack Pop nil is nothing in there
	notations []util.Notation
	// fnBuilder
	fnBuilder *FnBuilder
	// typeDefineBuilder
	typeDefineBuilder *TypeDefineBuilder
	// matchRuleBuilder
	matchRuleBuilder *MatchBuilder
	// exportThis markup the reference Name should be public or not.
	exportThis bool
	// variable default immutable.
	immutable bool
	// isGlobalDef, if is global level var
	isGlobalDef bool
	// inExternBlock mean in extern "C" {}
	inExternBlock bool
	// This one is fuzzy, it
	isStatExpr bool
}

// Module return the llvm.Module generate by parse process
func (s *ElzListener) Module() llvm.Module {
	if s.context.Reporter.HasNoError() {
		return s.context.Module
	} else {
		return llvm.Module{}
	}
}

// New create a new listener
func New() *ElzListener {
	return &ElzListener{
		context:   ast.NewContext(),
		immutable: true,
		exprStack: stack.New(),
		notations: make([]util.Notation, 0),
	}
}

// ExitProg is end of an elz file, we will check and report error at this stage
func (s *ElzListener) ExitProg(ctx *parser.ProgContext) {
	for _, ast := range s.AstList {
		ast.Check(s.context)
	}
	for _, ast := range s.AstList {
		ast.Codegen(s.context)
	}
	s.context.Reporter.Report()
}

func NewParse(source string) string {
	input := antlr.NewInputStream(source)
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	listener := /*listener.*/ New()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return fmt.Sprint(listener.Module())
}

func (s *ElzListener) TakeAllNotation() (result []util.Notation) {
	result = s.notations
	s.notations = make([]util.Notation, 0)
	return
}
