package types

import (
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"

	llvmtypes "github.com/llir/llvm/ir/types"
)

type Struct struct {
	isType
	*ast.TypeDefine
}

func (s *Struct) String() string {
	buf := strings.Builder{}
	buf.WriteString(s.Name)
	buf.WriteRune('(')
	for i, field := range s.Fields {
		buf.WriteString(field.Name)
		buf.WriteRune(':')
		buf.WriteRune(' ')
		buf.WriteString(field.Type.String())
		if i == len(s.Fields)-1 {
			buf.WriteRune(')')
		} else {
			buf.WriteRune(',')
		}
	}
	return buf.String()
}

func (s *Struct) LLVMType() llvmtypes.Type {
	return convert(s.TypeDefine)
}

func convert(def *ast.TypeDefine) llvmtypes.Type {
	newStruct := &llvmtypes.StructType{
		Fields: make([]llvmtypes.Type, 0),
		Opaque: false,
	}
	for _, field := range def.Fields {
		// FIXME: hard code
		newStruct.Fields = append(newStruct.Fields,
			FromString(field.Type.String()).LLVMType(),
		)
	}
	return newStruct
}
