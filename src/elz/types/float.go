package types

import (
	"github.com/llir/llvm/ir/types"
)

func (f *Float) String() string {
	return "f64"
}

func (f *Float) LLVMType() types.Type {
	return types.Double
}
