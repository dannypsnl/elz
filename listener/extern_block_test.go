package listener

import (
	"testing"
)

func TestExternDeclare(t *testing.T) {
	res := NewParse(`
	extern "C" {
		fn add_one(i32) -> i32
		fn printf(ref<i8>) -> i32
		fn printf(ref<i8>, i32, i32) -> i32
	}
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

declare i32 @add_one(i32)

declare i32 @printf(i8*)

declare i32 @printf.1(i8*, i32, i32)
`

	if res != expected {
		t.Errorf("expected: `%s`\nactual: `%s`", expected, res)
	}
}
