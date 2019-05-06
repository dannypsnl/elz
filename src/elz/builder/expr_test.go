package builder

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/ast"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpr(t *testing.T) {
	testCases := []struct {
		name         string
		code         string
		expectedExpr ast.Expr
	}{
		{
			name:         "no param bind",
			code:         `i = 1`,
			expectedExpr: ast.NewInt("1"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := New()
			builder.BuildFromCode(testCase.code)

			tree := builder.GetTree()
			b, err := tree.GetBinding("i")
			require.NoError(t, err)
			assert.Equal(t, testCase.expectedExpr, b.Binding.Expr)
		})
	}
}
