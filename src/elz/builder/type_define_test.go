package builder

import (
	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTypeDefineAst(t *testing.T) {
	testCases := []struct {
		name               string
		code               string
		expectedTypeDefine *ast.NewType
	}{
		{
			name: "basic type define",
			code: `
type Car = (
  name: str,
  price: int,
)`,
			expectedTypeDefine: &ast.NewType{
				Name: "Car",
				Fields: []*ast.Field{
					ast.NewField("name", &ast.ExistType{Name: "str"}),
					ast.NewField("price", &ast.ExistType{Name: "int"}),
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := New()
			builder.BuildFromCode(testCase.code)

			tree := builder.GetTree()
			b, err := tree.GetTypeDefine(testCase.expectedTypeDefine.Name)
			require.NoError(t, err)

			for i, expectedField := range testCase.expectedTypeDefine.Fields {
				actualField := b.Fields[i]
				assert.Equal(t, expectedField.Name, actualField.Name)
				assert.Equal(t, expectedField.Type, actualField.Type)
			}
		})
	}
}
