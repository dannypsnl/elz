package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

var ctx = &Context{
	module:  llvm.NewModule("main"),
	context: llvm.NewContext(),
	vars:    make(map[string]llvm.Value),
}

func TestNumber(t *testing.T) {
	num := &Number{
		Val: "3.1415926",
	}
	result := num.Codegen(ctx)
	if !result.IsConstant() {
		result.Dump()
		t.Error(`error`)
	}
}

func TestUnaryExpr(t *testing.T) {
	num := &Number{
		Val: "1.23",
	}
	ub := &UnaryExpr{
		E:  num,
		Op: "-",
	}
	result := ub.Codegen(ctx)
	if !result.IsConstant() {
		result.Dump()
		t.Error(`unary expression fail`)
	}
}
