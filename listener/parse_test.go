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

@a = global float 1.000000e+01
`

	if res != expected {
		t.Errorf("expected: `%s`\nactual: `%s`", expected, res)
	}
}
