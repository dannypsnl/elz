package ast

type (
	Node    interface{}
	Binding struct {
		Name      string
		ParamList []string
		Expr      Expr
	}
	Expr interface {
		IsConst() bool
	}
	FuncCall struct {
		Identifier string
		ExprList   []Expr
	}
	Int struct {
		Literal string
	}
	Float struct {
		Literal string
	}
	String struct {
		Literal string
	}
	Bool struct {
		IsTrue bool
	}
	BinaryExpr struct {
		LExpr    Expr
		RExpr    Expr
		Operator string
	}
	Ident struct {
		Value string
	}
)

func NewIdent(literal string) *Ident {
	return &Ident{
		Value: literal,
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

func (f *FuncCall) IsConst() bool   { return false }
func (i *Int) IsConst() bool        { return true }
func (f *Float) IsConst() bool      { return true }
func (b *Bool) IsConst() bool       { return true }
func (s *String) IsConst() bool     { return true }
func (b *BinaryExpr) IsConst() bool { return false }
func (i *Ident) IsConst() bool      { return false }
