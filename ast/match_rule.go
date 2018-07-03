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
	leave := llvm.InsertBasicBlock(bb, "")
	rest := llvm.InsertBasicBlock(bb, "")

	c.Builder.SetInsertPointAtEnd(rest)
	c.Builder.CreateBr(leave)

	c.Builder.SetInsertPointAtEnd(bb)
	switchBlock := c.Builder.CreateSwitch(expr, rest, len(m.patterns))
	prevPattern := bb
	for _, p := range m.patterns {
		c.Builder.SetInsertPointAtEnd(bb)

		pattern := llvm.InsertBasicBlock(bb, "")
		switchBlock.AddCase(p.E.Codegen(c), pattern)

		c.Builder.SetInsertPointAtEnd(pattern)

		// each pattern at least have to do
		p.S.Codegen(c)
		c.Builder.CreateBr(leave)

		pattern.MoveAfter(prevPattern)
		prevPattern = pattern
	}
	rest.MoveAfter(prevPattern)
	leave.MoveAfter(rest)

	c.Builder.SetInsertPointAtEnd(leave)
	return llvm.Value{}
}

func (m *Match) Type(*Context) string {
	return "missing type now"
}
