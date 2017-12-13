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
	if ctx.Vars["pi"].IsDeclaration() ||
		ctx.Vars["pi"].Type().String() != "PointerType(FloatType)" {
		t.Error(`error`)
	}
	str := &Str{`"a string"`}
	v = &VarDefination{
		Immutable:  true,
		Export:     false,
		Name:       "str",
		VarType:    str.Type(),
		Expression: str,
	}
	v.Codegen(ctx)
	if ctx.Vars["str"].IsDeclaration() ||
		ctx.Vars["str"].Type().String() != "PointerType("+str.Type().String()+")" {
		t.Error(`error`)
	}
	fmt.Println(ctx.Module)
}
