package main

import (
	"fmt"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
	"github.com/golang-collections/collections/stack"
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
func (s *ElzListener) ExitDefine(ctx *parser.DefineContext) {
	// FIXME: We need to get Expression's type, not give it a default value.
	typ := "f32"
	if ctx.TypePass() != nil {
		typ = ctx.TypePass().GetText()
	}
	fmt.Println(ctx.ID().GetText(), `:`, typ)
	s.AstList = append(s.AstList, &ast.VarDefination{
		Immutable:  s.immutable,
		Export:     s.exportThis,
		Name:       ctx.ID().GetText(),
		VarType:    typ,
		Expression: &ast.Number{Val: "1"},
	})
}
func (s *ElzListener) ExitExpr(ctx *parser.ExprContext) {
	exprs := ctx.AllExpr()
	if len(exprs) != 2 {
		//s.exprStack.Push()
	}
}
func (s *ElzListener) ExitVarDefine(*parser.VarDefineContext) {
	if !s.immutable {
		s.immutable = true
	}
}
func (s *ElzListener) ExitStr(ctx *parser.StrContext) {
	print(ctx.STRING().GetText())
	s.exprStack.Push(ctx.STRING().GetText())
}
func (s *ElzListener) ExitId(ctx *parser.IdContext) {
	print(ctx.ID().GetText())
}
func (s *ElzListener) ExitNum(ctx *parser.NumContext) {
	print(ctx.NUM().GetText())
}
