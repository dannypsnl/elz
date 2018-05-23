package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

type TypeDefineBuilder struct {
	attrs    []ast.TypeAttr
	typeName string
}

func (t *TypeDefineBuilder) generate() *ast.TypeDef {
	&ast.TypeDef{
		Name:  t.typeName,
		Attrs: t.attrs,
	}
}

func (s *ElzListener) EnterTypeDefine(c *parser.TypeDefineContext) {
	s.typeDefineBuilder = new(TypeDefineBuilder)
}

func (s *ElzListener) ExitAttr(c *parser.AttrContext) {
	if s.typeDefineBuilder == nil {
		panic("compiler bug, type define builder must be not nil in rule attr")
	}
	attr := ast.TypeAttr{
		Export: s.exportThis,
		Name:   c.ID().GetText(),
		Type:   c.TypeForm().GetText(),
	}
	s.typeDefineBuilder.attrs = append(s.typeDefineBuilder.attrs, attr)
}

func (s *ElzListener) ExitTypeDefine(c *parser.TypeDefineContext) {
	if s.typeDefineBuilder == nil {
		panic("compiler bug, type define builder must be not nil in rule type define")
	}

	s.typeDefineBuilder.typeName = c.ID().GetText()

	s.AstList = append(s.AstList, s.typeDefineBuilder.generate())

	// clean up after use
	s.typeDefineBuilder = nil
}
