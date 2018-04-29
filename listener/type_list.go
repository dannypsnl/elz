package listener

import (
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) EnterTypeList(c *parser.TypeListContext) {
	if s.fnBuilder == nil {
		panic("Compiler bug, declaration expect have function builder when parsing")
	}
	for _, typ := range c.AllTypeForm() {
		s.fnBuilder.PushParamName(".wont_use")
		s.fnBuilder.PushParamType(typ.GetText())
	}
}
