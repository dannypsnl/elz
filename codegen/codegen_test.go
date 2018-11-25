package codegen

import (
	"fmt"
	"testing"

	"github.com/elz-lang/elz/ast"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func TestBinaryFunction(t *testing.T) {
	c := NewGenerator()

	c.GenBinding(&ast.Binding{
		Name:      "add",
		ParamList: []string{"x", "y"},
		Expr: &ast.BinaryExpr{
			LExpr:    ast.NewIdent("x"),
			RExpr:    ast.NewIdent("y"),
			Operator: "+",
		},
	})
	mainFn := c.mod.NewFunction("main", types.I32)
	bb := mainFn.NewBlock("")
	builder := bb
	v := c.NewExpr(
		map[string]value.Value{},
		builder,
		&ast.FuncCall{
			Identifier: "add",
			ExprList: []ast.Expr{
				ast.NewInt("10"),
				ast.NewInt("10"),
			},
		},
	)
	builder.NewRet(v)
	fmt.Printf("%s", c.mod)
}
