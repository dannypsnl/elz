package ast

import (
	"fmt"
)

type Tree struct {
	bindings map[string][]*Binding
}

func NewTree() *Tree {
	return &Tree{
		bindings: make(map[string][]*Binding),
	}
}

func (t *Tree) InsertBinding(b *Binding) {
	if t.bindings[b.Name] == nil {
		t.bindings[b.Name] = make([]*Binding, 0)
	}
	t.bindings[b.Name] = append(t.bindings[b.Name], b)
}

func (t *Tree) GetBinding(bindName string) (*Binding, error) {
	binding, exist := t.bindings[bindName]
	if !exist {
		return nil, fmt.Errorf("no binding name: %s", bindName)
	}
	// FIXME: assuming only one binding for right now situation, should do complete discuss into how to deal with multiple implementations
	return binding[0], nil
}

type isExpr struct{}

func (isExpr) IsExpr() bool {
	return true
}

type (
	Expr interface {
		IsExpr() bool
	}

	Binding struct {
		Name      string
		ParamList []string
		Expr      Expr
		Type      []Type
	}
	FuncCall struct {
		isExpr
		FuncName string
		ExprList []*Arg
	}
	Arg struct {
		isExpr
		Ident string
		Expr  Expr
	}
	Int struct {
		isExpr
		Literal string
	}
	Float struct {
		isExpr
		Literal string
	}
	String struct {
		isExpr
		Literal string
	}
	Bool struct {
		isExpr
		IsTrue bool
	}
	BinaryExpr struct {
		isExpr
		LExpr Expr
		RExpr Expr
		Op    string
	}
	Ident struct {
		isExpr
		Literal string
	}
)

func NewArg(ident string, expr Expr) *Arg {
	return &Arg{
		Ident: ident,
		Expr:  expr,
	}
}

func NewIdent(literal string) *Ident {
	return &Ident{
		Literal: literal,
	}
}

func NewInt(literal string) *Int       { return &Int{Literal: literal} }
func NewFloat(literal string) *Float   { return &Float{Literal: literal} }
func NewString(literal string) *String { return &String{Literal: literal} }
func NewBool(literal string) *Bool {
	var value bool
	switch literal {
	case "true":
		value = true
	case "false":
		value = false
	default:
		panic("boolean syntax define must be wrong")
	}
	return &Bool{IsTrue: value}
}
