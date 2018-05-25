package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

// Array AST is a kind of Expr
//
// It is preparing for array literal
type Array struct {
	Elements []Expr
	Len      int

	elementType string
	dontCompile bool
}

func (a *Array) reportIfAnyElementDoNotMatchElementType(c *Context) {
	for _, e := range a.Elements {
		if e.Type(c) != a.elementType {
			a.dontCompile = true
			c.Reporter.Emit(
				fmt.Sprintf(
					"Array expected type: %s, but contains expression type: %s",
					a.elementType, e.Type(c),
				))
		}
	}
}

func (a *Array) Check(c *Context) {
	// Everyone should invoke it's sub node's Check first
	for _, e := range a.Elements {
		e.Check(c)
	}

	a.elementType = a.Elements[0].Type(c)

	a.reportIfAnyElementDoNotMatchElementType(c)
}

func (a *Array) Codegen(c *Context) llvm.Value {
	if a.dontCompile {
		return llvm.Value{}
	}

	values := make([]llvm.Value, 0)
	for i := 0; i < a.Len; i++ {
		var e Expr
		if len(a.Elements) == 1 {
			e = a.Elements[0]
		} else {
			e = a.Elements[i]
		}
		values = append(values, e.Codegen(c))
	}
	return llvm.ConstArray(c.Type(a.elementType), values)
}

func (a *Array) Type(*Context) string {
	return fmt.Sprintf("[%s;%d]", a.elementType, a.Len)
}
