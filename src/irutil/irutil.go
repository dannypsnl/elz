package irutil

import (
	"fmt"

	"github.com/llir/llvm/ir"
)

// FixDups fixes duplicates identifiers in the module by adding uniquely
// differentiating numerical suffixes.
func FixDups(m *ir.Module) {
	names := make(map[string]uint64)
	for _, g := range m.Globals {
		fixName(names, g)
	}
	for _, f := range m.Funcs {
		fixName(names, f)
	}
	for _, a := range m.Aliases {
		fixName(names, a)
	}
	for _, i := range m.IFuncs {
		fixName(names, i)
	}
}

type identifier interface {
	Name() string
	SetName(string)
}

func fixName(counter map[string]uint64, identifier identifier) {
	originName := identifier.Name()
	curCnt := counter[originName]
	if curCnt > 0 {
		newName := fmt.Sprintf("%s.%d", originName, curCnt)
		identifier.SetName(newName)
		counter[newName]++
	}
	counter[originName]++
}
