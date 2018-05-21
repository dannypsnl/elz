package listener

import (
	"strings"
	"testing"
)

func hasTestTemplate(t *testing.T, source, expectedIn string) {
	res := NewParse(source)

	if !strings.Contains(res, expectedIn) {
		t.Errorf("expected has: `%s`, actual: `%s`", expectedIn, res)
	}
}
