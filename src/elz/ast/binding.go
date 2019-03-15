package ast

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/types"
)

func (b *Binding) CheckArg(args ...*Arg) error {
	for i, arg := range args {
		argNameMustBe := b.ParamList[i]
		argName := arg.Ident
		// allow ignore argument name like: `add(1, 2)`
		if argName == "" {
			argName = argNameMustBe
		}
		if argNameMustBe != argName {
			return fmt.Errorf(`argument name must be parameter name(or empty), for example:
  assert that should_be = ...
  assert(that: 1+2, should_be: 3)
`)
		}
	}
	return nil
}

func (b *Binding) TypeCheck(actualTypeWhenCall string, typeList []types.Type) error {
	if b.Type == nil {
		return nil
	}
	var (
		builder        strings.Builder
		err            error
		variantTypeMap = map[string]string{}
	)
	for i, requireT := range b.Type[:len(b.Type)-1] {
		actualType := typeList[i]
		switch requireT := requireT.(type) {
		case *ExistType:
			if requireT.Name != actualType.String() {
				err = fmt.Errorf("")
			}
		case *VariantType:
			t, exist := variantTypeMap[requireT.Name]
			if !exist {
				variantTypeMap[requireT.Name] = actualType.String()
				t = actualType.String()
			}
			if t != actualType.String() {
				err = fmt.Errorf("")
			}
		case *VoidType:
		}
		builder.WriteString(requireT.String())
		builder.WriteString(" -> ")
	}
	requireT := builder.String()
	requireType := b.Name + " :: " + requireT[:len(requireT)-4]
	if err != nil {
		return fmt.Errorf("require type: `%s` but get: `%s`", requireType, actualTypeWhenCall)
	}
	return nil
}
