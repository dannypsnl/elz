package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestNumber(t *testing.T) {
	ctx := &Context{
		module:  llvm.NewModule("main"),
		context: llvm.NewContext(),
		vars:    make(map[string]llvm.Value),
	}
	num := &Number{
		Val: "3.1415926",
	}

	result := num.Codegen(ctx)
	if result.IsNil() {
		result.Dump()
		t.Error(`error`)
	}
}
