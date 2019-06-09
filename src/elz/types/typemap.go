package types

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"
)

type TypeMap struct {
	// parent is optional, when it's nil the type map is root entry
	parent  *TypeMap
	typeMap map[string]Type
}

func NewTypeMap(parent *TypeMap) *TypeMap {
	return &TypeMap{
		parent:  parent,
		typeMap: map[string]Type{},
	}
}

func (t *TypeMap) Add(identifier string, typ Type) {
	t.typeMap[identifier] = typ
}

func (t *TypeMap) GetTypeOfExpr(expr ast.Expr) (Type, error) {
	if ident, isIdentifier := expr.(*ast.Ident); isIdentifier {
		typ, ok := t.typeMap[ident.Literal]
		if !ok {
			if t.parent != nil {
				return t.parent.GetTypeOfExpr(expr)
			}
			return nil, fmt.Errorf("can't get type of identifier: %s", ident.Literal)
		}
		return typ, nil
	} else if funcCall, isFuncCall := expr.(*ast.FuncCall); isFuncCall {
		funcTyp, err := t.GetTypeOfExpr(funcCall.Func)
		if err != nil {
			return nil, err
		}
		return funcTyp, nil
	}
	return typeOfExpr(expr)
}

func (t *TypeMap) ConvertArgsToTypeList(args ...*ast.Arg) ([]Type, error) {
	typeList := make([]Type, 0)
	for _, arg := range args {
		typ, err := getTypeOfArg(t.typeMap, arg)
		if err != nil {
			return nil, err
		}
		typeList = append(typeList, typ)
	}
	return typeList, nil
}
func getTypeOfArg(typeMap map[string]Type, arg *ast.Arg) (Type, error) {
	if ident, isIdent := arg.Expr.(*ast.Ident); isIdent {
		return typeMap[ident.Literal], nil
	}
	return typeOfExpr(arg.Expr)
}

func typeOfExpr(e ast.Expr) (Type, error) {
	// where e := e.(type) can save the convert in case clause
	switch e := e.(type) {
	case *ast.Arg:
		return typeOfExpr(e.Expr)
	case *ast.Int:
		return &Int{}, nil
	case *ast.Float:
		return &Float{}, nil
	case *ast.String:
		return &String{}, nil
	case *ast.List:
		elemT, err := typeOfExpr(e.ExprList[0])
		if err != nil {
			return nil, err
		}
		return &List{
			ElemT: elemT,
		}, nil
	default:
		return nil, fmt.Errorf("you can't use expression: `%#v` to get type directly", e)
	}
}
