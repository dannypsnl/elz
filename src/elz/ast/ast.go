package ast

import (
	"github.com/sirupsen/logrus"
)

type isExpr struct{}

func (isExpr) IsExpr() bool {
	return true
}

type (
	Expr interface {
		IsExpr() bool
	}

	Field struct {
		Name string
		Type
	}
	NewType struct {
		Name   string
		Fields []*Field
	}
	// FIXME: merge this with the codegen tree(originally approach of imports)
	Import struct {
		AccessChain *Ident
	}
	BindingType struct {
		Name string
		Type []Type
	}
	Binding struct {
		Export    bool
		Name      string
		ParamList []string
		Expr      Expr
		Type      []Type
	}
	FuncCall struct {
		isExpr
		Func    Expr
		ArgList []*Arg
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
	List struct {
		isExpr
		ExprList []Expr
	}
	Bool struct {
		isExpr
		IsTrue bool
	}
	ExtractElement struct {
		isExpr
		X   Expr
		Key Expr
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

func NewField(name string, t Type) *Field {
	return &Field{
		Name: name,
		Type: t,
	}
}

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

func NewInt(literal string) *Int     { return &Int{Literal: literal} }
func NewFloat(literal string) *Float { return &Float{Literal: literal} }
func NewString(literal string) *String {
	return &String{Literal: literal}
}
func NewList(exprList ...Expr) *List {
	if len(exprList) == 0 {
		exprList = []Expr{}
	}
	return &List{
		ExprList: exprList,
	}
}
func NewBool(literal string) *Bool {
	var value bool
	switch literal {
	case "true":
		value = true
	case "false":
		value = false
	default:
		logrus.Fatalf("boolean syntax define must be wrong")
	}
	return &Bool{IsTrue: value}
}
func NewExtractElement(x, key Expr) *ExtractElement {
	return &ExtractElement{
		X:   x,
		Key: key,
	}
}
