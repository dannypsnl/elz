package listener

import (
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) EnterFnDefine(ctx *parser.FnDefineContext) {
	s.inFn = true

	s.fnBuilder = NewFnBuilder().
		Name(ctx.ID().GetText()).
		Export(s.exportThis)
}

func (s *ElzListener) EnterParamList(ctx *parser.ParamListContext) {
	// s.paramRecorder = recoder()
}

func (s *ElzListener) EnterTypeForm(ctx *parser.TypeFormContext) {
	s.fnBuilder.PushPType(ctx.GetText())
}

func (s *ElzListener) EnterParam(ctx *parser.ParamContext) {
	// TODO: complete fn generate
	// TODO: get type info
	// insert ID.Text into fnBuilder
	// set flag let fnBuilder record the type when pass type rule
	// clean flag
	s.fnBuilder.PushParamName(ctx.ID().GetText())
}

func (s *ElzListener) ExitFnDefine(ctx *parser.FnDefineContext) {
	s.inFn = false // Already leave function scope

	s.AstList = append(s.AstList,
		s.fnBuilder.generate(),
	)
	s.fnBuilder = nil
}
