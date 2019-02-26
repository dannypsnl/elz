package builder

import (
	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/parser"
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
func (b *Builder) ExitIdentifier(c *parser.IdentifierContext) {
	b.PushExpr(ast.NewIdent(c.GetText()))
}

func (b *Builder) ExitFnCall(c *parser.FnCallContext) {
	exprList := make([]*ast.Arg, 0)
	for e, hasExpr := b.PopExpr().(*ast.Arg); hasExpr; e, hasExpr = b.PopExpr().(*ast.Arg) {
		exprList = append([]*ast.Arg{e}, exprList...)
	}
	b.PushExpr(&ast.FuncCall{
		FuncName: c.IDENT().GetText(),
		ExprList: exprList,
	})
}
func (b *Builder) ExitArg(c *parser.ArgContext) {
	expr := b.PopExpr().(ast.Expr)
	ident := ""
	if c.IDENT() != nil {
		ident = c.IDENT().GetText()
	}
	b.PushExpr(&ast.Arg{
		Ident: ident,
		Expr:  expr,
	})
}
