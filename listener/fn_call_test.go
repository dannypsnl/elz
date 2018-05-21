package listener

import (
	"testing"
)

func TestFnCall(t *testing.T) {
	src := `
	fn add(l, r: i32) -> i32 { return l + r }
	fn main() {
		let a = add(1, 2)
	}
	`

	expected := `
  %0 = call i32 @add(i32 1, i32 2)
  %a = alloca i32
  store i32 %0, i32* %a`

	hasTestTemplate(t, src, expected)
}

func TestFnCallAsStatement(t *testing.T) {
	src := `
	fn foo() {
		// do nothing
	}
	fn main() {
		foo()
	}
	`

	expected := `call`

	hasTestTemplate(t, src, expected)
}

func TestCallTheFuncNoParam(t *testing.T) {
	src := `
	fn foo() -> i32 { return 10 }
	fn main() {
		let a = foo()
	}
	`

	expected := `call i32 @foo()`

	hasTestTemplate(t, src, expected)
}
