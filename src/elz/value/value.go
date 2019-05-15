package value

import (
	llvmtypes "github.com/llir/llvm/ir/types"
	llvmvalue "github.com/llir/llvm/ir/value"
)

type Value interface {
	llvmvalue.Value
}

type Wrapper struct {
	Value
	ElemT llvmtypes.Type
}

func NewWrap(v llvmvalue.Value, t llvmtypes.Type) Value {
	return &Wrapper{
		Value: v,
		ElemT: t,
	}
}
