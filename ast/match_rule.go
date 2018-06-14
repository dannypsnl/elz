package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Match struct{}

func (m *Match) Check(*Context) {
}

func (m *Match) Codegen(*Context) llvm.Value {
	return llvm.Value{}
}

func (m *Match) Type(*Context) string {
	return "missing type now"
}
