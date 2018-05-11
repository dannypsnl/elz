package ast

import (
	"testing"
)

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func TestBinaryExprAST(t *testing.T) {
	context := NewContext()
	num := &F32{
		Val: "1.23",
	}
	be0 := &BinaryExpr{
		RightE: num,
		LeftE:  num,
		Op:     "+",
	}
	be := &BinaryExpr{
		RightE: num,
		LeftE:  be0,
		Op:     "+",
	}
	result := be.Codegen(context)
	if !result.IsConstant() ||
		result != llvm.ConstFloat(llvm.FloatType(), 3.69) {
		result.Dump()
		t.Error(`binary expression fail`)
	}
}

func TestBinaryExprWithUnsupportOpShouldPanic(t *testing.T) {
	context := NewContext()
	defer func() {
		if p := recover(); p == nil {
			t.Error(`Should panic but not`)
		}
	}()
	num := &F32{
		Val: "1.23",
	}
	be := &BinaryExpr{
		RightE: num,
		LeftE:  num,
		Op:     "~",
	}
	be.Codegen(context)
}
