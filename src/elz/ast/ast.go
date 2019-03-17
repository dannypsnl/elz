package ast

type isExpr struct{}

func (isExpr) IsExpr() bool {
	return true
}

type (
	Expr interface {
		IsExpr() bool
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
		AccessChain string
		ArgList     []*Arg
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

func NewInt(literal string) *Int     { return &Int{Literal: literal} }
func NewFloat(literal string) *Float { return &Float{Literal: literal} }
func NewString(literal string) *String {
	return &String{Literal: literal}
}
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
