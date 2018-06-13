package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitReturnStat(ctx *parser.ReturnStatContext) {
	// We can return some thing in match expression
	// or in a function
	if s.matchRuleBuilder == nil && s.fnBuilder == nil {
		s.context.Reporter.Emit("return statement must in function")
	}
	// get expr
	expr := s.exprStack.Pop().(ast.Expr)
	stat := &ast.Return{
		Expr: expr,
	}
	s.stats = append(s.stats, stat)
}
