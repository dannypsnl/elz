package listener

import (
	"github.com/elz-lang/elz/parser"
	"github.com/elz-lang/elz/util"
)

func (s *ElzListener) EnterCompilerNotation(c *parser.CompilerNotationContext) {
	nota := util.Notation{
		Leading: c.ID(0).GetText(),
		Content: make([]string, 0),
	}
	content := c.AllID()
	if len(content) >= 2 {
		for _, c := range content[1:] {
			nota.Content = append(nota.Content, c.GetText())
		}
	}
	s.notations = append(s.notations, nota)
}
