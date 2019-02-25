package ast

type isType struct{}

func (isType) IsType() bool {
	return true
}

type (
	Type interface {
		IsType() bool
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
