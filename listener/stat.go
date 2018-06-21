package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitStat(c *parser.StatContext) {
	some := s.exprStack.Pop()
	if v, ok := some.(*ast.FnCall); ok {
		if s.matchRuleBuilder != nil {
			s.matchRuleBuilder.PushStat(v)
		} else if s.fnBuilder != nil {
			s.fnBuilder.Stat(v)
		}
	}

	if s.matchRuleBuilder != nil {
		if len(s.stats) != 0 {
			for _, st := range s.stats {
				s.matchRuleBuilder.PushStat(st)
			}
		}
	} else if s.fnBuilder != nil {
		if len(s.stats) != 0 {
			for _, st := range s.stats {
				s.fnBuilder.Stat(st)
			}
		}
	}

	s.stats = make([]ast.Stat, 0)
}
