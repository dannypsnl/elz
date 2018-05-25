package listener

import (
	"github.com/elz-lang/elz/parser"
	"github.com/elz-lang/elz/util"
)

func (s *ElzListener) EnterDeclareFn(c *parser.DeclareFnContext) {
	retTyp := ""
	if !c.ReturnType().IsEmpty() {
		retTyp = c.ReturnType().GetText()
	}
	s.fnBuilder = NewFnBuilder().
		Name(c.ID().GetText()).
		RetType(retTyp).
		Notation(s.notations)

	// take all notation in record
	s.notations = make([]util.Notation, 0)
}
func (s *ElzListener) ExitDeclareFn(c *parser.DeclareFnContext) {
	if s.fnBuilder == nil {
		panic("Compiler bug, declaration expect have function builder when parsing")
	}

	typ := ""
	escapeLevel := 0
	for _, r := range c.TypeList().GetText() {
		typ += string(r)
		if r == '<' {
			escapeLevel++
		} else if r == '>' {
			escapeLevel--
		}
		if escapeLevel == 0 && r == ',' {
			s.fnBuilder.PushParamName("")
			s.fnBuilder.PushParamType(typ[:len(typ)-1])
			typ = ""
		}
	}
	s.fnBuilder.PushParamName("")
	s.fnBuilder.PushParamType(typ)

	s.AstList = append(s.AstList,
		// generate extern declaration
		s.fnBuilder.generate(true),
	)
	s.fnBuilder = nil
}
