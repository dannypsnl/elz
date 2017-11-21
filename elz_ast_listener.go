package main

import (
	"fmt"
	_ "github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

type ElzListener struct {
	*parser.BaseElzListener
	exportThis bool
}

func NewElzListener() *ElzListener {
	return &ElzListener{
		exportThis: false,
	}
}

func (s *ElzListener) EnterProg(ctx *parser.ProgContext) {
	fmt.Println(`Elz prog`)
}
func (s *ElzListener) EnterExportor(*parser.ExportorContext) {
	s.exportThis = true
}
func (s *ElzListener) EnterVarDefine(ctx *parser.VarDefineContext) {
	fmt.Println("mmutable:", ctx.GetMut())
}
