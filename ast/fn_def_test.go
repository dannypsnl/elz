package ast

import (
	"testing"

	_ "llvm.org/llvm/bindings/go/llvm"
)

func TestFnDef(t *testing.T) {
	f := &FnDef{
		Export: false,
		Name:   "add",
		Params: []*Param{
			&Param{
				Name: "lv",
				Type: "num",
			},
			&Param{
				Name: "rv",
				Type: "num",
			},
		},
		Body:    []Stat{},
		RetType: "num",
	}
	f.Codegen(ctx)
}

func TestVarDefInFn(t *testing.T) {
	ctx := NewContext()

	f := &FnDef{
		Export: false,
		Name:   "Foo",
		Params: []*Param{},
		Body: []Stat{
			&LocalVarDef{
				Immutable:  true,
				Name:       "a",
				VarType:    "num",
				Expression: &Number{Val: "3.14"},
			},
		},
		RetType: "num",
	}
	f.Codegen(ctx)
	println(ctx.Module.String())
}
