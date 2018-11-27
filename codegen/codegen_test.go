package codegen

import (
	"fmt"
	"testing"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/lib/llvm"
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
	ft := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	mainFn := llvm.AddFunction(c.mod, "main", ft)
	bb := llvm.AddBasicBlock(mainFn, "")
	builder := llvm.NewBuilder()
	builder.SetInsertPointAtEnd(bb)
	c.NewExpr(
		map[string]llvm.Value{},
		builder,
		&ast.FuncCall{
			Identifier: "add",
			ExprList: []ast.Expr{
				ast.NewInt("10"),
				ast.NewInt("10"),
			},
		},
	)
	fmt.Printf("%s", c.mod)
}
