package builder

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/ast"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
				Expr:      ast.NewInt("1"),
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
					LExpr: ast.NewIdent("x"),
					RExpr: ast.NewIdent("y"),
				},
			},
		},
		{
			name: "function call",
			code: `addOne y = add(1, y)`,
			expectedBinding: &ast.Binding{
				Name:      "addOne",
				ParamList: []string{"y"},
				Expr: &ast.FuncCall{
					FuncName: "add",
					ExprList: []*ast.Arg{
						{Expr: ast.NewInt("1")},
						{Expr: ast.NewIdent("y")},
					},
				},
			},
		},
		{
			name: "function call with arg name",
			code: `tests = assert(that: 1, should_be: 1)`,
			expectedBinding: &ast.Binding{
				Name:      "tests",
				ParamList: []string{},
				Expr: &ast.FuncCall{
					FuncName: "assert",
					ExprList: []*ast.Arg{
						{Ident: "that", Expr: ast.NewInt("1")},
						{Ident: "should_be", Expr: ast.NewInt("1")},
					},
				},
			},
		},
		{
			name: "function call as argument",
			code: `main = printf("add(1, 2) = %d", add(1, 2))`,
			expectedBinding: &ast.Binding{
				Name:      "main",
				ParamList: []string{},
				Expr: &ast.FuncCall{
					FuncName: "printf",
					ExprList: []*ast.Arg{
						{Ident: "", Expr: ast.NewString(`"add(1, 2) = %d"`)},
						{Ident: "", Expr: &ast.FuncCall{
							FuncName: "add",
							ExprList: []*ast.Arg{
								{Expr: ast.NewInt("1")},
								{Expr: ast.NewInt("2")},
							},
						}},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := New()
			builder.BuildFromCode(testCase.code)
			if testCase.expectedBinding != nil {
				tree := builder.GetTree()
				b, err := tree.GetBinding(testCase.expectedBinding.Name)
				require.NoError(t, err)
				assert.Equal(t, testCase.expectedBinding, b)
			}
		})
	}
}
