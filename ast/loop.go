package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Loop struct {

}

type BreakStat struct {

}

func (b *BreakStat) Check(*Context) {}
func (b *BreakStat) Codegen(c *Context) llvm.Value {
	return llvm.Value{}
}