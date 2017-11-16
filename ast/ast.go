package ast

import (
	"github.com/elz-lang/elz/lexer"
)

type Ast interface{}

type Error struct {
	Msg string
}

type VarDefination struct {
	Immutable  bool
	Export     bool
	Name       string
	VarType    string
	Expression lexer.Item
}
