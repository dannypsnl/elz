package codegen_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/builder"
	"github.com/elz-lang/elz/src/elz/codegen"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	bindingsCode = `
	addOne y = add(1, y)
	add x y = x + y
	`

	tree *ast.Tree
)

func init() {
	b := builder.New()
	b.BuildFromCode(bindingsCode)

	tree = b.GetTree()
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
			expectErrorMsg: "can't infer return type",
		},
	}

	g := codegen.New(tree)
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			binding, err := tree.GetBinding(testCase.bindName)
			require.NoError(t, err)
			err = g.Call(binding, testCase.args...)
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
