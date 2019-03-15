package codegen

import "github.com/elz-lang/elz/src/elz/ast"

type Binding struct {
	*ast.Binding
}

func NewBinding(bind *ast.Binding) *Binding {
	return &Binding{
		Binding: bind,
	}
}
