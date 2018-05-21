package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitAssign(c *parser.AssignContext) {
	if s.fnBuilder != nil {
		s.fnBuilder.Stat(&ast.Assign{
			VarName: c.ID().GetText(),
			E:       s.exprStack.Pop().(ast.Expr),
		})
	} else {
		s.context.Reporter.Emit("assign statement should use in function")
	}
}
