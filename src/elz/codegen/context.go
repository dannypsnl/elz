package codegen

import (
	"github.com/elz-lang/elz/src/elz/types"

	"github.com/llir/llvm/ir"
)

type context struct {
	backupBlock *ir.Block
	*ir.Block
	binds   map[string]*ir.Param
	typeMap *types.TypeMap
}

func newContext(block *ir.Block, typeMap *types.TypeMap) *context {
	return &context{
		backupBlock: block,
		Block:       block,
		binds:       make(map[string]*ir.Param),
		typeMap:     typeMap,
	}
}

func (c *context) setBlock(b *ir.Block) {
	c.backupBlock = c.Block
	c.Block = b
}

func (c *context) restoreBlock() {
	c.Block = c.backupBlock
}
