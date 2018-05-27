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

	ctx.NewVar(g.Name, g.VarType)
}

func (g *GlobalVarDef) Codegen(ctx *Context) llvm.Value {
	expr := g.Expression.Codegen(ctx)
	val := llvm.AddGlobal(ctx.Module, expr.Type(), g.Name)
	if expr.IsConstant() {
		val.SetInitializer(expr)
	}
	ctx.Vars[g.Name] = val
	ctx.VarValue(g.Name, val)
	return val
}
