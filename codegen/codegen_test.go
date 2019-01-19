package codegen

import (
	"fmt"
	"testing"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/codegen/types"
	"github.com/llir/llvm/ir/value"
)

func TestBinaryFunction(t *testing.T) {
	c := NewGenerator()

	f := NewFunc(&ast.Func{
		Name:      "add",
		ParamList: []string{"x", "y"},
		Expr: &ast.BinaryExpr{
			LExpr:    ast.NewIdent("x"),
			RExpr:    ast.NewIdent("y"),
			Operator: "+",
		},
	})
	mainFn := c.mod.NewFunc("main", types.I32.LLVMT())
	mainBlock := mainFn.NewBlock("")
	f.CallWith(c, mainBlock, []value.Value{
		types.I64.NewInt("1"),
		types.I64.NewInt("1"),
	})

	mainBlock.NewRet(types.I32.NewInt("0"))
	fmt.Printf("%s\n", c.mod)
}
