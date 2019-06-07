package parser_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/parser"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseImport(t *testing.T) {
	testCases := []struct {
		code        string
		expectedAst *ast.Import
	}{
		{
			code: `import a`,
			expectedAst: &ast.Import{
				AccessChain: ast.NewIdent("a"),
			},
		},
		{
			code: `import a::b::c`,
			expectedAst: &ast.Import{
				AccessChain: ast.NewIdent("a::b::c"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.code, func(t *testing.T) {
			p := parser.NewParser("test", tc.code)
			actual, err := p.ParseImport()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedAst, actual)
		})
	}
}

func TestParseType(t *testing.T) {
	testCases := []struct {
		code        string
		expectedAst *ast.TypeDefine
	}{
		{
			code:        `type Foo = ()`,
			expectedAst: ast.NewTypeDefine(true, "Foo"),
		},
		{
			code:        `type _foo = ()`,
			expectedAst: ast.NewTypeDefine(false, "_foo"),
		},
		{
			code: `
type Foo = (
  a: i32,
  b: i32
)`,
			expectedAst: ast.NewTypeDefine(
				true,
				"Foo",
				ast.NewField("a", &ast.ExistType{Name: "i32"}),
				ast.NewField("b", &ast.ExistType{Name: "i32"}),
			),
		},
		{
			code: `
type Foo = (
  a: i32,
  b: i32,
)`,
			expectedAst: ast.NewTypeDefine(
				true,
				"Foo",
				ast.NewField("a", &ast.ExistType{Name: "i32"}),
				ast.NewField("b", &ast.ExistType{Name: "i32"}),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.code, func(t *testing.T) {
			p := parser.NewParser("test", tc.code)
			actual, err := p.ParseTypeDefine()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedAst, actual)
		})
	}
}

func TestParseBindingType(t *testing.T) {
	testCases := []struct {
		code        string
		expectedAst *ast.BindingType
	}{
		{
			code: `a :: i32`,
			expectedAst: &ast.BindingType{
				Name: "a",
				Type: []ast.Type{
					&ast.ExistType{Name: "i32"},
				},
			},
		},
		{
			code: `a :: i32 -> i32`,
			expectedAst: &ast.BindingType{
				Name: "a",
				Type: []ast.Type{
					&ast.ExistType{Name: "i32"},
					&ast.ExistType{Name: "i32"},
				},
			},
		},
		{
			code: `a :: 'a -> 'b`,
			expectedAst: &ast.BindingType{
				Name: "a",
				Type: []ast.Type{
					&ast.VariantType{Name: "a"},
					&ast.VariantType{Name: "b"},
				},
			},
		},
		{
			code: `a :: ()`,
			expectedAst: &ast.BindingType{
				Name: "a",
				Type: []ast.Type{
					&ast.VoidType{},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.code, func(t *testing.T) {
			p := parser.NewParser("test", tc.code)
			actual, err := p.ParseBindingType()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedAst, actual)
		})
	}
}

func TestParseBinding(t *testing.T) {
	testCases := []struct {
		code        string
		expectedAst *ast.Binding
	}{
		{
			code: `a = 1`,
			expectedAst: ast.NewBinding(
				false,
				true,
				"a",
				[]string{},
				ast.NewInt("1"),
			),
		},
		{
			code: `a() = 1`,
			expectedAst: ast.NewBinding(
				true,
				true,
				"a",
				[]string{},
				ast.NewInt("1"),
			),
		},
		{
			code: `add x y = x + y`,
			expectedAst: ast.NewBinding(
				true,
				true,
				"add",
				[]string{"x", "y"},
				&ast.BinaryExpr{
					LExpr: ast.NewIdent("x"),
					RExpr: ast.NewIdent("y"),
					Op:    "+",
				},
			),
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
		{
			code: `1 * (2 - 3)`,
			expectedExpr: &ast.BinaryExpr{
				LExpr: ast.NewInt("1"),
				RExpr: &ast.BinaryExpr{
					LExpr: ast.NewInt("2"),
					RExpr: ast.NewInt("3"),
					Op:    "-",
				},
				Op: "*",
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

func TestParseListLiteral(t *testing.T) {
	testCases := []struct {
		code         string
		expectedExpr *ast.List
	}{
		{
			code:         `[]`,
			expectedExpr: ast.NewList(),
		},
		{
			code: `[1, 2, 3]`,
			expectedExpr: ast.NewList(
				ast.NewInt("1"),
				ast.NewInt("2"),
				ast.NewInt("3"),
			),
		},
		{
			code: `[1, [2]]`,
			expectedExpr: ast.NewList(
				ast.NewInt("1"),
				ast.NewList(
					ast.NewInt("2"),
				),
			),
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

func TestParseElementAccess(t *testing.T) {
	testCases := []struct {
		code         string
		expectedExpr *ast.ExtractElement
	}{
		{
			code: "a[1]",
			expectedExpr: &ast.ExtractElement{
				X:   ast.NewIdent("a"),
				Key: ast.NewInt("1"),
			},
		},
		{
			code: "a[b[1]]",
			expectedExpr: &ast.ExtractElement{
				X: ast.NewIdent("a"),
				Key: &ast.ExtractElement{
					X:   ast.NewIdent("b"),
					Key: ast.NewInt("1"),
				},
			},
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

func TestParseFunctionCall(t *testing.T) {
	testCases := []struct {
		name         string
		code         string
		expectedExpr *ast.FuncCall
	}{
		{
			name: "no argument",
			code: "a()",
			expectedExpr: &ast.FuncCall{
				Func:    ast.NewIdent("a"),
				ArgList: []*ast.Arg{},
			},
		},
		{
			name: "2 arguments, with optional end comma",
			code: "a(1, 2,)",
			expectedExpr: &ast.FuncCall{
				Func: ast.NewIdent("a"),
				ArgList: []*ast.Arg{
					ast.NewArg("", ast.NewInt("1")),
					ast.NewArg("", ast.NewInt("2")),
				},
			},
		},
		{
			name: "2 arguments, one with name",
			code: "a(left: 1, 2)",
			expectedExpr: &ast.FuncCall{
				Func: ast.NewIdent("a"),
				ArgList: []*ast.Arg{
					ast.NewArg("left", ast.NewInt("1")),
					ast.NewArg("", ast.NewInt("2")),
				},
			},
		},
		{
			name: "2 arguments with name",
			code: "a(left: 1, right: 2)",
			expectedExpr: &ast.FuncCall{
				Func: ast.NewIdent("a"),
				ArgList: []*ast.Arg{
					ast.NewArg("left", ast.NewInt("1")),
					ast.NewArg("right", ast.NewInt("2")),
				},
			},
		},
		{
			name: "3 arguments, 2 with name",
			code: "a(first: 1, mid: 2, 3)",
			expectedExpr: &ast.FuncCall{
				Func: ast.NewIdent("a"),
				ArgList: []*ast.Arg{
					ast.NewArg("first", ast.NewInt("1")),
					ast.NewArg("mid", ast.NewInt("2")),
					ast.NewArg("", ast.NewInt("3")),
				},
			},
		},
		{
			name: "nesting",
			code: "a(b(), c())",
			expectedExpr: &ast.FuncCall{
				Func: ast.NewIdent("a"),
				ArgList: []*ast.Arg{
					ast.NewArg("", &ast.FuncCall{
						Func:    ast.NewIdent("b"),
						ArgList: []*ast.Arg{},
					}),
					ast.NewArg("", &ast.FuncCall{
						Func:    ast.NewIdent("c"),
						ArgList: []*ast.Arg{},
					}),
				},
			},
		},
		{
			name: "level 3 nesting",
			code: "a(b(c()))",
			expectedExpr: &ast.FuncCall{
				Func: ast.NewIdent("a"),
				ArgList: []*ast.Arg{
					ast.NewArg("", &ast.FuncCall{
						Func: ast.NewIdent("b"),
						ArgList: []*ast.Arg{
							ast.NewArg("", &ast.FuncCall{
								Func:    ast.NewIdent("c"),
								ArgList: []*ast.Arg{},
							}),
						},
					}),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := parser.NewParser("test", tc.code)
			actual, err := p.ParsePrimary()
			require.NoError(t, err)
			//actual, err := p.ParseExpression(primary, 0)
			//require.NoError(t, err)
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

func TestParseProgram(t *testing.T) {
	code := `import foo::bar
type Foo = ()
add :: i32 -> i32 -> i32
add x y = x + y
add_one y = add(1, y)
`
	p := parser.NewParser("test", code)
	program, err := p.ParseProgram()
	require.NoError(t, err)
	assert.Equal(t, 1, len(program.Imports))
	assert.Equal(t, 1, len(program.TypeDefines))
	assert.Equal(t, 1, len(program.BindingTypes))
	assert.Equal(t, 2, len(program.Bindings))
}
