package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

// pattern represents:
//
//  match E {
//    // pattern
//    expression => statement
//  }
type pattern struct {
	expr ast.Expr
	stat ast.Stat
}

type MatchBuilder struct {
	// match expr {}
	expr     ast.Expr
	patterns []*pattern
}

func NewMatchBuilder() *MatchBuilder {
	return &MatchBuilder{
		expr:     nil,
		patterns: make([]*pattern, 0),
	}
}

func (m *MatchBuilder) NewPattern(e ast.Expr) {
	m.patterns = append(m.patterns, &pattern{
		expr: e,
		stat: nil,
	})
}

func (m *MatchBuilder) PushStat(stat ast.Stat) {
	lastOne := len(m.patterns)
	m.patterns[lastOne-1].stat = stat
}

func (m *MatchBuilder) Generate() ast.Expr {
	// match rule could be a expression, so is Expr not Stat
	ps := make([]*ast.Pattern, 0)
	for _, p := range m.patterns {
		ps = append(ps, &ast.Pattern{
			E:p.expr,
			S:p.stat,
		})
	}
	return ast.NewMatch(m.expr, ps)
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
	s.exprStack.Push(s.matchRuleBuilder.Generate())
	s.matchRuleBuilder = nil
}
