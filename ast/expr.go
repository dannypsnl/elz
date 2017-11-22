package ast

import "github.com/llir/llvm/ir/constant"

type Expr interface {
	Codegen() constant.Constant
	Type() string
}

type Number struct {
	Val string
}

func (n *Number) Codegen() constant.Constant {
	return constant.NewFloatFromString(n.Val, f32)
}
func (n *Number) Type() string {
	return "num"
}

type UnaryExpr struct {
	E  Expr
	Op string
}

func (u *UnaryExpr) Codegen() constant.Constant {
	return constant.NewFSub(constant.NewFloatFromString("0", f32), u.E.Codegen())
}

func (u *UnaryExpr) Type() string {
	return u.E.Type()
}

type BinaryExpr struct {
	LeftE  Expr
	RightE Expr
	Op     string
}

func (b *BinaryExpr) Codegen() constant.Constant {
	switch b.Op {
	case "+":
		return constant.NewFAdd(b.LeftE.Codegen(), b.RightE.Codegen())
	case "-":
		return constant.NewFSub(b.LeftE.Codegen(), b.RightE.Codegen())
	case "*":
		return constant.NewFMul(b.LeftE.Codegen(), b.RightE.Codegen())
	case "/":
		return constant.NewFDiv(b.LeftE.Codegen(), b.RightE.Codegen())
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
