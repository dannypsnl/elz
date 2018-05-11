package ast

import (
	"testing"
)

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
