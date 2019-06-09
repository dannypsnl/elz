package lexer_test

import (
	"testing"

	"github.com/elz-lang/elz/src/elz/lexer"

	"github.com/stretchr/testify/assert"
)

func TestLexingUnit(t *testing.T) {
	testCases := map[string]lexer.ItemType{
		// keyword
		"import": lexer.ItemKwImport,
		"type":   lexer.ItemKwType,
		"true":   lexer.ItemKwTrue,
		"false":  lexer.ItemKwFalse,
		"case":   lexer.ItemKwCase,
		"of":     lexer.ItemKwOf,
		"else":   lexer.ItemKwElse,
		// identifier
		"name": lexer.ItemIdent,
		// operators
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
		// string
		`"string"`: lexer.ItemString,
		// number
		"1":    lexer.ItemNumber,
		"3.14": lexer.ItemNumber,
		"0x0":  lexer.ItemNumber,
		"1e7":  lexer.ItemNumber,
	}

	for code, itemType := range testCases {
		lex := lexer.Lex("test", code)
		item := lex.NextItem()
		assert.Equal(t, itemType.String(), item.Type.String())
		assert.Equal(t, code, item.Val)
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
