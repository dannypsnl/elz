package main

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	m := ir.NewModule()

	f64 := types.Double
	i32 := types.I32
	main := m.NewFunction("main", i32)
	entry := main.NewBlock("entry")
	entry.NewFMul(
		constant.NewFloatFromString("4.33", f64), constant.NewFloatFromString("2.1", f64),
	)
	entry.NewRet(constant.NewInt(0, i32))

	fmt.Println(m)
}
