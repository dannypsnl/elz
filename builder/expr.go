package builder

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (b *Builder) PushExpr(e interface{}) {
	b.exprStack.Push(e)
}
func (b *Builder) PopExpr() interface{} {
	return b.exprStack.Pop()
}

// ExitMulDiv:
//
//  1 * 2
//  3 / 4
//
// WorkFlow:
//
// e.g. leftExpr Op rightExpr
// 1. push leftExpr
// 2. push rightExpr
// MulDiv
// 3. pop rightExpr
// 4. pop leftExpr
func (b *Builder) ExitMulDiv(c *parser.MulDivContext) {
	rExpr := b.PopExpr().(ast.Expr)
	lExpr := b.PopExpr().(ast.Expr)
	// left expr, right expr, operator
	b.PushExpr(&ast.BinaryExpr{
		LExpr: lExpr,
		RExpr: rExpr,
		Op:    c.GetOp().GetText(),
	})
}
func (b *Builder) ExitAddSub(c *parser.AddSubContext) {
	rExpr := b.PopExpr().(ast.Expr)
	lExpr := b.PopExpr().(ast.Expr)
	// left expr, right expr, operator
	b.PushExpr(&ast.BinaryExpr{
		LExpr: lExpr,
		RExpr: rExpr,
		Op:    c.GetOp().GetText(),
	})
}
func (b *Builder) ExitInt(c *parser.IntContext) {
	b.PushExpr(ast.NewInt(c.INT().GetText()))
}
func (b *Builder) ExitFloat(c *parser.FloatContext) {
	b.PushExpr(ast.NewFloat(c.FLOAT().GetText()))
}
func (b *Builder) ExitString(c *parser.StringContext) {
	b.PushExpr(ast.NewString(c.STRING().GetText()))
}
func (b *Builder) ExitBoolean(c *parser.BooleanContext) {
	b.PushExpr(ast.NewBool(c.BOOLEAN().GetText()))
}

func (b *Builder) ExitFuncCall(c *parser.FuncCallContext) {
	if len(c.AllExpr()) == 0 {
		b.PushExpr(&ast.Ident{
			Value: c.IDENT().GetText(),
		})
		return
	}
	exprList := make([]ast.Expr, 0)
	for e, hasExpr := b.PopExpr().(ast.Expr); hasExpr; e, hasExpr = b.PopExpr().(ast.Expr) {
		exprList = append([]ast.Expr{e}, exprList...)
	}
	b.PushExpr(&ast.FuncCall{
		Identifier: c.IDENT().GetText(),
		ExprList:   exprList,
	})
}
