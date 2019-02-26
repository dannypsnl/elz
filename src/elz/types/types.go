package types

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"

	"github.com/llir/llvm/ir/types"
)

type isType struct{}

func (*isType) IsType() bool { return true }

type Type interface {
	IsType() bool
	fmt.Stringer
	LLVMType() types.Type
}

// Int represents Elz Integer type
type Int struct {
	isType
}

func (*Int) String() string {
	return "int"
}

// LLVMType is I64 by default, currently have no plan for supporting others platform
func (*Int) LLVMType() types.Type {
	return types.I64
}

type Float struct {
	isType
}

func (f *Float) String() string {
	return "f64"
}

func (f *Float) LLVMType() types.Type {
	return types.Double
}

type String struct {
	isType
}

func (s *String) String() string {
	return "string"
}

func (s *String) LLVMType() types.Type {
	t := types.NewStruct()
	t.Opaque = true
	return t
}

func TypeOf(e ast.Expr) Type {
	// where e := e.(type) can save the convert in case clause
	switch e := e.(type) {
	case *ast.Arg:
		return TypeOf(e.Expr)
	case *ast.Int:
		return &Int{}
	case *ast.Float:
		return &Float{}
	case *ast.String:
		return &String{}
	default:
		panic(fmt.Sprintf("you can't use expression: `%#v` to get type directly", e))
	}
}
