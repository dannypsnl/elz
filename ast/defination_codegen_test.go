package ast

import (
	"testing"

	"github.com/llir/llvm/ir"
)

func TestVarDefine(t *testing.T) {
	m := ir.NewModule()
	v := &VarDefination{
		Immutable: false,
		Export:    false,
		Name:      "acb",
		VarType:   "num",
		Expression: &BinaryExpr{
			LeftE: &BinaryExpr{
				&Number{"2.4"},
				&Number{"3.1"},
				"*",
			},
			RightE: &Number{"3.3"},
			Op:     "+",
		},
	}
	v.Codegen(m)
	expected := `@acb = global double fadd (double fmul (double 2.4, double 3.1), double 3.3)
`
	if m.String() != expected {
		t.Errorf("expected: '%s', actual: '%s'", expected, m.String())
	}
}
