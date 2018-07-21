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
			E: p.expr,
			S: p.stat,
		})
	}
	if ps[len(ps)-1].E == nil {
		return ast.NewMatch(m.expr, ps[:len(ps)-1], ps[len(ps)-1])
	}
	return ast.NewMatch(m.expr, ps, nil)
}

func (s *ElzListener) EnterMatchRule(c *parser.MatchRuleContext) {
	s.statBuilder.Push(NewMatchBuilder())
}

func (s *ElzListener) ExitMatchExpr(c *parser.MatchExprContext) {
	if matchRuleBuilder, ok := s.statBuilder.Last().(*MatchBuilder); !ok {
		panic("Match Rule's implementation has bug, matchRuleBuilder should not be nil")
	} else {
		expr := s.exprStack.Pop().(ast.Expr)
		if matchRuleBuilder.expr == nil {
			matchRuleBuilder.expr = expr
		} else {
			matchRuleBuilder.NewPattern(expr)
		}
	}

}

func (s *ElzListener) EnterRestPattern(c *parser.RestPatternContext) {
	if matchRuleBuilder, ok := s.statBuilder.Last().(*MatchBuilder); !ok {
		panic("Match Rule's implementation has bug, matchRuleBuilder should not be nil in rest pattern")
	} else {
		// FIXME: make rest pattern more explicated at here
		matchRuleBuilder.NewPattern(nil)
	}
}

func (s *ElzListener) ExitMatchRule(c *parser.MatchRuleContext) {
	if matchRuleBuilder, ok := s.statBuilder.Last().(*MatchBuilder); !ok {
		panic("Match Rule's implementation has bug, matchRuleBuilder should not be nil in rest pattern")
	} else {
		s.exprStack.Push(matchRuleBuilder.Generate())
		s.statBuilder.Pop()
	}
}
