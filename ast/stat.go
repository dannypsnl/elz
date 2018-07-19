package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Stat interface {
	Codegen(*Context) llvm.Value
	Check(*Context)
}

type ExprStat interface {
	Stat
	ExprStat()
}
