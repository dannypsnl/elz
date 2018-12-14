package experiment

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func declarePrintf(mod *ir.Module) *ir.Function {
	f := mod.NewFunc(
		"printf",
		types.I32,
		ir.NewParam("format", types.NewPointer(types.I8)),
	)
	f.Sig.Variadic = true
	return f
}
