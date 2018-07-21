package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitReturnStat(ctx *parser.ReturnStatContext) {
	// We can return some thing in match expression
	// or in a function
	lastBuilder := s.statBuilder.Last()
	if lastBuilder != nil {
		_, inFn := lastBuilder.(*FnBuilder)
		_, inMatchRule := lastBuilder.(*MatchBuilder)
		if !inFn && !inMatchRule {
			s.context.Reporter.Emit("return statement must in function")
		}
	} else {
		s.context.Reporter.Emit("return statement must in function")
	}

	// get expr
	expr := s.exprStack.Pop().(ast.Expr)
	stat := &ast.Return{
		Expr: expr,
	}
	s.stats = append(s.stats, stat)
}
