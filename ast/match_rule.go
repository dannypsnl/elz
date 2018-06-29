package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Pattern struct {
	E Expr
	S Stat
}

type Match struct {
	matchExpr Expr
	patterns  []*Pattern
}

func NewMatch(e Expr, ps []*Pattern) *Match {
	return &Match{
		matchExpr: e,
		patterns:  ps,
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
	leave := llvm.InsertBasicBlock(bb, "switch_break")
	rest := llvm.InsertBasicBlock(bb, "rest")

	c.Builder.SetInsertPointAtEnd(rest)
	c.Builder.CreateBr(leave)
	c.Builder.ClearInsertionPoint()

	c.Builder.SetInsertPointAtEnd(bb)
	switchBlock := c.Builder.CreateSwitch(expr, rest, len(m.patterns))
	var prevPattern *llvm.BasicBlock
	for _, p := range m.patterns {
		c.Builder.SetInsertPointAtEnd(bb)
		pattern := llvm.InsertBasicBlock(bb, "p")

		switchBlock.AddCase(p.E.Codegen(c), pattern)

		c.Builder.SetInsertPointAtEnd(pattern)
		p.S.Codegen(c)
		c.Builder.CreateBr(leave)
		c.Builder.ClearInsertionPoint()

		if prevPattern != nil {
			pattern.MoveAfter(*prevPattern)
		} else {
			pattern.MoveAfter(bb)
		}
		prevPattern = &pattern
	}
	rest.MoveAfter(*prevPattern)
	leave.MoveAfter(rest)
	return llvm.Value{}
}

func (m *Match) Type(*Context) string {
	return "missing type now"
}
