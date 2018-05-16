package researchllvm

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func assert(t *testing.T, r bool) {
	if !r {
		t.Error("expected true")
	}
}

func TestIsConstant(t *testing.T) {
	var v llvm.Value
	v = llvm.ConstInt(llvm.Int32Type(), 10, true) // i32
	assert(t, v.IsConstant())
	v = llvm.ConstInt(llvm.Int32Type(), 10, false) // u32
	assert(t, v.IsConstant())
	b := llvm.NewBuilder()
	v = b.CreateAdd(v, v, "")
	assert(t, v.IsConstant())
	v = b.CreateSub(v, v, "")
	assert(t, v.IsConstant())
	v = b.CreateMul(v, v, "")
	assert(t, v.IsConstant())
	v = b.CreateSDiv(v, v, "")
	assert(t, v.IsConstant())
}
