package listener

import (
	"testing"
)

func TestNumStrExpression(t *testing.T) {
	src := `
	a = 3 * 2
	b = "abcde"
	`

	expected := `; ModuleID = 'main'
source_filename = "main"

@a = global i32 6
@b = global [6 x i8] c"abcde\00"
`

	hasTestTemplate(t, src, expected)
}
