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
		Expression: &UnaryExpr{
			E:  &Number{"1.2"},
			Op: "-",
		},
	}
	v.Codegen(m)
	expected := `@acb = global float fsub (float 0.0, float 1.2)
`
	if m.String() != expected {
		t.Errorf("expected: '%s', actual: '%s'", expected, m.String())
	}
}
