package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type ItemType int
type Pos int

const EOF = -1

const (
	ItemError ItemType = iota

	ItemIdent  // ex. name
	ItemNumber // ex. 1.24 0.12 .23
	ItemString // ex. "I am a string"

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
	name    string
	input   string
	state   stateFn
	pos     Pos
	start   Pos
	width   Pos
	lastPos Pos
	items   chan Item

	parenDepth int
	vectDepth  int
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

func (l *Lexer) NextItem() Item {
	item := <-l.items
	l.lastPos = item.Pos
	return item
}
func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- Item{ItemError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

func (lex *Lexer) run() {
	for lex.state = lexWhiteSpace; lex.state != nil; {
		lex.state = lex.state(lex)
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

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// ignore drop not yet emitted parts
func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) scanNumber() bool {
	// Optional leading sign.
	l.accept("+-")
	// Is it hex?
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}
	if l.accept("eE") {
		l.accept("+-")
		l.acceptRun("0123456789")
	}
	// Is it imaginary?
	l.accept("i")
	// Next thing mustn't be alphanumeric.
	if r := l.peek(); isAlphaNumeric(r) {
		l.next()
		return false
	}
	return true
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// stateFn is a function need get info from Lexer and return next stateFn
type stateFn func(*Lexer) stateFn

func lexWhiteSpace(l *Lexer) stateFn {
	for r := l.next(); isSpace(r) || r == '\n'; l.next() {
		r = l.peek()
	}
	l.backup()
	l.ignore()

	switch r := l.next(); {
	case r == EOF:
		l.emit(ItemEOF)
		return nil
	case ('0' <= r && r <= '9') || r == '+' || r == '-':
		l.backup()
		return lexNumber
	case isAlphaNumeric(r):
		return lexIdent
	default:
		panic(fmt.Sprintf("don't know what to do with: %q", r))
	}
}

func lexNumber(l *Lexer) stateFn {
	r := l.peek()
	if r == '+' || r == '-' {
		l.next()
		r2 := l.peek()
		if isSpace(r2) || r2 == '\n' || r2 == EOF {
			if r == '+' {
				l.emit(ItemPlus)
				return lexWhiteSpace
			} else if r == '-' {
				l.emit(ItemMinus)
				return lexWhiteSpace
			}
		}
	}
	l.backup()
	r = l.peek()

	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
	}

	if l.start+1 == l.pos {
		return lexIdent
	}

	l.emit(ItemNumber)

	return lexWhiteSpace
}

func lexIdent(l *Lexer) stateFn {
	for r := l.next(); isAlphaNumeric(r); r = l.next() {
	}
	l.backup()

	l.emit(ItemIdent)
	return lexWhiteSpace
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is a valid rune for an identifier.
func isAlphaNumeric(r rune) bool {
	return r == '>' || r == '<' || r == '=' || r == '-' || r == '+' || r == '*' || r == '&' || r == '_' || r == '/' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
