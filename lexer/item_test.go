package lexer

import (
	"strings"
	"testing"
)

func TestItemStringFormat(t *testing.T) {
	it := Item{
		Type: ItemEOF,
	}
	if strings.Compare(it.String(), "EOF") != 0 {
		t.Error(`fail, EOF format is not correct`)
	}
	it = Item{
		Type: ItemNumber,
		Val:  "10.34234234842394",
	}
	expected := "\"10.3423423\"..."
	if strings.Compare(it.String(), expected) != 0 {
		t.Errorf(`fail, long expression format error, expected: %s, actual: %s`, expected, it.String())
	}
}
