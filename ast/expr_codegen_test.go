package ast

import (
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
)

func TestNumber(t *testing.T) {
	m := ir.NewModule()
	n := &Number{
		Val: "1.233",
	}
	expected := constant.NewFloatFromString("1.233", f64)
	if n.Codegen(m).Type() != expected.Type() {
		t.Error(`Ast[Number] should return a constant.`, expected, n.Codegen(m))
	}
}

func TestUnaryExpr(t *testing.T) {
	m := ir.NewModule()
	u := &UnaryExpr{
		E:  &Number{"1.2"},
		Op: "-",
	}
	expected := constant.NewFSub(constant.NewFloatFromString("0", f64), constant.NewFloatFromString("1.2", f64))
	if u.Codegen(m).Type() != expected.Type() {
		t.Error(`Ast[UnaryExpr] should return a FSub`)
	}
}

func TestBinaryExpr(t *testing.T) {
	m := ir.NewModule()
	b := &BinaryExpr{
		LeftE:  &Number{"1.2"},
		RightE: &Number{"3.5"},
		Op:     "-",
	}
	expected := constant.NewFSub(constant.NewFloatFromString("1.2", f64), constant.NewFloatFromString("3.5", f64))
	if b.Codegen(m).Type() != expected.Type() {
		t.Error(`Ast[BinaryExpr] should return a FSub`)
	}
}
