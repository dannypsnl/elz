package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Pattern struct {
	E Expr
	S Stat
}

func (p *Pattern) IsBr() bool {
	switch p.S.(type) {
	case *BreakStat:
		return true
	default:
		return false
	}
}

type Match struct {
	matchExpr   Expr
	patterns    []*Pattern
	restPattern *Pattern
}

func NewMatch(e Expr, ps []*Pattern, restPattern *Pattern) *Match {
	return &Match{
		matchExpr:   e,
		patterns:    ps,
		restPattern: restPattern,
	}
}

func (m *Match) Check(c *Context) {
	m.matchExpr.Check(c)

	for _, p := range m.patterns {
		p.E.Check(c)
	}
}

func (m *Match) Codegen(c *Context) llvm.Value {
	bb := c.Builder.GetInsertBlock()
	expr := m.matchExpr.Codegen(c)
	leave := llvm.InsertBasicBlock(bb, "match.end")
	rest := llvm.InsertBasicBlock(bb, "pattern.rest")

	c.Builder.SetInsertPointAtEnd(bb)
	switchBlock := c.Builder.CreateSwitch(expr, rest, len(m.patterns))
	prevPattern := bb
	for _, pattern := range m.patterns {
		c.Builder.SetInsertPointAtEnd(bb)

		patternBlock := llvm.InsertBasicBlock(bb, "pattern")
		switchBlock.AddCase(pattern.E.Codegen(c), patternBlock)

		c.Builder.SetInsertPointAtEnd(patternBlock)

		// each patternBlock at least have to do
		pattern.S.Codegen(c)
		if !pattern.IsBr() {
			c.Builder.CreateBr(leave)
		}

		patternBlock.MoveAfter(prevPattern)
		prevPattern = patternBlock
	}
	rest.MoveAfter(prevPattern)
	leave.MoveAfter(rest)

	c.Builder.SetInsertPointAtEnd(rest)

	if m.restPattern != nil {
		m.restPattern.S.Codegen(c)
	}
	c.Builder.CreateBr(leave)

	c.Builder.SetInsertPointAtEnd(leave)
	// FIXME: match should also be an expression
	return llvm.Value{}
}

func (m *Match) Type(*Context) string {
	// FIXME: match should also be an expression
	return "missing type now"
}

func (m *Match) ExprStat() {}
