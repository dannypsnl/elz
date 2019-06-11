package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type ItemType int
type Offset int
type Pos struct {
	Line, Pos Offset
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.Line, p.Pos)
}

const EOF = -1

const (
	ItemError ItemType = iota

	ItemIdent  // ex. name
	ItemNumber // ex. 1.24 0.12 .23
	ItemString // ex. "I am a string"
	// Keywords
	ItemKwTrue   // true
	ItemKwFalse  // false
	ItemKwCase   // case
	ItemKwOf     // of
	ItemKwElse   // else
	ItemKwType   // type
	ItemKwImport // import module
	// Operator
	ItemAssign       // =
	ItemColon        // :
	ItemAccessor     // ::
	ItemDot          // .
	ItemComma        // ,
	ItemLeftParen    // (
	ItemRightParen   // )
	ItemLeftBracket  // [
	ItemRightBracket // ]
	ItemGreaterThan  // >
	ItemPlus         // +
	ItemMinus        // -
	ItemMul          // *
	ItemDiv          // /
	ItemPrime        /// '
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
	itemTypeToString[ItemKwTrue] = "keyword:true"
	itemTypeToString[ItemKwFalse] = "keyword:false"
	itemTypeToString[ItemKwCase] = "keyword:case"
	itemTypeToString[ItemKwOf] = "keyword:of"
	itemTypeToString[ItemKwElse] = "keyword:else"
	itemTypeToString[ItemKwType] = "keyword:type"
	itemTypeToString[ItemKwImport] = "keyword:import"
	itemTypeToString[ItemAssign] = "operator:assign"
	itemTypeToString[ItemColon] = "operator:colon"
	itemTypeToString[ItemAccessor] = "operator:accessor"
	itemTypeToString[ItemDot] = "operator:dot"
	itemTypeToString[ItemComma] = "operator:comma"
	itemTypeToString[ItemLeftParen] = "operator:left_paren"
	itemTypeToString[ItemRightParen] = "operator:right_paren"
	itemTypeToString[ItemLeftBracket] = "operator:left_bracket"
	itemTypeToString[ItemRightBracket] = "operator:right_bracket"
	itemTypeToString[ItemGreaterThan] = "operator:greater_than"
	itemTypeToString[ItemPlus] = "operator:plus"
	itemTypeToString[ItemMinus] = "operator:minus"
	itemTypeToString[ItemMul] = "operator:mul"
	itemTypeToString[ItemDiv] = "operator:div"
	itemTypeToString[ItemPrime] = "operator:prime"
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
	Val  string
	Pos
}

func (i Item) String() string {
	return fmt.Sprintf(
		"(%s, value: `%s` at %s)",
		i.Type, i.Val, i.Pos,
	)
}

type Lexer struct {
	name  string
	input []rune
	state stateFn
	// for report position as format: (line, pos)
	line Offset
	pos  Offset
	// lexing helpers
	offset Offset
	start  Offset
	items  chan Item
}

func Lex(name, code string) *Lexer {
	input := []rune(code)
	line := Offset(1)
	if isEndOfLine(input[0]) {
		line++
	}
	l := &Lexer{
		name:   name,
		input:  input,
		items:  make(chan Item),
		state:  lexWhiteSpace,
		line:   line,
		pos:    0,
		offset: 0,
		start:  0,
	}
	go l.run()
	return l
}

func (l *Lexer) NextItem() Item {
	item := <-l.items
	return item
}
func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- l.newItem(ItemError, fmt.Sprintf(format, args...))
	return nil
}

func (l *Lexer) run() {
	for l.state != nil {
		l.state = l.state(l)
	}
	defer close(l.items)
}

// next returns the next rune in the input.
func (l *Lexer) next() rune {
	l.offset += 1
	if int(l.offset) >= len(l.input) {
		return EOF
	}
	r := l.input[l.offset]
	l.pos++
	if isEndOfLine(r) {
		// newline -> pos: (line+1, 0)
		l.line++
		l.pos = 0
	}
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *Lexer) peek() rune {
	if int(l.offset) >= len(l.input) {
		return EOF
	}
	return l.input[l.offset]
}

func (l *Lexer) newItem(typ ItemType, val string) Item {
	return Item{
		Type: typ,
		Val:  val,
		Pos:  Pos{l.line, l.pos},
	}
}

// emit passes an Item back to the client.
func (l *Lexer) emit(t ItemType) {
	value := string(l.input[l.start:l.offset])
	switch value {
	case "true":
		l.items <- l.newItem(ItemKwTrue, value)
	case "false":
		l.items <- l.newItem(ItemKwFalse, value)
	case "case":
		l.items <- l.newItem(ItemKwCase, value)
	case "of":
		l.items <- l.newItem(ItemKwOf, value)
	case "else":
		l.items <- l.newItem(ItemKwElse, value)
	case "type":
		l.items <- l.newItem(ItemKwType, value)
	case "import":
		l.items <- l.newItem(ItemKwImport, value)
	default:
		l.items <- l.newItem(t, value)
	}
	l.start = l.offset
}

// accept consumes the next rune if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.peek()) >= 0 {
		l.next()
		return true
	}
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.peek()) >= 0 {
		l.next()
	}
}

// ignore drop not yet emitted parts
func (l *Lexer) ignore() {
	l.start = l.offset
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
	// Next thing mustn't be alphanumeric.
	if r := l.peek(); isAlphaNumeric(r) {
		l.next()
		return false
	}
	return true
}

// stateFn is a function need get info from Lexer and return next stateFn
type stateFn func(*Lexer) stateFn

func lexWhiteSpace(l *Lexer) stateFn {
	for r := l.peek(); isSpace(r) || isEndOfLine(r); {
		r = l.next()
	}
	l.ignore()

	switch r := l.peek(); {
	case r == EOF:
		l.emit(ItemEOF)
		return nil
	case r == '"':
		l.next()
		return lexString
	case r == '*':
		l.next()
		l.emit(ItemMul)
		return lexWhiteSpace
	case r == '/':
		l.next()
		if l.peek() == '/' {
			return lexComment
		}
		l.emit(ItemDiv)
		return lexWhiteSpace
	case r == ':':
		l.next()
		if l.peek() == ':' {
			l.next()
			l.emit(ItemAccessor)
		} else {
			l.emit(ItemColon)
		}
		return lexWhiteSpace
	case r == '.':
		l.next()
		l.emit(ItemDot)
		return lexWhiteSpace
	case r == ',':
		l.next()
		l.emit(ItemComma)
		return lexWhiteSpace
	case r == '=':
		l.next()
		l.emit(ItemAssign)
		return lexWhiteSpace
	case r == '-':
		l.next()
		l.emit(ItemMinus)
		return lexWhiteSpace
	case r == '>':
		l.next()
		l.emit(ItemGreaterThan)
		return lexWhiteSpace
	case r == '+':
		l.next()
		l.emit(ItemPlus)
		return lexWhiteSpace
	case r == '(':
		l.next()
		l.emit(ItemLeftParen)
		return lexWhiteSpace
	case r == ')':
		l.next()
		l.emit(ItemRightParen)
		return lexWhiteSpace
	case r == '[':
		l.next()
		l.emit(ItemLeftBracket)
		return lexWhiteSpace
	case r == ']':
		l.next()
		l.emit(ItemRightBracket)
		return lexWhiteSpace
	case r == '\'':
		l.next()
		l.emit(ItemPrime)
		return lexWhiteSpace
	case '0' <= r && r <= '9':
		return lexNumber
	case isAlphaNumeric(r):
		return lexIdent
	default:
		panic(fmt.Sprintf("don't know what to do with: %q", r))
	}
}

func lexComment(l *Lexer) stateFn {
	for r := l.next(); !isEndOfLine(r); r = l.next() {
	}
	return lexWhiteSpace
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
	l.next()
	l.emit(ItemString)
	return lexWhiteSpace
}

func lexNumber(l *Lexer) stateFn {
	if !l.scanNumber() {
		return l.errorf("bad number syntax: %q", l.input[l.start:l.offset])
	}

	l.emit(ItemNumber)

	return lexWhiteSpace
}

func lexIdent(l *Lexer) stateFn {
	for r := l.peek(); isAlphaNumeric(r); {
		r = l.next()
	}

	l.emit(ItemIdent)
	return lexWhiteSpace
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is a valid rune for an identifier.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}
