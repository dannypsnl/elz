package parser

import (
	"github.com/elz-lang/elz/src/elz/lexer"
)

type Precedence uint8

const (
	Error Precedence = iota
	Minimum
	Plus
	Minus            = Plus
	Mul   Precedence = iota + Minus
	Div              = Mul
	Dot   Precedence = iota + Div
)

func precedence(token lexer.Item) Precedence {
	precedence, ok := precedenceOfOperator[token.Type]
	if !ok {
		return Error
	}
	return precedence
}

func isRightAssociative(token lexer.Item) bool {
	// no right associative operator now
	return false
}

var precedenceOfOperator = map[lexer.ItemType]Precedence{}

func init() {
	precedenceOfOperator[lexer.ItemPlus] = Plus
	precedenceOfOperator[lexer.ItemMinus] = Minus
	precedenceOfOperator[lexer.ItemMul] = Mul
	precedenceOfOperator[lexer.ItemDiv] = Div
	precedenceOfOperator[lexer.ItemDot] = Dot
}
