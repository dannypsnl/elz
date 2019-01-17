package main

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irTypes "github.com/llir/llvm/ir/types"
)

func main() {
	mod := prepareingModule()
	mainFn := mod.NewFunc("main", irTypes.I32)
	block := mainFn.NewBlock("")
	// var elzFuncT irTypes.Type
	// for _, t := range mod.TypeDefs {
	// 	if t.Name() == "elz_func" {
	// 		elzFuncT = t
	// 	}
	// }
	var (
		newElzFunc, appendArgToElzFunc *ir.Func
	)
	for _, f := range mod.Funcs {
		if f.Name() == "new_elz_func" {
			newElzFunc = f
		} else if f.Name() == "elz_func_append_argument" {
			appendArgToElzFunc = f
		}
	}
	elzFunc := block.NewCall(newElzFunc, constant.NewInt(irTypes.I64, 2))
	block.NewCall(appendArgToElzFunc, elzFunc)

	block.NewRet(constant.NewInt(irTypes.I32, 0))
	fmt.Printf("%s\n", mod)
}

func prepareingModule() *ir.Module {
	mod := ir.NewModule()
	elzFuncType := irTypes.NewStruct()
	elzFuncType.Opaque = true
	elzFT := irTypes.NewPointer(mod.NewTypeDef("elz_func", elzFuncType))
	mod.NewFunc("new_elz_func", elzFT, ir.NewParam("parameters_length", irTypes.I64))
	mod.NewFunc("elz_func_append_argument",
		irTypes.Void,
		ir.NewParam("f", elzFT),
		ir.NewParam("arg", irTypes.NewPointer(irTypes.Void)),
	)
	return mod
}
