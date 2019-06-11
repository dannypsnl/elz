package types

import (
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"

	"github.com/llir/llvm/ir/types"
)

type Module interface {
	InferTypeOf(expr ast.Expr, typeMap *TypeMap) (Type, error)
}

type BindingType struct {
	isType
	cacheOfType map[string]Type
	selfModule  Module
	*ast.Binding
}

func NewBindingType(m Module, binding *ast.Binding) Type {
	bt := &BindingType{
		cacheOfType: make(map[string]Type),
		selfModule:  m,
		Binding:     binding,
	}
	return bt
}

func (b *BindingType) GetReturnType(typeMap *TypeMap, typeListOfArgs ...Type) (Type, error) {
	key := TypeFormat(typeListOfArgs...)
	t, ok := b.cacheOfType[key]
	if ok {
		return t, nil
	}
	inferT, err := b.selfModule.InferTypeOf(b.Expr, typeMap)
	if err != nil {
		return nil, err
	}
	b.cacheOfType[key] = inferT
	return inferT, nil
}

func (b *BindingType) String() string {
	var buf strings.Builder
	buf.WriteRune('(')
	for i, t := range b.Binding.ParamList {
		if i != 0 {
			buf.WriteRune(',')
			buf.WriteRune(' ')
		}
		buf.WriteString(t.String())
	}
	buf.WriteRune(')')
	return buf.String()
}

// LLVMType of binding type shouldn't be used
func (b *BindingType) LLVMType() types.Type {
	panic("LLVMType of BindingType shouldn't be used, since it might not have real underlying type")
}

func TypeFormat(typeList ...Type) string {
	var b strings.Builder
	b.WriteRune('(')
	for i, t := range typeList {
		if i != 0 {
			b.WriteRune(',')
			b.WriteRune(' ')
		}
		b.WriteString(t.String())
	}
	b.WriteRune(')')
	return b.String()
}
