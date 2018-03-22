package ast

import (
	"testing"
)

func TestGlobalVarDef(t *testing.T) {
	v := &GlobalVarDef{
		Export:  false,
		Name:    "pi",
		VarType: "num",
		Expression: &BinaryExpr{
			&Number{"2.1215926"},
			&Number{"1.02"},
			"+",
		},
	}
	v.Codegen(ctx)
	if ctx.GlobalVars["pi"].v.Type().String() != "FloatType" {
		t.Error(`error`)
	}
}

func TestStrGlobalVarDef(t *testing.T) {
	str := &Str{`"a string"`}
	v := &GlobalVarDef{
		Export:     false,
		Name:       "string1",
		VarType:    "str",
		Expression: str,
	}
	v.Codegen(ctx)
	if ctx.GlobalVars["string1"].v.Type().String() != "ArrayType(IntegerType(8 bits)[10])" {
		t.Errorf("var: %s, expected: %s", ctx.GlobalVars["string1"].v.Type().String(),
			"ArrayType(IntegerType(8 bits)[10])")
	}

}
