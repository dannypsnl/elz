package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

func LLVMType(t string) llvm.Type {
	switch t {
	case "()":
		return llvm.VoidType()
	case "i8":
		return llvm.Int8Type()
	case "i16":
		return llvm.Int16Type()
	case "i32":
		return llvm.Int32Type()
	case "i64":
		return llvm.Int64Type()
	case "f32":
		return llvm.FloatType() // f32
	case "f64":
		return llvm.DoubleType() // f64
	default:
		if len(t) == 3 && string(t[:3]) == "any" {
			goto end
		}
		if len(t) > 4 && string(t[:4]) == "ref<" && string(t[len(t)-1:]) == ">" {
			return llvm.PointerType(LLVMType(t[4:len(t)-1]), 0)
		}
		if len(t) > 3 && string(t[:3]) == "..." {
			return LLVMType(t[3:])
		}
	end:
		panic(fmt.Sprintf("not support type: `%s` yet", t))
	}
}
