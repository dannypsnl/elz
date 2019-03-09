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

	astTree *ast.Tree

	implsOfBinding       map[string]*ir.Func
	typeOfBuiltInBinding map[string]types.Type
	typeOfBinding        map[string]types.Type
}

func New(astTree *ast.Tree) *Generator {
	typMap := make(map[string]types.Type)
	typMap["+ :: int -> int"] = &types.Int{}

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
		astTree:              astTree,
		implsOfBinding:       builtInImpl,
		typeOfBuiltInBinding: typMap,
		typeOfBinding:        typMap,
	}
}

func (g *Generator) String() string {
	return g.mod.String()
}

func (g *Generator) Generate() {
	bind, err := g.astTree.GetBinding("main")
	if err != nil {
		panic("no main function exist, no compile")
	}
	entryBinding := bind
	if len(entryBinding.ParamList) > 0 {
		panic("main function should not have any parameters")
	}
	impl := g.mod.NewFunc("main", llvmtypes.I64)
	b := impl.NewBlock("")
	_, err = g.genExpr(b, entryBinding.Expr, make(map[string]*ir.Param), make(map[string]types.Type))
	if err != nil {
		panic(fmt.Sprintf("report error: %s", err))
	}
	b.NewRet(constant.NewInt(llvmtypes.I64, 0))
}

func (g *Generator) Call(bind *ast.Binding, exprList ...*ast.Arg) error {
	_, err := g.mustGetImpl(bind, make(map[string]types.Type), exprList...)
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) mustGetImpl(bind *ast.Binding, typeMap map[string]types.Type, argList ...*ast.Arg) (*ir.Func, error) {
	bindName := bind.Name
	// FIXME: currently for convenience we skip all checking when it's a built-in function
	// it should be fix after we can do more type checking
	if bind.NoImplementation {
		return g.implsOfBinding[bindName], nil
	}
	typeList := getTypeListFrom(typeMap, argList...)
	actualTypeWhenCall := genKey(bindName, typeList...)
	err := typeCheck(bindName, actualTypeWhenCall, bind.Type, typeList)
	if err != nil {
		return nil, err
	}

	impl, getImpl := g.implsOfBinding[actualTypeWhenCall]
	if getImpl {
		return impl, nil
	}
	return g.generateNewImpl(bind, typeMap, argList...)
}

func (g *Generator) generateNewImpl(bind *ast.Binding, typeMap map[string]types.Type, argList ...*ast.Arg) (*ir.Func, error) {
	typeList := getTypeListFrom(typeMap, argList...)
	actualTypeWhenCall := genKey(bind.Name, typeList...)
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
		typeMap[argName] = t
	}

	inferT, err := g.inferReturnType(bind.Expr, typeMap)
	if err != nil {
		return nil, err
	}

	g.typeOfBinding[actualTypeWhenCall] = inferT
	f := g.mod.NewFunc(bind.Name, inferT.LLVMType(), params...)

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
func (g *Generator) inferReturnType(expr ast.Expr, typeMap map[string]types.Type) (types.Type, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, err := g.astTree.GetBinding(expr.FuncName)
		if err != nil {
			return nil, err
		}
		typeList := getTypeListFrom(typeMap, expr.ExprList...)
		typeMap := make(map[string]types.Type)
		for i, t := range typeList {
			argName := bind.ParamList[i]
			typeMap[argName] = t
		}
		inferT, err := g.inferReturnType(bind.Expr, typeMap)
		if err != nil {
			return nil, err
		}
		key := genKey(bind.Name, typeList...)
		g.typeOfBinding[key] = inferT
		t, exist := g.typeOfBind(genKey(expr.FuncName, typeList...))
		if exist {
			return t, nil
		}
		return nil, fmt.Errorf("can't find any binding call: %s", expr.FuncName)
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
		t, ok := typeMap[expr.Literal]
		if !ok {
			return nil, fmt.Errorf("can't get type of identifier: %s", expr.Literal)
		}
		return t, nil
	default:
		return nil, fmt.Errorf("unsupported type inference for expression: %#v yet", expr)
	}
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

func (g *Generator) funcBody(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap map[string]types.Type) error {
	v, err := g.genExpr(b, expr, binds, typeMap)
	if err != nil {
		return err
	}
	b.NewRet(v)
	return nil
}

func (g *Generator) genExpr(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap map[string]types.Type) (value.Value, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, err := g.astTree.GetBinding(expr.FuncName)
		if err != nil {
			return nil, err
		}
		f, err := g.mustGetImpl(bind, typeMap, expr.ExprList...)
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
		lt := getType(expr.LExpr, typeMap)
		rt := getType(expr.RExpr, typeMap)
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

func getType(e ast.Expr, typeMap map[string]types.Type) types.Type {
	if e, isIdentifier := e.(*ast.Ident); isIdentifier {
		return typeMap[e.Literal]
	}
	return types.TypeOf(e)
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
