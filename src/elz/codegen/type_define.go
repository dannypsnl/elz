package codegen

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"

	llvmtypes "github.com/llir/llvm/ir/types"
)

func (g *Generator) GenerateTypes() {
	mod := g.entryModule
	typDefs := mod.GetTypeDefines()
	for typName, def := range typDefs {
		g.mod.NewTypeDef(typName, convert(def))
	}
	for modName, mod := range g.allModule {
		typDefs := mod.GetTypeDefines()
		for typName, def := range typDefs {
			finalTypeName := fmt.Sprintf("%s::%s", modName, typName)
			g.mod.NewTypeDef(finalTypeName, convert(def))
		}
	}
}

func convert(def *ast.TypeDefine) llvmtypes.Type {
	newStruct := &llvmtypes.StructType{
		Fields: make([]llvmtypes.Type, 0),
		Opaque: false,
	}
	for _, field := range def.Fields {
		newStruct.Fields = append(newStruct.Fields,
			types.FromString(field.Type.String()).LLVMType(),
		)
	}
	return newStruct
}
