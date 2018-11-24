package codegen

import (
	"fmt"

	"github.com/elz-lang/elz/ast"

	"llvm.org/llvm/bindings/go/llvm"
)

type CodeGenerator struct {
	// For LLVM
	mod llvm.Module
	// Internal
	bindings map[string]*ast.Binding
}

func NewGenerator() *CodeGenerator {
	return &CodeGenerator{
		mod:      llvm.NewModule(""),
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

func (c *CodeGenerator) CallBindingWith(builder llvm.Builder, binding *ast.Binding, vs []llvm.Value) llvm.Value {
	if len(vs) == len(binding.ParamList) {
		typeList := make([]llvm.Type, len(vs))
		scope := map[string]llvm.Type{}
		valueScope := map[string]llvm.Value{}
		for i, v := range vs {
			paramName := binding.ParamList[i]
			scope[paramName] = v.Type()
			valueScope[paramName] = v
			typeList[i] = v.Type()
		}
		retT := c.GetExprType(scope, binding.Expr)
		ft := llvm.FunctionType(retT, typeList, false)
		newFn := llvm.AddFunction(c.mod, binding.Name, ft)
		newFunctionBuilder := llvm.NewBuilder()
		block := llvm.AddBasicBlock(newFn, "")
		newFunctionBuilder.SetInsertPointAtEnd(block)
		fnExpr := c.NewExpr(valueScope, newFunctionBuilder, binding.Expr)
		newFunctionBuilder.CreateRet(fnExpr)
		return builder.CreateCall(newFn, vs, "")
	}
	panic("not implement lambda yet")
}

func (c *CodeGenerator) NewExpr(scope map[string]llvm.Value, builder llvm.Builder, expr ast.Expr) llvm.Value {
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
		args := make([]llvm.Value, 0)
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
	RetType   llvm.Type
	Operation func(builder llvm.Builder, l, r llvm.Value) llvm.Value
}

var (
	binaryOpFormat = "%s(%s,%s)"
	i32            = llvm.Int32Type()
	opMap          = map[string]*Operator{
		fmt.Sprintf(binaryOpFormat, "+", i32, i32): {
			RetType: i32,
			Operation: func(builder llvm.Builder, l, r llvm.Value) llvm.Value {
				return builder.CreateAdd(l, r, "")
			},
		},
		fmt.Sprintf(binaryOpFormat, "-", i32, i32): {
			RetType: i32,
			Operation: func(builder llvm.Builder, l, r llvm.Value) llvm.Value {
				return builder.CreateAdd(l, r, "")
			},
		},
		fmt.Sprintf(binaryOpFormat, "*", i32, i32): {
			RetType: i32,
			Operation: func(builder llvm.Builder, l, r llvm.Value) llvm.Value {
				return builder.CreateAdd(l, r, "")
			},
		},
		fmt.Sprintf(binaryOpFormat, "/", i32, i32): {
			RetType: i32,
			Operation: func(builder llvm.Builder, l, r llvm.Value) llvm.Value {
				return builder.CreateAdd(l, r, "")
			},
		},
	}
)

func (c *CodeGenerator) SearchOperation(builder llvm.Builder, operator string, left, right llvm.Value) llvm.Value {
	generator := opMap[fmt.Sprintf(binaryOpFormat, operator, left.Type(), right.Type())]
	return generator.Operation(builder, left, right)
}

func (c *CodeGenerator) GetExprType(scope map[string]llvm.Type, expr ast.Expr) llvm.Type {
	switch expr := expr.(type) {
	case *ast.Int:
		return llvm.Int32Type()
	case *ast.Float:
		return llvm.DoubleType()
	case *ast.Bool:
		return llvm.Int1Type()
	case *ast.BinaryExpr:
		operator := opMap[fmt.Sprintf(binaryOpFormat, expr.Operator, c.GetExprType(scope, expr.LExpr), c.GetExprType(scope, expr.RExpr))]
		return operator.RetType
	case *ast.FuncCall:
		return llvm.Int32Type()
	case *ast.Ident:
		return scope[expr.Value]
	default:
		panic("unsupported type refer")
	}
}

func (c *CodeGenerator) NewConstExpr(expr ast.Expr) llvm.Value {
	switch expr := expr.(type) {
	case *ast.Int:
		// default is i32
		return llvm.ConstIntFromString(llvm.Int32Type(), expr.Literal, 10)
	case *ast.Float:
		// default is f64(double)
		return llvm.ConstFloatFromString(llvm.DoubleType(), expr.Literal)
	case *ast.Bool:
		if expr.IsTrue {
			return llvm.ConstInt(llvm.Int1Type(), 1, false)
		} else {
			return llvm.ConstInt(llvm.Int1Type(), 0, false)
		}
	case *ast.String:
		panic("unsupported string right now")
	case *ast.Ident:
		panic("unsupported identifier right now")
	default:
		panic("unsupported expression")
	}
}
