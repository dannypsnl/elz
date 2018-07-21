package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/util"
)

// FnBuilder is function's building recorder, help we Generate function much more easier
type FnBuilder struct {
	export    bool
	name      string
	returnTyp string
	params    []*ast.Param
	statments []ast.Stat
	notations []util.Notation
}

// create a null FnBuilder
func NewFnBuilder() *FnBuilder {
	return &FnBuilder{}
}

// Name setup function name by n
func (fb *FnBuilder) Name(n string) *FnBuilder {
	fb.name = n
	return fb
}

func (fb *FnBuilder) RetType(typ string) *FnBuilder {
	fb.returnTyp = typ
	return fb
}

func (fb *FnBuilder) Notation(ns []util.Notation) *FnBuilder {
	fb.notations = ns
	return fb
}

// Export setup final function access level by e
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

func (fb *FnBuilder) PushStat(s ast.Stat) {
	fb.statments = append(fb.statments, s)
}

// Generate return the final AST of function
func (fb *FnBuilder) Generate(extern bool) *ast.FnDef {
	return &ast.FnDef{
		Export:    fb.export,
		Name:      fb.name,
		Params:    fb.params,
		Body:      fb.statments,
		RetType:   fb.returnTyp,
		Notations: fb.notations,
	}
}
