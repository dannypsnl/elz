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
	// Keywords
	ItemKwTrue   // true
	ItemKwFalse  // false
	ItemKwType   // type
	ItemKwImport // import module
	// Operator
	ItemAssign     // =
	ItemColon      // :
	ItemLeftParen  // (
	ItemRightParen // )
	ItemPlus       // +
	ItemMinus      // -
	ItemMul        // *
	ItemDiv        // /
	// Meta
	ItemEOF
)

var itemTypeToString map[ItemType]string

func init() {
	itemTypeToString = make(map[ItemType]string)
	itemTypeToString[ItemError] = "error"
	itemTypeToString[ItemIdent] = "identifier"
	itemTypeToString[ItemNumber] = "number"
	itemTypeToString[ItemString] = "string"
	itemTypeToString[ItemKwType] = "keyword:type"
	itemTypeToString[ItemKwImport] = "keyword:import"
	itemTypeToString[ItemAssign] = "operator:assign"
	itemTypeToString[ItemColon] = "operator:colon"
	itemTypeToString[ItemLeftParen] = "operator:left_paren"
	itemTypeToString[ItemRightParen] = "operator:right_paren"
	itemTypeToString[ItemPlus] = "operator:plus"
	itemTypeToString[ItemMinus] = "operator:minus"
	itemTypeToString[ItemMul] = "operator:mul"
	itemTypeToString[ItemDiv] = "operator:div"
	itemTypeToString[ItemEOF] = "EOF"
}

func (t ItemType) String() string {
	s, ok := itemTypeToString[t]
	if !ok {
		return "unknown item type"
	}
	return s
}

type Item struct {
	Type ItemType
	Pos  Pos
	Val  string
}

func (i Item) String() string {
	return fmt.Sprintf(
		"type: %s, value: %s, pos: %d",
		i.Type, i.Val, i.Pos,
	)
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

func (l *Lexer) run() {
	for l.state = lexWhiteSpace; l.state != nil; {
		l.state = l.state(l)
	}
	defer close(l.items)
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
	case "true":
		l.items <- Item{ItemKwTrue, st, value}
	case "false":
		l.items <- Item{ItemKwFalse, st, value}
	case "type":
		l.items <- Item{ItemKwType, st, value}
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
	for r := l.next(); isSpace(r) || isEndOfLine(r); l.next() {
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
	case '0' <= r && r <= '9':
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
