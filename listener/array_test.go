package listener

import (
	"testing"
)

func TestArray(t *testing.T) {
	hasTestTemplate(t, `a = [1; 3]`, `@a = global [3 x i32] [i32 1, i32 1, i32 1]`)
	hasTestTemplate(t, `a = [1, 2, 3]`, `@a = global [3 x i32] [i32 1, i32 2, i32 3]`)
}
