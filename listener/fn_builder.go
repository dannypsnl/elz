package listener

import (
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
	params := []*ast.Param{}
	return &ast.FnDef{
		Export: fb.export,
		Name:   fb.name,
		Params: params,
		Body:   fb.statments,
		// FIXME: should decide by rule typeFrom
		RetType: "num",
	}
}
