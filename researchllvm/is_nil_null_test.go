package researchllvm

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestIsNilNull(t *testing.T) {
	v := llvm.Value{}

	println(v.IsNil())
}
