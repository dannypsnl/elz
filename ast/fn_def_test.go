package ast

import (
	"strings"
	"testing"

	_ "llvm.org/llvm/bindings/go/llvm"
)

func TestFnDef(t *testing.T) {
	context := NewContext()
	f := &FnDef{
		Export: false,
		Name:   "add",
		Params: []*Param{
			&Param{
				Name: "lv",
				Type: "f32",
			},
			&Param{
				Name: "rv",
				Type: "f32",
			},
		},
		Body:    []Stat{},
		RetType: "f32",
	}
	f.Check(context)
	f.Codegen(context)

	expected := `define float @add(float %lv, float %rv)`

	if !strings.Contains(context.Module.String(), expected) {
		t.Errorf("expected has: %s, but actual: %s", expected, context.Module.String())
	}
}

func TestVarDefInFn(t *testing.T) {
	context := NewContext()

	f := &FnDef{
		Export: false,
		Name:   "Foo",
		Params: []*Param{},
		Body: []Stat{
			&LocalVarDef{
				Immutable:  true,
				Name:       "a",
				VarType:    "f32",
				Expression: &F32{Val: "3.14"},
			},
		},
		RetType: "f32",
	}
	f.Check(context)
	f.Codegen(context)

	expected := `
define float @Foo() {
entry:
  %a = alloca float
  store float 0x40091EB860000000, float* %a
  %a1 = load float, float* %a
}
`

	if !strings.Contains(context.Module.String(), expected) {
		t.Errorf("expected has: %s, but actual: %s", expected, context.Module.String())
	}
}
