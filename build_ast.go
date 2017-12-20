package main

import (
	"fmt"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
	"github.com/golang-collections/collections/stack"
	_ "llvm.org/llvm/bindings/go/llvm"
)

type ElzListener struct {
	*parser.BaseElzListener
	AstList    []ast.Ast
	exprStack  *stack.Stack // Stack Pop nil is nothing in there
	exportThis bool
	immutable  bool
}

func NewElzListener() *ElzListener {
	return &ElzListener{
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
	typ := expr.(ast.Expr).Type()
	if ctx.TypePass() != nil {
		typ = ctx.TypePass().GetText()
	}
	fmt.Println(ctx.ID().GetText(), `:`, typ)
	s.AstList = append(s.AstList, &ast.VarDefination{
		Immutable:  s.immutable,
		Export:     s.exportThis,
		Name:       ctx.ID().GetText(),
		VarType:    expr.(ast.Expr).Type(),
		Expression: expr.(ast.Expr),
	})
	fmt.Println(expr)
}

func (s *ElzListener) ExitExpr(ctx *parser.ExprContext) {
	exprs := ctx.AllExpr()
	if len(exprs) != 2 {
	} else {
		le := s.exprStack.Pop()
		re := s.exprStack.Pop()
		if le != nil && re != nil {
			//fm := fmt.Sprintln(le, ctx.GetOp().GetText(), re)
			e := &ast.BinaryExpr{
				LeftE:  le.(ast.Expr),
				RightE: re.(ast.Expr),
				Op:     ctx.GetOp().GetText(),
			}
			s.exprStack.Push(e)
		}
	}
}

func (s *ElzListener) ExitStr(ctx *parser.StrContext) {
	s.exprStack.Push(&ast.Str{Val: ctx.STRING().GetText()})
}
func (s *ElzListener) ExitId(ctx *parser.IdContext) {
	// FIXME: Identifier still do not have Expr to represented it.
	s.exprStack.Push(ctx.ID().GetText())
}
func (s *ElzListener) ExitNum(ctx *parser.NumContext) {
	//s.exprStack.Push(ctx.NUM().GetText())
	s.exprStack.Push(&ast.Number{Val: "1"})
}
