package codegen

import (
	"github.com/elz-lang/elz/src/elz/internal/collection/stack"
	"github.com/elz-lang/elz/src/elz/types"

	"github.com/llir/llvm/ir"
)

type context struct {
	stack *stack.Stack
	*ir.Block
	binds   map[string]*ir.Param
	typeMap *types.TypeMap
}

func newContext(block *ir.Block, typeMap *types.TypeMap) *context {
	s := stack.New()
	s.Push(block)
	return &context{
		stack:   s,
		Block:   s.Last().(*ir.Block),
		binds:   make(map[string]*ir.Param),
		typeMap: typeMap,
	}
}

func (c *context) setBlock(b *ir.Block) {
	c.stack.Push(c.Block)
	c.Block = b
}

func (c *context) restoreBlock() {
	c.Block = c.stack.Pop().(*ir.Block)
}
