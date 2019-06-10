package codegen

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"
	"github.com/elz-lang/elz/src/elz/value"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/enum"
)

type Binding struct {
	value.Value
	*ast.Binding

	selfModule           *module
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

func (b *Binding) SetModule(m *module) {
	b.selfModule = m
}

func (b *Binding) getReturnType(typeMap *types.TypeMap, typeListOfArgs ...types.Type) (types.Type, error) {
	key := typeFormat(typeListOfArgs...)
	t, ok := b.cacheOfType[key]
	if ok {
		return t, nil
	}
	inferT, err := b.selfModule.InferTypeOf(b.Expr, typeMap)
	if err != nil {
		return nil, err
	}
	b.cacheOfType[key] = inferT
	return inferT, nil
}

func (b *Binding) GetImpl(typeMap *types.TypeMap, argList ...*ast.Arg) (*ir.Func, error) {
	// FIXME: currently for convenience we skip all checking when it's a built-in function
	// it should be fix after we can do more type checking
	//
	// the reason about this is because we add libc printf into our built-in function,
	// but didn't add `...` variadic into type system which make us can't do type check correctly
	if b.compilerProvidedImpl != nil {
		return b.compilerProvidedImpl, nil
	}
	typeListOfArgs, err := typeMap.ConvertArgsToTypeList(argList...)
	if err != nil {
		return nil, err
	}
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
		param := b.ParamList[i]
		typeMap.Add(param.Name, paramType)
		params = append(params, ir.NewParam(param.Name, paramType.LLVMType()))
	}
	returnType, err := b.getReturnType(typeMap, typeListOfArgs...)
	if err != nil {
		return nil, err
	}
	function, err := generateNewImpl(b, returnType, typeMap, params)
	if err != nil {
		return nil, err
	}
	b.cacheOfImpl[certainTypeFormatOfArgs] = function
	return function, nil
}

func generateNewImpl(bind *Binding, returnType types.Type, typeMap *types.TypeMap, params []*ir.Param) (*ir.Func, error) {
	if len(params) != len(bind.ParamList) {
		return nil, fmt.Errorf(`do not have enough arguments to evaluate binding: %s`, bind.Name)
	}
	function := bind.selfModule.generator.mod.NewFunc(
		bind.Name,
		returnType.LLVMType(),
		params...,
	)
	function.Linkage = enum.LinkageInternal
	block := function.NewBlock("")
	binds := make(map[string]*ir.Param)
	for i, p := range params {
		binds[bind.ParamList[i].Name] = p
	}
	err := funcBody(bind.selfModule, block, bind.Expr, binds, typeMap)
	if err != nil {
		return nil, err
	}
	return function, nil
}

func funcBody(m *module, b *ir.Block, expr ast.Expr, binds map[string]*ir.Param, typeMap *types.TypeMap) error {
	newCtx := newContext(b, typeMap)
	newCtx.binds = binds
	v, err := m.genExpr(newCtx, expr)
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
			argName = argNameMustBe.Name
		}
		if argNameMustBe.Name != argName {
			return fmt.Errorf(`argument name must be parameter name(or empty), for example:
  assert that should_be = ...
  assert(that: 1+2, should_be: 3)
`)
		}
	}
	return nil
}

func (b *Binding) typeCheck(typeList []types.Type) error {
	if len(b.TypeList) == 0 {
		return nil
	}
	var (
		builder        strings.Builder
		err            error
		variantTypeMap = map[string]string{}
	)
	for i, requireT := range b.TypeList[:len(b.TypeList)-1] {
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
