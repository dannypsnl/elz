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
		t.Errorf("expected has: `%s`, but actual: `%s`", expected, context.Module.String())
	}
}

func TestFnDefParamMissingType(t *testing.T) {
	c := NewContext()
	f := &FnDef{
		Export: false,
		Name:   "add",
		Params: []*Param{
			&Param{Name: "lv", Type: ""},
			&Param{Name: "rv", Type: "i32"},
		},
		RetType: "i32",
	}

	f.Check(c)
	f.Codegen(c)

	actual := c.Module.String()
	expected := `define i32 @add(i32 %lv, i32 %rv)`

	if !strings.Contains(actual, expected) {
		t.Errorf("expected has: `%s`, but actual: `%s`", expected, actual)
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
}
`

	if !strings.Contains(context.Module.String(), expected) {
		t.Errorf("expected has: `%s`, but actual: `%s`", expected, context.Module.String())
	}
}
