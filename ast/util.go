package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

func LLVMType(t string) llvm.Type {
	switch t {
	case "()":
		return llvm.VoidType()
	case "i32":
		return llvm.Int32Type()
	case "i64":
		return llvm.Int64Type()
	case "f32":
		return llvm.FloatType() // f32
	case "f64":
		return llvm.DoubleType() // f64
	default:
		panic(fmt.Sprintf("not support type: `%s` yet", t))
	}
}
