package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

type Scope struct {
	block *ir.Block
	scope map[string]value.Value
}

func NewScope(block *ir.Block, scope map[string]value.Value) *Scope {
	return &Scope{
		block: block,
		scope: scope,
	}
}

func (s *Scope) NewAdd(l, r value.Value) value.Value {
	return s.block.NewAdd(l, r)
}
func (s *Scope) Var(name string) value.Value {
	return s.scope[name]
}
