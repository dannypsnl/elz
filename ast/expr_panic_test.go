package ast

import (
	"testing"
)

func TestUnsupportOpShouldCausePanic(t *testing.T) {
	ctx := NewContext()
	defer func() {
		if p := recover(); p == nil {
			t.Error(`Didn't panic`)
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
	be.Codegen(ctx)
}
