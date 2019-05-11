package types

import (
	"github.com/llir/llvm/ir/types"
)

func NewList(elemT Type) Type {
	return &List{
		ElemT: elemT,
	}
}

func (l *List) String() string {
	return "list<" + l.ElemT.String() + ">"
}

// LLVMType is opaque type, from external implementation
func (*List) LLVMType() types.Type {
	t := types.NewStruct()
	t.TypeName = "list"
	t.Opaque = true
	return t
}
