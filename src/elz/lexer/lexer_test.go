package lexer_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/lexer"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		input string
		item  lexer.ItemType
	}{
		{input: "+", item: lexer.ItemPlus},
		{input: "name", item: lexer.ItemIdent},
		{input: `"string"`, item: lexer.ItemString},
		{input: "1", item: lexer.ItemNumber},
		{input: "3.14", item: lexer.ItemNumber},
		{input: "0x0", item: lexer.ItemNumber},
		{input: "1e7", item: lexer.ItemNumber},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			lex := lexer.Lex("test", testCase.input)
			assert.Equal(t, testCase.item, lex.NextItem().Type)
		})
	}
}
