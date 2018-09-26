package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitAccessChain(c *parser.AccessChainContext) {
	get := s.exprStack.Pop().(ast.Expr)
	from := s.exprStack.Pop().(ast.Expr)

	s.exprStack.Push(&ast.Access{
		From: from,
		Get:  get,
	})
}

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

// ExitCmp listen expression:
// - `expr < expr`
// - `expr > expr`
// - `expr <= expr`
// - `expr >= expr`
func (s *ElzListener) ExitCmp(c *parser.CmpContext) {
	re := s.exprStack.Pop().(ast.Expr)
	le := s.exprStack.Pop().(ast.Expr)
	e := &ast.BinaryExpr{
		LeftE:  le,
		RightE: re,
		Op:     c.GetOp().GetText(),
	}
	s.exprStack.Push(e)
}

// ExitEq listen expression `expr == expr`
func (s *ElzListener) ExitEq(c *parser.EqContext) {
	le := s.exprStack.Pop()
	re := s.exprStack.Pop()
	e := &ast.BinaryExpr{
		LeftE:  le.(ast.Expr),
		RightE: re.(ast.Expr),
		Op:     c.GetOp().GetText(),
	}
	s.exprStack.Push(e)
}

// ExitAs listen as expression, like: `3.1 as f64`
func (s *ElzListener) ExitAs(ctx *parser.AsContext) {
	e := s.exprStack.Pop()
	typ := ctx.TypeForm().GetText()
	s.exprStack.Push(&ast.As{E: e.(ast.Expr), T: typ})
}

// ExitRef listen ref expression, like: &a
func (s *ElzListener) ExitRef(ctx *parser.RefContext) {
	e := s.exprStack.Pop()
	s.exprStack.Push(&ast.Ref{E: e.(*ast.Id)})
}

// ExitDeRef
func (s *ElzListener) ExitDeRef(c *parser.DeRefContext) {
	e := s.exprStack.Pop()
	s.exprStack.Push(&ast.UnaryExpr{E: e.(ast.Expr), Op: "*"})
}

// ExitStr listen string literal, like: `"hello world"`
func (s *ElzListener) ExitStr(ctx *parser.StrContext) {
	value := ctx.STRING().GetText()
	value = value[1 : len(value)-1]
	s.exprStack.Push(&ast.Str{Val: value})
}

// ExitId listen identify rule
func (s *ElzListener) ExitId(ctx *parser.IdContext) {
	s.exprStack.Push(&ast.Id{Val: ctx.ID().GetText()})
}

// ExitFloat listen f32 literal, like: `0.1, .3, 3.14`
func (s *ElzListener) ExitFloat(ctx *parser.FloatContext) {
	suffix := ""
	if ctx.FloatSuffix() != nil {
		suffix = ctx.FloatSuffix().GetText()
	}
	switch suffix {
	case "'f64":
		s.exprStack.Push(&ast.F64{Val: ctx.FLOAT().GetText()})
	case "'f32":
		fallthrough
	default:
		s.exprStack.Push(&ast.F32{Val: ctx.FLOAT().GetText()})
	}
}

// ExitBool listen boolean literial like: `true`, `false`
func (s *ElzListener) ExitBool(c *parser.BoolContext) {
	s.exprStack.Push(&ast.Bool{Val: c.BOOLEAN().GetText()})
}

// ExitInt listen i32 literal, like: `1, 5, 321, 89`
func (s *ElzListener) ExitInt(ctx *parser.IntContext) {
	suffix := ""
	if ctx.IntSuffix() != nil {
		suffix = ctx.IntSuffix().GetText()
	}
	iVal := ctx.INT().GetText()
	switch suffix {
	case `'i8`:
		s.exprStack.Push(&ast.I8{Val: iVal})
	case `'i16`:
		s.exprStack.Push(&ast.I16{Val: iVal})
	case `'i32`:
		s.exprStack.Push(&ast.I32{Val: iVal})
	case `'i64`:
		s.exprStack.Push(&ast.I64{Val: iVal})
	case `'f32`:
		s.exprStack.Push(&ast.F32{Val: iVal})
	case `'f64`:
		s.exprStack.Push(&ast.F64{Val: iVal})
	default:
		s.exprStack.Push(&ast.I32{Val: iVal})
	}
}
