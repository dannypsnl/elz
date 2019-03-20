package codegen

import (
	"fmt"
	"strings"
)

type module struct {
	*Tree
	imports map[string]string
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
func newModule(tree *Tree) *module {
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
		Tree:    tree,
		imports: imports,
	}
}
