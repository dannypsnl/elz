package parser

import (
	"fmt"
	"strconv"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/lexer"
)

type Parser struct {
	lex   *lexer.Lexer
	cache *lexer.Item
}

func NewParser(name, code string) *Parser {
	return &Parser{
		lex: lexer.Lex(name, code),
	}
}

func (p *Parser) next() lexer.Item {
	item := p.lex.NextItem()
	p.cache = &item
	return *p.cache
}

func (p *Parser) peek() lexer.Item {
	if p.cache != nil {
		return *p.cache
	}
	return p.next()
}

func (p *Parser) ParseBinding() (*ast.Binding, error) {
	ident := p.peek()
	if ident.Type != lexer.ItemIdent {
		return nil, expectedError(lexer.ItemIdent, ident.Type, ident.Pos)
	}
	parameters := make([]string, 0)
	for tok := p.next(); tok.Type == lexer.ItemIdent; tok = p.next() {
		parameters = append(parameters, tok.Val)
	}
	assign := p.peek()
	if assign.Type != lexer.ItemAssign {
		return nil, expectedError(lexer.ItemAssign, assign.Type, assign.Pos)
	}
	// consume assign operator
	p.next()
	expr, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}
	return &ast.Binding{
		Export:    false,
		Name:      ident.Val,
		ParamList: parameters,
		Expr:      expr,
	}, nil
}

func (p *Parser) ParseExpression() (ast.Expr, error) {
	tok := p.peek()
	switch tok.Type {
	case lexer.ItemNumber:
		_, err := strconv.ParseInt(tok.Val, 10, 64)
		if err != nil {
			_, err := strconv.ParseFloat(tok.Val, 64)
			if err != nil {
				return nil, err
			}
			return ast.NewFloat(tok.Val), nil
		}
		return ast.NewInt(tok.Val), nil
	case lexer.ItemIdent:
	}
	return nil, fmt.Errorf("not implemented")
}

func expectedError(expected, actual fmt.Stringer, pos lexer.Pos) error {
	return fmt.Errorf("expected: %s but got: %s\npos: %d", expected, actual, pos)
}
