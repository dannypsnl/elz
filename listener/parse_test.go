package listener

import (
	"testing"
)

func TestParse(t *testing.T) {
	res := NewParse(`
	a = 10
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

@a = global i32 10
`

	if res != expected {
		t.Errorf("expected: `%s`\nactual: `%s`", expected, res)
	}
}
