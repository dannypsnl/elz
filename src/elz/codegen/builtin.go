package codegen

import (
	"github.com/elz-lang/elz/src/elz/ast"

	"github.com/llir/llvm/ir"
	llvmtypes "github.com/llir/llvm/ir/types"
)

func generateBuiltin(mod *ir.Module) map[string]*Binding {
	builtins := make(map[string]*Binding)

	printfBind := NewBinding(&ast.Binding{
		Name:      "printf",
		ParamList: []string{"format"},
	})
	printfImpl := mod.NewFunc("printf", llvmtypes.I64,
		ir.NewParam("format", llvmtypes.NewPointer(llvmtypes.I8)),
	)
	printfImpl.Sig.Variadic = true
	printfBind.compilerProvidedImpl = printfImpl
	builtins["printf"] = printfBind

	return builtins
}
