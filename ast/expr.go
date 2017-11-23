package ast

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
)

type Expr interface {
	Codegen(*ir.Module) constant.Constant
	Type() string
}

type Number struct {
	Val string
}

func (n *Number) Codegen(*ir.Module) constant.Constant {
	return constant.NewFloatFromString(n.Val, f64)
}
func (n *Number) Type() string {
	return "num"
}

type UnaryExpr struct {
	E  Expr
	Op string
}

func (u *UnaryExpr) Codegen(m *ir.Module) constant.Constant {
	return constant.NewFSub(constant.NewFloatFromString("0", f64), u.E.Codegen(m))
}

func (u *UnaryExpr) Type() string {
	return u.E.Type()
}

type BinaryExpr struct {
	LeftE  Expr
	RightE Expr
	Op     string
}

func (b *BinaryExpr) Codegen(m *ir.Module) constant.Constant {
	switch b.Op {
	case "+":
		return constant.NewFAdd(b.LeftE.Codegen(m), b.RightE.Codegen(m))
	case "-":
		return constant.NewFSub(b.LeftE.Codegen(m), b.RightE.Codegen(m))
	case "*":
		return constant.NewFMul(b.LeftE.Codegen(m), b.RightE.Codegen(m))
	case "/":
		return constant.NewFDiv(b.LeftE.Codegen(m), b.RightE.Codegen(m))
	default:
		panic(`Unsupport op`)
	}
}

func (b *BinaryExpr) Type() string {
	if b.LeftE.Type() == b.RightE.Type() {
		return b.LeftE.Type()
	} else {
		panic(`Type error`)
	}
}
