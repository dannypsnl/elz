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
)

type Generator struct {
	mod *ir.Module

	allModule   map[string]*module
	entryModule *module

	operatorTypeStore map[string]types.Type
}

func New(entryTree *Tree, allAstTree map[string]*Tree) *Generator {
	typMap := make(map[string]types.Type)
	typMap["+ :: int -> int"] = &types.Int{}
	typMap["+ :: f64 -> f64"] = &types.Float{}

	err := entryTree.InsertBinding(&ast.Binding{
		Name:      "printf",
		ParamList: []string{"format"},
	})
	if err != nil {
		panic("some function conflict with built-in function printf")
	}
	mod := ir.NewModule()
	g := &Generator{
		mod:               mod,
		operatorTypeStore: typMap,
	}
	allModule := make(map[string]*module)
	for name, tree := range allAstTree {
		allModule[name] = newModule(g, tree)
	}
	g.allModule = allModule
	g.entryModule = newModule(g, entryTree)

	printfImpl := mod.NewFunc("printf", llvmtypes.I64,
		ir.NewParam("format", llvmtypes.NewPointer(llvmtypes.I8)),
	)
	printfImpl.Sig.Variadic = true
	printfBind, err := entryTree.GetBinding("printf")
	if err != nil {
		panic(fmt.Errorf("can't get printf binding: %s", err))
	}
	printfBind.compilerProvidedImpl = printfImpl

	return g
}

func (g *Generator) String() string {
	irutil.FixDups(g.mod)
	return g.mod.String()
}

func (g *Generator) Generate() {
	entryBinding, err := g.entryModule.GetBinding("main")
	if err != nil {
		panic("no main function exist, no compile")
	}
	if len(entryBinding.ParamList) > 0 {
		panic("main function should not have any parameters")
	}
	impl := g.mod.NewFunc("main", llvmtypes.I64)
	b := impl.NewBlock("")
	_, err = g.entryModule.genExpr(b, entryBinding.Expr, make(map[string]*ir.Param), newTypeMap())
	if err != nil {
		panic(fmt.Sprintf("report error: %s", err))
	}
	b.NewRet(constant.NewInt(llvmtypes.I64, 0))
}

func (g *Generator) Call(bind *Binding, exprList ...*ast.Arg) error {
	_, err := bind.GetImpl(newTypeMap(), exprList...)
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
