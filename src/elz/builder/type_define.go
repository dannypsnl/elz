package builder

import (
	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/parser"

	"github.com/sirupsen/logrus"
)

func (b *Builder) ExitTypeField(c *parser.TypeFieldContext) {
	typ := b.bindTypeList[0]
	b.bindTypeList = make([]ast.Type, 0)
	newField := ast.NewField(c.IDENT().GetText(), typ)
	b.typeDefineFields = append(
		b.typeDefineFields,
		newField,
	)
}

// ExitTypeDefineBody would be ignored now
func (b *Builder) ExitTypeDefineBody(c *parser.TypeDefineBodyContext) {
}

// ExitTypeDefine handle:
//
// ```
// type Car = (
//   name: str,
//   price: int,
// )
// ```
func (b *Builder) ExitTypeDefine(c *parser.TypeDefineContext) {
	newTypeName := c.IDENT().GetText()
	newTyp := &ast.NewType{
		Name:   newTypeName,
		Fields: b.typeDefineFields,
	}
	err := b.astTree.InsertTypeDefine(newTyp)
	if err != nil {
		logrus.Fatalf("stop parsing, error: %s", err)
	}
}
