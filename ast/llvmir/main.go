package main

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	f64 := types.Double
	m := ir.NewModule()

	n_v := constant.NewFMul(
		constant.NewFloatFromString("4.33", f64),
		constant.NewFAdd(
			constant.NewFloatFromString("2.1", f64),
			constant.NewFloatFromString("1.2", f64),
		),
	)
	nest := m.NewGlobalDef("nest", n_v)
	main := m.NewFunction("nestedinstruction", f64)
	entry := main.NewBlock("entry")
	entry.NewRet(entry.NewLoad(nest))

	fmt.Println(m)
}
