package irutil_test

import (
	"testing"

	"github.com/elz-lang/elz/src/irutil"

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
