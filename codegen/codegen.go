package codegen

import (
	"fmt"

	"github.com/elz-lang/elz/ast"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type CodeGenerator struct {
	// For LLVM
	mod *ir.Module
	// Internal
	bindings map[string]*ast.Binding
}

func NewGenerator() *CodeGenerator {
	return &CodeGenerator{
		mod:      ir.NewModule(),
		bindings: make(map[string]*ast.Binding),
	}
}

func (c *CodeGenerator) GenBinding(binding *ast.Binding) {
	c.bindings[binding.Name] = binding
	// e.g. i = 1
	// But also should consider something like Main
	// e.g. main = println "Hello, World"
	// e.g. add x y = x + y
	// TODO: generate function template
}

func (c *CodeGenerator) CallBindingWith(builder *ir.BasicBlock, binding *ast.Binding, vs []value.Value) value.Value {
	if len(vs) == len(binding.ParamList) {
		params := make([]*ir.Param, len(vs))
		scope := map[string]types.Type{}
		valueScope := map[string]value.Value{}
		for i, v := range vs {
			paramName := binding.ParamList[i]
			scope[paramName] = v.Type()
			valueScope[paramName] = v
			params[i] = ir.NewParam(paramName, v.Type())
		}
		retT := c.GetExprType(scope, binding.Expr)
		newFn := c.mod.NewFunc(binding.Name, retT, params...)
		newFnBuilder := newFn.NewBlock("")
		fnExpr := c.NewExpr(valueScope, newFnBuilder, binding.Expr)
		newFnBuilder.NewRet(fnExpr)
		return builder.NewCall(newFn, vs...)
	}
	panic("not implement lambda yet")
}

func (c *CodeGenerator) BindingReturnType(binding *ast.Binding, typeList []types.Type) types.Type {
	paramTypes := map[string]types.Type{}
	for i, t := range typeList {
		paramTypes[binding.ParamList[i]] = t
	}
	return c.GetExprType(paramTypes, binding.Expr)
}

func (c *CodeGenerator) NewExpr(scope map[string]value.Value, builder *ir.BasicBlock, expr ast.Expr) value.Value {
	if expr.IsConst() {
		return c.NewConstExpr(expr)
	}
	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		left := c.NewExpr(scope, builder, expr.LExpr)
		right := c.NewExpr(scope, builder, expr.RExpr)
		return c.SearchOperation(builder, expr.Operator, left, right)
	case *ast.FuncCall:
		binding := c.bindings[expr.Identifier]
		args := make([]value.Value, 0)
		for _, expr := range expr.ExprList {
			args = append(args, c.NewExpr(scope, builder, expr))
		}
		return c.CallBindingWith(builder, binding, args)
	case *ast.Ident:
		return scope[expr.Value]
	}
	panic("unsupported expr")
}

type Operator struct {
	RetType   types.Type
	Operation func(builder *ir.BasicBlock, l, r value.Value) value.Value
}

var (
	binaryOpFormat = "%s(%s,%s)"
	i32            = types.I32
	opMap          = map[string]*Operator{
		fmt.Sprintf(binaryOpFormat, "+", i32, i32): {
			RetType: i32,
			Operation: func(builder *ir.BasicBlock, l, r value.Value) value.Value {
				return builder.NewAdd(l, r)
			},
		},
		fmt.Sprintf(binaryOpFormat, "-", i32, i32): {
			RetType: i32,
			Operation: func(builder *ir.BasicBlock, l, r value.Value) value.Value {
				return builder.NewSub(l, r)
			},
		},
		fmt.Sprintf(binaryOpFormat, "*", i32, i32): {
			RetType: i32,
			Operation: func(builder *ir.BasicBlock, l, r value.Value) value.Value {
				return builder.NewMul(l, r)
			},
		},
		fmt.Sprintf(binaryOpFormat, "/", i32, i32): {
			RetType: i32,
			Operation: func(builder *ir.BasicBlock, l, r value.Value) value.Value {
				return builder.NewSDiv(l, r)
			},
		},
	}
)

func (c *CodeGenerator) SearchOperation(builder *ir.BasicBlock, operator string, left, right value.Value) value.Value {
	generator := opMap[fmt.Sprintf(binaryOpFormat, operator, left.Type(), right.Type())]
	return generator.Operation(builder, left, right)
}

func (c *CodeGenerator) GetExprType(scope map[string]types.Type, expr ast.Expr) types.Type {
	switch expr := expr.(type) {
	case *ast.Int:
		return types.I32
	case *ast.Float:
		return types.Double
	case *ast.Bool:
		return types.I1
	case *ast.BinaryExpr:
		operator := opMap[fmt.Sprintf(binaryOpFormat, expr.Operator, c.GetExprType(scope, expr.LExpr), c.GetExprType(scope, expr.RExpr))]
		return operator.RetType
	case *ast.FuncCall:
		binding, found := c.bindings[expr.Identifier]
		if !found {
			panic("no this function")
		}
		paramTypes := make([]types.Type, 0)
		for _, e := range expr.ExprList {
			paramTypes = append(paramTypes, c.GetExprType(scope, e))
		}
		return c.BindingReturnType(binding, paramTypes)
	case *ast.Ident:
		return scope[expr.Value]
	default:
		panic("unsupported type refer")
	}
}

func (c *CodeGenerator) NewConstExpr(expr ast.Expr) value.Value {
	switch expr := expr.(type) {
	case *ast.Int:
		// default is i32
		v, err := constant.NewIntFromString(types.I32, expr.Literal)
		if err != nil {
			panic(fmt.Errorf("unable to parse integer literal %q; %v", expr.Literal, err))
		}
		return v
	case *ast.Float:
		// default is f64(double)
		v, err := constant.NewFloatFromString(types.Double, expr.Literal)
		if err != nil {
			panic(fmt.Errorf("unable to parse floating-point literal %q; %v", expr.Literal, err))
		}
		return v
	case *ast.Bool:
		return constant.NewBool(expr.IsTrue)
	case *ast.String:
		panic("unsupported string right now")
	case *ast.Ident:
		panic("unsupported identifier right now")
	default:
		panic("unsupported expression")
	}
}
