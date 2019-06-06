package parser

import (
	"fmt"
	"strings"

	"github.com/elz-lang/elz/src/elz/ast"
	"github.com/elz-lang/elz/src/elz/lexer"
)

type Parser struct {
	lex *lexer.Lexer
	// token system
	curToken  lexer.Item
	peekToken lexer.Item
	// error reports
	errors []error
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

func (p *Parser) ParseProgram() (*ast.Program, error) {
	program := &ast.Program{}
	for {
		switch p.curToken.Type {
		case lexer.ItemIdent:
			if p.peekToken.Type == lexer.ItemColon {
				bindingType, err := p.ParseBindingType()
				if err != nil {
					p.errors = append(p.errors, err)
				} else {
					program.AddBindingType(bindingType)
					p.next()
				}
			} else {
				binding, err := p.ParseBinding()
				if err != nil {
					p.errors = append(p.errors, err)
				} else {
					program.AddBinding(binding)
				}
				p.next()
			}
		case lexer.ItemKwImport:
			importStmt, err := p.ParseImport()
			if err != nil {
				p.errors = append(p.errors, err)
			} else {
				program.AddImport(importStmt)
				p.next()
			}
		case lexer.ItemKwType:
			// FIXME: complete the type define parsing
			return nil, fmt.Errorf("unimplemented type define yet")
		case lexer.ItemEOF:
			return program, p.reportErrors()
		default:
			return nil, fmt.Errorf("unexpected token: %s", p.curToken)
		}
	}
}

func (p *Parser) reportErrors() error {
	if len(p.errors) == 0 {
		return nil
	}
	buf := strings.Builder{}
	for i, err := range p.errors {
		buf.WriteString(fmt.Sprintf("error #%d\n", i+1))
		buf.WriteString(err.Error())
		buf.WriteRune('\n')
		buf.WriteRune('\n')
	}
	return fmt.Errorf(buf.String())
}

func (p *Parser) next() {
	p.curToken = p.peekToken
	p.peekToken = p.lex.NextItem()
}

func (p *Parser) want(wantType lexer.ItemType) error {
	if p.curToken.Type != wantType {
		return expectedError(wantType, p.curToken)
	}
	return nil
}

func (p *Parser) ParseImport() (*ast.Import, error) {
	if err := p.want(lexer.ItemKwImport); err != nil {
		return nil, err
	}
	p.next() // consume keyword import
	accessChain, err := p.ParseAccessChain()
	if err != nil {
		return nil, err
	}
	return &ast.Import{
		AccessChain: accessChain,
	}, nil
}

func (p *Parser) ParseBindingType() (*ast.BindingType, error) {
	if err := p.want(lexer.ItemIdent); err != nil {
		return nil, err
	}
	bindingName := p.curToken.Val
	p.next()
	if err := p.want(lexer.ItemColon); err != nil {
		return nil, err
	}
	p.next()
	if err := p.want(lexer.ItemColon); err != nil {
		return nil, err
	}
	p.next()

	bindingType := make([]ast.Type, 0)
	for {
		t, err := p.ParseType()
		if err != nil {
			return nil, err
		}
		bindingType = append(bindingType, t)
		if p.peekToken.Type == lexer.ItemMinus {
			p.next()
			p.next()
			if err := p.want(lexer.ItemGreaterThan); err != nil {
				return nil, err
			}
			p.next()
		} else {
			return &ast.BindingType{
				Name: bindingName,
				Type: bindingType,
			}, nil
		}
	}
}

func (p *Parser) ParseType() (ast.Type, error) {
	switch p.curToken.Type {
	case lexer.ItemIdent:
		accessChain, err := p.ParseAccessChain()
		if err != nil {
			return nil, err
		}
		return &ast.ExistType{Name: accessChain.Literal}, nil
	case lexer.ItemPrime:
		p.next()
		if err := p.want(lexer.ItemIdent); err != nil {
			return nil, err
		}
		return &ast.VariantType{Name: p.curToken.Val}, nil
	case lexer.ItemLeftParen:
		p.next()
		if err := p.want(lexer.ItemRightParen); err != nil {
			return nil, err
		}
		p.next()
		return &ast.VoidType{}, nil
	default:
		return nil, fmt.Errorf("unexpected token at parsing type: %s", p.curToken)
	}
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
	primary, err := p.ParsePrimary()
	if err != nil {
		return nil, err
	}
	expr, err := p.ParseExpression(primary, 0)
	if err != nil {
		return nil, err
	}
	return &ast.Binding{
		Export:    !strings.HasPrefix(bindingName, "_"),
		Name:      bindingName,
		ParamList: parameterList,
		Expr:      expr,
	}, nil
}

func (p *Parser) ParseExpression(leftHandSide ast.Expr, previousPrimary int) (ast.Expr, error) {
	lhs := leftHandSide
	lookahead := p.peekToken
	// precedence of lookahead would also break with not operator token
	// since the precedence would be -1
	for precedence(lookahead) >= previousPrimary {
		operator := lookahead
		p.next() // consume end of lhs
		p.next() // consume operator
		rhs, err := p.ParsePrimary()
		if err != nil {
			return nil, err
		}
		// update lookahead, else the loop would never end
		// and bring unexpected result
		lookahead = p.peekToken
		for precedence(lookahead) > precedence(operator) ||
			(isRightAssociative(lookahead) && precedence(lookahead) == precedence(operator)) {
			rhs, err = p.ParseExpression(rhs, precedence(lookahead))
			if err != nil {
				return nil, err
			}
			lookahead = p.peekToken
		}
		lhs = &ast.BinaryExpr{
			LExpr: lhs,
			RExpr: rhs,
			Op:    operator.Val,
		}
	}
	return lhs, nil
}

func (p *Parser) ParsePrimary() (ast.Expr, error) {
	expr, err := p.ParseUnary()
	if err != nil {
		return nil, err
	}
	switch p.peekToken.Type {
	case lexer.ItemLeftParen:
		p.next()
		return p.ParseArgument(expr)
	case lexer.ItemLeftBracket:
		p.next()
		return p.ParseElementAccess(expr)
	default:
		return expr, nil
	}
}

func (p *Parser) ParseUnary() (ast.Expr, error) {
	switch p.curToken.Type {
	case lexer.ItemNumber:
		if strings.ContainsRune(p.curToken.Val, '.') {
			return ast.NewFloat(p.curToken.Val), nil
		} else {
			return ast.NewInt(p.curToken.Val), nil
		}
	case lexer.ItemIdent:
		return p.ParseAccessChain()
	case lexer.ItemString:
		return ast.NewString(p.curToken.Val), nil
	case lexer.ItemKwTrue, lexer.ItemKwFalse:
		return ast.NewBool(p.curToken.Val), nil
	case lexer.ItemLeftParen:
		p.next() // consume left paren (
		primary, err := p.ParsePrimary()
		if err != nil {
			return nil, err
		}
		expr, err := p.ParseExpression(primary, 0)
		if err != nil {
			return nil, err
		}
		p.next() // consume end of expression
		if err := p.want(lexer.ItemRightParen); err != nil {
			return nil, err
		}
		return expr, nil
	case lexer.ItemLeftBracket:
		listElem := make([]ast.Expr, 0)
		for p.peekToken.Type != lexer.ItemRightBracket {
			p.next()
			primary, err := p.ParsePrimary()
			if err != nil {
				return nil, err
			}
			expr, err := p.ParseExpression(primary, 0)
			if err != nil {
				return nil, err
			}
			p.next() // consume end of expression
			listElem = append(listElem, expr)
			if err := p.want(lexer.ItemComma); err != nil {
				if err := p.want(lexer.ItemRightBracket); err != nil {
					return nil, err
				} else {
					break
				}
			}
		}
		return ast.NewList(listElem...), nil
	default:
		return nil, fmt.Errorf("unsupported primary token: %s", p.curToken)
	}
}

func (p *Parser) ParseElementAccess(expr ast.Expr) (*ast.ExtractElement, error) {
	if p.curToken.Type != lexer.ItemLeftBracket {
		return nil, expectedError(lexer.ItemLeftBracket, p.peekToken)
	}
	p.next() // consume left bracket [
	primary, err := p.ParsePrimary()
	if err != nil {
		return nil, err
	}
	keyExpr, err := p.ParseExpression(primary, 0)
	if err != nil {
		return nil, err
	}
	p.next() // consume end of expression
	if err := p.want(lexer.ItemRightBracket); err != nil {
		return nil, err
	}
	return ast.NewExtractElement(expr, keyExpr), nil
}

func (p *Parser) ParseArgument(expr ast.Expr) (*ast.FuncCall, error) {
	if p.curToken.Type != lexer.ItemLeftParen {
		return nil, expectedError(lexer.ItemLeftParen, p.peekToken)
	}
	p.next() // consume left paren (

	args := make([]*ast.Arg, 0)
	for p.curToken.Type != lexer.ItemRightParen {
		identifier := ""
		// `identifier:` is optional, so we don't check the error
		if err := p.want(lexer.ItemIdent); err == nil {
			if p.peekToken.Type == (lexer.ItemColon) {
				identifier = p.curToken.Val
				p.next() // consume identifier
				p.next() // consume colon :
			}
		}
		primary, err := p.ParsePrimary()
		if err != nil {
			return nil, err
		}
		expr, err := p.ParseExpression(primary, 0)
		if err != nil {
			return nil, err
		}
		p.next() // consume end of expression
		args = append(args, ast.NewArg(identifier, expr))
		if err := p.want(lexer.ItemComma); err != nil {
			if err := p.want(lexer.ItemRightParen); err != nil {
				return nil, err
			} else {
				break
			}
		}
		p.next() // consume comma
	}
	return &ast.FuncCall{
		X:       expr,
		ArgList: args,
	}, nil
}

func (p *Parser) ParseAccessChain() (*ast.Ident, error) {
	if err := p.want(lexer.ItemIdent); err != nil {
		return nil, err
	}
	var identifier strings.Builder
	identifier.WriteString(p.curToken.Val)
	for p.peekToken.Type == lexer.ItemColon {
		p.next() // consume identifier
		if p.peekToken.Type == lexer.ItemColon {
			// match :: now
			// consume ::
			p.next()
			p.next()
			if err := p.want(lexer.ItemIdent); err != nil {
				return nil, err
			}
			identifier.WriteString("::")
			identifier.WriteString(p.curToken.Val)
		}
	}
	return ast.NewIdent(identifier.String()), nil
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

func expectedError(expected, actual fmt.Stringer) error {
	return fmt.Errorf("expected: %s but got: %s", expected, actual)
}