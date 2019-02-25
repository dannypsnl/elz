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
	mod      *ir.Module
	bindType map[string]*ast.BindType
	bindMap  map[string]*ast.Binding

	implsOfBinding       map[string]*ir.Func
	typeOfBuiltInBinding map[string]types.Type
	typeOfBinding        map[string]types.Type
}

func New(bindMap map[string]*ast.Binding, bindType map[string]*ast.BindType) *Generator {
	typMap := make(map[string]types.Type)
	typMap["+(int,int)"] = &types.Int{}
	return &Generator{
		mod:                  ir.NewModule(),
		bindType:             bindType,
		bindMap:              bindMap,
		implsOfBinding:       make(map[string]*ir.Func),
		typeOfBuiltInBinding: typMap,
		typeOfBinding:        typMap,
	}
}

func (g *Generator) String() string {
	return g.mod.String()
}

func (g *Generator) Generate() {
	entryBinding := g.bindMap["main"]
	if len(entryBinding.ParamList) > 0 {
		panic("main function should not have any parameters")
	}
	impl := g.mod.NewFunc("main", llvmtypes.I64)
	b := impl.NewBlock("")
	g.genExpr(b, entryBinding.Expr, make(map[string]*ir.Param), make(map[string]types.Type))
	b.NewRet(constant.NewInt(llvmtypes.I64, 0))
}

func (g *Generator) Call(bind *ast.Binding, exprList ...*ast.Arg) {
	g.mustGetImpl(bind, make(map[string]types.Type), exprList...)
}

func (g *Generator) mustGetImpl(bind *ast.Binding, typeMap map[string]types.Type, argList ...*ast.Arg) *ir.Func {
	bindName := bind.Name
	typeList := getTypeListFrom(typeMap, argList...)
	key := genKey(bindName, typeList...)
	impl, getImpl := g.implsOfBinding[key]
	if getImpl {
		return impl
	}
	if len(argList) != len(bind.ParamList) {
		panic(`do not have enough arguments to call function`)
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
		params = append(params, ir.NewParam(e.Ident, typeList[i].LLVMType()))
	}
	for i, t := range typeList {
		argName := bind.ParamList[i]
		typeMap[argName] = t
	}
	inferT := g.inferReturnType(bind.Expr, typeMap)
	g.typeOfBinding[key] = inferT
	f := g.mod.NewFunc(bindName, inferT.LLVMType(), params...)

	b := f.NewBlock("")
	binds := make(map[string]*ir.Param)
	for i, p := range params {
		binds[bind.ParamList[i]] = p
	}
	g.funcBody(b, bind.Expr, binds, typeMap)

	g.implsOfBinding[key] = f
	return f
}

// inference the return type by the expression we going to execute and input types
func (g *Generator) inferReturnType(expr ast.Expr, typeMap map[string]types.Type) types.Type {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, hasBind := g.bindMap[expr.FuncName]
		if hasBind {
			typeList := getTypeListFrom(typeMap, expr.ExprList...)
			typeMap := make(map[string]types.Type)
			for i, t := range typeList {
				argName := bind.ParamList[i]
				typeMap[argName] = t
			}
			inferT := g.inferReturnType(bind.Expr, typeMap)
			key := genKey(bind.Name, typeList...)
			g.typeOfBinding[key] = inferT
			t, exist := g.typeOfBind(genKey(expr.FuncName, typeList...))
			if exist {
				return t
			}
		}
		panic(fmt.Sprintf("can't find any binding call: %s", expr.FuncName))
	case *ast.BinaryExpr:
		lt := g.inferReturnType(expr.LExpr, typeMap)
		rt := g.inferReturnType(expr.RExpr, typeMap)
		op := expr.Op
		key := genKey(op, lt, rt)
		t, ok := g.typeOfBind(key)
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

func (g *Generator) isBuiltIn(key string) bool {
	_, isBuiltIn := g.typeOfBuiltInBinding[key]
	return isBuiltIn
}

func (g *Generator) typeOfBind(key string) (types.Type, bool) {
	t, existed := g.typeOfBuiltInBinding[key]
	if existed {
		return t, true
	}
	t, existed = g.typeOfBinding[key]
	if existed {
		return t, true
	}
	return nil, false
}

func (g *Generator) funcBody(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap map[string]types.Type) {
	v := g.genExpr(b, expr, binds, typeMap)
	b.NewRet(v)
}

func (g *Generator) genExpr(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap map[string]types.Type) value.Value {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind := g.bindMap[expr.FuncName]
		f := g.mustGetImpl(bind, typeMap, expr.ExprList...)
		valueList := make([]value.Value, 0)
		for _, arg := range expr.ExprList {
			e := g.genExpr(b, arg.Expr, binds, typeMap)
			valueList = append(valueList, e)
		}
		return b.NewCall(f, valueList...)
	case *ast.BinaryExpr:
		x := g.genExpr(b, expr.LExpr, binds, typeMap)
		y := g.genExpr(b, expr.RExpr, binds, typeMap)
		lt := getType(expr.LExpr, typeMap)
		rt := getType(expr.RExpr, typeMap)
		key := genKey(expr.Op, lt, rt)
		if g.isBuiltIn(key) {
			if lt.String() == "int" && rt.String() == "int" {
				switch expr.Op {
				case "+":
					return b.NewAdd(x, y)
				case "-":
					return b.NewSub(x, y)
				case "*":
					return b.NewMul(x, y)
				case "/":
					return b.NewSDiv(x, y)
				}
			}
		}
		panic(fmt.Sprintf("unsupported operator: %s", expr.Op))
	case *ast.Ident:
		v, exist := binds[expr.Literal]
		if exist {
			return v
		}
		panic(fmt.Sprintf("can't find any identifier: %s", expr.Literal))
	case *ast.Int:
		v, err := constant.NewIntFromString(llvmtypes.I64, expr.Literal)
		if err != nil {
		}
		return v
	default:
		panic(fmt.Sprintf("failed at generate expression: %#v", expr))
	}
}

func getType(e ast.Expr, typeMap map[string]types.Type) types.Type {
	if e, isIdentifier := e.(*ast.Ident); isIdentifier {
		return typeMap[e.Literal]
	}
	return types.TypeOf(e)
}

func genKey(bindName string, typeList ...types.Type) string {
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

func getTypeListFrom(typeMap map[string]types.Type, args ...*ast.Arg) []types.Type {
	typeList := make([]types.Type, 0)
	for _, arg := range args {
		var t types.Type
		if ident, isIdent := arg.Expr.(*ast.Ident); isIdent {
			t = typeMap[ident.Literal]
		} else {
			t = types.TypeOf(arg.Expr)
		}
		typeList = append(typeList, t)
	}
	return typeList
}
