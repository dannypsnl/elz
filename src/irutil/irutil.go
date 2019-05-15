package irutil

import (
	"fmt"
	"strconv"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func SizeOf(t types.Type) int64 {
	switch t {
	case types.I64:
		return 64
	default:
		panic(fmt.Sprintf("size of unsupported type %s yet", t))
	}
}

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
		// if string is a int as 0, 1
		if v, err := strconv.Atoi(originName); err == nil {
			fixNumberName(v, counter, identifier)
			return
		}
		newName := fmt.Sprintf("%s.%d", originName, curCnt)
		identifier.SetName(newName)
		counter[newName]++
	}
	counter[originName]++
}

func fixNumberName(v int, counter map[string]uint64, identifier identifier) {
	newName := fmt.Sprintf("%d", v+1)
	_, ok := counter[newName]
	if ok {
		fixNumberName(v+1, counter, identifier)
	}
	identifier.SetName(newName)
	counter[newName]++
}
