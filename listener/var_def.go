package listener

import (
	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/parser"
)

// EnterExportor: + prefix
func (s *ElzListener) EnterExportor(*parser.ExportorContext) {
	s.exportThis = true
}

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
	// get expr
	expr := s.exprStack.Pop()
	// FIXME: this line cause we have to get type at parsing stage
	// get type from expression
	var typ string
	// get identifier
	// TODO: fix with scope rule, and some rule to detected fn, type, trait or what
	name := ctx.ID().GetText()
	// get type from source code, so we can find out the problem if expr != user_def type
	if ctx.TypeForm() != nil {
		typ = ctx.TypeForm().GetText()
	}

	if s.exportThis {
		s.exportThis = false
	}

	// FIXME: Need to classify global var & local var, because local var of course can't be export
	// FIXME: Need to classify heap & stack, and can find out the lifetime, else sending data by return will become bug
	if s.isGlobalDef {
		s.AstList = append(s.AstList, &ast.GlobalVarDef{
			Export:     s.exportThis,
			Name:       name,
			VarType:    typ,
			Expression: expr.(ast.Expr),
		})
		// Record type for compiler
	} else if s.fnBuilder != nil {
		s.fnBuilder.Stat(&ast.LocalVarDef{
			Immutable:  s.immutable,
			Name:       name,
			VarType:    typ,
			Expression: expr.(ast.Expr),
		})
	} else {
		panic("A define should be Global Var else is Local Var, but both of this isn't")
	}
}
