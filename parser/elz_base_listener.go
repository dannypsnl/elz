// Code generated from Elz.g4 by ANTLR 4.7.1. DO NOT EDIT.

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

// EnterIntSuffix is called when production intSuffix is entered.
func (s *BaseElzListener) EnterIntSuffix(ctx *IntSuffixContext) {}

// ExitIntSuffix is called when production intSuffix is exited.
func (s *BaseElzListener) ExitIntSuffix(ctx *IntSuffixContext) {}

// EnterFloatSuffix is called when production floatSuffix is entered.
func (s *BaseElzListener) EnterFloatSuffix(ctx *FloatSuffixContext) {}

// ExitFloatSuffix is called when production floatSuffix is exited.
func (s *BaseElzListener) ExitFloatSuffix(ctx *FloatSuffixContext) {}

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

// EnterImportMod is called when production importMod is entered.
func (s *BaseElzListener) EnterImportMod(ctx *ImportModContext) {}

// ExitImportMod is called when production importMod is exited.
func (s *BaseElzListener) ExitImportMod(ctx *ImportModContext) {}

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

// EnterTypeForm is called when production typeForm is entered.
func (s *BaseElzListener) EnterTypeForm(ctx *TypeFormContext) {}

// ExitTypeForm is called when production typeForm is exited.
func (s *BaseElzListener) ExitTypeForm(ctx *TypeFormContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseElzListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseElzListener) ExitTypeList(ctx *TypeListContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseElzListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseElzListener) ExitAnnotation(ctx *AnnotationContext) {}

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

// EnterGlobalVarDef is called when production globalVarDef is entered.
func (s *BaseElzListener) EnterGlobalVarDef(ctx *GlobalVarDefContext) {}

// ExitGlobalVarDef is called when production globalVarDef is exited.
func (s *BaseElzListener) ExitGlobalVarDef(ctx *GlobalVarDefContext) {}

// EnterDefine is called when production define is entered.
func (s *BaseElzListener) EnterDefine(ctx *DefineContext) {}

// ExitDefine is called when production define is exited.
func (s *BaseElzListener) ExitDefine(ctx *DefineContext) {}

// EnterLocalVarDef is called when production localVarDef is entered.
func (s *BaseElzListener) EnterLocalVarDef(ctx *LocalVarDefContext) {}

// ExitLocalVarDef is called when production localVarDef is exited.
func (s *BaseElzListener) ExitLocalVarDef(ctx *LocalVarDefContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseElzListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseElzListener) ExitParamList(ctx *ParamListContext) {}

// EnterParamType is called when production paramType is entered.
func (s *BaseElzListener) EnterParamType(ctx *ParamTypeContext) {}

// ExitParamType is called when production paramType is exited.
func (s *BaseElzListener) ExitParamType(ctx *ParamTypeContext) {}

// EnterParam is called when production param is entered.
func (s *BaseElzListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseElzListener) ExitParam(ctx *ParamContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseElzListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseElzListener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterFnDefine is called when production fnDefine is entered.
func (s *BaseElzListener) EnterFnDefine(ctx *FnDefineContext) {}

// ExitFnDefine is called when production fnDefine is exited.
func (s *BaseElzListener) ExitFnDefine(ctx *FnDefineContext) {}

// EnterDeclareFn is called when production declareFn is entered.
func (s *BaseElzListener) EnterDeclareFn(ctx *DeclareFnContext) {}

// ExitDeclareFn is called when production declareFn is exited.
func (s *BaseElzListener) ExitDeclareFn(ctx *DeclareFnContext) {}

// EnterExternBlock is called when production externBlock is entered.
func (s *BaseElzListener) EnterExternBlock(ctx *ExternBlockContext) {}

// ExitExternBlock is called when production externBlock is exited.
func (s *BaseElzListener) ExitExternBlock(ctx *ExternBlockContext) {}

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

// EnterNotEq is called when production NotEq is entered.
func (s *BaseElzListener) EnterNotEq(ctx *NotEqContext) {}

// ExitNotEq is called when production NotEq is exited.
func (s *BaseElzListener) ExitNotEq(ctx *NotEqContext) {}

// EnterSubExpr is called when production SubExpr is entered.
func (s *BaseElzListener) EnterSubExpr(ctx *SubExprContext) {}

// ExitSubExpr is called when production SubExpr is exited.
func (s *BaseElzListener) ExitSubExpr(ctx *SubExprContext) {}

// EnterMulOrDiv is called when production MulOrDiv is entered.
func (s *BaseElzListener) EnterMulOrDiv(ctx *MulOrDivContext) {}

// ExitMulOrDiv is called when production MulOrDiv is exited.
func (s *BaseElzListener) ExitMulOrDiv(ctx *MulOrDivContext) {}

// EnterCmp is called when production Cmp is entered.
func (s *BaseElzListener) EnterCmp(ctx *CmpContext) {}

// ExitCmp is called when production Cmp is exited.
func (s *BaseElzListener) ExitCmp(ctx *CmpContext) {}

// EnterEq is called when production Eq is entered.
func (s *BaseElzListener) EnterEq(ctx *EqContext) {}

// ExitEq is called when production Eq is exited.
func (s *BaseElzListener) ExitEq(ctx *EqContext) {}

// EnterAndOrOr is called when production AndOrOr is entered.
func (s *BaseElzListener) EnterAndOrOr(ctx *AndOrOrContext) {}

// ExitAndOrOr is called when production AndOrOr is exited.
func (s *BaseElzListener) ExitAndOrOr(ctx *AndOrOrContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseElzListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseElzListener) ExitInt(ctx *IntContext) {}

// EnterThreeOpCmp is called when production ThreeOpCmp is entered.
func (s *BaseElzListener) EnterThreeOpCmp(ctx *ThreeOpCmpContext) {}

// ExitThreeOpCmp is called when production ThreeOpCmp is exited.
func (s *BaseElzListener) ExitThreeOpCmp(ctx *ThreeOpCmpContext) {}

// EnterStr is called when production Str is entered.
func (s *BaseElzListener) EnterStr(ctx *StrContext) {}

// ExitStr is called when production Str is exited.
func (s *BaseElzListener) ExitStr(ctx *StrContext) {}

// EnterFloat is called when production Float is entered.
func (s *BaseElzListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production Float is exited.
func (s *BaseElzListener) ExitFloat(ctx *FloatContext) {}

// EnterRef is called when production Ref is entered.
func (s *BaseElzListener) EnterRef(ctx *RefContext) {}

// ExitRef is called when production Ref is exited.
func (s *BaseElzListener) ExitRef(ctx *RefContext) {}

// EnterAddOrSub is called when production AddOrSub is entered.
func (s *BaseElzListener) EnterAddOrSub(ctx *AddOrSubContext) {}

// ExitAddOrSub is called when production AddOrSub is exited.
func (s *BaseElzListener) ExitAddOrSub(ctx *AddOrSubContext) {}

// EnterAs is called when production As is entered.
func (s *BaseElzListener) EnterAs(ctx *AsContext) {}

// ExitAs is called when production As is exited.
func (s *BaseElzListener) ExitAs(ctx *AsContext) {}

// EnterStatExpr is called when production StatExpr is entered.
func (s *BaseElzListener) EnterStatExpr(ctx *StatExprContext) {}

// ExitStatExpr is called when production StatExpr is exited.
func (s *BaseElzListener) ExitStatExpr(ctx *StatExprContext) {}

// EnterPow is called when production Pow is entered.
func (s *BaseElzListener) EnterPow(ctx *PowContext) {}

// ExitPow is called when production Pow is exited.
func (s *BaseElzListener) ExitPow(ctx *PowContext) {}

// EnterId is called when production Id is entered.
func (s *BaseElzListener) EnterId(ctx *IdContext) {}

// ExitId is called when production Id is exited.
func (s *BaseElzListener) ExitId(ctx *IdContext) {}
