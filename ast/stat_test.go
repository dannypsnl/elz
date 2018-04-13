package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

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
	if context.GlobalVars["pi"].v.Type().String() != "FloatType" {
		t.Error(`error`)
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
	if context.GlobalVars["string1"].v.Type().String() != "ArrayType(IntegerType(8 bits)[10])" {
		t.Errorf("var: %s, expected: %s", context.GlobalVars["string1"].v.Type().String(),
			"ArrayType(IntegerType(8 bits)[10])")
	}
}

func TestLocalVarDef(t *testing.T) {
	context := NewContext()
	ft := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
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
	res := context.Builder.CreateLoad(v.Codegen(context), "x.load")
	context.Builder.CreateRet(res)

	engine, err := llvm.NewExecutionEngine(context.Module)
	if err != nil {
		t.Error("Build Engine Problem")
	}
	gres := engine.RunFunction(fn, []llvm.GenericValue{})
	if gres.Float(llvm.FloatType())-3.1415926 > 0.0000001 {
		t.Error("error", gres.Float(llvm.FloatType()))
	}
}
