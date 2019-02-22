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
	t.Run("ValueBind", func(t *testing.T) {
		builder := New()
		builder.BuildFromCode("i = 1")
		binding := builder.bindings[0]
		expected := &ast.Binding{
			Name:      "i",
			ParamList: []string{},
			Expr:      &ast.Int{Literal: "1"},
		}
		assert.Equal(t, expected, binding)
	})
	t.Run("WithParam", func(t *testing.T) {
		builder := New()
		builder.BuildFromCode("add x y = x + y")
		binding := builder.bindings[0]
		expected := &ast.Binding{
			Name:      "add",
			ParamList: []string{"x", "y"},
			Expr: &ast.BinaryExpr{
				Op:    "+",
				LExpr: &ast.Ident{Value: "x"},
				RExpr: &ast.Ident{Value: "y"},
			},
		}
		assert.Equal(t, expected, binding)
	})
}
