package codegen_test

import (
	"testing"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/codegen"

	"github.com/stretchr/testify/assert"
)

var (
	bindings = []*ast.Binding{
		{
			Name:      "addOne",
			ParamList: []string{"y"},
			Expr: &ast.FuncCall{
				FuncName: "add",
				ExprList: []*ast.Arg{
					ast.NewArg("", ast.NewInt("1")),
					ast.NewArg("", ast.NewIdent("y")),
				},
			},
		},
		{
			Name:      "add",
			ParamList: []string{"x", "y"},
			Expr: &ast.BinaryExpr{
				LExpr: ast.NewIdent("x"),
				RExpr: ast.NewIdent("y"),
				Op:    "+",
			},
		},
	}

	bindMap = map[string]*ast.Binding{}
)

func init() {
	for _, bind := range bindings {
		bindMap[bind.Name] = bind
	}
}

func TestBindingCodegen(t *testing.T) {
	testCases := []struct {
		name           string
		bindName       string
		args           []*ast.Arg
		expectContains string
	}{
		{
			name:     "call by generator",
			bindName: "add",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewInt("1")),
				ast.NewArg("", ast.NewInt("2")),
			},
			expectContains: `define i64 @add(i64, i64) {
; <label>:2
	%3 = add i64 1, 2
	ret i64 %3
}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			g := codegen.New()
			g.Call(bindMap[testCase.bindName], testCase.args...)
			assert.Contains(t, g.String(), testCase.expectContains)
		})
	}
}
