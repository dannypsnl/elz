package ast

type Ast interface{}
type Expr interface{}
type Stat interface{}

type StatList []Stat

type Error struct {
	Msg string
}

type Number struct {
	Val string
}

type UnaryExpr struct {
	E    Expr
	Op   string
	Type string
}

type BinaryExpr struct {
	LeftE  Expr
	RightE Expr
	Op     string
	Type   string
}

type VarDefination struct {
	Immutable  bool
	Export     bool
	Name       string
	VarType    string
	Expression Expr
}

type Param struct {
	Name string
	Type string
}
type FnDefination struct {
	Export bool
	Name   string
	Params []Param
	Body   StatList
}

type Argu struct {
	Val  string
	Type string
}
type FnCall struct {
	Name string
	Args []Argu
	Type string
}
