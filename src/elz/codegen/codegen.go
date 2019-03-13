package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Generator struct {
	mod *ir.Module

	allTree   map[string]*ast.Tree
	entryTree *ast.Tree

	implsOfBinding       map[string]*ir.Func
	typeOfBuiltInBinding map[string]types.Type
	typeOfBinding        map[string]types.Type
}

func New(entryTree *ast.Tree, astAllTree map[string]*ast.Tree) *Generator {
	typMap := make(map[string]types.Type)
	typMap["+ :: int -> int"] = &types.Int{}

	astTree := entryTree
	err := astTree.InsertBinding(&ast.Binding{
		Name:             "printf",
		ParamList:        []string{"format"},
		NoImplementation: true,
	})
	if err != nil {
		panic("some function conflict with built-in function printf")
	}
	mod := ir.NewModule()
	builtInImpl := make(map[string]*ir.Func)
	printfImpl := mod.NewFunc("printf", llvmtypes.I64,
		ir.NewParam("format", llvmtypes.NewPointer(llvmtypes.I8)),
	)
	printfImpl.Sig.Variadic = true
	builtInImpl["printf"] = printfImpl

	return &Generator{
		mod:                  mod,
		entryTree:            astTree,
		allTree:              astAllTree,
		implsOfBinding:       builtInImpl,
		typeOfBuiltInBinding: typMap,
		typeOfBinding:        typMap,
	}
}

func (g *Generator) String() string {
	return g.mod.String()
}

func (g *Generator) Generate() {
	entryBinding, err := g.entryTree.GetBinding("main")
	if err != nil {
		panic("no main function exist, no compile")
	}
	if len(entryBinding.ParamList) > 0 {
		panic("main function should not have any parameters")
	}
	impl := g.mod.NewFunc("main", llvmtypes.I64)
	b := impl.NewBlock("")
	_, err = g.genExpr(b, entryBinding.Expr, make(map[string]*ir.Param), newTypeMap())
	if err != nil {
		panic(fmt.Sprintf("report error: %s", err))
	}
	b.NewRet(constant.NewInt(llvmtypes.I64, 0))
}

func (g *Generator) Call(bind *ast.Binding, exprList ...*ast.Arg) error {
	_, err := g.mustGetImpl(bind.Name, bind, newTypeMap(), exprList...)
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) mustGetImpl(accessChain string, bind *ast.Binding, typeMap *typeMap, argList ...*ast.Arg) (*ir.Func, error) {
	bindName := bind.Name
	// FIXME: currently for convenience we skip all checking when it's a built-in function
	// it should be fix after we can do more type checking
	if bind.NoImplementation {
		return g.implsOfBinding[bindName], nil
	}
	typeList := typeMap.convertArgsToTypeList(argList...)
	actualTypeWhenCall := genKey(bindName, typeList...)
	err := typeCheck(bindName, actualTypeWhenCall, bind.Type, typeList)
	if err != nil {
		return nil, err
	}

	impl, getImpl := g.implsOfBinding[actualTypeWhenCall]
	if getImpl {
		return impl, nil
	}
	return g.generateNewImpl(accessChain, bind, typeMap, argList...)
}

func (g *Generator) generateNewImpl(accessChain string, bind *ast.Binding, typeMap *typeMap, argList ...*ast.Arg) (*ir.Func, error) {
	typeList := typeMap.convertArgsToTypeList(argList...)
	actualTypeWhenCall := genKey(accessChain, typeList...)
	if len(argList) != len(bind.ParamList) {
		return nil, fmt.Errorf(`do not have enough arguments to evaluate binding: %s, argList: %#v`, bind.Name, argList)
	}
	params := make([]*ir.Param, 0)
	for i, arg := range argList {
		argNameMustBe := bind.ParamList[i]
		argName := arg.Ident
		// allow ignore argument name like: `add(1, 2)`
		if argName == "" {
			argName = argNameMustBe
		}
		if argNameMustBe != argName {
			return nil, fmt.Errorf(`argument name must be parameter name(or empty), for example:
  assert that should_be = ...
  assert(that: 1+2, should_be: 3)
`)
		}
		params = append(params, ir.NewParam(arg.Ident, typeList[i].LLVMType()))
	}
	for i, t := range typeList {
		argName := bind.ParamList[i]
		typeMap.add(argName, t)
	}

	inferT, err := g.inferReturnType(bind.Expr, typeMap)
	if err != nil {
		return nil, err
	}

	g.typeOfBinding[actualTypeWhenCall] = inferT
	f := g.mod.NewFunc(accessChain, inferT.LLVMType(), params...)

	b := f.NewBlock("")
	binds := make(map[string]*ir.Param)
	for i, p := range params {
		binds[bind.ParamList[i]] = p
	}
	if err := g.funcBody(b, bind.Expr, binds, typeMap); err != nil {
		return nil, err
	}

	g.implsOfBinding[actualTypeWhenCall] = f
	return f, nil
}

func typeCheck(bindName, actualTypeWhenCall string, bindType []ast.Type, typeList []types.Type) error {
	if bindType == nil {
		return nil
	}
	var (
		b              strings.Builder
		err            error
		variantTypeMap = map[string]string{}
	)
	for i, requireT := range bindType[:len(bindType)-1] {
		actualType := typeList[i]
		switch requireT := requireT.(type) {
		case *ast.ExistType:
			if requireT.Name != actualType.String() {
				err = fmt.Errorf("")
			}
		case *ast.VariantType:
			t, exist := variantTypeMap[requireT.Name]
			if !exist {
				variantTypeMap[requireT.Name] = actualType.String()
				t = actualType.String()
			}
			if t != actualType.String() {
				err = fmt.Errorf("")
			}
		case *ast.VoidType:
		}
		b.WriteString(requireT.String())
		b.WriteString(" -> ")
	}
	requireT := b.String()
	requireType := bindName + " :: " + requireT[:len(requireT)-4]
	if err != nil {
		return fmt.Errorf("require type: `%s` but get: `%s`", requireType, actualTypeWhenCall)
	}
	return nil
}

// inference the return type by the expression we going to execute and input types
func (g *Generator) inferReturnType(expr ast.Expr, typeMap *typeMap) (types.Type, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, err := g.getBindingByAccessChain(expr.AccessChain)
		if err != nil {
			return nil, err
		}
		typeList := typeMap.convertArgsToTypeList(expr.ExprList...)
		typeMap := newTypeMap()
		for i, t := range typeList {
			argName := bind.ParamList[i]
			typeMap.add(argName, t)
		}
		inferT, err := g.inferReturnType(bind.Expr, typeMap)
		if err != nil {
			return nil, err
		}
		key := genKey(bind.Name, typeList...)
		g.typeOfBinding[key] = inferT
		t, exist := g.typeOfBind(genKey(expr.AccessChain, typeList...))
		if exist {
			return t, nil
		}
		return nil, fmt.Errorf("can't find any binding call: %s", expr.AccessChain)
	case *ast.BinaryExpr:
		lt, err := g.inferReturnType(expr.LExpr, typeMap)
		if err != nil {
			return nil, err
		}
		rt, err := g.inferReturnType(expr.RExpr, typeMap)
		if err != nil {
			return nil, err
		}
		op := expr.Op
		key := genKey(op, lt, rt)
		t, ok := g.typeOfBind(key)
		if !ok {
			return nil, fmt.Errorf("can't infer return type by %s", key)
		}
		return t, nil
	case *ast.Ident:
		t := typeMap.getTypeOfExpr(expr)
		if t == nil {
			return nil, fmt.Errorf("can't get type of identifier: %s", expr.Literal)
		}
		return t, nil
	default:
		return nil, fmt.Errorf("unsupported type inference for expression: %#v yet", expr)
	}
}

func (g *Generator) getBindingByAccessChain(accessChain string) (*ast.Binding, error) {
	chain := strings.Split(accessChain, "::")
	if len(chain) == 2 {
		moduleName := chain[0]
		funcName := chain[1]
		return g.allTree[moduleName].GetExportBinding(funcName)
	}
	if len(chain) == 1 {
		return g.entryTree.GetBinding(accessChain)
	}
	return nil, fmt.Errorf("not supported access chain: %s", accessChain)
}

func (g *Generator) isBuiltInFunction(key string) bool {
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

func (g *Generator) funcBody(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap *typeMap) error {
	v, err := g.genExpr(b, expr, binds, typeMap)
	if err != nil {
		return err
	}
	b.NewRet(v)
	return nil
}

func (g *Generator) genExpr(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap *typeMap) (value.Value, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, err := g.getBindingByAccessChain(expr.AccessChain)
		if err != nil {
			return nil, err
		}
		f, err := g.mustGetImpl(expr.AccessChain, bind, typeMap, expr.ExprList...)
		if err != nil {
			return nil, err
		}
		valueList := make([]value.Value, 0)
		for _, arg := range expr.ExprList {
			e, err := g.genExpr(b, arg.Expr, binds, typeMap)
			if err != nil {
				return nil, err
			}
			valueList = append(valueList, e)
		}
		return b.NewCall(f, valueList...), nil
	case *ast.BinaryExpr:
		x, err := g.genExpr(b, expr.LExpr, binds, typeMap)
		if err != nil {
			return nil, err
		}
		y, err := g.genExpr(b, expr.RExpr, binds, typeMap)
		if err != nil {
			return nil, err
		}
		lt := typeMap.getTypeOfExpr(expr.LExpr)
		rt := typeMap.getTypeOfExpr(expr.RExpr)
		key := genKey(expr.Op, lt, rt)
		if g.isBuiltInFunction(key) {
			if lt.String() == "int" && rt.String() == "int" {
				switch expr.Op {
				case "+":
					return b.NewAdd(x, y), nil
				case "-":
					return b.NewSub(x, y), nil
				case "*":
					return b.NewMul(x, y), nil
				case "/":
					return b.NewSDiv(x, y), nil
				}
			}
		}
		return nil, fmt.Errorf("unsupported operator: %s", expr.Op)
	case *ast.Ident:
		v, exist := binds[expr.Literal]
		if exist {
			return v, nil
		}
		return nil, fmt.Errorf("can't find any identifier: %s", expr.Literal)
	case *ast.Int:
		v, err := constant.NewIntFromString(llvmtypes.I64, expr.Literal)
		if err != nil {
			return nil, err
		}
		return v, nil
	case *ast.String:
		str := g.mod.NewGlobal("", llvmtypes.NewArray(uint64(len(expr.Literal)), llvmtypes.I8))
		str.Align = 1
		str.Init = constant.NewCharArrayFromString(expr.Literal)
		strGEP := b.NewGetElementPtr(str,
			constant.NewInt(llvmtypes.I64, 0),
			constant.NewInt(llvmtypes.I64, 0),
		)
		x := b.NewAlloca(llvmtypes.NewPointer(llvmtypes.I8))
		b.NewStore(strGEP, x)
		return b.NewLoad(x), nil
	default:
		return nil, fmt.Errorf("failed at generate expression: %#v", expr)
	}
}

func genKey(bindName string, typeList ...types.Type) string {
	var b strings.Builder
	b.WriteString(bindName)
	b.WriteString(" :: ")
	fmtTypes := make([]fmt.Stringer, 0)
	for _, t := range typeList {
		fmtTypes = append(fmtTypes, t)
	}
	b.WriteString(typeFormat(fmtTypes...))
	return b.String()
}

func typeFormat(typeList ...fmt.Stringer) string {
	var b strings.Builder
	if len(typeList) > 0 {
		for _, t := range typeList[:len(typeList)-1] {
			b.WriteString(t.String())
			b.WriteString(" -> ")
		}
		b.WriteString(typeList[len(typeList)-1].String())
	}
	return b.String()
}
