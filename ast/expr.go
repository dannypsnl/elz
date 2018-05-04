package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

// Expr required method that an expression node have to implement
type Expr interface {
	// Codegen return a llvm.Value
	Codegen(*Context) llvm.Value
	// Type return a type info by string format.
	// It help elz's type system working with AST.
	Type(*Context) string
	Check(*Context)
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
	exprType := b.Type(ctx)
	if exprType == "i32" || exprType == "i64" {
		switch b.Op {
		case "+":
			return ctx.Builder.CreateAdd(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx), ".add_tmp")
		case "-":
			return ctx.Builder.CreateSub(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx), ".sub_tmp")
		case "*":
			return ctx.Builder.CreateMul(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx), ".mul_tmp")
		case "/":
			return ctx.Builder.CreateSDiv(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx), ".div_tmp")
		default:
			// FIXME: wait for impl
			// TODO: If have function implement by @Op, it can be a operator at here
			panic(`Unsupport this operator`)
		}
	} else if exprType == "f32" || exprType == "f64" {
		switch b.Op {
		case "+":
			return ctx.Builder.CreateFAdd(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx), ".fadd_tmp")
		case "-":
			return ctx.Builder.CreateFSub(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx), ".fsub_tmp")
		case "*":
			return ctx.Builder.CreateFMul(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx), ".fmul_tmp")
		case "/":
			return ctx.Builder.CreateFDiv(b.LeftE.Codegen(ctx), b.RightE.Codegen(ctx), ".fdiv_tmp")
		default:
			// FIXME: wait for impl
			// TODO: If have function implement by @Op, it can be a operator at here
			panic(`Unsupport this operator`)
		}
	} else {
		panic(fmt.Sprintf("BinaryExpr not support this type: %s yet", exprType))
	}
}

func (b *BinaryExpr) Check(ctx *Context) {
	leftT, rightT := b.LeftE.Type(ctx), b.RightE.Type(ctx)
	b.LeftE.Check(ctx)
	b.RightE.Check(ctx)
	if leftT != rightT {
		// TODO: If have function implement by @Op, it can be a operator at here

		// TODO: if can't find Op-function support this operation, error report
		// check `rightT != "type error"` for sure the error won't report again
		if rightT != "type error" {
			ctx.Reporter.Emit(fmt.Sprintf("left expression type: %s, right expression type: %s", leftT, rightT))
		}
	}
}

func (b *BinaryExpr) Type(ctx *Context) string {
	leftT, rightT := b.LeftE.Type(ctx), b.RightE.Type(ctx)
	if leftT != rightT {
		// TODO: If have function implement by @Op, it can be a operator at here

		// TODO: if can't find Op-function support this operation, error report
		// check `rightT != "type error"` for sure the error won't report again
		return "type error"
	}
	return leftT
}
