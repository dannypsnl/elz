package ast

import (
	"strings"
	"testing"
)

func assertEq(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("expected: `%v`, actual: `%v`", expected, actual)
	}
}

func testHas(t *testing.T, res, expectedIn string) {
	if !strings.Contains(res, expectedIn) {
		t.Errorf("expected has: `%s`, actual: `%s`", expectedIn, res)
	}
}
