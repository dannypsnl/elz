package codegen

import (
	"fmt"

	"github.com/elz-lang/elz/src/elz/ast"
)

type Tree struct {
	imports     []string
	bindings    map[string]*Binding
	typeDefines map[string]*ast.NewType
}

func NewTree() *Tree {
	return &Tree{
		imports:     make([]string, 0),
		bindings:    make(map[string]*Binding),
		typeDefines: make(map[string]*ast.NewType),
	}
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

func (t *Tree) InsertTypeDefine(typDef *ast.NewType) error {
	_, exist := t.bindings[typDef.Name]
	if exist {
		return fmt.Errorf("type: %s already exist", typDef.Name)
	}
	t.typeDefines[typDef.Name] = typDef
	return nil
}

func (t *Tree) GetTypeDefine(typeName string) (*ast.NewType, error) {
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
