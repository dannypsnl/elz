package listener

import (
	"fmt"

	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) EnterFnDefine(ctx *parser.FnDefineContext) {
	// TODO: complete fn generate
	s.inFn = true

	s.fnBuilder = NewFnBuilder().
		Name(ctx.ID().GetText()).
		Export(s.exportThis)
}

func (s *ElzListener) EnterParam(ctx *parser.ParamContext) {
	// TODO: get type info
	// insert ID.Text into fnBuilder
	// set flag let fnBuilder record the type when pass type rule
	// clean flag
}

func (s *ElzListener) ExitFnDefine(ctx *parser.FnDefineContext) {
	s.inFn = false

	s.AstList = append(s.AstList,
		s.fnBuilder.generate(),
	)
	// TODO: fn need builder to create at here, because it will cross several rules
	s.fnBuilder = nil
}
