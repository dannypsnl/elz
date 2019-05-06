package types

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
)

type isType struct{}

func (*isType) IsType() bool { return true }

type Type interface {
	IsType() bool
	fmt.Stringer
	LLVMType() types.Type
}

// Int represents Elz Integer type
type (
	Int struct {
		isType
	}
	Float struct {
		isType
	}
	String struct {
		isType
	}
	List struct {
		isType
		elemT Type
	}
)

func FromString(typeName string) Type {
	switch typeName {
	case "string":
		return &String{}
	case "int":
		return &Int{}
	case "f64":
		return &Float{}
	default:
		panic(fmt.Errorf("unknown type: %s", typeName))
	}
}
