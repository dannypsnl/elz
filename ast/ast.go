package ast

import (
	"fmt"
	"llvm.org/llvm/bindings/go/llvm"
)

type Context struct {
	module  llvm.Module
	context llvm.Context
	vars    map[string]llvm.Value
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
