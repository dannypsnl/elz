package listener

import (
	"github.com/elz-lang/elz/ast"
)

type FnBuilder struct {
	export    bool
	name      string
	returnTyp string
	params    []*ast.Param
	statments []ast.Stat
}

func NewFnBuilder() *FnBuilder {
	return &FnBuilder{}
}

func (fb *FnBuilder) Name(n string) *FnBuilder {
	fb.name = n
	return fb
}

func (fb *FnBuilder) RetType(typ string) *FnBuilder {
	fb.returnTyp = typ
	return fb
}

func (fb *FnBuilder) Export(e bool) *FnBuilder {
	fb.export = e
	return fb
}

func (fb *FnBuilder) PushParamType(typ string) *FnBuilder {
	l := len(fb.params)
	if l > 0 {
		fb.params[l-1].Type = typ
	}
	return fb
}

func (fb *FnBuilder) PushParamName(name string) *FnBuilder {
	fb.params = append(fb.params, &ast.Param{
		Name: name,
		Type: "",
	})
	return fb
}

func (fb *FnBuilder) Stat(s ast.Stat) *FnBuilder {
	fb.statments = append(fb.statments, s)
	return fb
}

func (fb *FnBuilder) generate() *ast.FnDef {
	return &ast.FnDef{
		Export:  fb.export,
		Name:    fb.name,
		Params:  fb.params,
		Body:    fb.statments,
		RetType: fb.returnTyp,
	}
}
