package ast

import (
	"testing"
)

func assertEq(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("expected: `%v`, actual: `%v`", expected, actual)
	}
}
