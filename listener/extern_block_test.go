package listener

import (
	"testing"
)

func TestExternDeclare(t *testing.T) {
	src := `
	extern "C" {
		fn add_one(i32) -> i32
		fn printf(ref<i8>) -> i32
		fn printf(ref<i8>, i32, i32) -> i32
	}
	`

	expected := `
declare i32 @add_one(i32)

declare i32 @printf(i8*)

declare i32 @printf.1(i8*, i32, i32)
`

	hasTestTemplate(t, src, expected)
}
