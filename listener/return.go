package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitReturnStat(ctx *parser.ReturnStatContext) {
	// get expr
	expr := s.exprStack.Pop().(ast.Expr)
	stat := &ast.Return{
		Expr: expr,
	}
	s.stats = append(s.stats, stat)
	// matchRule has higher level to combine a statement
	if s.matchRuleBuilder != nil {
	} else if s.fnBuilder != nil {
	} else {
		s.context.Reporter.Emit("return statement must in function")
	}
}
