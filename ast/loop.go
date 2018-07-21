package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Loop struct {
	Stats []Stat
}

func (b *Loop) Check(c *Context) {
	for _, s := range b.Stats {
		s.Check(c)
	}
}
func (b *Loop) Codegen(c *Context) llvm.Value {
	return llvm.Value{}
}

type BreakStat struct {
}

func (b *BreakStat) Check(*Context) {}
func (b *BreakStat) Codegen(c *Context) llvm.Value {
	return llvm.Value{}
}
