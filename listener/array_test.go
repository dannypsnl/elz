package listener

import (
	"testing"
)

func TestArray(t *testing.T) {
	hasTestTemplate(t, `a = [1; 3]`, `@a = global [3 x i32] [i32 1, i32 1, i32 1]`)
	hasTestTemplate(t, `a = [1, 2, 3]`, `@a = global [3 x i32] [i32 1, i32 2, i32 3]`)
}

// FIXME: using execution engine to test access array element
func TestAccessArrayElem(t *testing.T) {
	src := `
	fn main() {
		let a = [1, 2, 3]
		let b = a[0]
		let c = a[1]
		let d = 1 + a[0]
	}
	`

	expected := `
  %0 = load [3 x i32], [3 x i32]* %a
  %1 = extractvalue [3 x i32] %0, 0
  %b = alloca i32
  store i32 %1, i32* %b
  %2 = load [3 x i32], [3 x i32]* %a
  %3 = extractvalue [3 x i32] %2, 1
  %c = alloca i32
  store i32 %3, i32* %c
  %4 = load [3 x i32], [3 x i32]* %a
  %5 = extractvalue [3 x i32] %4, 0
  %6 = add i32 1, %5
  %d = alloca i32
  store i32 %6, i32* %d
`

	hasTestTemplate(t, src, expected)
}
