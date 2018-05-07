package listener

import (
	"testing"

	"strings"
)

func TestFnCall(t *testing.T) {
	r := NewParse(`
	fn add(l, r: i32) -> i32 { return l + r }
	fn main() {
		let a = add(1, 2)
	}
	`)

	expected := `
  %0 = call i32 @add(i32 1, i32 2)
  %a = alloca i32
  store i32 %0, i32* %a`

	if !strings.Contains(r, expected) {
		t.Errorf("expected: %s, actual: %s", expected, r)
	}
}

func TestFnCallAsStatement(t *testing.T) {
	r := NewParse(`
	fn foo() {
		// do nothing
	}
	fn main() {
		foo()
	}
	`)

	expected := `call`

	if !strings.Contains(r, expected) {
		t.Errorf("expected: %s, actual: %s", expected, r)
	}
}

func TestCallTheFuncNoParam(t *testing.T) {
	r := NewParse(`
	fn foo() -> i32 { return 10 }
	fn main() {
		let a = foo()
	}
	`)

	expected := `call i32 @foo()`

	if !strings.Contains(r, expected) {
		t.Errorf("expected: %s, actual: %s", expected, r)
	}
}
