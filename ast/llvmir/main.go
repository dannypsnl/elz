package main

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	m := ir.NewModule()

	f32 := types.Float
	i32 := types.I32
	main := m.NewFunction("main", i32)
	entry := main.NewBlock("entry")
	entry.NewFMul(
		constant.NewFloatFromString("4.3300000e+00", f32), constant.NewFloatFromString("2.10000000e+00", f32),
	)
	entry.NewRet(constant.NewInt(0, i32))

	fmt.Println(m)
}
