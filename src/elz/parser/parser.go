package parser

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/lexer"
)

type Parser struct {
	lex *lexer.Lexer
	// token system
	curToken  lexer.Item
	peekToken lexer.Item
}

func NewParser(name, code string) *Parser {
	parser := &Parser{
		lex: lexer.Lex(name, code),
	}
	// init the peek token
	parser.next()
	// init the current token
	parser.next()

	return parser
}

func (p *Parser) next() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextItem()
}

func (p *Parser) want(wantType lexer.ItemType) error {
	if p.curToken.Type != wantType {
		return expectedError(wantType, p.curToken.Type, p.curToken.Pos)
	}
	return nil
}

func (p *Parser) ParseBinding() (*ast.Binding, error) {
	if err := p.want(lexer.ItemIdent); err != nil {
		return nil, err
	}
	bindingName := p.curToken.Val
	parameterList := make([]string, 0)
	// consume binding name then start parameters parsing
	for p.next(); p.curToken.Type == lexer.ItemIdent; p.next() {
		parameterList = append(parameterList, p.curToken.Val)
	}
	if err := p.want(lexer.ItemAssign); err != nil {
		return nil, err
	}
	// consume assign
	p.next()
	expr := p.ParseExpression(p.ParsePrimary(), 0)
	return &ast.Binding{
		Export:    false,
		Name:      bindingName,
		ParamList: parameterList,
		Expr:      expr,
	}, nil
}

func (p *Parser) ParseExpression(leftHandSide ast.Expr, previousPrimary int) ast.Expr {
	lhs := leftHandSide
	lookahead := p.peekToken
	// precedence of lookahead would also break with not operator token
	// since the precedence would be -1
	for precedence(lookahead) >= previousPrimary {
		operator := lookahead
		p.next() // consume end of lhs
		p.next() // consume operator
		rhs := p.ParsePrimary()
		// consume end of rhs
		// update lookahead, else the loop would never end
		// and bring unexpected result
		lookahead = p.peekToken
		for precedence(lookahead) > precedence(operator) ||
			(isRightAssociative(lookahead) && precedence(lookahead) == precedence(operator)) {
			rhs = p.ParseExpression(rhs, precedence(lookahead))
			lookahead = p.peekToken
		}
		lhs = &ast.BinaryExpr{
			LExpr: lhs,
			RExpr: rhs,
			Op:    operator.Val,
		}
	}
	return lhs
}

func (p *Parser) ParsePrimary() ast.Expr {
	switch p.curToken.Type {
	case lexer.ItemNumber:
		if strings.ContainsRune(p.curToken.Val, '.') {
			return ast.NewFloat(p.curToken.Val)
		} else {
			return ast.NewInt(p.curToken.Val)
		}
	case lexer.ItemIdent:
		return ast.NewIdent(p.curToken.Val)
	default:
		logrus.Fatalf("unsupported primary token: %s", p.curToken)
	}
	return nil
}

func precedence(token lexer.Item) int {
	precedence, ok := precedenceOfOperator[token.Type]
	if !ok {
		return -1
	}
	return precedence
}

func isRightAssociative(token lexer.Item) bool {
	// no right associative operator now
	return false
}

var precedenceOfOperator = map[lexer.ItemType]int{}

func init() {
	precedenceOfOperator[lexer.ItemPlus] = 1
	precedenceOfOperator[lexer.ItemMinus] = 1
	precedenceOfOperator[lexer.ItemMul] = 2
	precedenceOfOperator[lexer.ItemDiv] = 2
}

func expectedError(expected, actual fmt.Stringer, pos lexer.Pos) error {
	return fmt.Errorf("expected: %s but got: %s at: %d", expected, actual, pos)
}
