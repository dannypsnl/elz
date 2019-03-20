package codegen

import (
	"fmt"
	"strings"
)

type module struct {
	*Tree
	generator *Generator
	imports   map[string]string
}

// ```
// import math
// import mod::sub_mod
//
// // access
// math::abs(-1)
// sub_mod::xxx()
// ```
//
// import format is `lib::lib::lib`,
// but would only take last name as local name of the module
func newModule(g *Generator, tree *Tree) *module {
	imports := map[string]string{}
	for _, importPath := range tree.imports {
		accessChain := strings.Split(importPath, "::")
		lastOne := len(accessChain) - 1
		accessKey := accessChain[lastOne]
		if mod1, exist := imports[accessKey]; exist {
			panic(fmt.Sprintf(`import %s
and
import %s
has the same name in the module`, mod1, importPath))
		}
		imports[accessKey] = importPath
	}
	return &module{
		Tree:      tree,
		generator: g,
		imports:   imports,
	}
}

func (m *module) getBindingByAccessChain(accessChain string) (*Binding, error) {
	chain := strings.Split(accessChain, "::")
	if len(chain) >= 2 {
		localModuleName := chain[len(chain)-2]
		funcName := chain[len(chain)-1]
		moduleName := m.imports[localModuleName]
		return m.generator.allModule[moduleName].GetExportBinding(funcName)
	}
	if len(chain) == 1 {
		bind, err := m.GetBinding(accessChain)
		if err != nil {
			return m.generator.entryModule.GetBinding(accessChain)
		}
		return bind, nil
	}
	return nil, fmt.Errorf("not supported access chain: %s", accessChain)
}
