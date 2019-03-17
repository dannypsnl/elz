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

	if err := b.checkArg(argList...); err != nil {
		return nil, err
	}
	if err := b.typeCheck(typeList); err != nil {
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

	certainTypeFormatOfArgs := typeFormat(typeList...)
	impl, getImpl := b.cacheOfImpl[certainTypeFormatOfArgs]
	if getImpl {
		return impl, nil
	}
	f, err := generateNewImpl(g, b, inferT, typeMap, argList...)
	if err != nil {
		return nil, err
	}
	b.cacheOfImpl[certainTypeFormatOfArgs] = f
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
