package codegen

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"
)

type typeMap struct {
	typeMap map[string]types.Type
}

func newTypeMap() *typeMap {
	return &typeMap{
		typeMap: map[string]types.Type{},
	}
}

func (t *typeMap) add(identifier string, typ types.Type) {
	t.typeMap[identifier] = typ
}

func (t *typeMap) getTypeOfExpr(expr ast.Expr) types.Type {
	if e, isIdentifier := expr.(*ast.Ident); isIdentifier {
		return t.typeMap[e.Literal]
	}
	return typeOfExpr(expr)
}

func (t *typeMap) convertArgsToTypeList(args ...*ast.Arg) []types.Type {
	typeList := make([]types.Type, 0)
	for _, arg := range args {
		typeList = append(typeList, getTypeOfArg(t.typeMap, arg))
	}
	return typeList
}
func getTypeOfArg(typeMap map[string]types.Type, arg *ast.Arg) types.Type {
	if ident, isIdent := arg.Expr.(*ast.Ident); isIdent {
		return typeMap[ident.Literal]
	}
	return typeOfExpr(arg.Expr)
}

func typeOfExpr(e ast.Expr) types.Type {
	// where e := e.(type) can save the convert in case clause
	switch e := e.(type) {
	case *ast.Arg:
		return typeOfExpr(e.Expr)
	case *ast.Int:
		return &types.Int{}
	case *ast.Float:
		return &types.Float{}
	case *ast.String:
		return &types.String{}
	default:
		panic(fmt.Sprintf("you can't use expression: `%#v` to get type directly", e))
	}
}
