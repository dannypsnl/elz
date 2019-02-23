package codegen_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/codegen"
)

func TestBindingCodegen(t *testing.T) {
	g := codegen.New()

	g.Call(&ast.Binding{
		Name:      "add",
		ParamList: []string{"x", "y"},
		Expr: &ast.BinaryExpr{
			LExpr: ast.NewIdent("x"),
			RExpr: ast.NewIdent("y"),
			Op:    "+",
		},
	},
		ast.NewArg("", ast.NewInt("1")),
		ast.NewArg("", ast.NewInt("2")),
	)

	assert.Contains(t, g.String(), `define i64 @add(i64, i64) {
; <label>:2
	%3 = add i64 1, 2
	ret i64 %3
}`)
}
