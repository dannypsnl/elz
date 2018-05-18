package listener

import (
	"testing"

	"strings"
)

func TestArray(t *testing.T) {
	arrayTestTemplate(t, `a = [1; 3]`, `@a = global [3 x i32] [i32 1, i32 1, i32 1]`)
	arrayTestTemplate(t, `a = [1, 2, 3]`, `@a = global [3 x i32] [i32 1, i32 2, i32 3]`)
}

func arrayTestTemplate(t *testing.T, source, expectedIn string) {
	res := NewParse(source)

	if !strings.Contains(res, expectedIn) {
		t.Errorf("expected has: `%s`, actual: `%s`", expectedIn, res)
	}
}
