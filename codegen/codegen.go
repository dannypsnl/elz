package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/types"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Generator struct {
	mod *ir.Module

	implsOfBinding map[string]*ir.Func
}

func New() *Generator {
	return &Generator{
		mod:            ir.NewModule(),
		implsOfBinding: make(map[string]*ir.Func),
	}
}

func (g *Generator) String() string {
	return g.mod.String()
}

func (g *Generator) Call(bind *ast.Binding, exprList ...*ast.Arg) {
	g.mustGetImpl(bind, exprList...)
}

func (g *Generator) mustGetImpl(bind *ast.Binding, exprList ...*ast.Arg) *ir.Func {
	bindName := bind.Name
	key := genKey(bindName, exprList...)
	impl, getImpl := g.implsOfBinding[key]
	if getImpl {
		return impl
	}
	if len(exprList) != len(bind.ParamList) {
		panic(`do not have enough arguments to call function`)
	}
	binds := make(map[string]ast.Expr)
	params := make([]*ir.Param, 0)
	for i, e := range exprList {
		argNameMustBe := bind.ParamList[i]
		argName := e.Ident
		if argName == "" {
			argName = argNameMustBe
		}
		if argNameMustBe != argName {
			panic(`argument name must be parameter name(or empty), for example:
  assert that should_be = ...
  assert(that: 1+2, should_be: 3)
`)
		}
		t := types.TypeOf(e)
		binds[argName] = e.Expr
		params = append(params, ir.NewParam(e.Ident, t.LLVMType()))
	}
	// FIXME: Hard code return type first, please fix it by calling type infer function
	f := g.mod.NewFunc(bindName, llvmtypes.I64, params...)

	b := f.NewBlock("")
	g.funcBody(b, bind.Expr, binds)

	g.implsOfBinding[key] = f
	return f
}

func (g *Generator) funcBody(b *ir.Block, expr ast.Expr, binds map[string]ast.Expr) {
	v := genExpr(b, expr, binds)
	b.NewRet(v)
}

func genExpr(b *ir.Block, expr ast.Expr, binds map[string]ast.Expr) value.Value {
	switch expr := expr.(type) {
	case *ast.BinaryExpr:
		if expr.Op == "+" {
			x := genExpr(b, expr.LExpr, binds)
			y := genExpr(b, expr.RExpr, binds)
			return b.NewAdd(x, y)
		}
		panic(fmt.Sprintf("unsupported operator: %s", expr.Op))
	case *ast.Ident:
		e := binds[expr.Literal]
		return genExpr(b, e, binds)
	case *ast.Int:
		v, err := constant.NewIntFromString(llvmtypes.I64, expr.Literal)
		if err != nil {
		}
		return v
	default:
		panic(fmt.Sprintf("failed at generate expression: %#v", expr))
	}
}

func genKey(bindName string, exprList ...*ast.Arg) string {
	var b strings.Builder
	b.WriteString(bindName)
	if len(exprList) > 0 {
		b.WriteRune('(')
		for _, e := range exprList[:len(exprList)-1] {
			typeOfExpr := types.TypeOf(e)
			b.WriteString(typeOfExpr.String())
			b.WriteRune(',')
		}
		b.WriteString(types.TypeOf(exprList[len(exprList)-1]).String())
		b.WriteRune(')')
	}
	return b.String()
}
