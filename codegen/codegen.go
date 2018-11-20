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
	// From Parser
	bindings []*ast.Binding
}

func NewGenerator() *CodeGenerator {
	return &CodeGenerator{
		mod:      ir.NewModule(),
		bindings: make([]*ast.Binding, 0),
	}
}

func (c *CodeGenerator) GenAllBinding() {
	for _, binding := range c.bindings {
		c.GenBinding(binding)
	}
}

func (c *CodeGenerator) GenBinding(binding *ast.Binding) {
	// e.g. i = 1
	// But also should consider something like Main
	// e.g. main = println "Hello, World"
	if len(binding.ParamList) == 0 {
		if binding.Name == "main" {
			// generate binary
			mainFn := c.mod.NewFunction("main", types.I32)
			entryBlock := mainFn.NewBlock("")
			entryBlock.NewRet(constant.NewInt(0, types.I32))
			c.mod.AppendFunction(mainFn)
			return
		}
		if binding.Expr.IsConst() {
			llvmValue := c.NewConstExpr(binding.Expr)
			c.mod.NewGlobalDef(binding.Name, llvmValue)
		} else {
			newFn := c.GetFuncWithoutParam(binding)
			c.mod.AppendFunction(newFn)
		}
		return
	}
	// e.g. add x y = x + y
	// TODO: generate function template
}

func (c *CodeGenerator) GetFuncWithoutParam(binding *ast.Binding) *ir.Function {
	exprType := c.GetExprType(binding.Expr)
	f := c.mod.NewFunction(binding.Name, exprType)
	entryBlock := f.NewBlock("")
	v := c.NewExpr(entryBlock, binding.Expr)
	entryBlock.NewRet(v)
	return f
}

func (c *CodeGenerator) NewExpr(block *ir.BasicBlock, expr ast.Expr) value.Value {
	if expr.IsConst() {
		return c.NewConstExpr(expr)
	}
	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		left := c.NewExpr(block, expr.LExpr)
		right := c.NewExpr(block, expr.RExpr)
		return c.SearchOperation(block, expr.Operator, left, right)
	}
	panic("unsupported expr")
}

type Operator struct {
	RetType   types.Type
	Operation func(block *ir.BasicBlock, l, r value.Value) value.Value
}

var (
	binaryOpFormat = "%s(%s,%s)"
	i32            = types.I32
	opMap          = map[string]*Operator{
		fmt.Sprintf(binaryOpFormat, "+", i32, i32): {
			RetType: i32,
			Operation: func(block *ir.BasicBlock, l, r value.Value) value.Value {
				return block.NewAdd(l, r)
			},
		},
		fmt.Sprintf(binaryOpFormat, "-", i32, i32): {
			RetType: i32,
			Operation: func(block *ir.BasicBlock, l, r value.Value) value.Value {
				return block.NewSub(l, r)
			},
		},
		fmt.Sprintf(binaryOpFormat, "*", i32, i32): {
			RetType: i32,
			Operation: func(block *ir.BasicBlock, l, r value.Value) value.Value {
				return block.NewMul(l, r)
			},
		},
		fmt.Sprintf(binaryOpFormat, "/", i32, i32): {
			RetType: i32,
			Operation: func(block *ir.BasicBlock, l, r value.Value) value.Value {
				return block.NewSDiv(l, r)
			},
		},
	}
)

func (c *CodeGenerator) SearchOperation(block *ir.BasicBlock, operator string, left, right value.Value) value.Value {
	generator := opMap[fmt.Sprintf(binaryOpFormat, operator, left.Type(), right.Type())]
	return generator.Operation(block, left, right)
}

func (c *CodeGenerator) GetExprType(expr ast.Expr) types.Type {
	switch expr := expr.(type) {
	case *ast.Int:
		return types.I32
	case *ast.Float:
		return types.Double
	case *ast.Bool:
		return types.I1
	case *ast.BinaryExpr:
		operator := opMap[fmt.Sprintf(binaryOpFormat, expr.Operator, c.GetExprType(expr.LExpr), c.GetExprType(expr.RExpr))]
		return operator.RetType
	default:
		panic("unsupported type refer")
	}
}

func (c *CodeGenerator) NewConstExpr(expr ast.Expr) constant.Constant {
	switch expr := expr.(type) {
	case *ast.Int:
		// default is i32
		return constant.NewIntFromString(expr.Literal, types.I32)
	case *ast.Float:
		// default is f64(double)
		return constant.NewFloatFromString(expr.Literal, types.Double)
	case *ast.Bool:
		if expr.IsTrue {
			return constant.NewInt(1, types.I1)
		} else {
			return constant.NewInt(0, types.I1)
		}
	case *ast.String:
		panic("unsupported string right now")
	case *ast.Ident:
		panic("unsupported identifier right now")
	default:
		panic("unsupported expression")
	}
}
