package listener

import (
	_ "strings"
	"testing"
)

func TestNumStrExpression(t *testing.T) {
	res := NewParse(`
	a = 3 * 2
	b = "abcde"
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

@a = global i32 6
@b = global [7 x i8] c"\22abcde\22"
`

	if res != expected {
		t.Errorf("expected: `%s`\nactual: `%s`", expected, res)
	}
}
