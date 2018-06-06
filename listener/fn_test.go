package listener

import (
	"testing"
)

func TestConstComputingInFunction(t *testing.T) {
	src := `
	fn foo() {
		let a: i32 = 1 + 3
	}
	`

	expected := `
define void @foo() {
entry:
  %a = alloca i32
  store i32 4, i32* %a
  ret void
}
`

	hasTestTemplate(t, src, expected)
}

func TestMissTypeComplete(t *testing.T) {
	src := `
	fn foo(a, b, c: i32) {}
	`

	expected := `
define void @foo(i32 %a, i32 %b, i32 %c) {
entry:
  ret void
}
`

	hasTestTemplate(t, src, expected)
}

func TestOverloadingFunction(t *testing.T) {
	src := `
	fn add(l: f32, r: f32) -> f32 { return l + r }
	fn add(l: i32, r: i32) -> i32 { return l + r }
	`

	expected := `
define float @add(float %l, float %r) {
entry:
  %0 = fadd float %l, %r
  ret float %0
}

define i32 @add.1(i32 %l, i32 %r) {
entry:
  %0 = add i32 %l, %r
  ret i32 %0
}
`

	hasTestTemplate(t, src, expected)
}

func TestMainFunction(t *testing.T) {
	src := `
	fn main() {}
	`

	expected := `
define i32 @main() {
entry:
  ret i32 0
}
`

	hasTestTemplate(t, src, expected)
}

func TestBasicComputingFunction(t *testing.T) {
	src := `
	fn add(lv, rv: i32) -> i32 { return lv + rv }
	`

	expected := `
define i32 @add(i32 %lv, i32 %rv) {
entry:
  %0 = add i32 %lv, %rv
  ret i32 %0
}
`

	hasTestTemplate(t, src, expected)
}

func TestAccessGlobalVarInFunction(t *testing.T) {
	src := `
	one = 1
	fn add_one(v: i32) -> i32 { return v + one }
	`

	expected := `
@one = global i32 1

define i32 @add_one(i32 %v) {
entry:
  %0 = load i32, i32* @one
  %1 = add i32 %v, %0
  ret i32 %1
}
`

	hasTestTemplate(t, src, expected)
}
