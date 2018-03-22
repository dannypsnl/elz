package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Stat interface {
	Codegen(*Context) llvm.Value
}

type GlobalVarDef struct {
	Export bool
	// let a = 1, a is Name
	Name string
	// let a: num = 1, num is VarType, but expression could not have the same type, we have to check it.
	VarType    string
	Expression Expr
}

func (varDef *GlobalVarDef) Codegen(ctx *Context) llvm.Value {
	// Parser should insert Type if user didn't define it.
	// So we should not get null string
	if varDef.VarType == "" || varDef.Expression.Type(ctx) != varDef.VarType {
		panic(`expr type != var type`)
	}
	expr := varDef.Expression.Codegen(ctx)
	val := llvm.AddGlobal(ctx.Module, expr.Type(), varDef.Name)
	val.SetInitializer(expr)
	ctx.GlobalVars[varDef.Name] = &VarNode{
		v:    expr,
		Type: varDef.VarType,
	}
	return val
}
