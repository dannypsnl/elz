package ast

import (
	"testing"

	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestAsZExt(t *testing.T) {
	testI32ToI64(t)
}

func getStructBy(typ string) Expr {
	v := "10"
	switch typ {
	case "i8":
		return &I8{Val: v}
	case "i16":
		return &I16{Val: v}
	case "i32":
		return &I32{Val: v}
	case "i64":
		return &I64{Val: v}
	case "f32":
		return &F32{Val: v}
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

	from = convertToLLVMType(from)
	targetType = convertToLLVMType(targetType)

	expected := `; ModuleID = 'main'
source_filename = "main"

define i32 @main() {
entry:
  %a = alloca ` + from + `
  store ` + from + ` ` + constValueDecideBy(from) + `, ` + from + `* %a
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
	case llvm.SExt:
		return "sext"
	case llvm.Trunc:
		return "trunc"
	case llvm.FPExt:
		return "fpext"
	case llvm.FPTrunc:
		return "fptrunc"
	default:
		panic("unknown opcode")
	}
}

func testI32ToI64(t *testing.T) {
	convertTemplate(t, "i8", "i16")
	convertTemplate(t, "i8", "i32")
	convertTemplate(t, "i8", "i64")

	convertTemplate(t, "i16", "i8")
	convertTemplate(t, "i16", "i32")
	convertTemplate(t, "i16", "i64")

	convertTemplate(t, "i32", "i64")
	convertTemplate(t, "i32", "i16")
	convertTemplate(t, "i32", "i8")

	convertTemplate(t, "i64", "i32")
	convertTemplate(t, "i64", "i16")
	convertTemplate(t, "i64", "i8")

	convertTemplate(t, "f32", "f64")
	convertTemplate(t, "f64", "f32")
}

func convertToLLVMType(t string) string {
	switch t {
	case "i8":
		fallthrough
	case "i16":
		fallthrough
	case "i32":
		fallthrough
	case "i64":
		return t
	case "f32":
		return "float"
	case "f64":
		return "double"
	default:
		panic(fmt.Sprintf("Unsupport type %s yet", t))
	}
}

func constValueDecideBy(t string) string {
	switch t {
	case "i8":
		fallthrough
	case "i16":
		fallthrough
	case "i32":
		fallthrough
	case "i64":
		return "10"
	case "float":
		fallthrough
	case "double":
		return "1.000000e+01"
	default:
		panic(fmt.Sprintf("Unsupport type %s yet", t))
	}
}
