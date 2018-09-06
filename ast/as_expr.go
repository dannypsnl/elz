package ast

import (
	"fmt"
	"strconv"
	"strings"

	"llvm.org/llvm/bindings/go/llvm"
)

type As struct {
	E  Expr
	T  string
	op llvm.Opcode
}

func makeOp(exprType, toType string) llvm.Opcode {
	if exprType[0] == 'i' && toType[0] == 'i' {
		src := exprType[1:]
		from, err := strconv.ParseInt(src, 10, 32)
		panicIfErr(err, exprType, toType)
		dest := toType[1:]
		to, err := strconv.ParseInt(dest, 10, 32)
		panicIfErr(err, exprType, toType)
		if from > to {
			return llvm.Trunc
		} else {
			return llvm.SExt
		}
	}

	if exprType[0] == 'f' && toType[0] == 'f' {
		src := exprType[1:]
		from, err := strconv.ParseInt(src, 10, 32)
		panicIfErr(err, exprType, toType)
		dest := toType[1:]
		to, err := strconv.ParseInt(dest, 10, 32)
		panicIfErr(err, exprType, toType)
		if from > to {
			return llvm.FPTrunc
		} else {
			return llvm.FPExt
		}
	}

	// This part need to call function to complete
	panic(fmt.Sprintf("Not yet impl other as expr, %s, %s", exprType, toType))
}

func panicIfErr(err error, exprType, toType string) {
	if err != nil {
		panic(fmt.Sprintf("Not yet impl other as expr, %s, %s", exprType, toType))
	}
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
