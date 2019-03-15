package codegen

import (
	"fmt"
	"github.com/elz-lang/elz/src/elz/ast"
	"path/filepath"
	"strings"
)

type Tree struct {
	imports  []string
	bindings map[string]*Binding
}

func NewTree() *Tree {
	return &Tree{
		imports:  make([]string, 0),
		bindings: make(map[string]*Binding),
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
	// FIXME: assuming only one binding for right now situation, should do complete discuss into how to deal with multiple implementations
	return binding, nil
}

func (t *Tree) InsertImport(s string) {
	t.imports = append(t.imports, s)
}

func (t *Tree) GetDependencies() []string {
	deps := make([]string, 0)
	for _, importStr := range t.imports {
		dep := filepath.Join(strings.Split(importStr, "::")...)
		deps = append(deps, dep+".elz")
	}
	return deps
}
