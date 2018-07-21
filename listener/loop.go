package listener

import (
	"github.com/elz-lang/elz/parser"
	"github.com/elz-lang/elz/ast"
)

func (s *ElzListener) ExitBreakStat(c *parser.BreakStatContext) {
	s.stats = append(s.stats, &ast.BreakStat{})
}
