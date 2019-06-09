package irutil_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/internal/irutil"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"

	"github.com/stretchr/testify/assert"
)

func TestIRFixDuplicate(t *testing.T) {
	testCases := []struct {
		name         string
		originalName string
		newName      string
	}{
		{
			name:         "simple",
			originalName: "a",
			newName:      "a.1",
		},
		{
			name:         "number should be increased since `<num>.<num>` is invalid",
			originalName: "",
			newName:      "1",
		},
	}

	for _, testCase := range testCases {
		mod := ir.NewModule()
		mod.NewGlobalDef(testCase.originalName, constant.NewZeroInitializer(types.I8))
		def := mod.NewGlobalDef(testCase.originalName, constant.NewZeroInitializer(types.I8))
		irutil.FixDups(mod)
		assert.Equal(t, testCase.newName, def.GlobalName, testCase.name)
	}
}

func TestSizeOf(t *testing.T) {
	testCases := []struct {
		name     string
		llvmType types.Type
		size     int64
	}{
		{
			name:     "int64",
			llvmType: types.I64,
			size:     64,
		},
		{
			name: "structure",
			llvmType: types.NewStruct(
				types.I64,
				types.I64,
			),
			size: 128,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.size, irutil.SizeOf(testCase.llvmType))
	}
}
