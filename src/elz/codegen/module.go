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
	initFunc := g.mod.NewFunc("init", llvmtypes.Void)
	m := &module{
		Tree:          tree,
		generator:     g,
		initFunc:      initFunc,
		initFuncBlock: initFunc.NewBlock(""),
		imports:       imports,
	}
	for _, bind := range tree.bindings {
		bind.SetModule(m)
	}
	return m
}

// inference the return type by the expression we going to execute and input types
func (m *module) inferTypeOf(expr ast.Expr, typeMap *typeMap) (types.Type, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, err := m.getBindingByAccessChain(expr.X.(*ast.Ident).Literal)
		if err != nil {
			return nil, err
		}
		typeList := typeMap.convertArgsToTypeList(expr.ArgList...)
		typeMap := newTypeMap()
		for i, paramType := range typeList {
			paramName := bind.ParamList[i]
			typeMap.add(paramName, paramType)
		}
		t, err := bind.GetReturnType(typeMap, typeList...)
		if err != nil {
			return nil, err
		}
		return t, nil
	case *ast.BinaryExpr:
		lt, err := m.inferTypeOf(expr.LExpr, typeMap)
		if err != nil {
			return nil, err
		}
		rt, err := m.inferTypeOf(expr.RExpr, typeMap)
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
		t := typeMap.getTypeOfExpr(expr)
		if t == nil {
			return nil, fmt.Errorf("can't get type of identifier: %s", expr.Literal)
		}
		return t, nil
	case *ast.List:
		elemT, err := m.inferTypeOf(expr.ExprList[0], typeMap)
		if err != nil {
			return nil, err
		}
		return types.NewList(elemT), nil
	case *ast.ExtractElement:
		list, ok := expr.X.(*ast.List)
		if !ok {
			return nil, fmt.Errorf("extract element do not support ast: %#v", expr.X)
		}
		listT, err := m.inferTypeOf(list, typeMap)
		if err != nil {
			return nil, err
		}
		keyT, err := m.inferTypeOf(expr.Key, typeMap)
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

func (m *module) genExpr(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap *typeMap) (value.Value, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, err := m.getBindingByAccessChain(expr.X.(*ast.Ident).Literal)
		if err != nil {
			return nil, err
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
		lt := typeMap.getTypeOfExpr(expr.LExpr)
		rt := typeMap.getTypeOfExpr(expr.RExpr)
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
			return nil, err
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
		// rely on infer type checking the X and Key type already,
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

func (m *module) getBindingByAccessChain(accessChain string) (*Binding, error) {
	chain := strings.Split(accessChain, "::")
	if len(chain) >= 2 {
		localModuleName := chain[len(chain)-2]
		funcName := chain[len(chain)-1]
		moduleName := m.imports[localModuleName]
		return m.generator.allModule[moduleName].GetExportBinding(funcName)
	}
	if len(chain) == 1 {
		bind, err := m.GetBinding(accessChain)
		if err != nil {
			return m.generator.getBuiltin(accessChain)
		}
		return bind, nil
	}
	return nil, fmt.Errorf("not supported access chain: %s", accessChain)
}
