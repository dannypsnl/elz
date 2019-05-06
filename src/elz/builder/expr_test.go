package builder

import (
	"encoding/json"
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
		{
			name:         "empty list",
			code:         `i = []`,
			expectedExpr: ast.NewList(),
		},
		{
			name: "list with elements",
			code: `i = [1, 2, 3]`,
			expectedExpr: ast.NewList(
				ast.NewInt("1"),
				ast.NewInt("2"),
				ast.NewInt("3"),
			),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := New()
			builder.BuildFromCode(testCase.code)

			tree := builder.GetTree()
			b, err := tree.GetBinding("i")
			require.NoError(t, err)
			expected, err := json.Marshal(testCase.expectedExpr)
			require.NoError(t, err)
			actual, err := json.Marshal(b.Binding.Expr)
			require.NoError(t, err)
			assert.Equal(t, string(expected), string(actual))
		})
	}
}
