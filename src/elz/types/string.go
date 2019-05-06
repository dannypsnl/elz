package types

import (
	"github.com/llir/llvm/ir/types"
)

func (s *String) String() string {
	return "string"
}

func (s *String) LLVMType() types.Type {
	t := types.NewStruct()
	t.Opaque = true
	return t
}
