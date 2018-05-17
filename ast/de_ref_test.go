package ast

import (
	"testing"

	"strings"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestDereference(t *testing.T) {
	c := NewContext()

	ft := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	main := llvm.AddFunction(c.Module, "main", ft)
	entry := llvm.AddBasicBlock(main, "entry")
	c.Builder.SetInsertPointAtEnd(entry)

	varA := &LocalVarDef{
		Immutable:  true,
		Name:       "a",
		VarType:    "i32",
		Expression: &I32{Val: "10"},
	}
	varB := &LocalVarDef{
		Immutable:  true,
		Name:       "b",
		VarType:    "ref<i32>",
		Expression: &Ref{E: &Id{Val: "a"}},
	}
	varC := &LocalVarDef{
		Immutable:  true,
		Name:       "c",
		VarType:    "i32",
		Expression: &DeRef{E: &Id{Val: "b"}},
	}

	varA.Check(c)
	varB.Check(c)
	varC.Check(c)

	varA.Codegen(c)
	varB.Codegen(c)
	varC.Codegen(c)

	actual := c.Module.String()

	expected := `
  %b = alloca i32*
  store i32* %0, i32** %b
  %1 = load i32*, i32** %b
  %2 = load i32, i32* %1
  %c = alloca i32
  store i32 %2, i32* %c
`

	if !strings.Contains(actual, expected) {
		t.Errorf("expected: `%s`, actual: `%s`", expected, actual)
	}
}
