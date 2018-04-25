package ast

import (
	"testing"
)

func TestI32(t *testing.T) {
	context := NewContext()
	iVal := &I32{Val: "10"}
	iVal.Check(context)
	llvmIR := iVal.Codegen(context)
	if llvmIR.Type() != LLVMType(iVal.Type(context)) {
		t.Error("Bug in ast.I32")
	}
}

func TestF32(t *testing.T) {
	context := NewContext()
	fVal := &F32{Val: "3.14"}
	fVal.Check(context)
	llvmIR := fVal.Codegen(context)
	if llvmIR.Type() != LLVMType(fVal.Type(context)) {
		t.Error("Bug in ast.F32")
	}
}
