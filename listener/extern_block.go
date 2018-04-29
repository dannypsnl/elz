package listener

import (
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) EnterExternBlock(c *parser.ExternBlockContext) {
	// support only C ffi now
	if c.STRING().GetText() == `"C"` {
		s.inExternBlock = true
	} else {
		s.context.Reporter.Emit("support extern C only now")
	}
}

func (s *ElzListener) ExitExternBlock(c *parser.ExternBlockContext) {
	if !s.inExternBlock {
		panic("Some code broke the extern block expect, inExternBlock must be true when exit externBlock rule")
	}
	s.inExternBlock = false
}

func (s *ElzListener) EnterDeclareFn(c *parser.DeclareFnContext) {
	s.fnBuilder = NewFnBuilder().
		Name(c.ID().GetText()).
		RetType(c.ReturnType().GetText())
}
func (s *ElzListener) ExitDeclareFn(c *parser.DeclareFnContext) {
	if s.fnBuilder == nil {
		panic("Compiler bug, declaration expect have function builder when parsing")
	}
	s.AstList = append(s.AstList,
		// generate extern declaration
		s.fnBuilder.generate(true),
	)
	s.fnBuilder = nil
}
