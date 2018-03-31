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
	paramsName []string
	paramsType []string
	statments  []ast.Stat
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

func (fb *FnBuilder) PushPType(typ string) *FnBuilder {
	return fb
}

func (fb *FnBuilder) PushParamName(name string) *FnBuilder {
	return fb
}

func (fb *FnBuilder) Stat(s ast.Stat) *FnBuilder {
	fb.statments = append(fb.statments, s)
	return fb
}

func (fb *FnBuilder) generate() *ast.FnDef {
	if fb.export {
		fmt.Print("public ")
	}
	params := []*ast.Param{}
	// FIXME: This is let result show a fn, not correct impl
	fmt.Printf("fn %s\n", fb.name)
	return &ast.FnDef{
		Export: fb.export,
		Name:   fb.name,
		Params: params,
		// TODO: implement statments
		Body: fb.statments,
		// FIXME: should decide by rule typePass
		RetType: "num",
	}
	// TODO: local var def need spec_name
}
