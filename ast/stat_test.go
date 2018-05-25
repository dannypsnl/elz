package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestLocalTypeError(t *testing.T) {
	context := NewContext()

	v := &LocalVarDef{
		Immutable:  true,
		Name:       "x",
		VarType:    "f32",
		Expression: &I32{"10"},
	}

	v.Check(context)
	v.Codegen(context)

	if context.Reporter.HasNoError() {
		t.Error("Local Var Ast should receive an error about var type is not equal to expression type")
	}
}

func TestTypeError(t *testing.T) {
	context := NewContext()

	v := &GlobalVarDef{
		Export:     false,
		Name:       "x",
		VarType:    "f32",
		Expression: &I32{"10"},
	}

	v.Check(context)
	v.Codegen(context)

	if context.Reporter.HasNoError() {
		t.Error("Global Var Ast should receive an error about var type is not equal to expression type")
	}
}

func TestGlobalVarDef(t *testing.T) {
	context := NewContext()
	v := &GlobalVarDef{
		Export:  false,
		Name:    "pi",
		VarType: "f32",
		Expression: &BinaryExpr{
			&F32{"2.1215926"},
			&F32{"1.02"},
			"+",
		},
	}
	v.Codegen(context)
	expected := "PointerType(FloatType)"
	if context.Vars["pi"].Type().String() != expected {
		t.Errorf("var: %s, expected: %s",
			context.Vars["pi"].Type().String(),
			expected,
		)
	}
}

func TestStrGlobalVarDef(t *testing.T) {
	context := NewContext()
	str := &Str{`"a string"`}
	v := &GlobalVarDef{
		Export:     false,
		Name:       "string1",
		VarType:    "str",
		Expression: str,
	}
	v.Codegen(context)

	expected := "PointerType(ArrayType(IntegerType(8 bits)[10]))"
	if context.Vars["string1"].Type().String() != expected {
		t.Errorf("var: %s, expected: %s",
			context.Vars["string1"].Type().String(),
			expected,
		)
	}
}

func TestLocalVarDef(t *testing.T) {
	context := NewContext()
	ft := llvm.FunctionType(llvm.FloatType(), []llvm.Type{}, false)
	fn := llvm.AddFunction(context.Module, "test", ft)
	block := llvm.AddBasicBlock(fn, "entry")
	context.Builder.SetInsertPointAtEnd(block)
	v := &LocalVarDef{
		Immutable: true,
		Name:      "x",
		VarType:   "f32",
		Expression: &BinaryExpr{
			&F32{"2.1215926"},
			&F32{"1.02"},
			"+",
		},
	}
	v.Check(context)
	res := context.Builder.CreateLoad(v.Codegen(context), "x.load")
	context.Builder.CreateRet(res)

	engine, err := llvm.NewExecutionEngine(context.Module)
	if err != nil {
		t.Error("Build Engine Problem")
	}
	gres := engine.RunFunction(fn, []llvm.GenericValue{})
	if gres.Float(context.Type("f32"))-3.1415926 > 0.0000001 {
		t.Error("error", gres.Float(llvm.FloatType()))
	}
}
