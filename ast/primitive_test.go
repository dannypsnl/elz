package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestI32(t *testing.T) {
	iVal := &I32{Val: "10"}
	llvmIR := iVal.Codegen(ctx)
	if llvmIR.Type() != llvm.Int32Type() {
		t.Error("Bug in ast.I32")
	}
}

func TestF32(t *testing.T) {
	fVal := &F32{Val: "3.14"}
	llvmIR := fVal.Codegen(ctx)
	if llvmIR.Type() != llvm.FloatType() {
		t.Error("Bug in ast.F32")
	}
}
