package main

import (
	"fmt"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

type ElzListener struct {
	*parser.BaseElzListener
	AstList    []ast.Ast
	exportThis bool
	immutable  bool
}

func NewElzListener() *ElzListener {
	return &ElzListener{
		exportThis: false,
		immutable:  true,
	}
}

func (s *ElzListener) EnterProg(ctx *parser.ProgContext) {
	fmt.Println(`Elz prog`)
}
func (s *ElzListener) EnterExportor(*parser.ExportorContext) {
	s.exportThis = true
}
func (s *ElzListener) EnterVarDefine(ctx *parser.VarDefineContext) {
	if ctx.GetMut() != nil {
		s.immutable = false
	}
}
func (s *ElzListener) ExitDefine(ctx *parser.DefineContext) {
	// FIXME: We need to get Expression's type, not give it a default value.
	typ := "f32"
	if ctx.TypePass() != nil {
		typ = ctx.TypePass().GetText()
	}
	s.AstList = append(s.AstList, &ast.VarDefination{
		Immutable:  s.immutable,
		Export:     s.exportThis,
		Name:       ctx.ID().GetText(),
		VarType:    typ,
		Expression: &ast.Number{"1"},
	})
}
func (s *ElzListener) ExitVarDefine(*parser.VarDefineContext) {
	if !s.immutable {
		s.immutable = true
	}
}
