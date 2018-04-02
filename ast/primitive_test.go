package ast

import (
	"testing"
)

func TestI32(t *testing.T) {
	iVal := &I32{Val: "10"}
	llvmIR := iVal.Codegen(ctx)
	if !llvmIR.IsConstant() {
		t.Error("Bug in ast.I32")
	}
}
