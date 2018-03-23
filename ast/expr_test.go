package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

var ctx = NewContext()

func TestNumberAST(t *testing.T) {
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

func TestUnaryExprAST(t *testing.T) {
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

func TestBinaryExprAST(t *testing.T) {
	num := &Number{
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
	result := be.Codegen(ctx)
	if !result.IsConstant() ||
		result != llvm.ConstFloat(llvm.FloatType(), 3.69) {
		result.Dump()
		t.Error(`binary expression fail`)
	}
}

func TestStringAST(t *testing.T) {
	str := &Str{
		Val: `"a string with 中文"`,
	}
	result := str.Codegen(ctx)
	if !result.IsConstant() ||
		// string is [i8] in LLVM IR
		result.Type().String() != "ArrayType(IntegerType(8 bits)[22])" ||
		result != llvm.ConstString(`"a string with 中文"`, false) {
		result.Dump()
		t.Error(`error`)
	}
}
