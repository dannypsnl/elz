package ast

import (
	"testing"
)

func TestStr(t *testing.T) {
	context := NewContext()
	str := &Str{Val: "Hello"}
	str.Check(context)
	ir := str.Codegen(context)
	if str.Type(context) != "str" {
		t.Error("str::Type return type is wrong")
	}
	if ir.Type().String() != "ArrayType(IntegerType(8 bits)[5])" {
		println(ir.Type().String())
		t.Error("string ir type is incorrect")
	}
}

func TestI8(t *testing.T) {
	context := NewContext()
	iVal := &I8{Val: "10"}
	iVal.Check(context)
	llvmIR := iVal.Codegen(context)
	if llvmIR.Type() != LLVMType(iVal.Type(context)) {
		t.Error("Bug in ast.I8")
	}
}

func TestI16(t *testing.T) {
	context := NewContext()
	iVal := &I16{Val: "10"}
	iVal.Check(context)
	llvmIR := iVal.Codegen(context)
	if llvmIR.Type() != LLVMType(iVal.Type(context)) {
		t.Error("Bug in ast.I16")
	}
}

func TestI32(t *testing.T) {
	context := NewContext()
	iVal := &I32{Val: "10"}
	iVal.Check(context)
	llvmIR := iVal.Codegen(context)
	if llvmIR.Type() != LLVMType(iVal.Type(context)) {
		t.Error("Bug in ast.I32")
	}
}

func TestI64(t *testing.T) {
	context := NewContext()
	iVal := &I64{Val: "10"}
	iVal.Check(context)
	llvmIR := iVal.Codegen(context)
	if llvmIR.Type() != LLVMType(iVal.Type(context)) {
		t.Error("Bug in ast.I64")
	}
}

func TestF32(t *testing.T) {
	context := NewContext()
	fVal := &F32{Val: "3.14"}
	fVal.Check(context)
	llvmIR := fVal.Codegen(context)
	if llvmIR.Type() != LLVMType(fVal.Type(context)) {
		t.Error("Bug in ast.F32")
	}
}

func TestF64(t *testing.T) {
	context := NewContext()
	fVal := &F64{Val: "3.14"}
	fVal.Check(context)
	llvmIR := fVal.Codegen(context)
	if llvmIR.Type() != LLVMType(fVal.Type(context)) {
		t.Error("Bug in ast.F64")
	}
}
