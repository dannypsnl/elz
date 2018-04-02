package listener

import (
	"github.com/elz-lang/elz/ast"
)

type FnBuilder struct {
	export     bool
	name       string
	returnTyp  string
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

func (fb *FnBuilder) RetType(typ string) *FnBuilder {
	fb.returnTyp = typ
	return fb
}

func (fb *FnBuilder) Export(e bool) *FnBuilder {
	fb.export = e
	return fb
}

func (fb *FnBuilder) PushPType(typ string) *FnBuilder {
	fb.paramsType = append(fb.paramsType, typ)
	return fb
}

func (fb *FnBuilder) PushParamName(name string) *FnBuilder {
	fb.paramsName = append(fb.paramsName, name)
	return fb
}

func (fb *FnBuilder) Stat(s ast.Stat) *FnBuilder {
	fb.statments = append(fb.statments, s)
	return fb
}

func (fb *FnBuilder) generate() *ast.FnDef {
	params := []*ast.Param{}

	for i, name := range fb.paramsName {
		t := fb.paramsType[i]
		params = append(params, &ast.Param{name, t})
	}

	return &ast.FnDef{
		Export:  fb.export,
		Name:    fb.name,
		Params:  params,
		Body:    fb.statments,
		RetType: fb.returnTyp,
	}
}
