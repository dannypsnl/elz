package experiment

import (
	"testing"

	asserter "github.com/dannypsnl/assert"
	"github.com/llir/llvm/ir"
)

func TestHowToCreateVARGFunction(t *testing.T) {
	assert := asserter.NewTester(t)
	mod := ir.NewModule()
	f := declarePrintf(mod)
	assert.Eq(f.Def(), "declare i32 @printf(i8* %format, ...)")
}
