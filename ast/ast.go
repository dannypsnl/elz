package ast

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

var Module = ir.NewModule()
var f32 = types.Float

type Ast interface {
	Codegen()
}
type Expr interface {
	Codegen() constant.Constant
}
type Stat interface{}

type StatList []Stat

type Error struct {
	Msg string
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

type Number struct {
	Val string
}

func (n *Number) Codegen() constant.Constant {
	return constant.NewFloatFromString(n.Val, f32)
}

type UnaryExpr struct {
	E    Expr
	Op   string
	Type string
}

type BinaryExpr struct {
	LeftE  Expr
	RightE Expr
	Op     string
	Type   string
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
