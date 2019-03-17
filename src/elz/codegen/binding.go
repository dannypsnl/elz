package codegen

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"

	"github.com/llir/llvm/ir"
)

type Binding struct {
	*ast.Binding

	compilerProvidedImpl *ir.Func
	cacheOfImpl          map[string]*ir.Func
}

func NewBinding(bind *ast.Binding) *Binding {
	return &Binding{
		Binding: bind,
		// unexported part
		cacheOfImpl: map[string]*ir.Func{},
	}
}

func (b *Binding) GetImpl(g *Generator, typeMap *typeMap, argList ...*ast.Arg) (*ir.Func, error) {
	// FIXME: currently for convenience we skip all checking when it's a built-in function
	// it should be fix after we can do more type checking
	//
	// the reason about this is because we add libc printf into our built-in function,
	// but didn't add `...` variadic into type system which make us can't do type check correctly
	if b.compilerProvidedImpl != nil {
		return b.compilerProvidedImpl, nil
	}
	typeList := typeMap.convertArgsToTypeList(argList...)
	actualTypeWhenCall := genKey(b.Name, typeList...)
	if err := b.CheckArg(argList...); err != nil {
		return nil, err
	}
	if err := b.TypeCheck(actualTypeWhenCall, typeList); err != nil {
		return nil, err
	}
	for i, t := range typeList {
		argName := b.ParamList[i]
		typeMap.add(argName, t)
	}
	inferT, err := g.inferReturnType(b.Expr, typeMap)
	if err != nil {
		return nil, err
	}

	impl, getImpl := b.cacheOfImpl[actualTypeWhenCall]
	if getImpl {
		return impl, nil
	}
	f, err := generateNewImpl(g, b, inferT, typeMap, argList...)
	if err != nil {
		return nil, err
	}
	b.cacheOfImpl[actualTypeWhenCall] = f
	return f, nil
}

func generateNewImpl(g *Generator, bind *Binding, inferT types.Type, typeMap *typeMap, argList ...*ast.Arg) (*ir.Func, error) {
	typeList := typeMap.convertArgsToTypeList(argList...)
	if len(argList) != len(bind.ParamList) {
		return nil, fmt.Errorf(`do not have enough arguments to evaluate binding: %s, argList: %#v`, bind.Name, argList)
	}
	params := make([]*ir.Param, 0)
	for i, arg := range argList {
		params = append(params, ir.NewParam(arg.Ident, typeList[i].LLVMType()))
	}

	f := g.mod.NewFunc(bind.Name, inferT.LLVMType(), params...)

	block := f.NewBlock("")
	binds := make(map[string]*ir.Param)
	for i, p := range params {
		binds[bind.ParamList[i]] = p
	}
	if err := funcBody(g, block, bind.Expr, binds, typeMap); err != nil {
		return nil, err
	}
	return f, nil
}

func funcBody(g *Generator, b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap *typeMap) error {
	v, err := g.genExpr(b, expr, binds, typeMap)
	if err != nil {
		return err
	}
	b.NewRet(v)
	return nil
}
