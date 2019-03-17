package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"
	"github.com/elz-lang/elz/src/irutil"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Generator struct {
	mod *ir.Module

	allTree   map[string]*Tree
	entryTree *Tree

	operatorTypeStore map[string]types.Type
}

func New(entryTree *Tree, astAllTree map[string]*Tree) *Generator {
	typMap := make(map[string]types.Type)
	typMap["+ :: int -> int"] = &types.Int{}

	astTree := entryTree
	err := astTree.InsertBinding(&ast.Binding{
		Name:      "printf",
		ParamList: []string{"format"},
	})
	if err != nil {
		panic("some function conflict with built-in function printf")
	}
	mod := ir.NewModule()
	printfImpl := mod.NewFunc("printf", llvmtypes.I64,
		ir.NewParam("format", llvmtypes.NewPointer(llvmtypes.I8)),
	)
	printfImpl.Sig.Variadic = true
	printfBind, err := astTree.GetBinding("printf")
	if err != nil {
		panic(fmt.Errorf("can't get printf binding: %s", err))
	}
	printfBind.compilerProvidedImpl = printfImpl

	return &Generator{
		mod:               mod,
		entryTree:         astTree,
		allTree:           astAllTree,
		operatorTypeStore: typMap,
	}
}

func (g *Generator) String() string {
	irutil.FixDups(g.mod)
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

func (g *Generator) Call(bind *Binding, exprList ...*ast.Arg) error {
	_, err := bind.GetImpl(g, newTypeMap(), exprList...)
	if err != nil {
		return err
	}
	return nil
}

// inference the return type by the expression we going to execute and input types
func (g *Generator) inferTypeOf(expr ast.Expr, typeMap *typeMap) (types.Type, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, err := g.getBindingByAccessChain(expr.AccessChain)
		if err != nil {
			return nil, err
		}
		typeList := typeMap.convertArgsToTypeList(expr.ArgList...)
		typeMap := newTypeMap()
		for i, paramType := range typeList {
			paramName := bind.ParamList[i]
			typeMap.add(paramName, paramType)
		}
		t, err := bind.GetReturnType(g, typeMap, typeList...)
		if err != nil {
			return nil, err
		}
		return t, nil
	case *ast.BinaryExpr:
		lt, err := g.inferTypeOf(expr.LExpr, typeMap)
		if err != nil {
			return nil, err
		}
		rt, err := g.inferTypeOf(expr.RExpr, typeMap)
		if err != nil {
			return nil, err
		}
		op := expr.Op
		t, err := g.typeOfOperator(op, lt, rt)
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
	default:
		return nil, fmt.Errorf("unsupported type inference for expression: %#v yet", expr)
	}
}

func (g *Generator) getBindingByAccessChain(accessChain string) (*Binding, error) {
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

func (g *Generator) isOperator(key string) bool {
	_, isBuiltIn := g.operatorTypeStore[key]
	return isBuiltIn
}

func (g *Generator) typeOfOperator(op string, typeList ...types.Type) (types.Type, error) {
	key := genKey(op, typeList...)
	t, existed := g.operatorTypeStore[key]
	if !existed {
		return nil, fmt.Errorf("can't infer return type by %s", key)
	}
	return t, nil
}

func (g *Generator) genExpr(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap *typeMap) (value.Value, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, err := g.getBindingByAccessChain(expr.AccessChain)
		if err != nil {
			return nil, err
		}
		f, err := bind.GetImpl(g, typeMap, expr.ArgList...)
		if err != nil {
			return nil, err
		}
		valueList := make([]value.Value, 0)
		for _, arg := range expr.ArgList {
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
		if g.isOperator(key) {
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
	b.WriteString(typeFormat(typeList...))
	return b.String()
}

func typeFormat(typeList ...types.Type) string {
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
