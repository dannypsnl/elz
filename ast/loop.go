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
	bb := c.Builder.GetInsertBlock()

	loop := llvm.InsertBasicBlock(bb, "")
	loop.MoveAfter(bb)
	c.Builder.SetInsertPointAtEnd(loop)
	// statements
	for _, s := range b.Stats {
		s.Codegen(c)
	}
	c.Builder.CreateBr(loop)

	return llvm.Value{}
}

type BreakStat struct {
}

func (b *BreakStat) Check(*Context) {}
func (b *BreakStat) Codegen(c *Context) llvm.Value {
	// FIXME: help break statement knows jump to where
	return llvm.Value{}
}
