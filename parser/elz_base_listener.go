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

// EnterReturnStat is called when production returnStat is entered.
func (s *BaseElzListener) EnterReturnStat(ctx *ReturnStatContext) {}

// ExitReturnStat is called when production returnStat is exited.
func (s *BaseElzListener) ExitReturnStat(ctx *ReturnStatContext) {}

// EnterLoopStat is called when production loopStat is entered.
func (s *BaseElzListener) EnterLoopStat(ctx *LoopStatContext) {}

// ExitLoopStat is called when production loopStat is exited.
func (s *BaseElzListener) ExitLoopStat(ctx *LoopStatContext) {}

// EnterExprStat is called when production exprStat is entered.
func (s *BaseElzListener) EnterExprStat(ctx *ExprStatContext) {}

// ExitExprStat is called when production exprStat is exited.
func (s *BaseElzListener) ExitExprStat(ctx *ExprStatContext) {}

// EnterMatchRule is called when production matchRule is entered.
func (s *BaseElzListener) EnterMatchRule(ctx *MatchRuleContext) {}

// ExitMatchRule is called when production matchRule is exited.
func (s *BaseElzListener) ExitMatchRule(ctx *MatchRuleContext) {}

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

// EnterTypeList is called when production typeList is entered.
func (s *BaseElzListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseElzListener) ExitTypeList(ctx *TypeListContext) {}

// EnterMethodList is called when production methodList is entered.
func (s *BaseElzListener) EnterMethodList(ctx *MethodListContext) {}

// ExitMethodList is called when production methodList is exited.
func (s *BaseElzListener) ExitMethodList(ctx *MethodListContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseElzListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseElzListener) ExitMethod(ctx *MethodContext) {}

// EnterImplBlock is called when production implBlock is entered.
func (s *BaseElzListener) EnterImplBlock(ctx *ImplBlockContext) {}

// ExitImplBlock is called when production implBlock is exited.
func (s *BaseElzListener) ExitImplBlock(ctx *ImplBlockContext) {}

// EnterExportor is called when production exportor is entered.
func (s *BaseElzListener) EnterExportor(ctx *ExportorContext) {}

// ExitExportor is called when production exportor is exited.
func (s *BaseElzListener) ExitExportor(ctx *ExportorContext) {}

// EnterDefine is called when production define is entered.
func (s *BaseElzListener) EnterDefine(ctx *DefineContext) {}

// ExitDefine is called when production define is exited.
func (s *BaseElzListener) ExitDefine(ctx *DefineContext) {}

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

// EnterNum is called when production num is entered.
func (s *BaseElzListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BaseElzListener) ExitNum(ctx *NumContext) {}

// EnterId is called when production id is entered.
func (s *BaseElzListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseElzListener) ExitId(ctx *IdContext) {}

// EnterStr is called when production str is entered.
func (s *BaseElzListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production str is exited.
func (s *BaseElzListener) ExitStr(ctx *StrContext) {}
