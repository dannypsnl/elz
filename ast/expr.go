package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Expr interface {
	Codegen(*Context) llvm.Value
	Type() string
}

type Number struct {
	Val string
}

func (n *Number) Codegen(*Context) llvm.Value {
	return llvm.ConstFloatFromString(llvm.FloatType(), n.Val)
}
func (n *Number) Type() string {
	return "num"
}

type UnaryExpr struct {
	E  Expr
	Op string
}

func (u *UnaryExpr) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstFNeg(u.E.Codegen(ctx))
}

func (u *UnaryExpr) Type() string {
	return u.E.Type()
}

type BinaryExpr struct {
	LeftE  Expr
	RightE Expr
	Op     string
}

func (b *BinaryExpr) Codegen(ctx *Context) llvm.Value {
	switch b.Op {
	case "+":
		return llvm.ConstFAdd(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx))
	case "-":
		return llvm.ConstFSub(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx))
	case "*":
		return llvm.ConstFMul(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx))
	case "/":
		return llvm.ConstFDiv(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx))
	default:
		// FIXME: wait for impl
		panic(`Unsupport this operator`)
	}
}

func (b *BinaryExpr) Type() string {
	if b.LeftE.Type() == b.RightE.Type() {
		return b.LeftE.Type()
	} else {
		panic(`Type error`)
	}
}

type Argu struct {
	E Expr
}

func (a *Argu) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstFloatFromString(llvm.FloatType(), "3.14")
}

func (a *Argu) Type() string {
	return a.E.Type()
}

type FnCall struct {
	Name    string
	Args    []Argu
	RetType string
}

func (fc *FnCall) Type() string {
	return fc.RetType
}

type Str struct {
	Val string
}

func (s *Str) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstString(s.Val, false)
}
func (s *Str) Type() string {
	return "str"
}
