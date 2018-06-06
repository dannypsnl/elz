package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

type BinaryExpr struct {
	LeftE  Expr
	RightE Expr
	Op     string
}

func (b *BinaryExpr) Codegen(ctx *Context) llvm.Value {
	exprType := b.LeftE.Type(ctx)
	if exprType == "i32" || exprType == "i64" || exprType == "f32" || exprType == "f64" {
		return ctx.Call(b.Op, b.LeftE, b.RightE)
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
	if b.Op == "==" {
		return "bool"
	}
	return leftT
}
