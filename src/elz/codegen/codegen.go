package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/internal/irutil"
	"github.com/elz-lang/elz/src/elz/types"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmtypes "github.com/llir/llvm/ir/types"

	"github.com/sirupsen/logrus"
)

type Generator struct {
	mod *ir.Module

	allModule   map[string]*module
	entryModule *module

	builtin map[string]*Binding

	operatorTypeStore map[string]types.Type
}

func New(entryTree *Tree, allAstTree map[string]*Tree) *Generator {
	typMap := make(map[string]types.Type)
	typMap["+(int, int)"] = &types.Int{}
	typMap["+(f64, f64)"] = &types.Float{}

	mod := ir.NewModule()
	builtin := generateBuiltin(mod)
	g := &Generator{
		mod:               mod,
		operatorTypeStore: typMap,
		builtin:           builtin,
	}
	allModule := make(map[string]*module)
	for name, tree := range allAstTree {
		allModule[name] = newModule(g, name, tree)
	}
	g.allModule = allModule
	g.entryModule = newModule(g, "", entryTree)
	return g
}

func (g *Generator) String() string {
	irutil.FixDups(g.mod)
	return g.mod.String()
}

func (g *Generator) Generate() {
	entryBinding, err := g.entryModule.GetBinding("main")
	if err != nil {
		logrus.Fatalf("no main function exist, no compile")
		return // compiler notation
	}
	if len(entryBinding.ParamList) > 0 {
		logrus.Fatalf("main function should not have any parameters")
	}
	impl := g.mod.NewFunc("main", llvmtypes.I64)
	// call init function of module to sure global variable would be initialized
	b := impl.NewBlock("")

	g.entryModule.initFuncBlock.NewRet(nil)
	b.NewCall(g.entryModule.initFunc)
	for _, mod := range g.allModule {
		mod.initFuncBlock.NewRet(nil)
		b.NewCall(mod.initFunc)
	}

	ctx := newContext(b, g.entryModule.typeMap)
	_, err = g.entryModule.genExpr(ctx, entryBinding.Expr)
	if err != nil {
		logrus.Fatalf("report error: %s", err)
	}
	ctx.NewRet(constant.NewInt(llvmtypes.I64, 0))
}

func (g *Generator) Call(bind *Binding, exprList ...*ast.Arg) error {
	g.entryModule.initFuncBlock.NewRet(nil)
	_, err := bind.GetImpl(g.entryModule.typeMap, exprList...)
	if err != nil {
		return err
	}
	return nil
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

func (g *Generator) getBuiltin(name string) (*Binding, error) {
	b, ok := g.builtin[name]
	if !ok {
		return nil, fmt.Errorf("no builtin binding call: `%s`", name)
	}
	return b, nil
}

func genKey(bindName string, typeList ...types.Type) string {
	var b strings.Builder
	b.WriteString(bindName)
	b.WriteString(typeFormat(typeList...))
	return b.String()
}

func typeFormat(typeList ...types.Type) string {
	var b strings.Builder
	b.WriteRune('(')
	if len(typeList) > 0 {
		for _, t := range typeList[:len(typeList)-1] {
			b.WriteString(t.String())
			b.WriteString(", ")
		}
		b.WriteString(typeList[len(typeList)-1].String())
	}
	b.WriteRune(')')
	return b.String()
}
