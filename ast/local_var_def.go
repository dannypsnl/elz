package ast

import (
	"fmt"
	"llvm.org/llvm/bindings/go/llvm"
)

type LocalVarDef struct {
	Immutable      bool
	Name           string
	VarType        string
	Expression     Expr
	generateIsSafe bool
}

func (lv *LocalVarDef) Check(ctx *Context) {
	lv.Expression.Check(ctx)

	lv.generateIsSafe = true
	exprType := lv.Expression.Type(ctx)
	if lv.VarType == "" {
		lv.VarType = exprType
	}
	if lv.VarType != exprType {
		ctx.Reporter.Emit(fmt.Sprintf("var: %s, it's type is: %s, but receive: %s", lv.Name, lv.VarType, exprType))
		lv.generateIsSafe = false
	}
	ctx.NewVar(lv.Name, lv.VarType)
}

func (lv *LocalVarDef) Codegen(ctx *Context) llvm.Value {
	if lv.generateIsSafe {
		expr := lv.Expression.Codegen(ctx)
		typ := ctx.Type(lv.VarType)
		val := ctx.Builder.CreateAlloca(typ, lv.Name)
		ctx.Builder.CreateStore(expr, val)
		ctx.VarValue(lv.Name, val)
		return val
	}
	return llvm.Value{}
}
