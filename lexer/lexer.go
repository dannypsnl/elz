package lexer

import (
	"fmt"
	_ "strings"
	_ "unicode"
	_ "unicode/utf8"
)

type ItemType int
type Pos int

const (
	ItemError ItemType = iota

	ItemIdentifier // ex. name
	ItemNumber     // ex. 1.24 0.12 .23
	ItemString     // ex. "I am a string"

	// Here are keywords
	ItemKwMatch  // match xxx { ... }
	ItemKwLoop   // loop { statement }
	ItemKwBreak  // break
	ItemKwReturn // return expression
	ItemKwLet    // let a = 1
	ItemKwMut    // let mut a = 1
	ItemKwFn     // fn add(lv: num, rv: num) -> num {return lv+rv}
	ItemKwType   // type
	ItemKwImpl   // impl type { methods }
	ItemKwTrait  // trait xxx { methodTypes }
	ItemKwAct    // act functionCall

	// Here are operator
	ItemLeftParam  // (
	ItemRightParam // )
	ItemAdd        // +
	ItemMinus      // -
	ItemMul        // *
	ItemDiv        // /
	ItemDot        // .

	ItemEOF
)

type Item struct {
	Type ItemType
	Pos  Pos
	Val  string
}

func (i Item) String() string {
	switch i.Type {
	case ItemEOF:
		return "EOF"
	case ItemError:
		return i.Val
	}
	if len(i.Val) > 10 {
		return fmt.Sprintf("%.10q...", i.Val)
	}
	return fmt.Sprintf("%q", i.Val)
}

type Lexer struct {
	name  string
	input string
	start int
	pos   int
	width int
	items chan Item
}

func Lex(name, input string) *Lexer {
	l := &Lexer{
		name:  name,
		input: input,
		items: make(chan Item),
	}
	go l.run()
	return l
}

func (lex *Lexer) run() {
	for state := lexWhiteSpace; state != nil; {
		state = state(lex)
	}
	close(lex.items)
}

type stateFn func(*Lexer) stateFn

func lexWhiteSpace(l *Lexer) stateFn {}
