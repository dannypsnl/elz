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
	if ctx.Vars["pi"].v.IsDeclaration() ||
		ctx.Vars["pi"].v.Type().String() != "PointerType(FloatType)" {
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
	if ctx.Vars["string1"].v.IsDeclaration() ||
		ctx.Vars["string1"].v.Type().String() != "PointerType(ArrayType(IntegerType(8 bits)[10]))" {
		t.Errorf("var: %s, expected: %s", ctx.Vars["string1"].v.Type().String(), "PointerType(ArrayType(IntegerType(8 bits)[10]))")
	}
	fn := &FnDefination{
		Export: true,
		Name:   "foo",
		Params: []*Param{
			&Param{"x", "num"},
			&Param{"y", "num"},
		},
		Body:    StatList{v},
		RetType: "num",
	}
	fn.Codegen(ctx)
	fmt.Println(ctx.Module)
}

func TestFnDef(t *testing.T) {
	fn := &FnDefination{
		Name:    "t",
		RetType: "num",
	}
	fn.Codegen(ctx)
	fmt.Println(ctx.Module)
}
