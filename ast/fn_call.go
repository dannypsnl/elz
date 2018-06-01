package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

// FnCall is AST for code like:
//
// ```
// add(1, 2)
// ```
type FnCall struct {
	Name string
	Args []Expr
}

// Check each arguments in FnCall
func (fc *FnCall) Check(c *Context) {
	for _, at := range fc.Args {
		at.Check(c)
	}

}

// Codegen invoke Context Call by itself Name & Args
func (fc *FnCall) Codegen(c *Context) llvm.Value {
	return c.Call(fc.Name, fc.Args...)
}

// Type invoke c funcRetTyp & signature to get return type
func (fc *FnCall) Type(c *Context) string {
	return c.funcRetTyp(c.signature(fc.Name, fc.Args...)).retType
}
