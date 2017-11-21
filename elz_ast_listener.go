package main

import (
	"fmt"

	_ "github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

type ElzListener struct {
	*parser.BaseElzListener
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
func (s *ElzListener) ExitVarDefine(*parser.VarDefineContext) {
	if !s.immutable {
		s.immutable = true
	}
}
