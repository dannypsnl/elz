package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestAsZExt(t *testing.T) {
	testI32ToI64(t)
}

func assertEq(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func getStructBy(typ string) Expr {
	v := "10"
	switch typ {
	case "i32":
		return &I32{Val: v}
	case "f32":
		return &F32{Val: v}
	case "i64":
		return &I64{Val: v}
	case "f64":
		return &F64{Val: v}
	}
	panic("We don't support yet")
}

func convertTemplate(t *testing.T, from, targetType string) {
	context := NewContext()
	ft := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	main := llvm.AddFunction(context.Module, "main", ft)
	entry := llvm.AddBasicBlock(main, "entry")

	context.Builder.SetInsertPointAtEnd(entry)

	local := &LocalVarDef{
		Immutable:  true,
		Name:       "a",
		VarType:    from,
		Expression: getStructBy(from),
	}
	as := &As{E: &Id{Val: "a"}, T: targetType}
	target := &LocalVarDef{
		Immutable:  true,
		Name:       "b",
		VarType:    targetType,
		Expression: as,
	}
	local.Check(context)
	target.Check(context)

	local.Codegen(context)
	target.Codegen(context)

	actual := context.Module.String()

	expected := `; ModuleID = 'main'
source_filename = "main"

define ` + from + ` @main() {
entry:
  %a = alloca ` + from + `
  store ` + from + ` 10, ` + from + `* %a
  %a1 = load ` + from + `, ` + from + `* %a
  %.as_tmp = ` + opcode2String(as.op) + ` ` + from + ` %a1 to ` + targetType + `
  %b = alloca ` + targetType + `
  store ` + targetType + ` %.as_tmp, ` + targetType + `* %b
  %b2 = load ` + targetType + `, ` + targetType + `* %b
}
`

	assertEq(t, actual, expected)
}

func opcode2String(op llvm.Opcode) string {
	switch op {
	case llvm.ZExt:
		return "zext"
	default:
		panic("unknown opcode")
	}
}

func testI32ToI64(t *testing.T) {
	convertTemplate(t, "i32", "i64")
}
