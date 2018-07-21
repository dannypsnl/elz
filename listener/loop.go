package listener

import (
	"github.com/elz-lang/elz/parser"
	"github.com/elz-lang/elz/ast"
)

type LoopBuilder struct {
	stats []ast.Stat
}

func NewLoopBuilder() *LoopBuilder {
	return &LoopBuilder{
		stats: make([]ast.Stat, 0),
	}
}

func (block *LoopBuilder) PushStat(stat ast.Stat) {
	block.stats = append(block.stats, stat)
}

func (block *LoopBuilder) Generate() ast.Stat {
	return &ast.Loop{Stats: block.stats}
}

func (s *ElzListener) EnterLoopStat(c *parser.LoopStatContext) {
	s.statBuilder.Push(NewLoopBuilder())
}
func (s *ElzListener) ExitLoopStat(c *parser.LoopStatContext) {
	s.statBuilder.Pop()
}

func (s *ElzListener) ExitBreakStat(c *parser.BreakStatContext) {
	s.stats = append(s.stats, &ast.BreakStat{})
}
