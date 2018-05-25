package listener

import (
	"testing"
)

func TestTypeDefine(t *testing.T) {
	src := `
	type Foo (
		a: i32,
		b: i8
	)
	`

	expected := `%Foo = type <{ i32, i8 }>`

	hasTestTemplate(t, src, expected)
}
