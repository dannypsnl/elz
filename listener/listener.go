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
	// exportThis markup the reference Name should be public or not.
	exportThis bool
	// variable default immutable.
	immutable bool
}

func (s *ElzListener) Module() llvm.Module {
	for _, ast := range s.AstList {
		ast.Codegen(s.context)
	}
	return s.context.Module
}

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

func (s *ElzListener) EnterExportor(*parser.ExportorContext) {
	fmt.Print(`public `)
	s.exportThis = true
}

func (s *ElzListener) EnterVarDefine(ctx *parser.VarDefineContext) {
	fmt.Print(`var `)
	if ctx.GetMut() != nil {
		s.immutable = false
	}
}

func (s *ElzListener) ExitVarDefine(*parser.VarDefineContext) {
	if !s.immutable {
		s.immutable = true
	}
}

func (s *ElzListener) ExitDefine(ctx *parser.DefineContext) {
	expr := s.exprStack.Pop()
	typ := expr.(ast.Expr).Type(s.context)
	// TODO: fix with scope rule, and some rule to detected fn, type, trait or what
	name := ctx.ID().GetText()
	if ctx.TypePass() != nil {
		typ = ctx.TypePass().GetText()
	}

	fmt.Printf("%s: %s = %s\n", ctx.ID().GetText(), typ, expr)

	s.AstList = append(s.AstList, &ast.VarDefination{
		Immutable:  s.immutable,
		Export:     s.exportThis,
		Name:       name,
		VarType:    typ,
		Expression: expr.(ast.Expr),
	})
	s.context.VarsType[name] = typ
}

func (s *ElzListener) ExitAddOrSub(ctx *parser.AddOrSubContext) {
	le := s.exprStack.Pop()
	re := s.exprStack.Pop()
	e := &ast.BinaryExpr{
		LeftE:  le.(ast.Expr),
		RightE: re.(ast.Expr),
		Op:     ctx.GetOp().GetText(),
	}
	s.exprStack.Push(e)
}

func (s *ElzListener) ExitMulOrDiv(ctx *parser.MulOrDivContext) {
	le := s.exprStack.Pop()
	re := s.exprStack.Pop()
	e := &ast.BinaryExpr{
		LeftE:  le.(ast.Expr),
		RightE: re.(ast.Expr),
		Op:     ctx.GetOp().GetText(),
	}
	s.exprStack.Push(e)
}

func (s *ElzListener) ExitStr(ctx *parser.StrContext) {
	s.exprStack.Push(&ast.Str{Val: ctx.STRING().GetText()})
}
func (s *ElzListener) ExitId(ctx *parser.IdContext) {
	// BUG: ast.Id implementation is buggy
	s.exprStack.Push(&ast.Id{Val: ctx.ID().GetText()})
}
func (s *ElzListener) ExitNum(ctx *parser.NumContext) {
	s.exprStack.Push(&ast.Number{Val: ctx.NUM().GetText()})
}
