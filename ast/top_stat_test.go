package ast

import (
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
	if ctx.Vars["pi"].v.Type().String() != "FloatType" {
		t.Error(`error`)
	}
}

func TestStrVarDef(t *testing.T) {
	str := &Str{`"a string"`}
	v := &VarDefination{
		Immutable:  true,
		Export:     false,
		Name:       "string1",
		VarType:    "str",
		Expression: str,
	}
	v.Codegen(ctx)
	if ctx.Vars["string1"].v.Type().String() != "ArrayType(IntegerType(8 bits)[10])" {
		t.Errorf("var: %s, expected: %s", ctx.Vars["string1"].v.Type().String(),
			"ArrayType(IntegerType(8 bits)[10])")
	}

}
