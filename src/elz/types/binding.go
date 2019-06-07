package types

import (
	"fmt"
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
	fs := make([]fmt.Stringer, 0)
	for _, t := range typeListOfArgs {
		fs = append(fs, t)
	}
	key := typeFormat(fs...)
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
	fs := make([]fmt.Stringer, 0)
	for _, t := range b.Binding.Type {
		fs = append(fs, t)
	}
	return typeFormat(fs...)
}

// LLVMType of binding type shouldn't be used
func (b *BindingType) LLVMType() types.Type {
	panic("LLVMType of BindingType shouldn't be used, since it might not have real underlying type")
}

func typeFormat(typeList ...fmt.Stringer) string {
	var b strings.Builder
	if len(typeList) > 0 {
		for _, t := range typeList[:len(typeList)-1] {
			b.WriteString(t.String())
			b.WriteString(" -> ")
		}
		b.WriteString(typeList[len(typeList)-1].String())
	}
	return b.String()
}
