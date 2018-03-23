package ast

import (
	"testing"

	_ "llvm.org/llvm/bindings/go/llvm"
)

func TestFnDef(t *testing.T) {
	f := FnDef{
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
