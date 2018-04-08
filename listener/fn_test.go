package listener

import (
	"testing"
)

func TestConstComputingInFunction(t *testing.T) {
	res := NewParse(`
	fn foo() {
		let a: i32 = 1 + 3
	}
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

define void @foo() {
entry:
  %a = alloca i32
  store i32 4, i32* %a
}
`
	if res != expected {
		t.Errorf("Function Error; expected: `%s`\nactual: `%s`", expected, res)
	}
}

func TestOverloadingFunction(t *testing.T) {
	res := NewParse(`
	fn add(l: num, r: num) -> num {}
	fn add(l: i32, r: i32) -> i32 {}
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

define float @add(float %l, float %r) {
entry:
}

define i32 @add.1(i32 %l, i32 %r) {
entry:
}
`
	if res != expected {
		t.Errorf("Function Error; expected: `%s`\nactual: `%s`", expected, res)
	}
}

func TestMainFunction(t *testing.T) {
	res := NewParse(`
	fn main() {}
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

define i32 @main() {
entry:
  ret i32 0
}
`
	if res != expected {
		t.Errorf("Main Function Error; expected: `%s`\nactual: `%s`", expected, res)
	}
}
