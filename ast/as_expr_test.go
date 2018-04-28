package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestAsZExt(t *testing.T) {
	context := NewContext()
	ft := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	main := llvm.AddFunction(context.Module, "main", ft)
	entry := llvm.AddBasicBlock(main, "entry")

	context.Builder.SetInsertPointAtEnd(entry)

	local := &LocalVarDef{
		Immutable:  true,
		Name:       "a",
		VarType:    "i32",
		Expression: &I32{Val: "10"},
	}
	target := &LocalVarDef{
		Immutable:  true,
		Name:       "b",
		VarType:    "i64",
		Expression: &As{E: &Id{Val: "a"}, T: "i64"},
	}
	local.Check(context)
	target.Check(context)

	local.Codegen(context)
	target.Codegen(context)

	expected := `; ModuleID = 'main'
source_filename = "main"

define i32 @main() {
entry:
  %a = alloca i32
  store i32 10, i32* %a
  %a1 = load i32, i32* %a
  %.as_tmp = zext i32 %a1 to i64
  %b = alloca i64
  store i64 %.as_tmp, i64* %b
  %b2 = load i64, i64* %b
}
`

	actual := context.Module

	if context.Module.String() != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}
