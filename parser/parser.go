package parser

import (
	"fmt"
	"github.com/elz-lang/elz/lexer"
)

type Ast interface{}

type Error struct {
	msg string
}

type parser struct {
	lex       *lexer.Lexer
	cur_token lexer.Item
	tree      []Ast
}

func (p *parser) Next() lexer.Item { p.cur_token = p.lex.NextItem(); return p.cur_token }

func Parse(filename, source_code string) *parser {
	return &parser{
		lex:       lexer.Lex(filename, source_code),
		cur_token: lexer.Item{lexer.ItemForInit, 0, ""},
	}
}

func (p *parser) parseProgram() []Ast {
	for p.Next(); p.cur_token.Type != lexer.ItemEOF; p.Next() {
		switch p.cur_token.Type {
		case lexer.ItemKwLet:
			p.parserVarDefination()
		case lexer.ItemKwFn:
		case lexer.ItemKwType:
		case lexer.ItemKwTrait:
		case lexer.ItemKwImport:
		default:
			p.tree = append(p.tree, Error{fmt.Sprintf("At(%d), token '%s' is not allow at top level\n", p.cur_token.Pos, p.cur_token.Val)})
		}
	}
	return p.tree
}

type VarDefination struct {
	immutable  bool
	export     bool
	name       string
	varType    string
	expression lexer.Item
}

func (p *parser) parserVarDefination() {
	immutable := true
	export := false
	Type := ""
	tk := p.Next()
	// Because we allow let mut `ident`, let `ident`, but others are error
	if tk.Type == lexer.ItemPlus {
		tk = p.Next()
		export = true
	}
	if tk.Type != lexer.ItemKwMut && tk.Type != lexer.ItemIdent {
		p.tree = append(p.tree, Error{fmt.Sprintf("At(%d), Expected a keyword[mut] or a identifier, but is '%s'\n", p.cur_token.Pos, p.cur_token.Val)})
	} else if tk.Type == lexer.ItemKwMut {
		immutable = false
		tk = p.Next()
	}
	name := tk.Val // identifier's value
	if res := p.parseType(); res {
		Type = p.cur_token.Val
	}
	p.Next()
	if p.cur_token.Type == lexer.ItemAssign {
		p.tree = append(p.tree, &VarDefination{
			immutable:  immutable,
			export:     export,
			name:       name,
			varType:    Type,
			expression: p.Next(),
		})
	} else {
		p.tree = append(p.tree, Error{fmt.Sprintf("At(%d), Expected a assign symbol, but is '%s'\n", p.cur_token.Pos, p.cur_token.Val)})
	}
}

func (p *parser) parseType() bool {
	p.Next()
	if p.cur_token.Val != ":" {
		return false
	} else {
		p.Next()
		if p.cur_token.Type == lexer.ItemIdent {
			return true
		}
		p.tree = append(p.tree, Error{fmt.Sprintf("At(%d), Expected a identifier, but is '%s'\n", p.cur_token.Pos, p.cur_token.Val)})
		return false
	}
}
