package listener

import (
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) EnterMatchRule(c *parser.MatchRuleContext) {
	println("match " + c.MatchExpr().GetText())
}

func (s *ElzListener) ExitMatchExpr(c *parser.MatchExprContext) {
	s.exprStack.Pop()
}

func (s *ElzListener) ExitMatchRule(c *parser.MatchRuleContext) {
}
