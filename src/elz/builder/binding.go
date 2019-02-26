package builder

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/parser"
)

func (b *Builder) ExitBindType(c *parser.BindTypeContext) {
	bindName := c.IDENT().GetText()
	// FIXME: nil dereference
	fmt.Printf("new bind spec type of %s: %#v\n", bindName, b.bindTypeList)
	b.bindTypeList = make([]ast.Type, 0)
}

// ExitExistType listen format: `int` represents existing type
// You would get compile fatal if the type isn't existed
func (b *Builder) ExitExistType(c *parser.ExistTypeContext) {
	b.bindTypeList = append(b.bindTypeList, &ast.ExistType{Name: c.IDENT().GetText()})
}

func (b *Builder) ExitVoidType(c *parser.VoidTypeContext) {
	b.bindTypeList = append(b.bindTypeList, &ast.VoidType{})
}

// ExitVariantType listen format: `'a` as type hole
func (b *Builder) ExitVariantType(c *parser.VariantTypeContext) {
	b.bindTypeList = append(b.bindTypeList, &ast.VariantType{Name: c.IDENT().GetText()})
}

// ExitCombineType listen format: `int -> int -> int`
func (b *Builder) ExitCombineType(c *parser.CombineTypeContext) {
	// ignore, just help we know has this syntax
}

func (b *Builder) ExitBinding(c *parser.BindingContext) {
	bindingTo := b.PopExpr().(ast.Expr)
	paramList := make([]string, 0)
	for _, paramName := range c.AllIDENT() {
		paramList = append(paramList, paramName.GetText())
	}
	err := b.astTree.InsertBinding(&ast.Binding{
		Name:      c.IDENT(0).GetText(),
		ParamList: paramList[1:],
		Expr:      bindingTo,
	})
	if err != nil {
		err := fmt.Errorf("stop parsing, error: %s", err)
		panic(err)
	}
}
