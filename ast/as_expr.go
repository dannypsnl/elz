package ast

import (
	"fmt"
	"strings"

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
			return llvm.SExt
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
			return llvm.SExt
		}
	} else if exprType == "i16" {
		switch toType {
		case "i8":
			return llvm.Trunc
		case "i32":
			fallthrough
		case "i64":
			return llvm.SExt
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
}
func (a *As) Codegen(c *Context) llvm.Value {
	v := a.E.Codegen(c)
	if isArrayType(a.E.Type(c)) && isRefType(a.T) {
		_, ok := a.E.(*Id)
		if elemType(a.E.Type(c)) == elemType(a.T) && ok {
			v, _ = c.Var(a.E.(*Id).Val)
			return c.Builder.CreateGEP(v, []llvm.Value{
				llvm.ConstInt(llvm.Int32Type(), 0, false),
				llvm.ConstInt(llvm.Int32Type(), 0, false),
			}, "")
		}
		return v // FIXME: tmp solution
	}
	a.op = makeOp(a.E.Type(c), a.T)
	return c.Builder.CreateCast(v, a.op, c.Type(a.T), ".as_tmp")
}
func (a *As) Type(ctx *Context) string {
	return a.T
}

func isArrayType(t string) bool {
	if t[0] == '[' && t[len(t)-1] == ']' {
		for _, r := range t {
			if r == ';' {
				return true
			}
		}
		return false
	}
	return false
}

func isRefType(t string) bool {
	if strings.HasPrefix(t, "ref<") && t[len(t)-1] == '>' && len(t) != 5 {
		return true
	}
	return false
}

func elemType(t string) string {
	if isArrayType(t) {
		for i, v := range t {
			if v == ';' {
				return string(t[1:i])
			}
		}
	} else if isRefType(t) {
		return t[4 : len(t)-1]
	}
	return "error"
}
