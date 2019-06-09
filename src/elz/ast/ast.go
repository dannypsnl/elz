package ast

import (
	"strconv"

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
	TypeDefine struct {
		Export bool
		Name   string
		Fields []*Field
	}
	Import struct {
		AccessChain *Ident
	}
	BindingType struct {
		Name string
		Type []Type
	}
	Binding struct {
		IsFunc    bool
		Export    bool
		Name      string
		ParamList []string
		Expr      Expr
		TypeList  []Type
	}
	CaseOf struct {
		isExpr
		Case, Else Expr
		CaseOf     []*Of
	}
	Of struct {
		Pattern Expr
		Expr
	}
	AccessField struct {
		isExpr
		From   Expr
		ByName string
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

func NewOf(pattern, do Expr) *Of {
	return &Of{
		Pattern: pattern,
		Expr:    do,
	}
}
func NewCaseOf(caseExpr Expr, elseExpr Expr, ofs []*Of) *CaseOf {
	caseOf := ofs
	if caseOf == nil {
		caseOf = make([]*Of, 0)
	}
	return &CaseOf{
		Case:   caseExpr,
		Else:   elseExpr,
		CaseOf: caseOf,
	}
}

func NewAccessField(from Expr, name string) *AccessField {
	return &AccessField{
		From:   from,
		ByName: name,
	}
}

func NewTypeDefine(export bool, name string, fields ...*Field) *TypeDefine {
	if fields == nil {
		fields = make([]*Field, 0)
	}
	return &TypeDefine{
		Export: export,
		Name:   name,
		Fields: fields,
	}
}

func NewBinding(isFunc, export bool, name string, params []string, expr Expr) *Binding {
	return &Binding{
		IsFunc:    isFunc,
		Export:    export,
		Name:      name,
		ParamList: params,
		Expr:      expr,
		TypeList:  []Type{},
	}
}

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
	str, err := strconv.Unquote(literal)
	if err != nil {
		logrus.Fatalf("parser bug, string can't be unquoted: %s", err)
	}
	return &String{Literal: str}
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
