package listener

import (
	"testing"
)

func TestF32(t *testing.T) {
	res := NewParse(`
	x = 3.2
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

@x = global float 0x40099999A0000000
`

	if expected != res {
		t.Errorf("f32 parse Error; expected: `%s`\nactual: `%s`", expected, res)
	}
}
