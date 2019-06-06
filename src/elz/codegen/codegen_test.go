package codegen_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/codegen"
	"github.com/elz-lang/elz/src/elz/parser"

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
	}{
		{
			name:     "call by generator",
			code:     `add x y = x + y`,
			bindName: "add",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewInt("1")),
				ast.NewArg("", ast.NewInt("2")),
			},
			expectContains: []string{`define internal i64 @add(i64 %x, i64 %y) {
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
				`define internal i64 @addOne(i64 %y) {
; <label>:0
	%1 = call i64 @add(i64 1, i64 %y)
	ret i64 %1
}`,
				`define internal i64 @add(i64 %x, i64 %y) {
; <label>:0
	%1 = add i64 %x, %y
	ret i64 %1
}`,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			p := parser.NewParser("codegen-test", testCase.code)
			program, err := p.ParseProgram()
			require.NoError(t, err)
			tree, err := codegen.NewTree(program)
			require.NoError(t, err)

			binding, err := tree.GetBinding(testCase.bindName)
			require.NoError(t, err)
			generator := codegen.New(tree, nil)
			err = generator.Call(binding, testCase.args...)
			require.NoError(t, err)
			for _, expectedContain := range testCase.expectContains {
				assert.Contains(t, generator.String(), expectedContain)
			}
		})
	}
}

func TestErrorReporting(t *testing.T) {
	testCases := []struct {
		name           string
		code           string
		bindName       string
		args           []*ast.Arg
		expectErrorMsg string
	}{
		{
			name:     "can not inference return type would report error",
			code:     `add x y = x + y`,
			bindName: "add",
			args: []*ast.Arg{
				ast.NewArg("", ast.NewFloat("3.3")),
				ast.NewArg("", ast.NewInt("3")),
			},
			expectErrorMsg: "can't infer return type",
		},
		{
			name: "certain type limiter would reject type mismatched arguments",
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
			name: "variant type limiter would use first exact type as it's real type",
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
			p := parser.NewParser("codegen-test", testCase.code)
			program, err := p.ParseProgram()
			require.NoError(t, err)
			tree, err := codegen.NewTree(program)
			require.NoError(t, err)

			binding, err := tree.GetBinding(testCase.bindName)
			require.NoError(t, err)
			generator := codegen.New(tree, nil)
			err = generator.Call(binding, testCase.args...)
			require.Error(t, err, testCase.name)
			require.Contains(t, err.Error(), testCase.expectErrorMsg)
		})
	}
}
