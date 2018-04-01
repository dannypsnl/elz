package listener

import (
	"fmt"

	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) EnterFnDefine(ctx *parser.FnDefineContext) {
	fmt.Println("fn", ctx.ID().GetText())

	s.fnBuilder = NewFnBuilder().
		Name(ctx.ID().GetText()).
		Export(s.exportThis)
}

func (s *ElzListener) EnterTypeForm(ctx *parser.TypeFormContext) {
	if s.fnBuilder != nil {
		s.fnBuilder.PushPType(ctx.GetText())
	}
}

func (s *ElzListener) EnterParam(ctx *parser.ParamContext) {
	s.fnBuilder.PushParamName(ctx.ID().GetText())
}

func (s *ElzListener) EnterReturnType(ctx *parser.ReturnTypeContext) {
	fmt.Println("return", ctx.TypeForm().GetText())
	s.fnBuilder.RetType(ctx.TypeForm().GetText())
}

func (s *ElzListener) ExitFnDefine(ctx *parser.FnDefineContext) {
	fmt.Println("fn", ctx.ID().GetText(), "end")
	s.AstList = append(s.AstList,
		s.fnBuilder.generate(),
	)
	// Already leave function scope
	s.fnBuilder = nil
}
