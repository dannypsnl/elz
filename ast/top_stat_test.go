package ast

import (
	"fmt"
	"llvm.org/llvm/bindings/go/llvm"
	"testing"
)

func TestVarDefination(t *testing.T) {
	v := &VarDefination{
		Immutable: true,
		Export:    false,
		Name:      "pi",
		VarType:   llvm.FloatType(),
		Expression: &BinaryExpr{
			&Number{"2.1215926"},
			&Number{"1.02"},
			"+",
		},
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
