package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitAssign(c *parser.AssignContext) {
	if s.statBuilder.Last() == nil {
		s.context.Reporter.Emit("assign statement must in one of statement container that could be pushed into listener.statBuilder")
	}
	s.stats = append(s.stats, &ast.Assign{
		VarName: c.ID().GetText(),
		E:       s.exprStack.Pop().(ast.Expr),
	})
}
