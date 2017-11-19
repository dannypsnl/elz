// Generated from Elz.g4 by ANTLR 4.7.

package parser // Elz

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseElzListener is a complete listener for a parse tree produced by ElzParser.
type BaseElzListener struct{}

var _ ElzListener = &BaseElzListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseElzListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseElzListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseElzListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseElzListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseElzListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseElzListener) ExitProg(ctx *ProgContext) {}

// EnterTopStatList is called when production topStatList is entered.
func (s *BaseElzListener) EnterTopStatList(ctx *TopStatListContext) {}

// ExitTopStatList is called when production topStatList is exited.
func (s *BaseElzListener) ExitTopStatList(ctx *TopStatListContext) {}

// EnterTopStat is called when production topStat is entered.
func (s *BaseElzListener) EnterTopStat(ctx *TopStatContext) {}

// ExitTopStat is called when production topStat is exited.
func (s *BaseElzListener) ExitTopStat(ctx *TopStatContext) {}

// EnterImportStat is called when production importStat is entered.
func (s *BaseElzListener) EnterImportStat(ctx *ImportStatContext) {}

// ExitImportStat is called when production importStat is exited.
func (s *BaseElzListener) ExitImportStat(ctx *ImportStatContext) {}

// EnterStatList is called when production statList is entered.
func (s *BaseElzListener) EnterStatList(ctx *StatListContext) {}

// ExitStatList is called when production statList is exited.
func (s *BaseElzListener) ExitStatList(ctx *StatListContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseElzListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseElzListener) ExitStat(ctx *StatContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseElzListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseElzListener) ExitAssign(ctx *AssignContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseElzListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseElzListener) ExitExprList(ctx *ExprListContext) {}

// EnterFnCall is called when production fnCall is entered.
func (s *BaseElzListener) EnterFnCall(ctx *FnCallContext) {}

// ExitFnCall is called when production fnCall is exited.
func (s *BaseElzListener) ExitFnCall(ctx *FnCallContext) {}

// EnterTypePass is called when production typePass is entered.
func (s *BaseElzListener) EnterTypePass(ctx *TypePassContext) {}

// ExitTypePass is called when production typePass is exited.
func (s *BaseElzListener) ExitTypePass(ctx *TypePassContext) {}

// EnterExportor is called when production exportor is entered.
func (s *BaseElzListener) EnterExportor(ctx *ExportorContext) {}

// ExitExportor is called when production exportor is exited.
func (s *BaseElzListener) ExitExportor(ctx *ExportorContext) {}

// EnterVarDefine is called when production varDefine is entered.
func (s *BaseElzListener) EnterVarDefine(ctx *VarDefineContext) {}

// ExitVarDefine is called when production varDefine is exited.
func (s *BaseElzListener) ExitVarDefine(ctx *VarDefineContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseElzListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseElzListener) ExitParamList(ctx *ParamListContext) {}

// EnterParam is called when production param is entered.
func (s *BaseElzListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseElzListener) ExitParam(ctx *ParamContext) {}

// EnterFnDefine is called when production fnDefine is entered.
func (s *BaseElzListener) EnterFnDefine(ctx *FnDefineContext) {}

// ExitFnDefine is called when production fnDefine is exited.
func (s *BaseElzListener) ExitFnDefine(ctx *FnDefineContext) {}

// EnterAttrList is called when production attrList is entered.
func (s *BaseElzListener) EnterAttrList(ctx *AttrListContext) {}

// ExitAttrList is called when production attrList is exited.
func (s *BaseElzListener) ExitAttrList(ctx *AttrListContext) {}

// EnterAttr is called when production attr is entered.
func (s *BaseElzListener) EnterAttr(ctx *AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *BaseElzListener) ExitAttr(ctx *AttrContext) {}

// EnterTypeDefine is called when production typeDefine is entered.
func (s *BaseElzListener) EnterTypeDefine(ctx *TypeDefineContext) {}

// ExitTypeDefine is called when production typeDefine is exited.
func (s *BaseElzListener) ExitTypeDefine(ctx *TypeDefineContext) {}

// EnterTmethodList is called when production tmethodList is entered.
func (s *BaseElzListener) EnterTmethodList(ctx *TmethodListContext) {}

// ExitTmethodList is called when production tmethodList is exited.
func (s *BaseElzListener) ExitTmethodList(ctx *TmethodListContext) {}

// EnterTmethod is called when production tmethod is entered.
func (s *BaseElzListener) EnterTmethod(ctx *TmethodContext) {}

// ExitTmethod is called when production tmethod is exited.
func (s *BaseElzListener) ExitTmethod(ctx *TmethodContext) {}

// EnterTraitDefine is called when production traitDefine is entered.
func (s *BaseElzListener) EnterTraitDefine(ctx *TraitDefineContext) {}

// ExitTraitDefine is called when production traitDefine is exited.
func (s *BaseElzListener) ExitTraitDefine(ctx *TraitDefineContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseElzListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseElzListener) ExitExpr(ctx *ExprContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseElzListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseElzListener) ExitFactor(ctx *FactorContext) {}
