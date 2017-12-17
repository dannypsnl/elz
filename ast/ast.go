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
	Immutable  bool
	Export     bool
	Name       string
	VarType    llvm.Type
	Expression Expr
}

func (v *VarDefination) Codegen(ctx *Context) {
	val := llvm.AddGlobal(ctx.Module, v.VarType, v.Name)
	val.SetInitializer(v.Expression.Codegen(ctx))
	ctx.Vars[v.Name] = val
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
