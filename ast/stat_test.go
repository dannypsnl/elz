package ast

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestGlobalVarDef(t *testing.T) {
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
	v.Codegen(ctx)
	if ctx.GlobalVars["pi"].v.Type().String() != "FloatType" {
		t.Error(`error`)
	}
}

func TestStrGlobalVarDef(t *testing.T) {
	str := &Str{`"a string"`}
	v := &GlobalVarDef{
		Export:     false,
		Name:       "string1",
		VarType:    "str",
		Expression: str,
	}
	v.Codegen(ctx)
	if ctx.GlobalVars["string1"].v.Type().String() != "ArrayType(IntegerType(8 bits)[10])" {
		t.Errorf("var: %s, expected: %s", ctx.GlobalVars["string1"].v.Type().String(),
			"ArrayType(IntegerType(8 bits)[10])")
	}
}

func TestLocalVarDef(t *testing.T) {
	ft := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	fn := llvm.AddFunction(ctx.Module, "test", ft)
	block := llvm.AddBasicBlock(fn, "entry")
	ctx.Builder.SetInsertPointAtEnd(block)
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
	res := ctx.Builder.CreateLoad(v.Codegen(ctx), "x.load")
	ctx.Builder.CreateRet(res)

	engine, err := llvm.NewExecutionEngine(ctx.Module)
	if err != nil {
		t.Error("Build Engine Problem")
	}
	gres := engine.RunFunction(fn, []llvm.GenericValue{})
	if gres.Float(llvm.FloatType())-3.1415926 > 0.0000001 {
		t.Error("error", gres.Float(llvm.FloatType()))
	}
}
