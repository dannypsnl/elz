package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitReturnStat(ctx *parser.ReturnStatContext) {
	// get expr
	expr := s.exprStack.Pop()
	if s.fnBuilder != nil {
		s.fnBuilder.Stat(&ast.Return{
			Expr: expr.(ast.Expr),
		})
	} else {
		// FIXME: please use error reporter handle error report
		// return statement only allow in function, so it's an error if no fnBuilder at here
		panic("return error")
	}
}
