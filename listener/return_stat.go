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
		s.context.Reporter.Emit("return statement must in a function")
	}
}
