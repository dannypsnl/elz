package listener

import (
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) EnterDeclareFn(c *parser.DeclareFnContext) {
	retTyp := ""
	if c.ReturnType() != nil {
		retTyp = c.ReturnType().GetText()
	}
	fnBuilder := NewFnBuilder().
		Name(c.ID().GetText()).
		RetType(retTyp).
		Notation(s.TakeAllNotation())
	s.statBuilder.Push(fnBuilder)
}
func (s *ElzListener) ExitDeclareFn(c *parser.DeclareFnContext) {
	if fnBuilder, ok := s.statBuilder.Last().(*FnBuilder); !ok {
		panic("Compiler bug, function define expect have function builder when parsing")
	} else {
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
				fnBuilder.PushParamName("")
				fnBuilder.PushParamType(typ[:len(typ)-1])
				typ = ""
			}
		}
		fnBuilder.PushParamName("")
		fnBuilder.PushParamType(typ)

		s.AstList = append(s.AstList,
			// Generate extern declaration
			fnBuilder.Generate(true),
		)
		s.statBuilder.Pop()
	}
}
