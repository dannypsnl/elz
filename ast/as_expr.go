package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

type As struct {
	E  Expr
	T  string
	op llvm.Opcode
}

func makeOp(exprType, toType string) llvm.Opcode {
	if exprType == "i32" {
		switch toType {
		case "i64":
			return llvm.ZExt
		case "i8":
			fallthrough
		case "i16":
			return llvm.Trunc
		}
	} else if exprType == "i8" {
		switch toType {
		case "i16":
			fallthrough
		case "i32":
			fallthrough
		case "i64":
			return llvm.ZExt
		}
	} else if exprType == "i16" {
		switch toType {
		case "i8":
			return llvm.Trunc
		case "i32":
			fallthrough
		case "i64":
			return llvm.ZExt
		}
	} else if exprType == "f32" && toType == "f64" {
		return llvm.FPExt
	} else if exprType == "f64" && toType == "f32" {
		return llvm.FPTrunc
	} else if exprType == "i64" {
		switch toType {
		case "i32":
			fallthrough
		case "i16":
			fallthrough
		case "i8":
			return llvm.Trunc
		}
	}
	// This part need to call function to complete
	panic(fmt.Sprintf("Not yet impl other as expr, %s, %s", exprType, toType))
}

func (a *As) Check(ctx *Context) {
	a.E.Check(ctx)
	a.op = makeOp(a.E.Type(ctx), a.T)
}
func (a *As) Codegen(ctx *Context) llvm.Value {
	v := a.E.Codegen(ctx)
	return ctx.Builder.CreateCast(v, a.op, LLVMType(a.T), ".as_tmp")
}
func (a *As) Type(ctx *Context) string {
	return a.T
}
