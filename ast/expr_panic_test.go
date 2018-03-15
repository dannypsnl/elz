package ast

import (
	"testing"
)

func TestUnsupportOpShouldCausePanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error(`Didn't panic`)
		}
	}()
	num := &Number{
		Val: "1.23",
	}
	be := &BinaryExpr{
		RightE: num,
		LeftE:  num,
		Op:     "~",
	}
	be.Codegen(ctx)
}
