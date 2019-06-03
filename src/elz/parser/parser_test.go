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

func TestParseBinaryExpression(t *testing.T) {
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
			primary, err := p.ParsePrimary()
			require.NoError(t, err)
			actual, err := p.ParseExpression(primary, 0)
			require.NoError(t, err)
			assert.Equal(t, tc.expectedExpr, actual)
		})
	}
}

func TestParseAccessChain(t *testing.T) {
	testCases := []struct {
		code         string
		expectedExpr *ast.Ident
	}{
		{
			code:         "a",
			expectedExpr: ast.NewIdent("a"),
		},
		{
			code:         "a::b",
			expectedExpr: ast.NewIdent("a::b"),
		},
		{
			code:         "a::b::c",
			expectedExpr: ast.NewIdent("a::b::c"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.code, func(t *testing.T) {
			p := parser.NewParser("test", tc.code)
			actual, err := p.ParseAccessChain()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedExpr, actual)
		})
	}
}

func TestParseStringLiteral(t *testing.T) {
	testCases := []struct {
		code         string
		expectedExpr *ast.String
	}{
		{
			code:         `"a"`,
			expectedExpr: ast.NewString(`"a"`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.code, func(t *testing.T) {
			p := parser.NewParser("test", tc.code)
			actual, err := p.ParsePrimary()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedExpr, actual)
		})
	}
}

func TestParseBooleanLiteral(t *testing.T) {
	testCases := []struct {
		code         string
		expectedExpr *ast.Bool
	}{
		{
			code:         `true`,
			expectedExpr: ast.NewBool("true"),
		},
		{
			code:         `false`,
			expectedExpr: ast.NewBool("false"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.code, func(t *testing.T) {
			p := parser.NewParser("test", tc.code)
			actual, err := p.ParsePrimary()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedExpr, actual)
		})
	}
}
