package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

// VarDef:
//   let (mut) $var_name = $expr
func (s *ElzListener) EnterLocalVarDef(ctx *parser.LocalVarDefContext) {
	if ctx.GetMut() != nil {
		s.immutable = false
	}
}

func (s *ElzListener) ExitLocalVarDef(*parser.LocalVarDefContext) {
	if !s.immutable {
		s.immutable = true
	}
}

func (s *ElzListener) EnterGlobalVarDef(*parser.GlobalVarDefContext) {
	s.isGlobalDef = true
}
func (s *ElzListener) ExitGlobalVarDef(*parser.GlobalVarDefContext) {
	s.isGlobalDef = false
}

// Def:
//   $var_name = $expr
func (s *ElzListener) ExitDefine(ctx *parser.DefineContext) {
	// get identifier
	// TODO: fix with scope rule, and some rule to detected fn, type, trait or what
	name := ctx.ID().GetText()
	// get type from expression
	var typ string
	// get type from source code, so we can find out the problem if expr != user_def type
	if ctx.TypeForm() != nil {
		typ = ctx.TypeForm().GetText()
	}

	// get expr
	expr := s.exprStack.Pop().(ast.Expr)

	// FIXME: Need to classify heap & stack, and can find out the lifetime, else return something in stack will become bug
	// TODO: match rule although cause consume statement, have to preparing this
	if s.isGlobalDef {
		s.AstList = append(s.AstList, &ast.GlobalVarDef{
			Export:     s.exportThis,
			Name:       name,
			VarType:    typ,
			Expression: expr,
		})
		// consume export effect
		if s.exportThis {
			s.exportThis = false
		}
		// Record type for compiler
	} else if s.matchRuleBuilder != nil {
		s.matchRuleBuilder.PushStat(&ast.LocalVarDef{
			Immutable:  s.immutable,
			Name:       name,
			VarType:    typ,
			Expression: expr,
		})
	} else if s.fnBuilder != nil {
		// local variable won't be exported
		s.fnBuilder.Stat(&ast.LocalVarDef{
			Immutable:  s.immutable,
			Name:       name,
			VarType:    typ,
			Expression: expr,
		})
	} else {
		panic("A define should be Global Var else is Local Var, but both of this isn't")
	}
}
