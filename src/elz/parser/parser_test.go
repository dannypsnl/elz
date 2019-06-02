package parser_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/parser"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBinding(t *testing.T) {
	testCases := []struct {
		code        string
		expectedAst *ast.Binding
	}{
		{
			code: `a = 1`,
			expectedAst: &ast.Binding{
				Name:      "a",
				ParamList: []string{},
				Expr:      ast.NewInt("1"),
			},
		},
		{
			code: `add x y = x + y`,
			expectedAst: &ast.Binding{
				Name:      "add",
				ParamList: []string{"x", "y"},
				Expr: &ast.BinaryExpr{
					LExpr: ast.NewIdent("x"),
					RExpr: ast.NewIdent("y"),
					Op:    "+",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.code, func(t *testing.T) {
			p := parser.NewParser("test", tc.code)
			actual, err := p.ParseBinding()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedAst, actual)
		})
	}
}

func TestParseExpression(t *testing.T) {
	testCases := []struct {
		code         string
		expectedExpr ast.Expr
	}{
		{
			code: `1 * 2`,
			expectedExpr: &ast.BinaryExpr{
				LExpr: ast.NewInt("1"),
				RExpr: ast.NewInt("2"),
				Op:    "*",
			},
		},
		{
			code: `1 / 2`,
			expectedExpr: &ast.BinaryExpr{
				LExpr: ast.NewInt("1"),
				RExpr: ast.NewInt("2"),
				Op:    "/",
			},
		},
		{
			code: `1 - 2`,
			expectedExpr: &ast.BinaryExpr{
				LExpr: ast.NewInt("1"),
				RExpr: ast.NewInt("2"),
				Op:    "-",
			},
		},
		{
			code: `1 * 2 - 3`,
			expectedExpr: &ast.BinaryExpr{
				LExpr: &ast.BinaryExpr{
					LExpr: ast.NewInt("1"),
					RExpr: ast.NewInt("2"),
					Op:    "*",
				},
				RExpr: ast.NewInt("3"),
				Op:    "-",
			},
		},
		{
			code: `1 * 2 - 3 / 4`,
			expectedExpr: &ast.BinaryExpr{
				LExpr: &ast.BinaryExpr{
					LExpr: ast.NewInt("1"),
					RExpr: ast.NewInt("2"),
					Op:    "*",
				},
				RExpr: &ast.BinaryExpr{
					LExpr: ast.NewInt("3"),
					RExpr: ast.NewInt("4"),
					Op:    "/",
				},
				Op: "-",
			},
		},
		{
			code: `1 * 2 - 3 / 4 + 5`,
			expectedExpr: &ast.BinaryExpr{
				LExpr: &ast.BinaryExpr{
					LExpr: &ast.BinaryExpr{
						LExpr: ast.NewInt("1"),
						RExpr: ast.NewInt("2"),
						Op:    "*",
					},
					RExpr: &ast.BinaryExpr{
						LExpr: ast.NewInt("3"),
						RExpr: ast.NewInt("4"),
						Op:    "/",
					},
					Op: "-",
				},
				RExpr: ast.NewInt("5"),
				Op:    "+",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.code, func(t *testing.T) {
			p := parser.NewParser("test", tc.code)
			actual := p.ParseExpression(p.ParsePrimary(), 0)
			require.NotNil(t, actual)
			assert.Equal(t, tc.expectedExpr, actual)
		})
	}
}
