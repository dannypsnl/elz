package builder

import (
	"github.com/elz-lang/elz/ast"
	"testing"

	asserter "github.com/dannypsnl/assert"
)

func TestNewBuilder(t *testing.T) {
	assert := asserter.NewTester(t)

	builder := New()
	builder.debug = true
	builder.BuildFromCode(`
main =
  println "Hello"
  |> add 1 2
  |> println "World" 1+2
  |> println [1, 2, 3]
`)

	assert.Eq(len(builder.bindings), 1)
	b0 := builder.bindings[0]
	assert.Eq(b0.Name, "main")
}

func TestBindingRule(t *testing.T) {
	assert := asserter.NewTester(t)
	t.Run("ValueBind", func(t *testing.T) {
		builder := New()
		builder.BuildFromCode("i = 1")
		binding := builder.bindings[0]
		assert.Eq(binding.Name, "i")
		assert.Eq(binding.ParamList, []string{})
		assert.Eq(binding.Expr.(*ast.Int), &ast.Int{Literal: "1"})
	})
	t.Run("WithParam", func(t *testing.T) {
		builder := New()
		builder.BuildFromCode("add x y = x + y")
		binding := builder.bindings[0]
		assert.Eq(binding.Name, "add")
		assert.Eq(binding.ParamList, []string{"x", "y"})
		assert.Eq(binding.Expr.(*ast.BinaryExpr).LExpr.(*ast.Ident),
			&ast.Ident{Value: "x"})
		assert.Eq(binding.Expr.(*ast.BinaryExpr).RExpr.(*ast.Ident),
			&ast.Ident{Value: "y"})
	})
}
