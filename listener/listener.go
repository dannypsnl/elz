package listener

import (
	"fmt"

	"github.com/elz-lang/elz/ast"
	"github.com/elz-lang/elz/collection/stack"
	"github.com/elz-lang/elz/parser"
	"llvm.org/llvm/bindings/go/llvm"
)

type ElzListener struct {
	*parser.BaseElzListener
	context *ast.Context
	// AstList contain top level's ast
	AstList []ast.Ast
	// exprStack help we implement expression percedence table.
	exprStack *stack.Stack // Stack Pop nil is nothing in there
	// fnBuilder
	fnBuilder *FnBuilder
	// exportThis markup the reference Name should be public or not.
	exportThis bool
	// variable default immutable.
	immutable bool
}

// Module return the llvm.Module generate by parse process
func (s *ElzListener) Module() llvm.Module {
	for _, ast := range s.AstList {
		ast.Codegen(s.context)
	}
	return s.context.Module
}

// New create a new listener
func New() *ElzListener {
	return &ElzListener{
		context:   ast.NewContext(),
		immutable: true,
		exprStack: stack.New(),
	}
}

func (s *ElzListener) EnterProg(ctx *parser.ProgContext) {
	fmt.Println(`Elz prog`)
}

// EnterExportor: + prefix
func (s *ElzListener) EnterExportor(*parser.ExportorContext) {
	s.exportThis = true
}

// VarDef:
//   let (mut) $var_name = $expr
func (s *ElzListener) EnterVarDefine(ctx *parser.VarDefineContext) {
	fmt.Print(`var `)
	if ctx.GetMut() != nil {
		s.immutable = false
	}
}

func (s *ElzListener) ExitVarDefine(*parser.VarDefineContext) {
	if !s.immutable {
		s.immutable = true
	}
}

// Def:
//   $var_name = $expr
func (s *ElzListener) ExitDefine(ctx *parser.DefineContext) {
	// get expr
	expr := s.exprStack.Pop()
	// get type from expression
	typ := expr.(ast.Expr).Type(s.context)
	// get identifier
	// TODO: fix with scope rule, and some rule to detected fn, type, trait or what
	name := ctx.ID().GetText()
	// get type from source code, so we can find out the problem if expr != user_def type
	if ctx.TypePass() != nil {
		typ = ctx.TypePass().GetText()
	}

	if s.exportThis {
		fmt.Print("public ")
		s.exportThis = false
	}
	fmt.Printf("%s: %s = %s\n", ctx.ID().GetText(), typ, expr)

	// FIXME: Need to classify global var & local var, because local var of course can't be export
	// FIXME: Need to classify heap & stack, and can find out the lifetime, else sending data by return will become bug
	s.AstList = append(s.AstList, &ast.VarDefination{
		// TODO: immutable should be put in an array, and don't need to be knew by LLVM Module, because LLVM is SSA form
		Immutable:  s.immutable,
		Export:     s.exportThis,
		Name:       name,
		VarType:    typ,
		Expression: expr.(ast.Expr),
	})
	// Record type for compiler
	s.context.VarsType[name] = typ
}

type FnBuilder struct {
	export    bool
	name      string
	returnTyp string
	// TODO: How to record?
	// Type: Param = { Name, Type }
	params []*ast.Param
}

func NewFnBuilder() *FnBuilder {
	return &FnBuilder{}
}

func (fb *FnBuilder) Name(n string) *FnBuilder {
	fb.name = n
	return fb
}
func (fb *FnBuilder) Export(e bool) *FnBuilder {
	fb.export = e
	return fb
}
func (fb *FnBuilder) generate() *ast.FnDef {
	if fb.export {
		fmt.Print("public ")
	}
	// FIXME: This is let result show a fn, not correct impl
	fmt.Printf("fn %s\n", fb.name)
	return &ast.FnDef{
		Export: fb.export,
		Name:   fb.name,
		Params: []*ast.Param{},
		// TODO: implement statments
		// FIXME: should decide by rule typePass
		RetType: "num",
	}
	// TODO: local var def need spec_name
}

func (s *ElzListener) EnterFnDefine(ctx *parser.FnDefineContext) {
	// TODO: complete fn generate
	s.fnBuilder = NewFnBuilder().
		Name(ctx.ID().GetText()).
		Export(s.exportThis)
}
func (s *ElzListener) ExitFnDefine(ctx *parser.FnDefineContext) {
	s.AstList = append(s.AstList,
		s.fnBuilder.generate(),
	)
	// TODO: fn need builder to create at here, because it will cross several rules
	s.fnBuilder = nil
}
