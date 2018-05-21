package listener

import (
	"testing"
)

func TestAssign(t *testing.T) {
	src := `
	fn main() {
		let mut num = 1
		num = 2
	}
	`

	expected := `
  %num = alloca i32
  store i32 1, i32* %num
  store i32 2, i32* %num`

	hasTestTemplate(t, src, expected)
}
