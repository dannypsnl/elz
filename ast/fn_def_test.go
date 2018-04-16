package ast

import (
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
	f.Codegen(context)
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
	f.Codegen(context)
}
