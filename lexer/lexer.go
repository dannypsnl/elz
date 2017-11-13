package lexer

import (
	"fmt"
	_ "strings"
	_ "unicode"
	_ "unicode/utf8"
)

type ItemType int
type Pos int

const EOF = -1

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
	ItemLeftParen  // (
	ItemRightParen // )
	ItemPlus       // +
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

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return EOF
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *Lexer) backup() {
	l.pos -= l.width
}

// emit passes an Item back to the client.
func (l *Lexer) emit(t ItemType) {
	// Item {ItemType, Pos, Val}
	l.items <- Item{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

// ignore drop not yet emitted parts
func (l *Lexer) ignore() {
	l.start = l.pos
}

// stateFn is a function need get info from Lexer and return next stateFn
type stateFn func(*Lexer) stateFn

func lexWhiteSpace(l *Lexer) stateFn {
	for {
		if l.next() == EOF {
			break
		}
	}
	if l.pos > l.start {
		l.emit(ItemError)
	}
	l.emit(ItemEOF)
	return nil
}
