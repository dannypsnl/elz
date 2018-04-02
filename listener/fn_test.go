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

func TestMainFunction(t *testing.T) {
	res := NewParse(`
	fn main() {}
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

define i32 @main() {
entry:
  ret i32 0
}
`

	if res != expected {
		t.Errorf("Main Function Error; expected: `%s`\nactual: `%s`", expected, res)
	}
}
