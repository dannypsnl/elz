package listener

import (
	"testing"
)

func TestNumAddFunction(t *testing.T) {
	res := NewParse(`
	fn add(l: num, r: num) -> num {}
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

define float @add(float %l, float %r) {
entry:
}
`

	if res != expected {
		t.Errorf("Function Error; expected: `%s`\nactual: `%s`", expected, res)
	}
}
