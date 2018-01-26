package listener

import (
	_ "strings"
	"testing"
)

func TestOperatorExpression(t *testing.T) {
	res := NewParse(`
	let a = 3 * 2
	let mut b = "abcde"
	`)
	expected := `; ModuleID = 'main'
source_filename = "main"

@a = global float 6.000000e+00
@b = global [7 x i8] c"\22abcde\22"
`
	//if !strings.Contains(res, expected) {
	if res != expected {
		t.Errorf("expected: `%s`\nactual: `%s`", expected, res)
	}

}
