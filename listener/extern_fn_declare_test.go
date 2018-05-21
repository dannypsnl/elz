package listener

import (
	"testing"
)

func TestFnDelcare(t *testing.T) {
	src := `
	#[extern(c)]
	fn printf(ref<i8>) -> i32
	`

	expected := `
declare i32 @printf(i8*)
`

	hasTestTemplate(t, src, expected)
}
