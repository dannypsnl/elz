package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"
	"github.com/elz-lang/elz/src/elz/value"
	"github.com/elz-lang/elz/src/irutil"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmtypes "github.com/llir/llvm/ir/types"
	llvmvalue "github.com/llir/llvm/ir/value"

	"github.com/sirupsen/logrus"
)

type module struct {
	*Tree
	generator     *Generator
	typeMap       *types.TypeMap
	initFunc      *ir.Func
	initFuncBlock *ir.Block
	imports       map[string]string
}

// ```
// import math
// import mod::sub_mod
//
// // access
// math::abs(-1)
// sub_mod::xxx()
// ```
//
// import format is `lib::lib::lib`,
// but would only take last name as local name of the module
func newModule(g *Generator, tree *Tree) *module {
	imports := map[string]string{}
	for _, importPath := range tree.imports {
		accessChain := strings.Split(importPath, "::")
		lastOne := len(accessChain) - 1
		accessKey := accessChain[lastOne]
		if mod1, exist := imports[accessKey]; exist {
			logrus.Fatalf(`import %s
and
import %s
has the same name in the module`, mod1, importPath)
		}
		imports[accessKey] = importPath
	}

	// init func declaration of module
	initFunc := g.mod.NewFunc("init", llvmtypes.Void)
	m := &module{
		Tree:          tree,
		generator:     g,
		typeMap:       types.NewTypeMap(nil),
		initFunc:      initFunc,
		initFuncBlock: initFunc.NewBlock(""),
		imports:       imports,
	}
	// init type map, we would add binding type at here
	for _, b := range tree.bindings {
		m.typeMap.Add(b.Name, types.NewBindingType(m, b.Binding))
	}
	for _, bind := range tree.bindings {
		bind.SetModule(m)
	}
	return m
}

// inference the return type by the expression we going to execute and input types
func (m *module) InferTypeOf(expr ast.Expr, typeMap *types.TypeMap) (types.Type, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		funcType, err := m.InferTypeOf(expr.Func, typeMap)
		if err != nil {
			return nil, err
		}
		bindingType, ok := funcType.(*types.BindingType)
		if !ok {
			return nil, fmt.Errorf("call a non-callable expression: %s", expr.Func)
		}
		typeList, err := typeMap.ConvertArgsToTypeList(expr.ArgList...)
		if err != nil {
			return nil, err
		}
		typeMap := types.NewTypeMap(typeMap)
		for i, paramType := range typeList {
			paramName := bindingType.ParamList[i]
			typeMap.Add(paramName, paramType)
		}
		return bindingType.GetReturnType(typeMap, typeList...)
	case *ast.BinaryExpr:
		lt, err := m.InferTypeOf(expr.LExpr, typeMap)
		if err != nil {
			return nil, err
		}
		rt, err := m.InferTypeOf(expr.RExpr, typeMap)
		if err != nil {
			return nil, err
		}
		op := expr.Op
		t, err := m.generator.typeOfOperator(op, lt, rt)
		if err != nil {
			return nil, err
		}
		return t, nil
	case *ast.Ident:
		return typeMap.GetTypeOfExpr(expr)
	case *ast.List:
		elemT, err := m.InferTypeOf(expr.ExprList[0], typeMap)
		if err != nil {
			return nil, err
		}
		return types.NewList(elemT), nil
	case *ast.ExtractElement:
		list, ok := expr.X.(*ast.List)
		if !ok {
			return nil, fmt.Errorf("extract element do not support ast: %#v", expr.X)
		}
		listT, err := m.InferTypeOf(list, typeMap)
		if err != nil {
			return nil, err
		}
		keyT, err := m.InferTypeOf(expr.Key, typeMap)
		if err != nil {
			return nil, err
		}
		kt, ok := keyT.(*types.Int)
		if !ok {
			return nil, fmt.Errorf("extract element do not support key type: %s", kt)
		}
		return listT.(*types.List).ElemT, nil
	default:
		return nil, fmt.Errorf("unsupported type inference for expression: %#v yet", expr)
	}
}

func (m *module) genExpr(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap *types.TypeMap) (value.Value, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		v, err := m.genExpr(b, expr.Func, binds, typeMap)
		if err != nil {
			return nil, err
		}
		bind, isBinding := v.(*Binding)
		if !isBinding {
			return nil, fmt.Errorf("call a non-callable expression: %s", expr.Func)
		}
		f, err := bind.GetImpl(typeMap, expr.ArgList...)
		if err != nil {
			return nil, err
		}
		valueList := make([]llvmvalue.Value, 0)
		for _, arg := range expr.ArgList {
			e, err := m.genExpr(b, arg.Expr, binds, typeMap)
			if err != nil {
				return nil, err
			}
			valueList = append(valueList, e)
		}
		return b.NewCall(f, valueList...), nil
	case *ast.BinaryExpr:
		x, err := m.genExpr(b, expr.LExpr, binds, typeMap)
		if err != nil {
			return nil, err
		}
		y, err := m.genExpr(b, expr.RExpr, binds, typeMap)
		if err != nil {
			return nil, err
		}
		lt, err := typeMap.GetTypeOfExpr(expr.LExpr)
		if err != nil {
			return nil, err
		}
		rt, err := typeMap.GetTypeOfExpr(expr.RExpr)
		if err != nil {
			return nil, err
		}
		key := genKey(expr.Op, lt, rt)
		if m.generator.isOperator(key) {
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
			if lt.String() == "f64" && rt.String() == "f64" {
				switch expr.Op {
				case "+":
					return b.NewFAdd(x, y), nil
				case "-":
					return b.NewFSub(x, y), nil
				case "*":
					return b.NewFMul(x, y), nil
				case "/":
					return b.NewFDiv(x, y), nil
				}
			}
		}
		return nil, fmt.Errorf("unsupported operator: %s", expr.Op)
	case *ast.Ident:
		v, ok := binds[expr.Literal]
		if ok {
			return v, nil
		}
		bind, err := m.GetBinding(expr.Literal)
		if err != nil {
			chain := strings.Split(expr.Literal, "::")
			if len(chain) >= 2 {
				localModuleName := chain[len(chain)-2]
				funcName := chain[len(chain)-1]
				moduleName := m.imports[localModuleName]
				return m.generator.allModule[moduleName].GetExportBinding(funcName)
			}
			return m.generator.getBuiltin(expr.Literal)
		}
		if bind.IsFunc {
			return bind, err
		}
		globalVarValue, err := m.genExpr(b, bind.Expr, binds, typeMap)
		if err != nil {
			return nil, err
		}
		// declare global variable
		globalVar := m.generator.mod.NewGlobal(bind.Name, globalVarValue.Type())
		globalVar.Init = &constant.ZeroInitializer{}
		globalVar.Align = 1
		// store the temporary value into global variable in init function
		m.initFuncBlock.NewStore(globalVarValue, globalVar)
		if v, ok := globalVarValue.(*value.Wrapper); ok {
			return value.NewWrap(
				b.NewLoad(globalVar),
				v.ElemT,
			), nil
		}
		return b.NewLoad(globalVar), nil
	case *ast.Int:
		return constant.NewIntFromString(llvmtypes.I64, expr.Literal)
	case *ast.Float:
		return constant.NewFloatFromString(llvmtypes.Double, expr.Literal)
	case *ast.String:
		str := m.generator.mod.NewGlobal("", llvmtypes.NewArray(uint64(len(expr.Literal)), llvmtypes.I8))
		str.Align = 1
		str.Init = constant.NewCharArrayFromString(expr.Literal)
		strGEP := b.NewGetElementPtr(str,
			constant.NewInt(llvmtypes.I64, 0),
			constant.NewInt(llvmtypes.I64, 0),
		)
		x := b.NewAlloca(llvmtypes.NewPointer(llvmtypes.I8))
		b.NewStore(strGEP, x)
		return b.NewLoad(x), nil
	case *ast.List:
		newList, err := m.generator.getBuiltin("new_list")
		if err != nil {
			return nil, err
		}
		newListImpl, err := newList.GetImpl(typeMap)
		if err != nil {
			return nil, err
		}
		elzMalloc, err := m.generator.getBuiltin("elz_malloc")
		if err != nil {
			return nil, err
		}
		elzMallocImpl, err := elzMalloc.GetImpl(typeMap)
		if err != nil {
			return nil, err
		}
		// make a temporary global array for storing arguments of new_list
		tmpListPtr := m.generator.mod.NewGlobalDef("",
			constant.NewZeroInitializer(
				llvmtypes.NewArray(
					uint64(len(expr.ExprList)),
					llvmtypes.NewPointer(llvmtypes.I8),
				),
			),
		)
		tmpListPtr.Align = ir.Align(1)
		var elemT llvmtypes.Type
		// storing ast list into tmp list
		initBlock := m.initFuncBlock
		for i, e := range expr.ExprList {
			llvmExpr, err := m.genExpr(initBlock, e, binds, typeMap)
			if err != nil {
				return nil, err
			}
			if i == 0 {
				elemT = llvmExpr.Type()
			}
			indexI := initBlock.NewGetElementPtr(
				tmpListPtr,
				constant.NewInt(llvmtypes.I64, 0),
				constant.NewInt(llvmtypes.I64, int64(i)),
			)
			size := constant.NewInt(llvmtypes.I64, irutil.SizeOf(elemT))
			exprMalloca := initBlock.NewCall(elzMallocImpl, size)
			storeTo := initBlock.NewBitCast(exprMalloca, llvmtypes.NewPointer(llvmExpr.Type()))

			initBlock.NewStore(llvmExpr, storeTo)
			initBlock.NewStore(exprMalloca, indexI)
		}
		// get this temporary array's address
		elems := initBlock.NewGetElementPtr(tmpListPtr,
			constant.NewInt(llvmtypes.I64, 0),
			constant.NewInt(llvmtypes.I64, 0),
		)
		return value.NewWrap(
			initBlock.NewCall(newListImpl,
				// size
				constant.NewInt(llvmtypes.I64, int64(len(expr.ExprList))),
				// elements
				elems,
			),
			elemT,
		), nil
	case *ast.ExtractElement:
		listIndex, err := m.generator.getBuiltin("list_index")
		if err != nil {
			return nil, err
		}
		listIndexImpl, err := listIndex.GetImpl(typeMap)
		if err != nil {
			return nil, err
		}
		// rely on infer type checking the Func and Key type already,
		// we don't check it again
		x, err := m.genExpr(b, expr.X, binds, typeMap)
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(x.Type().String(), "%list") {
			key, err := m.genExpr(b, expr.Key, binds, typeMap)
			if err != nil {
				return nil, err
			}
			elemPtr := b.NewCall(listIndexImpl, x, key)
			// x value of extract element must be a wrapper, so we using type assertion here
			convertedPtr := b.NewBitCast(elemPtr, llvmtypes.NewPointer(x.(*value.Wrapper).ElemT))
			return b.NewLoad(convertedPtr), nil
		}
		return nil, fmt.Errorf("unknown x value: %s", x)
	default:
		return nil, fmt.Errorf("[Unsupport Yet] failed at generate expression: %#v", expr)
	}
}
