package ast

import (
	"fmt"
	"llvm.org/llvm/bindings/go/llvm"
	"testing"
)

func TestVarDefination(t *testing.T) {
	v := &VarDefination{
		Immutable:  true,
		Export:     false,
		Name:       "pi",
		VarType:    llvm.FloatType(),
		Expression: &Number{"3.1415926"},
	}
	v.Codegen(ctx)
	if ctx.Vars["pi"].IsDeclaration() {
		t.Error(`error`)
	}
	fmt.Println(ctx.Module)
}
