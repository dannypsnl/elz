package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

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
	VarType    string
	Expression Expr
}

func (v *VarDefination) Codegen(ctx *Context) {
	val := llvm.AddGlobal(ctx.Module, llvm.FloatType(), v.Name)
	val.SetInitializer(
		llvm.ConstFloat(llvm.FloatType(), 3.14),
	//	v.Expression.Codegen(ctx)
	)
	ctx.Vars[v.Name] = val
}

type Param struct {
	Name string
	Type string
}
type FnDefination struct {
	Export bool
	Name   string
	Params []Param
	Body   StatList
}
