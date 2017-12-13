package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

var ctx = &Context{
	Module:  llvm.NewModule("main"),
	Context: llvm.NewContext(),
	Vars:    make(map[string]llvm.Value),
}

func TestNumber(t *testing.T) {
	num := &Number{
		Val: "3.1415926",
	}
	result := num.Codegen(ctx)
	if !result.IsConstant() ||
		result.Type().String() != "FloatType" {
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

func TestBinaryExpr(t *testing.T) {
	num := &Number{
		Val: "1.23",
	}
	be := &BinaryExpr{
		RightE: num,
		LeftE:  num,
		Op:     "+",
	}
	result := be.Codegen(ctx)
	if !result.IsConstant() {
		result.Dump()
		t.Error(`binary expression fail`)
	}
}

func TestBinaryForUnsupportOp(t *testing.T) {
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

func TestString(t *testing.T) {
	str := &Str{
		Val: `"a string with 中文"`,
	}
	result := str.Codegen(ctx)
	if !result.IsConstant() ||
		// string is [i8] in LLVM IR
		result.Type().String() != "ArrayType(IntegerType(8 bits)[22])" {
		result.Dump()
		t.Error(`error`)
	}
}
