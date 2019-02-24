package types

import (
	"fmt"

	"github.com/elz-lang/elz/ast"

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

func TypeOf(e ast.Expr) Type {
	// where e := e.(type) can save the convert in case clause
	switch e := e.(type) {
	case *ast.Arg:
		return TypeOf(e.Expr)
	case *ast.Int:
		return &Int{}
	default:
		panic(fmt.Sprintf("unsupport this expression: `%#v` yet", e))
	}
}
