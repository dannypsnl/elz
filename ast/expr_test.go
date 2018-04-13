package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestF32AST(t *testing.T) {
	context := NewContext()
	num := &F32{
		Val: "3.1415926",
	}
	result := num.Codegen(context)
	if !result.IsConstant() ||
		result.Type().String() != "FloatType" {
		result.Dump()
		t.Error(`error`)
	}
}

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

func TestStringAST(t *testing.T) {
	context := NewContext()
	str := &Str{
		Val: `"a string with 中文"`,
	}
	result := str.Codegen(context)
	if !result.IsConstant() ||
		// string is [i8] in LLVM IR
		result.Type().String() != "ArrayType(IntegerType(8 bits)[22])" ||
		result != llvm.ConstString(`"a string with 中文"`, false) {
		result.Dump()
		t.Error(`error`)
	}
}
