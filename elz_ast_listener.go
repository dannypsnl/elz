package main

import (
	"fmt"
	_ "github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

type ElzListener struct {
	*parser.BaseElzListener
}

func NewElzListener() *ElzListener {
	return new(ElzListener)
}

func (s *ElzListener) EnterProg(ctx *parser.ProgContext) {
	fmt.Println(`Elz prog`)
}
