package lexer

import (
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	var results []Item
	lexer := Lex("lex", "123")
	for item := lexer.NextItem(); item.Type != ItemEOF; item = lexer.NextItem() {
		results = append(results, item)
	}
	if strings.Compare(lexer.name, "lex") != 0 {
		t.Error("Lexer name is wrong!")
	}
	if results[0].Type != ItemInt {
		t.Error("First token should be a Int Literal!")
	}
}
