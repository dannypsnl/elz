package listener

import (
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) EnterFnDefine(ctx *parser.FnDefineContext) {
	s.fnBuilder = NewFnBuilder().
		Name(ctx.ID().GetText()).
		Export(s.exportThis)
}

func (s *ElzListener) EnterParamType(ctx *parser.ParamTypeContext) {
	s.fnBuilder.PushParamType(ctx.TypeForm().GetText())
}

// EnterParam listen parameter rule, emit name of parameter to FnBuilder
func (s *ElzListener) EnterParam(ctx *parser.ParamContext) {
	s.fnBuilder.PushParamName(ctx.ID().GetText())
}

func (s *ElzListener) EnterReturnType(ctx *parser.ReturnTypeContext) {
	s.fnBuilder.RetType(ctx.TypeForm().GetText())
}

func (s *ElzListener) ExitFnDefine(ctx *parser.FnDefineContext) {
	s.AstList = append(s.AstList,
		s.fnBuilder.generate(),
	)
	// Already leave function scope
	s.fnBuilder = nil
}
