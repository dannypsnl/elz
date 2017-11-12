package codegen

import (
	"testing"

	"github.com/elz-lang/elz/ast"

	"github.com/llir/llvm/ir/types"
	"github.com/stretchr/testify/assert"
)

func TestGlobalValue(t *testing.T) {
	c := NewGenerator()

	c.GenBinding(&ast.Binding{
		Name:      "i",
		ParamList: append(make([]string, 0)),
		Expr:      &ast.Int{Literal: "1"},
	})
	assert.Equal(t, len(c.mod.Globals), 1)
	g := c.mod.Globals[0]
	assert.Equal(t, "i", g.Name)
	assert.Equal(t, types.I32, g.Content)
	assert.Equal(t, types.NewPointer(types.I32), g.Typ)
}

func TestFuncNoParameter(t *testing.T) {
	c := NewGenerator()

	newFn := c.GetFuncWithoutParam(&ast.Binding{
		Name:      "x",
		ParamList: []string{},
		Expr: &ast.BinaryExpr{
			LExpr:    &ast.Int{Literal: "1"},
			RExpr:    &ast.Int{Literal: "1"},
			Operator: "+",
		},
	})
	assert.Equal(t, "x", newFn.Name)
}

func TestCodegenerationMain(t *testing.T) {
	c := NewGenerator()

	c.GenBinding(&ast.Binding{
		Name:      "main",
		ParamList: []string{},
		Expr: &ast.FuncCall{
			Identifier: "println",
			ExprList:   []ast.Expr{&ast.String{Literal: "Hello, World"}},
		},
	})
	mainFn := c.mod.Funcs[0]
	assert.Equal(t, "main", mainFn.Name)
}
