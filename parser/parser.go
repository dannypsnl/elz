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

func (p *parser) Next() { p.cur_token = p.lex.NextItem() }

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
		case lexer.ItemKwFn:
		case lexer.ItemKwType:
		case lexer.ItemKwTrait:
		case lexer.ItemKwImport:
		default:
			p.tree = append(p.tree, Error{fmt.Sprintf("Pos %d, token '%s' is not allow at top level\n", p.cur_token.Pos, p.cur_token.Val)})
		}
	}
	return p.tree
}
