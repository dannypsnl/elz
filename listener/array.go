package listener

import (
	"strconv"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/collection/stack"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitArrWithLen(c *parser.ArrWithLenContext) {
	expr := s.exprStack.Pop()
	length := c.INT().GetText()
	var (
		i   int
		err error
	)
	if i, err = strconv.Atoi(length); err != nil {
		panic("parser INT rule is not a integer, must is error")
	}
	s.exprStack.Push(&ast.Array{
		Elements: []ast.Expr{expr.(ast.Expr)},
		Len:      i,
	})
}

// '[' expr ';' INT ']'             # ArrWithLen
// '[' exprList ']'                 # ArrWithList
//for expr := s.exprStack.Pop(); expr != nil; expr = s.exprStack.Pop() {}

func (s *ElzListener) ExitArrWithList(c *parser.ArrWithListContext) {
	exprs := make([]ast.Expr, 0)
	tmp := stack.New()
	for e := s.exprStack.Pop(); e != nil; e = s.exprStack.Pop() {
		tmp.Push(e)
	}
	for e := tmp.Pop(); e != nil; e = tmp.Pop() {
		exprs = append(exprs, e.(ast.Expr))
	}
	s.exprStack.Push(&ast.Array{
		Elements: exprs,
		Len:      len(exprs),
	})
}
