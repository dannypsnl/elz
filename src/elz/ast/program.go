package ast

type Program struct {
	Bindings     []*Binding
	BindingTypes []*BindingType
	Imports      []*Import
}

func (p *Program) AddBinding(b *Binding) {
	p.Bindings = append(p.Bindings, b)
}

func (p *Program) AddBindingType(b *BindingType) {
	p.BindingTypes = append(p.BindingTypes, b)
}

func (p *Program) AddImport(i *Import) {
	p.Imports = append(p.Imports, i)
}
