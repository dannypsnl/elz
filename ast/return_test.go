package ast

import (
	"testing"
)

func TestReturnType(t *testing.T) {
	context := NewContext()
	ret := &Return{
		Expr: &I32{Val: "1"},
	}
	retTyp := ret.Type(context)
	if retTyp != "i32" {
		t.Error("return type is wrong")
	}
}
