package lexer_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/lexer"

	"github.com/stretchr/testify/assert"
)

func TestLexingUnit(t *testing.T) {
	testCases := map[string]lexer.ItemType{
		"=":  lexer.ItemAssign,
		":":  lexer.ItemColon,
		"::": lexer.ItemAccessor,
		".":  lexer.ItemDot,
		",":  lexer.ItemComma,
		"(":  lexer.ItemLeftParen,
		")":  lexer.ItemRightParen,
		"[":  lexer.ItemLeftBracket,
		"]":  lexer.ItemRightBracket,
		">":  lexer.ItemGreaterThan,
		"+":  lexer.ItemPlus,
		"-":  lexer.ItemMinus,
		"*":  lexer.ItemMul,
		"/":  lexer.ItemDiv,
		"'":  lexer.ItemPrime,
	}

	for code, itemType := range testCases {
		lex := lexer.Lex("test", code)
		item := lex.NextItem()
		assert.Equal(t, itemType.String(), item.Type.String())
		assert.Equal(t, code, item.Val)
	}
}

func TestLexer(t *testing.T) {
	testCases := []struct {
		input string
		val   string
		item  lexer.ItemType
	}{
		{
			input: "name",
			item:  lexer.ItemIdent,
		},
		{
			input: `"string"`,
			item:  lexer.ItemString,
		},
		{
			input: "1",
			item:  lexer.ItemNumber,
		},
		{
			input: "3.14",
			item:  lexer.ItemNumber,
		},
		{
			input: "0x0",
			item:  lexer.ItemNumber,
		},
		{
			input: "1e7",
			item:  lexer.ItemNumber,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			lex := lexer.Lex("test", testCase.input)
			item := lex.NextItem()
			assert.Equal(t, testCase.item.String(), item.Type.String())
			assert.Equal(t, testCase.input, item.Val)
		})
	}
}

func TestLexerPosition(t *testing.T) {
	testCases := []struct {
		input    string
		position lexer.Pos
	}{
		{` 1`, lexer.Pos{Line: 1, Pos: 1}},
		{`  1`, lexer.Pos{Line: 1, Pos: 2}},
		{`
1`, lexer.Pos{Line: 2, Pos: 1}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			lex := lexer.Lex("test", testCase.input)
			assert.Equal(t, testCase.position, lex.NextItem().Pos)
		})
	}
}
