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

	expected := `yee`

	hasTestTemplate(t, src, expected)
}
