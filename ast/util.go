package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func convertToLLVMType(t string) llvm.Type {
	switch t {
	case "num":
		return llvm.FloatType() // f32
	default:
		panic(`not support this type yet`)
	}
}
