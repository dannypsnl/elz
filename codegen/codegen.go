package codegen

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/codegen/types"

	"github.com/llir/llvm/ir"
	irTypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type CodeGenerator struct {
	// For LLVM
	mod *ir.Module
	// Internal
	functions map[string]*Func
}

func prepareingModule() *ir.Module {
	mod := ir.NewModule()
	elzFuncType := irTypes.NewStruct()
	elzFuncType.Opaque = true
	elzFT := irTypes.NewPointer(mod.NewTypeDef("elz_func", elzFuncType))
	mod.NewFunc("new_elz_func", elzFT, ir.NewParam("parameters_length", irTypes.I64))
	mod.NewFunc("elz_func_append_argument",
		irTypes.Void,
		ir.NewParam("f", elzFT),
		ir.NewParam("arg", irTypes.NewPointer(irTypes.Void)),
	)
	return mod
}

func NewGenerator() *CodeGenerator {
	return &CodeGenerator{
		mod:       prepareingModule(),
		functions: make(map[string]*Func),
	}
}

func (c *CodeGenerator) AddFunc(fn *ast.Func) {
	c.functions[fn.Name] = NewFunc(fn)
}

type Func struct {
	*ast.Func

	params []*ir.Param
}

func NewFunc(f *ast.Func) *Func {
	return &Func{
		Func:   f,
		params: make([]*ir.Param, 0),
	}
}

func (f *Func) CallWith(c *CodeGenerator, block *ir.Block, args []value.Value) value.Value {
	scope := make(map[string]value.Value)
	for i, arg := range args {
		scope[f.Func.ParamList[i]] = arg
		f.params = append(f.params, ir.NewParam(f.Func.ParamList[i], arg.Type()))
	}
	if len(f.params) > len(f.Func.ParamList) {
		panic("too many argument in call")
	}
	if len(f.params) == len(f.Func.ParamList) {
		newFn := c.mod.NewFunc(f.Func.Name, types.I64.LLVMT(), f.params...)
		fnBlock := newFn.NewBlock("")

		fnBlock.NewRet(c.NewExpr(NewScope(fnBlock, scope), f.Func.Expr))
		return block.NewCall(newFn, args...)
	}
	panic("not implement lambda yet")
}

func (c *CodeGenerator) NewExpr(scope *Scope, expr ast.Expr) value.Value {
	switch expr := expr.(type) {
	case *ast.Ident:
		return scope.Var(expr.Value)
	case *ast.BinaryExpr:
		lexpr, rexpr := c.NewExpr(scope, expr.LExpr), c.NewExpr(scope, expr.RExpr)
		return scope.NewAdd(lexpr, rexpr)
	}
	panic("Bomb")
	// return types.I64.NewInt("1")
}
