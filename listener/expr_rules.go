package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

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
