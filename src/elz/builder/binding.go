package builder

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/parser"

	"github.com/sirupsen/logrus"
)

func (b *Builder) ExitBindType(c *parser.BindTypeContext) {
	bindName := c.IDENT().GetText()
	_, exist := b.bindingType[bindName]
	if exist {
		logrus.Fatalf("bind type existed")
	}
	b.bindingType[bindName] = b.bindTypeList
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
	bindName := c.IDENT(0).GetText()
	export := c.KEYWORD_EXPORT() != nil
	binding := &ast.Binding{
		Export:    export,
		Name:      bindName,
		ParamList: paramList[1:],
		Expr:      bindingTo,
	}
	if t, exist := b.bindingType[bindName]; exist {
		binding.Type = t
	}
	err := b.astTree.InsertBinding(binding)
	if err != nil {
		err := fmt.Errorf("stop parsing, error: %s", err)
		logrus.Fatal(err)
	}
}
