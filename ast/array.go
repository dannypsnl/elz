package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

// Array AST is a kind of Expr
//
// It is preparing for array literal
type Array struct {
	Elements    []Expr
	ElementType string
	Len         int
	dontCompile bool
}

func (a *Array) Check(c *Context) {
	// Everyone should invoke it's sub node's Check first
	for _, e := range a.Elements {
		e.Check(c)
	}

	for _, e := range a.Elements {
		if e.Type(c) != a.ElementType {
			a.dontCompile = true
			c.Reporter.Emit(
				fmt.Sprintf(
					"Array expected type: %s, but contains expression type: %s",
					a.ElementType, e.Type(c),
				))
		}
	}
}

func (a *Array) Codegen(c *Context) llvm.Value {
	if a.dontCompile {
		return llvm.Value{}
	}

	values := make([]llvm.Value, 0)
	for _, e := range a.Elements {
		values = append(values, e.Codegen(c))
	}
	return llvm.ConstArray(LLVMType(a.ElementType), values)
}

func (a *Array) Type(*Context) string {
	return fmt.Sprintf("array<%s,%d>", a.ElementType, a.Len)
}
