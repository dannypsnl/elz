package codegen

import (
	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"

	"github.com/llir/llvm/ir"
	llvmtypes "github.com/llir/llvm/ir/types"
)

func generateBuiltin(mod *ir.Module) map[string]*Binding {
	listType := mod.NewTypeDef("list", (&types.List{}).LLVMType())
	listPointerType := llvmtypes.NewPointer(listType)

	builtins := make(map[string]*Binding)

	printfBind := NewBinding(&ast.Binding{
		Name: "printf",
		// format: ...
		ParamList: []*ast.Param{ast.NewParam("format", &ast.VariantType{})},
	})
	printfImpl := mod.NewFunc("printf", llvmtypes.I64,
		ir.NewParam("format", llvmtypes.NewPointer(llvmtypes.I8)),
	)
	printfImpl.Sig.Variadic = true
	printfBind.compilerProvidedImpl = printfImpl
	builtins["printf"] = printfBind

	elzMallocBind := NewBinding(&ast.Binding{
		Name: "elz_malloc",
		// size: size_t(aka int)
		ParamList: []*ast.Param{ast.NewParam("size", &ast.VariantType{})},
	})
	elzMallocBind.compilerProvidedImpl = mod.NewFunc("elz_malloc", llvmtypes.NewPointer(llvmtypes.I8),
		ir.NewParam("size", llvmtypes.I64),
	)
	builtins["elz_malloc"] = elzMallocBind

	newListBind := NewBinding(&ast.Binding{
		Name: "new_list",
		// size: int, elements: void **
		ParamList: []*ast.Param{
			ast.NewParam("size", &ast.VariantType{}),
			ast.NewParam("elements", &ast.VariantType{}),
		},
	})
	newListBind.compilerProvidedImpl = mod.NewFunc("new_list", listPointerType,
		ir.NewParam("size", llvmtypes.I64),
		ir.NewParam("elements", llvmtypes.NewPointer(llvmtypes.NewPointer(llvmtypes.I8))),
	)
	builtins["new_list"] = newListBind

	listIndexBind := NewBinding(&ast.Binding{
		Name: "list_index",
		// list: list, index: i64
		ParamList: []*ast.Param{
			ast.NewParam("list", &ast.VariantType{}),
			ast.NewParam("index", &ast.VariantType{}),
		},
	})
	listIndexBind.compilerProvidedImpl = mod.NewFunc("list_index", llvmtypes.NewPointer(llvmtypes.I8),
		ir.NewParam("list", listPointerType),
		ir.NewParam("index", llvmtypes.I64),
	)
	builtins["list_index"] = listIndexBind

	listLengthBind := NewBinding(&ast.Binding{
		Name: "list_length",
		ParamList: []*ast.Param{
			ast.NewParam("list", &ast.VariantType{}),
		},
	})
	listLengthBind.compilerProvidedImpl = mod.NewFunc("list_length", llvmtypes.I64,
		ir.NewParam("list", listPointerType),
	)
	builtins["list_length"] = listLengthBind

	return builtins
}
