package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

// Pattern represents:
//
//  match E {
//    // Pattern
//    expression => statement
//  }
type Pattern struct {
	expr ast.Expr
	stat ast.Stat
}

type MatchBuilder struct {
	// match expr {}
	expr     ast.Expr
	patterns []Pattern
}

func (s *ElzListener) EnterMatchRule(c *parser.MatchRuleContext) {
	println("match " + c.MatchExpr().GetText())
}

func (s *ElzListener) ExitMatchExpr(c *parser.MatchExprContext) {
	s.exprStack.Pop()
}

func (s *ElzListener) ExitMatchRule(c *parser.MatchRuleContext) {
}
