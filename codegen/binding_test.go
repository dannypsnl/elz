package codegen_test

import (
	"testing"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/codegen"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	bindings = []*ast.Binding{
		// addOne y = add(1, y)
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
		// add x y = x + y
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

	bindTypes = []*ast.BindType{
		// add :: int -> int -> int
		{
			Name: "add",
			Type: []ast.Type{
				&ast.ExistType{Name: "int"},
				&ast.ExistType{Name: "int"},
				&ast.ExistType{Name: "int"},
			},
		},
	}

	bindMap  = map[string]*ast.Binding{}
	bindType = map[string]*ast.BindType{}
)

func init() {
	for _, bind := range bindings {
		bindMap[bind.Name] = bind
	}
	for _, bindT := range bindTypes {
		bindType[bindT.Name] = bindT
	}
}

func TestBindingCodegen(t *testing.T) {
	testCases := []struct {
		name           string
		bindName       string
		args           []*ast.Arg
		expectContains []string
		expectErrorMsg string
	}{
		{
			name:     "call by generator",
			bindName: "add",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewInt("1")),
				ast.NewArg("", ast.NewInt("2")),
			},
			expectContains: []string{`define i64 @add(i64, i64) {
; <label>:2
	%3 = add i64 %0, %1
	ret i64 %3
}`},
		},
		{
			name:     "call function in function",
			bindName: "addOne",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewInt("2")),
			},
			expectContains: []string{
				`define i64 @addOne(i64) {
; <label>:1
	%2 = call i64 @add(i64 1, i64 %0)
	ret i64 %2
}`,
				`define i64 @add(i64, i64) {
; <label>:2
	%3 = add i64 %0, %1
	ret i64 %3
}`,
			},
		},
		{
			name:     "using wrong argument to call function should failed",
			bindName: "add",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewFloat("3.3")),
				ast.NewArg("", ast.NewFloat("3.4")),
			},
			expectErrorMsg: "the type of argument doesn't match bind type requirement",
		},
	}

	g := codegen.New(bindMap, bindType)
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := g.Call(bindMap[testCase.bindName], testCase.args...)
			if testCase.expectErrorMsg != "" {
				require.Contains(t, err.Error(), testCase.expectErrorMsg)
				return
			}
			require.NoError(t, err)
			for _, expectedContain := range testCase.expectContains {
				assert.Contains(t, g.String(), expectedContain)
			}
		})
	}
}
