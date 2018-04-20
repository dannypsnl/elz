package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/collection/stack"
	"github.com/elz-lang/elz/parser"

	"llvm.org/llvm/bindings/go/llvm"
)

type ElzListener struct {
	*parser.BaseElzListener

	context *ast.Context
	// AstList contain top level's ast
	AstList []ast.Ast
	// exprStack help we implement expression percedence table.
	exprStack *stack.Stack // Stack Pop nil is nothing in there
	// fnBuilder
	fnBuilder *FnBuilder
	// exportThis markup the reference Name should be public or not.
	exportThis bool
	// variable default immutable.
	immutable bool
	// isGlobalDef, if is global level var
	isGlobalDef bool
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
	}
}

func (s *ElzListener) ExitProg(ctx *parser.ProgContext) {
	for _, ast := range s.AstList {
		ast.Check(s.context)
	}
	for _, ast := range s.AstList {
		ast.Codegen(s.context)
	}
	s.context.Reporter.Report()
}
