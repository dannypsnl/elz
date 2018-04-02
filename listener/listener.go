package listener

import (
	"fmt"

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
	for _, ast := range s.AstList {
		ast.Codegen(s.context)
	}
	return s.context.Module
}

// New create a new listener
func New() *ElzListener {
	return &ElzListener{
		context:   ast.NewContext(),
		immutable: true,
		exprStack: stack.New(),
	}
}

func (s *ElzListener) EnterProg(ctx *parser.ProgContext) {
	fmt.Println(`Elz prog`)
}

// EnterExportor: + prefix
func (s *ElzListener) EnterExportor(*parser.ExportorContext) {
	s.exportThis = true
}

// VarDef:
//   let (mut) $var_name = $expr
func (s *ElzListener) EnterLocalVarDef(ctx *parser.LocalVarDefContext) {
	if ctx.GetMut() != nil {
		s.immutable = false
	}
}

func (s *ElzListener) ExitLocalVarDef(*parser.LocalVarDefContext) {
	if !s.immutable {
		s.immutable = true
	}
}

func (s *ElzListener) EnterGlobalVarDef(*parser.GlobalVarDefContext) {
	s.isGlobalDef = true
}
func (s *ElzListener) ExitGlobalVarDef(*parser.GlobalVarDefContext) {
	s.isGlobalDef = false
}

// Def:
//   $var_name = $expr
func (s *ElzListener) ExitDefine(ctx *parser.DefineContext) {
	// get expr
	expr := s.exprStack.Pop()
	// get type from expression
	typ := expr.(ast.Expr).Type(s.context)
	// get identifier
	// TODO: fix with scope rule, and some rule to detected fn, type, trait or what
	name := ctx.ID().GetText()
	// get type from source code, so we can find out the problem if expr != user_def type
	if ctx.TypeForm() != nil {
		typ = ctx.TypeForm().GetText()
	}

	if s.exportThis {
		s.exportThis = false
	}

	// FIXME: Need to classify global var & local var, because local var of course can't be export
	// FIXME: Need to classify heap & stack, and can find out the lifetime, else sending data by return will become bug
	if s.isGlobalDef {
		s.AstList = append(s.AstList, &ast.GlobalVarDef{
			Export:     s.exportThis,
			Name:       name,
			VarType:    typ,
			Expression: expr.(ast.Expr),
		})
		s.context.VarsType[name] = typ
		// Record type for compiler
	} else if s.fnBuilder != nil {
		s.fnBuilder.Stat(&ast.LocalVarDef{
			Immutable:  s.immutable,
			Name:       name,
			VarType:    typ,
			Expression: expr.(ast.Expr),
		})
	} else {
		panic("A define should be Global Var else is Local Var, but both of this isn't")
	}
}
