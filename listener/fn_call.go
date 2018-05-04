package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/collection/stack"
	"github.com/elz-lang/elz/parser"
)

func (s *ElzListener) ExitFnCall(ctx *parser.FnCallContext) {
	if s.fnBuilder == nil {
		panic("Not support constant function call yet!")
	}
	stack := stack.New()
	for e := s.exprStack.Pop(); e != nil; e = s.exprStack.Pop() {
		stack.Push(e)
	}
	args := []ast.Expr{}
	for e := stack.Pop(); e != nil; e = stack.Pop() {
		args = append(args, e.(ast.Expr))
	}
	s.exprStack.Push(&ast.FnCall{
		Name: ctx.ID().GetText(),
		Args: args,
	})
}
