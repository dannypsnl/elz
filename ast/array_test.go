package ast

import (
	"fmt"
	"strings"
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestArray(t *testing.T) {
	arr := &Array{
		Elements:    []Expr{&I32{Val: "10"}},
		ElementType: "i32",
		Len:         1,
	}

	testCodegenResult(t, arr)
	testArrayType(t, arr)
	testLocalVarCodegenResult(t, arr)
}

func testArrayType(t *testing.T, arr *Array) {
	c := NewContext()
	actual := arr.Type(c)
	expected := "[i32;1]"

	if actual != expected {
		t.Errorf(fmt.Sprintf("expected: `%s`, actual: `%s`", expected, actual))
	}
}

func testLocalVarCodegenResult(t *testing.T, arr *Array) {
	c := NewContext()

	arrl := &LocalVarDef{
		Immutable:  true,
		Name:       "local_arr",
		VarType:    "[i32;1]",
		Expression: arr,
	}
	main := &FnDef{
		Export:      false,
		Name:        "main",
		Params:      []*Param{},
		Body:        []Stat{arrl},
		RetType:     "i32",
		IsExternDef: false,
	}
	main.Check(c)
	main.Codegen(c)

	expected := `yeee`

	if !strings.Contains(c.Module.String(), expected) {
		t.Errorf(fmt.Sprintf("expected contains: `%s`, actual module is: `%s`", expected, c.Module))
	}
}

func testCodegenResult(t *testing.T, arr *Array) {
	c := NewContext()
	arr.Check(c)
	arrv := arr.Codegen(c)
	garr := llvm.AddGlobal(c.Module, llvm.ArrayType(llvm.Int32Type(), 1), "arr")
	garr.SetInitializer(arrv)

	expected := `@arr = global [1 x i32] [i32 10]`

	if !strings.Contains(c.Module.String(), expected) {
		t.Errorf(fmt.Sprintf("expected contains: `%s`, actual module is: `%s`", expected, c.Module))
	}
}

func TestBadArrayShouldReport(t *testing.T) {
	c := NewContext()
	arr := &Array{
		Elements:    []Expr{&I32{Val: "10"}},
		ElementType: "i64",
		Len:         1,
	}
	arr.Check(c)
	report := c.Reporter.ErrMsgs[0]

	expectedReport := "Array expected type: i64, but contains expression type: i32"

	if report != expectedReport {
		t.Errorf(fmt.Sprintf("expectedReport: `%s`, report: `%s`", expectedReport, report))
	}
}
