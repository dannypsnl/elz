package listener

import (
	"testing"
)

func TestAccessArrayElem(t *testing.T) {
	src := `
	fn main() {
		let a = [1, 2, 3]
		let b = a[0]
		let c = a[1]
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
`

	hasTestTemplate(t, src, expected)
}
