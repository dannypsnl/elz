package codegen

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"
)

type Tree struct {
	imports     []string
	bindings    map[string]*Binding
	typeDefines map[string]*ast.TypeDefine
}

func NewTree(program *ast.Program) (*Tree, error) {
	newTree := &Tree{
		imports:     make([]string, 0),
		bindings:    make(map[string]*Binding),
		typeDefines: make(map[string]*ast.TypeDefine),
	}
	for _, im := range program.Imports {
		newTree.InsertImport(im.AccessChain.Literal)
	}
	for _, typeDef := range program.TypeDefines {
		err := newTree.InsertTypeDefine(typeDef)
		if err != nil {
			return nil, err
		}
	}
	for _, binding := range program.Bindings {
		err := newTree.InsertBinding(binding)
		if err != nil {
			return nil, err
		}
	}
	for _, bindingT := range program.BindingTypes {
		err := newTree.InsertBindingType(bindingT)
		if err != nil {
			return nil, err
		}
	}
	return newTree, nil
}

func (t *Tree) InsertBindingType(bt *ast.BindingType) error {
	_, exist := t.bindings[bt.Name]
	if !exist {
		return fmt.Errorf("no binding: %s exist", bt.Name)
	}
	t.bindings[bt.Name].TypeList = bt.Type
	return nil
}

func (t *Tree) InsertBinding(b *ast.Binding) error {
	_, exist := t.bindings[b.Name]
	if exist {
		return fmt.Errorf("binding: %s already exist", b.Name)
	}
	t.bindings[b.Name] = NewBinding(b)
	return nil
}

func (t *Tree) GetExportBinding(bindName string) (*Binding, error) {
	b, exist := t.bindings[bindName]
	if !exist || !b.Export {
		return nil, fmt.Errorf("no export binding call: `%s`", bindName)
	}
	return b, nil
}

func (t *Tree) GetBinding(bindName string) (*Binding, error) {
	binding, exist := t.bindings[bindName]
	if !exist {
		return nil, fmt.Errorf("no binding call: `%s`", bindName)
	}
	return binding, nil
}

func (t *Tree) InsertTypeDefine(typDef *ast.TypeDefine) error {
	_, exist := t.typeDefines[typDef.Name]
	if exist {
		return fmt.Errorf("type: %s already exist", typDef.Name)
	}
	t.typeDefines[typDef.Name] = typDef
	return nil
}

func (t *Tree) GetTypeDefines() map[string]*ast.TypeDefine {
	return t.typeDefines
}

func (t *Tree) GetExportTypeDefine(typeName string) (*ast.TypeDefine, error) {
	typeDef, err := t.GetTypeDefine(typeName)
	if err != nil {
		return nil, err
	}
	if !typeDef.Export {
		return nil, fmt.Errorf("no export type call: %s", typeName)
	}
	return typeDef, nil
}

func (t *Tree) GetTypeDefine(typeName string) (*ast.TypeDefine, error) {
	typDef, exist := t.typeDefines[typeName]
	if !exist {
		return nil, fmt.Errorf("no type call: `%s`", typeName)
	}
	return typDef, nil
}

func (t *Tree) InsertImport(s string) {
	t.imports = append(t.imports, s)
}

func (t *Tree) GetDependencies() []string {
	return t.imports
}
