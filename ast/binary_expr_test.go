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

func TestBinaryReportLeftRightTypeIsDifferent(t *testing.T) {
	c := NewContext()

	b := &BinaryExpr{
		RightE: &F32{Val: "3.134"},
		LeftE:  &I32{Val: "10"},
		Op:     "-",
	}

	b.Check(c)

	actual := c.Reporter.ErrMsgs[0]
	expected := `left expression type: i32, right expression type: f32`

	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
