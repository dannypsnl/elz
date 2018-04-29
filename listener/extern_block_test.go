package listener

import (
	"testing"
)

func TestExternDeclare(t *testing.T) {
	res := NewParse(`
	extern "C" {
		fn add_one(i32) -> i32
	}
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

declare i32 @add_one(i32)
`

	if res != expected {
		t.Errorf("expected: `%s`\nactual: `%s`", expected, res)
	}
}
