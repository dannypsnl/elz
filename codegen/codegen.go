package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/types"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Generator struct {
	mod *ir.Module

	implsOfBinding map[string]*ir.Func
	typeOfBinding  map[string]types.Type
}

func New() *Generator {
	typMap := make(map[string]types.Type)
	typMap["+(int,int)"] = &types.Int{}
	return &Generator{
		mod:            ir.NewModule(),
		implsOfBinding: make(map[string]*ir.Func),
		typeOfBinding:  typMap,
	}
}

func (g *Generator) String() string {
	return g.mod.String()
}

func (g *Generator) Call(bind *ast.Binding, exprList ...*ast.Arg) {
	g.mustGetImpl(bind, exprList...)
}

func (g *Generator) mustGetImpl(bind *ast.Binding, argList ...*ast.Arg) *ir.Func {
	bindName := bind.Name
	key := genKeyByArg(bindName, argList...)
	impl, getImpl := g.implsOfBinding[key]
	if getImpl {
		return impl
	}
	if len(argList) != len(bind.ParamList) {
		panic(`do not have enough arguments to call function`)
	}
	binds := make(map[string]ast.Expr)
	typeList := make([]types.Type, 0)
	for _, e := range argList {
		typeList = append(typeList, types.TypeOf(e))
	}
	params := make([]*ir.Param, 0)
	for i, e := range argList {
		argNameMustBe := bind.ParamList[i]
		argName := e.Ident
		// allow ignore argument name like: `add(1, 2)`
		if argName == "" {
			argName = argNameMustBe
		}
		if argNameMustBe != argName {
			panic(`argument name must be parameter name(or empty), for example:
  assert that should_be = ...
  assert(that: 1+2, should_be: 3)
`)
		}
		binds[argName] = e.Expr
		params = append(params, ir.NewParam(e.Ident, typeList[i].LLVMType()))
	}
	typeMap := make(map[string]types.Type)
	for i, t := range typeList {
		argName := bind.ParamList[i]
		typeMap[argName] = t
	}
	inferT := g.inferReturnType(bind.Expr, typeMap)
	f := g.mod.NewFunc(bindName, inferT.LLVMType(), params...)

	b := f.NewBlock("")
	funcBody(b, bind.Expr, binds)

	g.implsOfBinding[key] = f
	return f
}

// inference the return type by the expression we going to execute and input types
func (g *Generator) inferReturnType(expr ast.Expr, typeMap map[string]types.Type) types.Type {
	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		lt := g.inferReturnType(expr.LExpr, typeMap)
		rt := g.inferReturnType(expr.RExpr, typeMap)
		op := expr.Op
		key := genKeyByTypes(op, lt, rt)
		t, ok := g.typeOfBinding[key]
		if !ok {
			panic(fmt.Sprintf("can't infer return type by %s", key))
		}
		return t
	case *ast.Ident:
		t, ok := typeMap[expr.Literal]
		if !ok {
			panic(fmt.Sprintf("can't get type of identifier: %s", expr.Literal))
		}
		return t
	default:
		panic(fmt.Sprintf("unsupported type inference for expression: %#v yet", expr))
	}
}

func funcBody(b *ir.Block, expr ast.Expr, binds map[string]ast.Expr) {
	v := genExpr(b, expr, binds)
	b.NewRet(v)
}

func genExpr(b *ir.Block, expr ast.Expr, binds map[string]ast.Expr) value.Value {
	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		if expr.Op == "+" {
			x := genExpr(b, expr.LExpr, binds)
			y := genExpr(b, expr.RExpr, binds)
			return b.NewAdd(x, y)
		}
		panic(fmt.Sprintf("unsupported operator: %s", expr.Op))
	case *ast.Ident:
		e := binds[expr.Literal]
		return genExpr(b, e, binds)
	case *ast.Int:
		v, err := constant.NewIntFromString(llvmtypes.I64, expr.Literal)
		if err != nil {
		}
		return v
	default:
		panic(fmt.Sprintf("failed at generate expression: %#v", expr))
	}
}

func genKeyByTypes(bindName string, typeList ...types.Type) string {
	var b strings.Builder
	b.WriteString(bindName)
	if len(typeList) > 0 {
		b.WriteRune('(')
		for _, t := range typeList[:len(typeList)-1] {
			b.WriteString(t.String())
			b.WriteRune(',')
		}
		b.WriteString(typeList[len(typeList)-1].String())
		b.WriteRune(')')
	}
	return b.String()
}

func genKey(bindName string, exprList ...ast.Expr) string {
	typeList := make([]types.Type, 0)
	for _, e := range exprList {
		typeList = append(typeList, types.TypeOf(e))
	}
	return genKeyByTypes(bindName, typeList...)
}

func genKeyByArg(bindName string, argList ...*ast.Arg) string {
	exprList := make([]ast.Expr, len(argList))
	for i, arg := range argList {
		exprList[i] = arg
	}
	return genKey(bindName, exprList...)
}
