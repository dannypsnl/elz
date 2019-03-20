package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type module struct {
	*Tree
	generator *Generator
	imports   map[string]string
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
			panic(fmt.Sprintf(`import %s
and
import %s
has the same name in the module`, mod1, importPath))
		}
		imports[accessKey] = importPath
	}
	return &module{
		Tree:      tree,
		generator: g,
		imports:   imports,
	}
}

func (m *module) genExpr(b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap *typeMap) (value.Value, error) {
	switch expr := expr.(type) {
	case *ast.FuncCall:
		bind, err := m.getBindingByAccessChain(expr.AccessChain)
		if err != nil {
			return nil, err
		}
		f, err := bind.GetImpl(m, typeMap, expr.ArgList...)
		if err != nil {
			return nil, err
		}
		valueList := make([]value.Value, 0)
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
	default:
		return nil, fmt.Errorf("failed at generate expression: %#v", expr)
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
			return m.generator.entryModule.GetBinding(accessChain)
		}
		return bind, nil
	}
	return nil, fmt.Errorf("not supported access chain: %s", accessChain)
}
