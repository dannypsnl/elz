package lexer

import (
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	var results []Item
	lexer := Lex("lex", "+123 name -12.3 ")
	for item := lexer.NextItem(); item.Type != ItemEOF; item = lexer.NextItem() {
		results = append(results, item)
	}
	if strings.Compare(lexer.name, "lex") != 0 {
		t.Error("Lexer name is wrong!")
	}
	expected := []Item{
		Item{ItemNumber, 0, "+123"},
		Item{ItemIdent, 0, "name"},
		Item{ItemNumber, 0, "-12.3"},
	}
	for i, result := range results {
		if result.Type != expected[i].Type {
			t.Errorf("expected: %q, actual: %q", expected[i].Type, result.Type)
		}
		if result.Val != expected[i].Val {
			t.Errorf("expected: %q, actual: %q", expected[i].Val, result.Val)
		}
	}
}
