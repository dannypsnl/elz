package types

import (
	"fmt"

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
