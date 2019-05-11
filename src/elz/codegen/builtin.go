package codegen

import (
	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/types"

	"github.com/llir/llvm/ir"
	llvmtypes "github.com/llir/llvm/ir/types"
)

func generateBuiltin(mod *ir.Module) map[string]*Binding {
	listType := mod.NewTypeDef("list", (&types.List{}).LLVMType())

	builtins := make(map[string]*Binding)

	printfBind := NewBinding(&ast.Binding{
		Name: "printf",
		// format: ...
		ParamList: []string{"format"},
	})
	printfImpl := mod.NewFunc("printf", llvmtypes.I64,
		ir.NewParam("format", llvmtypes.NewPointer(llvmtypes.I8)),
	)
	printfImpl.Sig.Variadic = true
	printfBind.compilerProvidedImpl = printfImpl
	builtins["printf"] = printfBind

	newListBind := NewBinding(&ast.Binding{
		Name: "new_list",
		// size: int, elements: void **
		ParamList: []string{"size", "elements"},
	})
	newListBind.compilerProvidedImpl = mod.NewFunc("new_list", listType,
		ir.NewParam("size", llvmtypes.I64),
		ir.NewParam("elements", llvmtypes.NewPointer(llvmtypes.NewPointer(llvmtypes.I8))),
	)
	builtins["new_list"] = newListBind

	listIndexBind := NewBinding(&ast.Binding{
		Name: "list_index",
		// list: list, index: i64
		ParamList: []string{"list", "index"},
	})
	listIndexBind.compilerProvidedImpl = mod.NewFunc("list_index", llvmtypes.NewPointer(llvmtypes.I8),
		ir.NewParam("list", listType),
		ir.NewParam("index", llvmtypes.I64),
	)
	builtins["list_index"] = listIndexBind

	listLengthBind := NewBinding(&ast.Binding{
		Name:      "list_length",
		ParamList: []string{"list"},
	})
	listLengthBind.compilerProvidedImpl = mod.NewFunc("list_length", llvmtypes.I64,
		ir.NewParam("list", listType),
	)
	builtins["list_length"] = listLengthBind

	return builtins
}
