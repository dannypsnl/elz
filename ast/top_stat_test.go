package ast

import (
	"fmt"
	_ "llvm.org/llvm/bindings/go/llvm"
	"testing"
)

func TestVarDefination(t *testing.T) {
	v := &VarDefination{
		Immutable: true,
		Export:    false,
		Name:      "pi",
		VarType:   "num",
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
		Name:       "string1",
		VarType:    "str",
		Expression: str,
	}
	v.Codegen(ctx)
	if ctx.Vars["string1"].IsDeclaration() ||
		ctx.Vars["string1"].Type().String() != "PointerType(ArrayType(IntegerType(8 bits)[10]))" {
		t.Errorf("var: %s, expected: %s", ctx.Vars["string1"].Type().String(), "PointerType(ArrayType(IntegerType(8 bits)[10]))")
	}
	fmt.Println(ctx.Module)
}
