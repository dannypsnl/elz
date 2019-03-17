package codegen_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/builder"
	"github.com/elz-lang/elz/src/elz/codegen"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBindingCodegen(t *testing.T) {
	testCases := []struct {
		name           string
		code           string
		bindName       string
		args           []*ast.Arg
		expectContains []string
		expectErrorMsg string
	}{
		{
			name:     "call by generator",
			code:     `add x y = x + y`,
			bindName: "add",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewInt("1")),
				ast.NewArg("", ast.NewInt("2")),
			},
			expectContains: []string{`define i64 @add(i64 %x, i64 %y) {
; <label>:0
	%1 = add i64 %x, %y
	ret i64 %1
}`},
		},
		{
			name: "call function in function",
			code: `
add x y = x + y
addOne y = add(1, y)
`,
			bindName: "addOne",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewInt("2")),
			},
			expectContains: []string{
				`define i64 @addOne(i64 %y) {
; <label>:0
	%1 = call i64 @add(i64 1, i64 %y)
	ret i64 %1
}`,
				`define i64 @add(i64 %x, i64 %y) {
; <label>:0
	%1 = add i64 %x, %y
	ret i64 %1
}`,
			},
		},
		{
			name:     "using wrong argument to call function should failed",
			code:     `add x y = x + y`,
			bindName: "add",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewFloat("3.3")),
				ast.NewArg("", ast.NewFloat("3.4")),
			},
			expectErrorMsg: "can't infer return type",
		},
		{
			name: "with certain type limiter",
			code: `
add :: int -> int -> int
add x y = x + y
`,
			bindName: "add",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewFloat("3.3")),
				ast.NewArg("", ast.NewFloat("3.4")),
			},
			expectErrorMsg: "require type: `add :: int -> int` but get: `add :: f64 -> f64",
		},
		{
			name: "with variant type limiter",
			code: `
add :: 'a -> 'a -> 'a
add x y = x + y
`,
			bindName: "add",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewInt("3")),
				ast.NewArg("", ast.NewFloat("3.4")),
			},
			expectErrorMsg: "require type: `add :: 'a -> 'a` but get: `add :: int -> f64`",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			b := builder.New()
			b.BuildFromCode(testCase.code)
			tree := b.GetTree()

			g := codegen.New(tree, nil)
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
