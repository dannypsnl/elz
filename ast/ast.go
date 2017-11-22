package ast

import (
	"fmt"
)

type Ast interface {
	Codegen()
}

type Stat interface{}

type StatList []Stat

type Error struct {
	Msg string
}

func (e Error) Codegen() {
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

func (v *VarDefination) Codegen() {
	Module.NewGlobalDef(v.Name, v.Expression.Codegen())
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

type Argu struct {
	Val  string
	Type string
}
type FnCall struct {
	Name string
	Args []Argu
	Type string
}
