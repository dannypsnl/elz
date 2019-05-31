package types

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/sirupsen/logrus"
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
		ElemT Type
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
		logrus.Fatalf("unknown type: %s", typeName)
		// dead code return for compiler
		return nil
	}
}
