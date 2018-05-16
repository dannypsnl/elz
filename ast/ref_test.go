package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestRefI32(t *testing.T) {
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

	varA.Check(c)
	varB.Check(c)

	varA.Codegen(c)
	varB.Codegen(c)

	actual := c.Module.String()

	expected := `; ModuleID = 'main'
source_filename = "main"

define i32 @main() {
entry:
  %a = alloca i32
  store i32 10, i32* %a
  %0 = getelementptr i32, i32* %a, i32 0
  %b = alloca i32*
  store i32* %0, i32** %b
}
`

	assertEq(t, actual, expected)
}
