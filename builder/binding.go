package builder

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (b *Builder) NewBinding(binding *ast.Binding) {
	b.bindings = append(b.bindings, binding)
}

func (b *Builder) ExitBinding(c *parser.BindingContext) {
	bindingTo := b.PopExpr().(ast.Expr)
	paramList := make([]string, 0)
	for _, paramName := range c.AllIDENT() {
		paramList = append(paramList, paramName.GetText())
	}
	b.NewBinding(&ast.Binding{
		Name:      c.IDENT(0).GetText(),
		ParamList: paramList[1:],
		Expr:      bindingTo,
	})
}
