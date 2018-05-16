package ast

import (
	"fmt"
	"strings"
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestArrayType(t *testing.T) {
	c := NewContext()

	arr := &Array{
		Elements: []Expr{&I32{Val: "10"}},
		Len:      1,
	}

	arr.Check(c)

	actual := arr.Type(c)
	expected := "[i32;1]"

	if actual != expected {
		t.Errorf(fmt.Sprintf("expected: `%s`, actual: `%s`", expected, actual))
	}
}

func TestCodegenResult(t *testing.T) {
	c := NewContext()

	arr := &Array{
		Elements: []Expr{&I32{Val: "10"}},
		Len:      1,
	}

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
		Elements: []Expr{&I32{Val: "10"}, &I64{Val: "20"}},
		Len:      2,
	}
	arr.Check(c)
	report := c.Reporter.ErrMsgs[0]

	expectedReport := "Array expected type: i32, but contains expression type: i64"

	if report != expectedReport {
		t.Errorf(fmt.Sprintf("expectedReport: `%s`, report: `%s`", expectedReport, report))
	}
}
