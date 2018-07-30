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

	loop := llvm.InsertBasicBlock(bb, "loop")
	loop.MoveAfter(bb)

	c.Builder.SetInsertPointAtEnd(bb)
	c.Builder.CreateBr(loop)

	leave := llvm.InsertBasicBlock(bb, "loop.end")
	leave.MoveAfter(loop)
	c.breaks.Push(leave)

	c.Builder.SetInsertPointAtEnd(loop)
	// statements
	for _, s := range b.Stats {
		s.Codegen(c)
	}
	c.Builder.CreateBr(loop)

	c.breaks.Pop()
	c.Builder.SetInsertPointAtEnd(leave)

	return llvm.Value{}
}

type BreakStat struct {
}

func (b *BreakStat) Check(*Context) {}
func (b *BreakStat) Codegen(c *Context) llvm.Value {
	// FIXME: help break statement knows jump to where
	bb := c.breaks.Last().(llvm.BasicBlock)
	c.Builder.CreateBr(bb)
	return llvm.Value{}
}
