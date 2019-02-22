package builder

import (
	"testing"

	"github.com/elz-lang/elz/ast"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	builder := New()
	builder.debug = true
	builder.BuildFromCode(`
main =
  println "Hello"
  |> add 1 2
  |> println "World" 1+2
  |> println [1, 2, 3]
`)

	assert.Equal(t, len(builder.bindings), 1)
	b0 := builder.bindings[0]
	assert.Equal(t, b0.Name, "main")
}

func TestBindingRule(t *testing.T) {
	testCases := []struct {
		name            string
		code            string
		expectedBinding *ast.Binding
	}{
		{
			name: "no param bind",
			code: `i = 1`,
			expectedBinding: &ast.Binding{
				Name:      "i",
				ParamList: []string{},
				Expr:      &ast.Int{Literal: "1"},
			},
		},
		{
			name: "with param",
			code: `add x y = x + y`,
			expectedBinding: &ast.Binding{
				Name:      "add",
				ParamList: []string{"x", "y"},
				Expr: &ast.BinaryExpr{
					Op:    "+",
					LExpr: &ast.Ident{Value: "x"},
					RExpr: &ast.Ident{Value: "y"},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := New()
			builder.BuildFromCode(testCase.code)
			b := builder.bindings[0]
			assert.Equal(t, testCase.expectedBinding, b)
		})
	}
}
