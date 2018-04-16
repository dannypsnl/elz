package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

// We still need a type system represent. Not focus on llvm's type.
// Else the high level type will be hard to represent.

type Ast interface {
	Codegen(*Context) llvm.Value
	Check(*Context)
}
