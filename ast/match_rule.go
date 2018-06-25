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

func (m *Match) Codegen(*Context) llvm.Value {
	return llvm.Value{}
}

func (m *Match) Type(*Context) string {
	return "missing type now"
}
