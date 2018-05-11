package ast

import (
	"testing"
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
