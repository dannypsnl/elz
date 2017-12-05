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
	ItemKwImport // import module

	// Here are operator
	ItemAssign     // =
	ItemColon      // :
	ItemLeftParen  // (
	ItemRightParen // )
	ItemPlus       // +
	ItemMinus      // -
	ItemMul        // *
	ItemDiv        // /
	ItemDot        // .

	ItemEOF
	ItemForInit // This ItemType is for parser initial it's current token
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
	defer close(lex.items)
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
	st := l.start
	value := l.input[l.start:l.pos]
	switch value {
	case "let":
		l.items <- Item{ItemKwLet, st, value}
	case "mut":
		l.items <- Item{ItemKwMut, st, value}
	case "match":
		l.items <- Item{ItemKwMatch, st, value}
	case "loop":
		l.items <- Item{ItemKwLoop, st, value}
	case "break":
		l.items <- Item{ItemKwBreak, st, value}
	case "return":
		l.items <- Item{ItemKwReturn, st, value}
	case "fn":
		l.items <- Item{ItemKwFn, st, value}
	case "type":
		l.items <- Item{ItemKwType, st, value}
	case "impl":
		l.items <- Item{ItemKwImpl, st, value}
	case "trait":
		l.items <- Item{ItemKwTrait, st, value}
	case "act":
		l.items <- Item{ItemKwAct, st, value}
	case "import":
		l.items <- Item{ItemKwImport, st, value}
	default:
		l.items <- Item{t, st, value}
	}
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
	case r == '"':
		return lexString
	case r == '*':
		l.emit(ItemMul)
		return lexWhiteSpace
	case r == '/':
		l.emit(ItemDiv)
		return lexWhiteSpace
	case r == ':':
		l.emit(ItemColon)
		return lexWhiteSpace
	case r == '=':
		l.emit(ItemAssign)
		return lexWhiteSpace
	case r == '-':
		l.emit(ItemMinus)
		return lexWhiteSpace
	case r == '+':
		l.emit(ItemPlus)
		return lexWhiteSpace
	case ('0' <= r && r <= '9'):
		l.backup()
		return lexNumber
	case isAlphaNumeric(r):
		return lexIdent
	default:
		panic(fmt.Sprintf("don't know what to do with: %q", r))
	}
}

func lexString(l *Lexer) stateFn {
	for r := l.next(); r != '"'; r = l.next() {
		if r == '\\' {
			r = l.next()
		}
		if r == EOF {
			return l.errorf("unterminated quoted string")
		}
	}
	l.emit(ItemString)
	return lexWhiteSpace
}

func lexNumber(l *Lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
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
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
