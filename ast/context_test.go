package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestContextType(t *testing.T) {
	c := NewContext()

	typ := c.Type("i32")

	if typ != llvm.Int32Type() {
		t.Error("context type implementation has bug")
	}
}
