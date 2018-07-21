package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

type PushStatement interface {
	PushStat(ast.Stat)
}

func (s *ElzListener) ExitStat(c *parser.StatContext) {
	// if fail, let it panic, that's compiler bug
	statContainer := s.statBuilder.Last()
	if statContainer != nil {
		statContainer := statContainer.(PushStatement)
		some := s.exprStack.Pop()
		switch some.(type) {
		case ast.ExprStat:
			v := some.(ast.ExprStat)
			statContainer.PushStat(v)
		}

		if len(s.stats) != 0 {
			for _, st := range s.stats {
				statContainer.PushStat(st)
			}
		}

		s.stats = make([]ast.Stat, 0)
	}
}
