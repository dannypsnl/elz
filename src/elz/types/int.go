package types

import (
	"github.com/llir/llvm/ir/types"
)

func (*Int) String() string {
	return "int"
}

// LLVMType is I64 by default, currently have no plan for supporting others platform
func (*Int) LLVMType() types.Type {
	return types.I64
}
