package listener

import (
	"strconv"

	"github.com/elz-lang/elz/ast"
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
		Elements:    []ast.Expr{expr.(ast.Expr)},
		ElementType: "",
		Len:         i,
	})
}

// '[' expr ';' INT ']'             # ArrWithLen
// '[' exprList ']'                 # ArrWithList
//for expr := s.exprStack.Pop(); expr != nil; expr = s.exprStack.Pop() {}
