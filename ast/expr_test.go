package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestUnaryExprAST(t *testing.T) {
	context := NewContext()
	num := &F32{
		Val: "1.23",
	}
	ub := &UnaryExpr{
		E:  num,
		Op: "-",
	}
	result := ub.Codegen(context)
	if !result.IsConstant() {
		result.Dump()
		t.Error(`unary expression fail`)
	}
}

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
