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
	// matchRule has higher level to combine a statement
	if s.matchRuleBuilder != nil {
		s.matchRuleBuilder.PushStat(stat)
	} else if s.fnBuilder != nil {
		s.fnBuilder.Stat(stat)
	} else {
		s.context.Reporter.Emit("return statement must in function")
	}
}
