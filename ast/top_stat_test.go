package ast

import (
	"fmt"
	"testing"
)

func TestVarDefination(t *testing.T) {
	v := &VarDefination{
		Immutable:  true,
		Export:     false,
		Name:       "pi",
		VarType:    "num",
		Expression: &Number{"3.1415926"},
	}
	v.Codegen(ctx)
	if ctx.Vars["pi"].IsDeclaration() {
		t.Error(`error`)
	}
	fmt.Println(ctx.Module)
}
