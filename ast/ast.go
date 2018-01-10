package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

// We still need a type system represent. Not focus on llvm's type.
// Else the high level type will be hard to represent.

type Ast interface {
	Codegen(*Context) llvm.Value
}

type Stat interface {
	Codegen(*Context) llvm.Value
}

type StatList []Stat

type VarDefination struct {
	Immutable bool
	Export    bool
	// let a = 1, a is Name
	Name string
	// let a: num = 1, num is VarType, but expression could not have the same type, we have to check it.
	VarType    string
	Expression Expr
}

func (varDef *VarDefination) Codegen(ctx *Context) llvm.Value {
	// Parser should insert Type if user didn't define it.
	// So we should not get null string
	if varDef.VarType != "" && varDef.Expression.Type(ctx) == varDef.VarType {
		expr := varDef.Expression.Codegen(ctx)
		val := llvm.AddGlobal(ctx.Module, expr.Type(), varDef.Name)
		val.SetInitializer(expr)
		ctx.Vars[varDef.Name] = &VarNode{
			v:    expr,
			Type: varDef.VarType,
		}
		return val
	} else {
		panic(`expr type != var type`)
	}
}

type Param struct {
	Name string
	Type string
}
type FnDefination struct {
	Export  bool
	Name    string
	Params  []*Param
	Body    StatList
	RetType string
}

func (f *FnDefination) Codegen(ctx *Context) llvm.Value {
	var paramsT []llvm.Type
	for _, v := range f.Params {
		paramsT = append(paramsT, convertToLLVMType(v.Type))
	}
	retT := convertToLLVMType(f.RetType)
	ft := llvm.FunctionType(retT, paramsT, false)
	fn := llvm.AddFunction(ctx.Module, f.Name, ft)
	fBlock := llvm.AddBasicBlock(fn, "entry")
	ctx.Builder.SetInsertPointAtEnd(fBlock)
	// TODO: for _, stat := range f.Body {
	//ctx.Builder.Insert(stat.Codegen(ctx))
	//}
	ctx.Builder.CreateRet(llvm.ConstFloat(llvm.FloatType(), 3.14))
	ctx.Builder.ClearInsertionPoint()
	return fn
}
