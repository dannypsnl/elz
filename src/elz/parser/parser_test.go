package parser_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/parser"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	p := parser.NewParser("test", "a = 1")
	binding, err := p.ParseBinding()
	require.NoError(t, err)
	assert.Equal(t, "a", binding.Name)
	assert.Equal(t, &ast.Int{Literal: "1"}, binding.Expr)
}
