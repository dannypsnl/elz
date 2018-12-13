package experiment

import (
	"fmt"
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func TestHowToCreateVARGFunction(t *testing.T) {
	mod := ir.NewModule()
	f := mod.NewFunc(
		"printf",
		types.I32,
		ir.NewParam("format", types.NewPointer(types.I8)),
	)
	fmt.Printf("%s\n", f.Def())
}
