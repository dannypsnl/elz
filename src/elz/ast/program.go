package ast

type Program struct {
	TypeDefines []*TypeDefine
	Bindings    []*Binding
	Imports     []*Import
}

func (p *Program) AddTypeDefine(typeDef *TypeDefine) {
	p.TypeDefines = append(p.TypeDefines, typeDef)
}

func (p *Program) AddBinding(b *Binding) {
	p.Bindings = append(p.Bindings, b)
}

func (p *Program) AddImport(i *Import) {
	p.Imports = append(p.Imports, i)
}
