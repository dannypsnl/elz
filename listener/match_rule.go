package listener

import (
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) EnterMatchRule(c *parser.MatchRuleContext) {
	println("match " + c.Expr(0).GetText())
}

func (s *ElzListener) ExitMatchRule(c *parser.MatchRuleContext) {
	for i := 0; i < len(c.AllExpr()); i++ {
		s.exprStack.Pop()
	}
}
