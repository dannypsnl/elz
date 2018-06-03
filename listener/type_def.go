package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

// EnterTypeDefine create a type define builder for following recording task
func (s *ElzListener) EnterTypeDefine(c *parser.TypeDefineContext) {
	s.typeDefineBuilder = new(TypeDefineBuilder)
}

// ExitAttr insert (export, field name, field type) infos into type define builder
func (s *ElzListener) ExitAttr(c *parser.AttrContext) {
	if s.typeDefineBuilder == nil {
		panic("compiler bug, type define builder must be not nil in rule attr")
	}
	s.typeDefineBuilder.newAttr(
		s.exportThis,
		c.ID().GetText(),
		c.TypeForm().GetText(),
	)
}

func (s *ElzListener) ExitTypeDefine(c *parser.TypeDefineContext) {
	if s.typeDefineBuilder == nil {
		panic("compiler bug, type define builder must be not nil in rule type define")
	}
	s.typeDefineBuilder.typeName = c.ID().GetText()
	s.AstList = append(s.AstList, s.typeDefineBuilder.generate())

	// clean up type define builder  after use
	s.typeDefineBuilder = nil
}

type TypeDefineBuilder struct {
	attrs    []ast.TypeAttr
	typeName string
}

func (t *TypeDefineBuilder) newAttr(exportOrNot bool, fieldName, fieldTyp string) {
	t.attrs = append(t.attrs, ast.TypeAttr{
		Export: exportOrNot,
		Name:   fieldName,
		Typ:    fieldTyp,
	})
}

func (t *TypeDefineBuilder) generate() *ast.TypeDef {
	return &ast.TypeDef{
		Name:  t.typeName,
		Attrs: t.attrs,
	}
}
