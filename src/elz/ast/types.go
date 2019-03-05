package ast

import (
	"fmt"
)

type isType struct{}

func (isType) IsType() bool {
	return true
}

type (
	Type interface {
		IsType() bool
		fmt.Stringer
	}

	ExistType struct {
		isType
		Name string
	}
	VoidType struct {
		isType
	}
	VariantType struct {
		isType
		Name string
	}
)

func (t *ExistType) String() string {
	return t.Name
}
func (t *VoidType) String() string {
	return "()"
}
func (t *VariantType) String() string {
	return "'" + t.Name
}
