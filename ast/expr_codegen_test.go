package ast

import (
	"testing"

	"github.com/llir/llvm/ir/constant"
)

func TestNumber(t *testing.T) {
	n := &Number{
		Val: "1.233",
	}
	expected := constant.NewFloatFromString("1.233", f32)
	if n.Codegen().Type() != expected.Type() {
		t.Error(`Ast[Number] should return a constant.`, expected, n.Codegen())
	}
}

func TestUnaryExpr(t *testing.T) {
	u := &UnaryExpr{
		E:  &Number{"1.2"},
		Op: "-",
	}
	expected := constant.NewFSub(constant.NewFloatFromString("0", f32), constant.NewFloatFromString("1.2", f32))
	if u.Codegen().Type() != expected.Type() {
		t.Error(`Ast[UnaryExpr] should return a FSub`)
	}
}

func TestBinaryExpr(t *testing.T) {
	b := &BinaryExpr{
		LeftE:  &Number{"1.2"},
		RightE: &Number{"3.5"},
		Op:     "-",
	}
	expected := constant.NewFSub(constant.NewFloatFromString("1.2", f32), constant.NewFloatFromString("3.5", f32))
	if b.Codegen().Type() != expected.Type() {
		t.Error(`Ast[BinaryExpr] should return a FSub`)
	}
}
