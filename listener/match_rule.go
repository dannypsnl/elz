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

func NewMatchBuilder() *MatchBuilder {
	return &MatchBuilder{
		expr:     nil,
		patterns: make([]Pattern, 0),
	}
}

func (m *MatchBuilder) NewPattern(e ast.Expr) {
	m.patterns = append(m.patterns, Pattern{
		expr: e,
		stat: nil,
	})
}

func (m *MatchBuilder) PushStat(stat ast.Stat) {
	lastOne := len(m.patterns)
	m.patterns[lastOne].stat = stat
}

func (s *ElzListener) EnterMatchRule(c *parser.MatchRuleContext) {
	s.matchRuleBuilder = NewMatchBuilder()
	println("match " + c.MatchExpr().GetText())
}

func (s *ElzListener) ExitMatchExpr(c *parser.MatchExprContext) {
	if s.matchRuleBuilder == nil {
		panic("Match Rule's implementation has bug, matchRuleBuilder should not be nil")
	}
	expr := s.exprStack.Pop().(ast.Expr)
	if s.matchRuleBuilder.expr == nil {
		s.matchRuleBuilder.expr = expr
	} else {
		s.matchRuleBuilder.NewPattern(expr)
	}
}

func (s *ElzListener) ExitMatchRule(c *parser.MatchRuleContext) {
	s.matchRuleBuilder = nil
}
