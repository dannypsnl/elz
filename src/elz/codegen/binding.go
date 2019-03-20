package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"

	"github.com/llir/llvm/ir"
)

type Binding struct {
	*ast.Binding

	compilerProvidedImpl *ir.Func
	cacheOfImpl          map[string]*ir.Func
	cacheOfType          map[string]types.Type
}

func NewBinding(bind *ast.Binding) *Binding {
	return &Binding{
		Binding: bind,
		// unexported part
		cacheOfImpl: map[string]*ir.Func{},
		cacheOfType: map[string]types.Type{},
	}
}

func (b *Binding) GetReturnType(g *Generator, typeMap *typeMap, typeListOfArgs ...types.Type) (types.Type, error) {
	key := typeFormat(typeListOfArgs...)
	t, ok := b.cacheOfType[key]
	if ok {
		return t, nil
	}
	inferT, err := g.inferTypeOf(b.Expr, typeMap)
	if err != nil {
		return nil, err
	}
	b.cacheOfType[key] = inferT
	return inferT, nil
}

func (b *Binding) GetImpl(m *module, typeMap *typeMap, argList ...*ast.Arg) (*ir.Func, error) {
	// FIXME: currently for convenience we skip all checking when it's a built-in function
	// it should be fix after we can do more type checking
	//
	// the reason about this is because we add libc printf into our built-in function,
	// but didn't add `...` variadic into type system which make us can't do type check correctly
	if b.compilerProvidedImpl != nil {
		return b.compilerProvidedImpl, nil
	}
	typeListOfArgs := typeMap.convertArgsToTypeList(argList...)
	certainTypeFormatOfArgs := typeFormat(typeListOfArgs...)
	// if we could get the implementation that we don't have to do any checking
	// because it must already be checked
	impl, getImpl := b.cacheOfImpl[certainTypeFormatOfArgs]
	if getImpl {
		return impl, nil
	}

	if err := b.checkArg(argList...); err != nil {
		return nil, err
	}
	if err := b.typeCheck(typeListOfArgs); err != nil {
		return nil, err
	}
	// - record parameters' type
	// - create parameters of IR
	params := make([]*ir.Param, 0)
	for i, paramType := range typeListOfArgs {
		paramName := b.ParamList[i]
		typeMap.add(paramName, paramType)
		params = append(params, ir.NewParam(paramName, paramType.LLVMType()))
	}
	returnType, err := b.GetReturnType(m.generator, typeMap, typeListOfArgs...)
	if err != nil {
		return nil, err
	}
	function, err := generateNewImpl(m, b, returnType, typeMap, params)
	if err != nil {
		return nil, err
	}
	b.cacheOfImpl[certainTypeFormatOfArgs] = function
	return function, nil
}

func generateNewImpl(m *module, bind *Binding, returnType types.Type, typeMap *typeMap, params []*ir.Param) (*ir.Func, error) {
	if len(params) != len(bind.ParamList) {
		return nil, fmt.Errorf(`do not have enough arguments to evaluate binding: %s`, bind.Name)
	}
	function := m.generator.mod.NewFunc(bind.Name, returnType.LLVMType(), params...)
	block := function.NewBlock("")
	binds := make(map[string]*ir.Param)
	for i, p := range params {
		binds[bind.ParamList[i]] = p
	}
	err := funcBody(m, block, bind.Expr, binds, typeMap)
	if err != nil {
		return nil, err
	}
	return function, nil
}

func funcBody(m *module, b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap *typeMap) error {
	v, err := m.genExpr(b, expr, binds, typeMap)
	if err != nil {
		return err
	}
	b.NewRet(v)
	return nil
}

func (b *Binding) checkArg(args ...*ast.Arg) error {
	for i, arg := range args {
		argNameMustBe := b.ParamList[i]
		argName := arg.Ident
		// allow ignore argument name like: `add(1, 2)`
		if argName == "" {
			argName = argNameMustBe
		}
		if argNameMustBe != argName {
			return fmt.Errorf(`argument name must be parameter name(or empty), for example:
  assert that should_be = ...
  assert(that: 1+2, should_be: 3)
`)
		}
	}
	return nil
}

func (b *Binding) typeCheck(typeList []types.Type) error {
	if b.Type == nil {
		return nil
	}
	var (
		builder        strings.Builder
		err            error
		variantTypeMap = map[string]string{}
	)
	for i, requireT := range b.Type[:len(b.Type)-1] {
		actualType := typeList[i]
		switch requireT := requireT.(type) {
		case *ast.ExistType:
			if requireT.Name != actualType.String() {
				err = fmt.Errorf("")
			}
		case *ast.VariantType:
			t, exist := variantTypeMap[requireT.Name]
			if !exist {
				variantTypeMap[requireT.Name] = actualType.String()
				t = actualType.String()
			}
			if t != actualType.String() {
				err = fmt.Errorf("")
			}
		case *ast.VoidType:
		}
		builder.WriteString(requireT.String())
		builder.WriteString(" -> ")
	}
	requireT := builder.String()
	requireType := b.Name + " :: " + requireT[:len(requireT)-4]
	if err != nil {
		// Format would like: `bindName :: type -> type -> type`
		actualCallTypeFormat := b.Name + " :: " + typeFormat(typeList...)
		return fmt.Errorf("require type: `%s` but get: `%s`", requireType, actualCallTypeFormat)
	}
	return nil
}
