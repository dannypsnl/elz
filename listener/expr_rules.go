package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitAddOrSub(ctx *parser.AddOrSubContext) {
	le := s.exprStack.Pop()
	re := s.exprStack.Pop()
	defer func() {
		// Only miss right expression can cause panic
		if r := recover(); r != nil {
			s.context.Reporter.Emit("expression miss error")
			s.exprStack.Push(le.(ast.Expr))
		}
	}()
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
	defer func() {
		// Only miss right expression can cause panic
		if r := recover(); r != nil {
			s.context.Reporter.Emit("expression miss error")
			s.exprStack.Push(le.(ast.Expr))
		}
	}()
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

func (s *ElzListener) ExitFloat(ctx *parser.FloatContext) {
	s.exprStack.Push(&ast.F32{Val: ctx.FLOAT().GetText()})
}

func (s *ElzListener) ExitInt(ctx *parser.IntContext) {
	s.exprStack.Push(&ast.I32{Val: ctx.INT().GetText()})
}
