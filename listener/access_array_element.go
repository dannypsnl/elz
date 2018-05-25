package listener

import (
	"strconv"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitAccessArrayElement(c *parser.AccessArrayElementContext) {
	index, _ := strconv.Atoi(c.INT().GetText())
	s.exprStack.Push(&ast.ArrayElement{
		E:     s.exprStack.Pop().(ast.Expr),
		Index: index,
	})
}
