package types

import (
	"fmt"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Type struct {
	t types.Type
}

func NewType(t types.Type) *Type {
	return &Type{
		t: t,
	}
}

func (t *Type) LLVMT() types.Type {
	return t.t
}

func (t *Type) NewInt(literal string) value.Value {
	v, err := constant.NewIntFromString(t.t.(*types.IntType), literal)
	if err != nil {
		panic(fmt.Errorf("unable to parse integer literal %s; %s", literal, err))
	}
	return v
}
func (t *Type) NewFloat(literal string) value.Value {
	v, err := constant.NewFloatFromString(t.t.(*types.FloatType), literal)
	if err != nil {
		panic(fmt.Errorf("unable to parse floating-point literal %s; %s", literal, err))
	}
	return v
}

func (t *Type) String() string {
	return t.t.String()
}

var (
	Bool = NewType(types.I1)
	I32  = NewType(types.I32)
	I64  = NewType(types.I64)

	F32 = NewType(types.Float)
	F64 = NewType(types.Double)
)
