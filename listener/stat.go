package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitStat(c *parser.StatContext) {
	some := s.exprStack.Pop()
	if v, ok := some.(*ast.FnCall); ok {
		if s.fnBuilder != nil {
			s.fnBuilder.Stat(v)
		}
	}
}
