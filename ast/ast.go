package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

// We still need a type system represent. Not focus on llvm's type.
// Else the high level type will be hard to represent.

type Context struct {
	Module  llvm.Module
	Context llvm.Context
	Vars    map[string]llvm.Value
}

type Ast interface {
	Codegen(*Context)
}

type Stat interface{}

type StatList []Stat

type Error struct {
	Msg string
}

func (e Error) Codegen(*Context) {
	// FIXME: position is hard coding
	fmt.Printf("At(0,0), error: %s\n", e.Msg)
}

type VarDefination struct {
	Immutable bool
	Export    bool
	// let a = 1, a is Name
	Name string
	// let a: num = 1, num is VarType, but expression could not have the same type, we have to check it.
	VarType    string
	Expression Expr
}

func (v *VarDefination) Codegen(ctx *Context) {
	expr := v.Expression.Codegen(ctx)
	if v.VarType != "" && v.Expression.Type() == v.VarType {
		val := llvm.AddGlobal(ctx.Module, expr.Type(), v.Name)
		val.SetInitializer(expr)
		ctx.Vars[v.Name] = val
	} else {
		panic(`expr type != var type`)
	}
}

type Param struct {
	Name string
	Type llvm.Type
}
type FnDefination struct {
	Export  bool
	Name    string
	Params  []Param
	Body    StatList
	RetType llvm.Type
}

func (f *FnDefination) Codegen(ctx *Context) {
	var paramsT []llvm.Type
	for _, v := range f.Params {
		paramsT = append(paramsT, v.Type)
	}
	ft := llvm.FunctionType(f.RetType, paramsT, false)
	llvm.AddFunction(ctx.Module, f.Name, ft)
}
