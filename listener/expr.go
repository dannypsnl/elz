package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

// ExitAddOrSub listen expression `expr + expr` or `expr - expr`
func (s *ElzListener) ExitAddOrSub(ctx *parser.AddOrSubContext) {
	re := s.exprStack.Pop()
	le := s.exprStack.Pop()
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

// ExitMulOrDiv listen expression `expr * expr` or `expr / expr`
func (s *ElzListener) ExitMulOrDiv(ctx *parser.MulOrDivContext) {
	re := s.exprStack.Pop()
	le := s.exprStack.Pop()
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

// ExitAs listen as expression, like: `3.1 as f64`
func (s *ElzListener) ExitAs(ctx *parser.AsContext) {
	e := s.exprStack.Pop()
	typ := ctx.TypeForm().GetText()
	s.exprStack.Push(&ast.As{E: e.(ast.Expr), T: typ})
}

// ExitStr listen string literal, like: `"hello world"`
func (s *ElzListener) ExitStr(ctx *parser.StrContext) {
	s.exprStack.Push(&ast.Str{Val: ctx.STRING().GetText()})
}

// ExitId listen identify rule
func (s *ElzListener) ExitId(ctx *parser.IdContext) {
	s.exprStack.Push(&ast.Id{Val: ctx.ID().GetText()})
}

// ExitFloat listen f32 literal, like: `0.1, .3, 3.14`
func (s *ElzListener) ExitFloat(ctx *parser.FloatContext) {
	s.exprStack.Push(&ast.F32{Val: ctx.FLOAT().GetText()})
}

// ExitInt listen i32 literal, like: `1, 5, 321, 89`
func (s *ElzListener) ExitInt(ctx *parser.IntContext) {
	suffix := ""
	if ctx.IntSuffix() != nil {
		suffix = ctx.IntSuffix().GetText()
	}
	switch suffix {
	case "'i8":
		s.exprStack.Push(&ast.I8{Val: ctx.INT().GetText()})
	case "'i16":
		s.exprStack.Push(&ast.I16{Val: ctx.INT().GetText()})
	case "'i32":
		s.exprStack.Push(&ast.I32{Val: ctx.INT().GetText()})
	case "'i64":
		s.exprStack.Push(&ast.I64{Val: ctx.INT().GetText()})
	default:
		s.exprStack.Push(&ast.I32{Val: ctx.INT().GetText()})
	}
}
