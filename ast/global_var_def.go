package ast

import (
	"fmt"
	"llvm.org/llvm/bindings/go/llvm"
)

type GlobalVarDef struct {
	Export bool
	// a = 1, a is Name
	Name string
	// a: num = 1, num is VarType, but expression could not have the same type, we have to check it.
	VarType    string
	Expression Expr
}

func (g *GlobalVarDef) Check(ctx *Context) {
	g.Expression.Check(ctx)

	if g.VarType == "" {
		g.VarType = g.Expression.Type(ctx)
	}
	exprType := g.Expression.Type(ctx)
	if g.VarType != exprType {
		ctx.Reporter.Emit(fmt.Sprintf("global var: %s, it's type is: %s, but receive: %s", g.Name, g.VarType, exprType))
	}
	ctx.VarsType[g.Name] = g.VarType
}

func (varDef *GlobalVarDef) Codegen(ctx *Context) llvm.Value {
	expr := varDef.Expression.Codegen(ctx)
	val := llvm.AddGlobal(ctx.Module, expr.Type(), varDef.Name)
	val.SetInitializer(expr)
	ctx.Vars["&"+varDef.Name] = val
	ctx.Vars[varDef.Name] = expr // Let's use load to access all var
	return val
}
