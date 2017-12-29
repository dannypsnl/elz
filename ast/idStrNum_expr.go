package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Str struct {
	Val string
}

func (s *Str) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstString(s.Val, false)
}

func (s *Str) Type() string {
	return "str"
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
