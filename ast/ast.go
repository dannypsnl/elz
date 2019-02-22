package ast

type (
	Node interface{}
	Expr interface{}

	Binding struct {
		Name      string
		ParamList []string
		Expr      Expr
	}
	FuncCall struct {
		FuncName string
		ExprList []Expr
	}
	Arg struct {
		Ident string
		Expr  Expr
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
		LExpr Expr
		RExpr Expr
		Op    string
	}
	Ident string
)

func NewIdent(literal string) Ident {
	return Ident(literal)
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
