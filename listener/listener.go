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
	if ctx.TypePass() != nil {
		typ = ctx.TypePass().GetText()
	}

	if s.exportThis {
		fmt.Print("public ")
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

func (s *ElzListener) EnterFnDefine(ctx *parser.FnDefineContext) {
	// FIXME: implement fn generate
	if s.exportThis {
		fmt.Print("public ")
	}
	fmt.Printf("fn %s\n", ctx.ID().GetText())
	// TODO: local var def should be spec
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
	s.exprStack.Push(&ast.Id{Val: ctx.ID().GetText()})
}
func (s *ElzListener) ExitNum(ctx *parser.NumContext) {
	s.exprStack.Push(&ast.Number{Val: ctx.NUM().GetText()})
}
