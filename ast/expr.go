package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

// Expr required method that an expression node have to implement
type Expr interface {
	// Codegen return a llvm.Value
	Codegen(*Context) llvm.Value
	// Type return a type info by string format.
	// It help elz's type system working with AST.
	Type(*Context) string
}

type UnaryExpr struct {
	E  Expr
	Op string
}

func (u *UnaryExpr) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstFNeg(u.E.Codegen(ctx))
}
func (u *UnaryExpr) Type(ctx *Context) string {
	return u.E.Type(ctx)
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
		// TODO: If have function implement by @Op, it can be a operator at here
		panic(`Unsupport this operator`)
	}
}

func (b *BinaryExpr) Type(ctx *Context) string {
	if b.LeftE.Type(ctx) == b.RightE.Type(ctx) {
		return b.LeftE.Type(ctx)
	} else {
		panic(`Type error`)
	}
}

type Argu struct {
	E Expr
}

func (a *Argu) Codegen(ctx *Context) llvm.Value {
	return a.E.Codegen(ctx)
}

func (a *Argu) Type(ctx *Context) string {
	return a.E.Type(ctx)
}

type FnCall struct {
	Name    string
	Args    []Argu
	RetType string
}

func (fc *FnCall) Type(*Context) string {
	return fc.RetType
}
