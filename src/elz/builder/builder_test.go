package builder

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/ast"

	"github.com/stretchr/testify/assert"
)

func TestBindingRule(t *testing.T) {
	testCases := []struct {
		name             string
		code             string
		expectedBinding  *ast.Binding
		expectedBindType *ast.BindType
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
			name: "type bind with existing types",
			code: `add :: int -> int -> int`,
			expectedBindType: &ast.BindType{
				Name: "add",
				Type: []ast.Type{
					&ast.ExistType{Name: "int"},
					&ast.ExistType{Name: "int"},
					&ast.ExistType{Name: "int"},
				},
			},
		},
		{
			name: "type bind with variant types",
			code: `assert :: 'that -> 'should_be -> ()`,
			expectedBindType: &ast.BindType{
				Name: "assert",
				Type: []ast.Type{
					&ast.VariantType{Name: "that"},
					&ast.VariantType{Name: "should_be"},
					&ast.VoidType{},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := New()
			builder.BuildFromCode(testCase.code)
			if testCase.expectedBinding != nil {
				bindMap := builder.GetBindMap()
				b := bindMap[testCase.expectedBinding.Name]
				assert.Equal(t, testCase.expectedBinding, b)
			}
			if testCase.expectedBindType != nil {
				bindTypes := builder.GetBindTypes()
				bt := bindTypes[testCase.expectedBindType.Name]
				assert.Equal(t, testCase.expectedBindType, bt)
			}
		})
	}
}
