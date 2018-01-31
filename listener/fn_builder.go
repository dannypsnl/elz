package listener

import (
	"fmt"
	"github.com/elz-lang/elz/ast"
)

type FnBuilder struct {
	export    bool
	name      string
	returnTyp string
	// TODO: How to record?
	// Type: Param = { Name, Type }
	params []*ast.Param
}

func NewFnBuilder() *FnBuilder {
	return &FnBuilder{}
}

func (fb *FnBuilder) Name(n string) *FnBuilder {
	fb.name = n
	return fb
}

func (fb *FnBuilder) Export(e bool) *FnBuilder {
	fb.export = e
	return fb
}

func (fb *FnBuilder) generate() *ast.FnDef {
	if fb.export {
		fmt.Print("public ")
	}
	// FIXME: This is let result show a fn, not correct impl
	fmt.Printf("fn %s\n", fb.name)
	return &ast.FnDef{
		Export: fb.export,
		Name:   fb.name,
		Params: []*ast.Param{},
		// TODO: implement statments
		// FIXME: should decide by rule typePass
		RetType: "num",
	}
	// TODO: local var def need spec_name
}
