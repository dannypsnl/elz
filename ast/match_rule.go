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
	rest := llvm.InsertBasicBlock(bb, "rest")
	switchBlock := c.Builder.CreateSwitch(expr, rest, len(m.patterns))
	for _, p := range m.patterns {
		pattern := llvm.InsertBasicBlock(bb, "p")
		c.Builder.SetInsertPointAtEnd(pattern)
		p.S.Codegen(c)
		c.Builder.ClearInsertionPoint()
		switchBlock.AddCase(p.E.Codegen(c), pattern)
	}
	return llvm.Value{}
}

func (m *Match) Type(*Context) string {
	return "missing type now"
}
