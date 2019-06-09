package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/internal/irutil"
	"github.com/elz-lang/elz/src/elz/types"
	"github.com/elz-lang/elz/src/elz/value"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irenum "github.com/llir/llvm/ir/enum"
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
func newModule(g *Generator, name string, tree *Tree) *module {
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
	for _, b := range m.bindings {
		if b.IsFunc {
			m.typeMap.Add(b.Name, types.NewBindingType(m, b.Binding))
		}
	}
	params := make([]string, 0)
	for _, typeDef := range m.typeDefines {
		for _, field := range typeDef.Fields {
			params = append(params, field.Name)
		}
		newStruct := value.NewStruct(
			m.generator.mod,
			name,
			m.generator.builtin["elz_malloc"].compilerProvidedImpl,
			typeDef,
		)
		err := m.InsertBinding(ast.NewBinding(
			true,
			typeDef.Export,
			typeDef.Name,
			params,
			newStruct,
		))
		if err != nil {
			logrus.Fatalf("failed at insert constructor of type: %s", typeDef.Name)
		}
		bind, _ := m.GetBinding(typeDef.Name)
		// FIXME: this would be a problem when we build generic constructor
		bind.compilerProvidedImpl = newStruct.Constructor()
		m.typeMap.Add(typeDef.Name, newStruct.ElzType())
	}
	for _, bind := range m.bindings {
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
	case *value.Struct:
		return expr.ElzType(), nil
	default:
		return nil, fmt.Errorf("unsupported type inference for expression: %#v yet", expr)
	}
}

type context struct {
	*ir.Block
	binds   map[string]*ir.Param
	typeMap *types.TypeMap
}

func newContext(block *ir.Block, typeMap *types.TypeMap) *context {
	return &context{
		Block:   block,
		binds:   make(map[string]*ir.Param),
		typeMap: typeMap,
	}
}

func (m *module) genExpr(c *context, expr ast.Expr) (value.Value, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		v, err := m.genExpr(c, expr.Func)
		if err != nil {
			return nil, err
		}
		bind, isBinding := v.(*Binding)
		if !isBinding {
			return nil, fmt.Errorf("call a non-callable expression: %s", expr.Func)
		}
		f, err := bind.GetImpl(c.typeMap, expr.ArgList...)
		if err != nil {
			return nil, err
		}
		valueList := make([]llvmvalue.Value, 0)
		for _, arg := range expr.ArgList {
			e, err := m.genExpr(c, arg.Expr)
			if err != nil {
				return nil, err
			}
			valueList = append(valueList, e)
		}
		return c.NewCall(f, valueList...), nil
	case *ast.BinaryExpr:
		x, err := m.genExpr(c, expr.LExpr)
		if err != nil {
			return nil, err
		}
		y, err := m.genExpr(c, expr.RExpr)
		if err != nil {
			return nil, err
		}
		lt, err := c.typeMap.GetTypeOfExpr(expr.LExpr)
		if err != nil {
			return nil, err
		}
		rt, err := c.typeMap.GetTypeOfExpr(expr.RExpr)
		if err != nil {
			return nil, err
		}
		key := genKey(expr.Op, lt, rt)
		if m.generator.isOperator(key) {
			if lt.String() == "int" && rt.String() == "int" {
				switch expr.Op {
				case "+":
					return c.NewAdd(x, y), nil
				case "-":
					return c.NewSub(x, y), nil
				case "*":
					return c.NewMul(x, y), nil
				case "/":
					return c.NewSDiv(x, y), nil
				}
			}
			if lt.String() == "f64" && rt.String() == "f64" {
				switch expr.Op {
				case "+":
					return c.NewFAdd(x, y), nil
				case "-":
					return c.NewFSub(x, y), nil
				case "*":
					return c.NewFMul(x, y), nil
				case "/":
					return c.NewFDiv(x, y), nil
				}
			}
		}
		return nil, fmt.Errorf("unsupported operator: %s", expr.Op)
	case *ast.AccessField:
		le, err := m.genExpr(c, expr.From)
		if err != nil {
			return nil, err
		}
		lt, err := c.typeMap.GetTypeOfExpr(expr.From)
		if err != nil {
			return nil, err
		}
		st, ok := lt.(*types.Struct)
		if !ok {
			return nil, fmt.Errorf("access expression is not a structure type: %s", lt)
		}
		indices := -1
		for i, field := range st.Fields {
			if field.Name == expr.ByName {
				indices = i
			}
		}
		if indices == -1 {
			return nil, fmt.Errorf("no field called: %s in type: %s", expr.ByName, lt)
		}
		firstField := c.NewExtractValue(le,
			uint64(indices),
		)
		return firstField, nil
	case *ast.CaseOf:
		caseExpr, err := m.genExpr(c, expr.Case)
		if err != nil {
			return nil, err
		}

		caseStartBlock := c.Parent.NewBlock("")
		c.NewBr(caseStartBlock)

		caseEndBlock := caseStartBlock.Parent.NewBlock("")

		assignV := caseStartBlock.NewAlloca(llvmtypes.I64)
		caseElseBlock := caseStartBlock.Parent.NewBlock("")
		c.Block = caseElseBlock
		elseExpr, err := m.genExpr(c, expr.Else)
		if err != nil {
			return nil, err
		}
		caseElseBlock.NewStore(elseExpr, assignV)
		caseElseBlock.NewBr(caseEndBlock)
		for _, of := range expr.CaseOf {
			c.Block = caseStartBlock
			pattern, err := m.genExpr(c, of.Pattern)
			if err != nil {
				return nil, err
			}
			cond := caseStartBlock.NewICmp(irenum.IPredEQ, pattern, caseExpr)
			caseOfBlock := caseStartBlock.Parent.NewBlock("")
			caseStartBlock.NewCondBr(cond, caseOfBlock, caseElseBlock)
			c.Block = caseOfBlock
			do, err := m.genExpr(c, of.Expr)
			if err != nil {
				return nil, err
			}
			caseOfBlock.NewStore(do, assignV)
			caseOfBlock.NewBr(caseEndBlock)
		}
		c.Block = caseEndBlock
		return c.NewLoad(assignV), nil
	case *ast.Ident:
		v, ok := c.binds[expr.Literal]
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
		b := c.Block
		c.Block = m.initFuncBlock
		globalVarValue, err := m.genExpr(c, bind.Expr)
		c.Block = b
		if err != nil {
			return nil, err
		}
		globalT, err := c.typeMap.GetTypeOfExpr(bind.Expr)
		if err != nil {
			return nil, err
		}
		c.typeMap.Add(expr.Literal, globalT)
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
		strGEP := c.NewGetElementPtr(str,
			constant.NewInt(llvmtypes.I64, 0),
			constant.NewInt(llvmtypes.I64, 0),
		)
		x := c.NewAlloca(llvmtypes.NewPointer(llvmtypes.I8))
		c.NewStore(strGEP, x)
		return c.NewLoad(x), nil
	case *ast.List:
		newList, err := m.generator.getBuiltin("new_list")
		if err != nil {
			return nil, err
		}
		newListImpl, err := newList.GetImpl(c.typeMap)
		if err != nil {
			return nil, err
		}
		elzMalloc, err := m.generator.getBuiltin("elz_malloc")
		if err != nil {
			return nil, err
		}
		elzMallocImpl, err := elzMalloc.GetImpl(c.typeMap)
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
		backup := c.Block
		c.Block = m.initFuncBlock
		// storing ast list into tmp list
		for i, e := range expr.ExprList {
			llvmExpr, err := m.genExpr(c, e)
			if err != nil {
				return nil, err
			}
			if i == 0 {
				elemT = llvmExpr.Type()
			}
			indexI := c.NewGetElementPtr(
				tmpListPtr,
				constant.NewInt(llvmtypes.I64, 0),
				constant.NewInt(llvmtypes.I64, int64(i)),
			)
			size := constant.NewInt(llvmtypes.I64, irutil.SizeOf(elemT))
			exprMalloca := c.NewCall(elzMallocImpl, size)
			storeTo := c.NewBitCast(exprMalloca, llvmtypes.NewPointer(llvmExpr.Type()))

			c.NewStore(llvmExpr, storeTo)
			c.NewStore(exprMalloca, indexI)
		}
		// get this temporary array's address
		elems := c.NewGetElementPtr(tmpListPtr,
			constant.NewInt(llvmtypes.I64, 0),
			constant.NewInt(llvmtypes.I64, 0),
		)
		call := c.NewCall(newListImpl,
			// size
			constant.NewInt(llvmtypes.I64, int64(len(expr.ExprList))),
			// elements
			elems,
		)
		c.Block = backup
		return value.NewWrap(
			call,
			elemT,
		), nil
	case *ast.ExtractElement:
		listIndex, err := m.generator.getBuiltin("list_index")
		if err != nil {
			return nil, err
		}
		listIndexImpl, err := listIndex.GetImpl(c.typeMap)
		if err != nil {
			return nil, err
		}
		// rely on infer type checking the Func and Key type already,
		// we don't check it again
		x, err := m.genExpr(c, expr.X)
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(x.Type().String(), "%list") {
			key, err := m.genExpr(c, expr.Key)
			if err != nil {
				return nil, err
			}
			elemPtr := c.NewCall(listIndexImpl, x, key)
			// x value of extract element must be a wrapper, so we using type assertion here
			convertedPtr := c.NewBitCast(elemPtr, llvmtypes.NewPointer(x.(*value.Wrapper).ElemT))
			return c.NewLoad(convertedPtr), nil
		}
		return nil, fmt.Errorf("unknown x value: %s", x)
	default:
		return nil, fmt.Errorf("[Unsupport Yet] failed at generate expression: %#v", expr)
	}
}
