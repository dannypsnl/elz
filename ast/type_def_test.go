package ast

import (
	"testing"
)

func TestTypeDef(t *testing.T) {
	c := NewContext()

	typeDef := &TypeDef{
		Name: "Foo",
		Attrs: []TypeAttr{
			TypeAttr{
				Name: "i",
				Typ:  "i32",
			},
		},
	}

	typeDef.Check(c)
	typeDef.Codegen(c)

	res := c.Module.String()
	expectedHas := "%Foo = type <{ i32 }>"

	testHas(t, res, expectedHas)
}
