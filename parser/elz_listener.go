// Code generated from Elz.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Elz

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ElzListener is a complete listener for a parse tree produced by ElzParser.
type ElzListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterTopStatList is called when entering the topStatList production.
	EnterTopStatList(c *TopStatListContext)

	// EnterTopStat is called when entering the topStat production.
	EnterTopStat(c *TopStatContext)

	// EnterImportMod is called when entering the importMod production.
	EnterImportMod(c *ImportModContext)

	// EnterImportStat is called when entering the importStat production.
	EnterImportStat(c *ImportStatContext)

	// EnterStatList is called when entering the statList production.
	EnterStatList(c *StatListContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterReturnStat is called when entering the returnStat production.
	EnterReturnStat(c *ReturnStatContext)

	// EnterLoopStat is called when entering the loopStat production.
	EnterLoopStat(c *LoopStatContext)

	// EnterExprStat is called when entering the exprStat production.
	EnterExprStat(c *ExprStatContext)

	// EnterMatchRule is called when entering the matchRule production.
	EnterMatchRule(c *MatchRuleContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterFnCall is called when entering the fnCall production.
	EnterFnCall(c *FnCallContext)

	// EnterTypeForm is called when entering the typeForm production.
	EnterTypeForm(c *TypeFormContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterMethodList is called when entering the methodList production.
	EnterMethodList(c *MethodListContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterImplBlock is called when entering the implBlock production.
	EnterImplBlock(c *ImplBlockContext)

	// EnterExportor is called when entering the exportor production.
	EnterExportor(c *ExportorContext)

	// EnterGlobalVarDefine is called when entering the globalVarDefine production.
	EnterGlobalVarDefine(c *GlobalVarDefineContext)

	// EnterDefine is called when entering the define production.
	EnterDefine(c *DefineContext)

	// EnterLocalVarDefine is called when entering the localVarDefine production.
	EnterLocalVarDefine(c *LocalVarDefineContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterFnDefine is called when entering the fnDefine production.
	EnterFnDefine(c *FnDefineContext)

	// EnterAttrList is called when entering the attrList production.
	EnterAttrList(c *AttrListContext)

	// EnterAttr is called when entering the attr production.
	EnterAttr(c *AttrContext)

	// EnterTypeDefine is called when entering the typeDefine production.
	EnterTypeDefine(c *TypeDefineContext)

	// EnterTmethodList is called when entering the tmethodList production.
	EnterTmethodList(c *TmethodListContext)

	// EnterTmethod is called when entering the tmethod production.
	EnterTmethod(c *TmethodContext)

	// EnterTraitDefine is called when entering the traitDefine production.
	EnterTraitDefine(c *TraitDefineContext)

	// EnterNotEq is called when entering the NotEq production.
	EnterNotEq(c *NotEqContext)

	// EnterSubExpr is called when entering the SubExpr production.
	EnterSubExpr(c *SubExprContext)

	// EnterNum is called when entering the Num production.
	EnterNum(c *NumContext)

	// EnterMulOrDiv is called when entering the MulOrDiv production.
	EnterMulOrDiv(c *MulOrDivContext)

	// EnterCmp is called when entering the Cmp production.
	EnterCmp(c *CmpContext)

	// EnterEq is called when entering the Eq production.
	EnterEq(c *EqContext)

	// EnterAndOrOr is called when entering the AndOrOr production.
	EnterAndOrOr(c *AndOrOrContext)

	// EnterThreeOpCmp is called when entering the ThreeOpCmp production.
	EnterThreeOpCmp(c *ThreeOpCmpContext)

	// EnterStr is called when entering the Str production.
	EnterStr(c *StrContext)

	// EnterAddOrSub is called when entering the AddOrSub production.
	EnterAddOrSub(c *AddOrSubContext)

	// EnterStatExpr is called when entering the StatExpr production.
	EnterStatExpr(c *StatExprContext)

	// EnterPow is called when entering the Pow production.
	EnterPow(c *PowContext)

	// EnterId is called when entering the Id production.
	EnterId(c *IdContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitTopStatList is called when exiting the topStatList production.
	ExitTopStatList(c *TopStatListContext)

	// ExitTopStat is called when exiting the topStat production.
	ExitTopStat(c *TopStatContext)

	// ExitImportMod is called when exiting the importMod production.
	ExitImportMod(c *ImportModContext)

	// ExitImportStat is called when exiting the importStat production.
	ExitImportStat(c *ImportStatContext)

	// ExitStatList is called when exiting the statList production.
	ExitStatList(c *StatListContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitReturnStat is called when exiting the returnStat production.
	ExitReturnStat(c *ReturnStatContext)

	// ExitLoopStat is called when exiting the loopStat production.
	ExitLoopStat(c *LoopStatContext)

	// ExitExprStat is called when exiting the exprStat production.
	ExitExprStat(c *ExprStatContext)

	// ExitMatchRule is called when exiting the matchRule production.
	ExitMatchRule(c *MatchRuleContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitFnCall is called when exiting the fnCall production.
	ExitFnCall(c *FnCallContext)

	// ExitTypeForm is called when exiting the typeForm production.
	ExitTypeForm(c *TypeFormContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitMethodList is called when exiting the methodList production.
	ExitMethodList(c *MethodListContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitImplBlock is called when exiting the implBlock production.
	ExitImplBlock(c *ImplBlockContext)

	// ExitExportor is called when exiting the exportor production.
	ExitExportor(c *ExportorContext)

	// ExitGlobalVarDefine is called when exiting the globalVarDefine production.
	ExitGlobalVarDefine(c *GlobalVarDefineContext)

	// ExitDefine is called when exiting the define production.
	ExitDefine(c *DefineContext)

	// ExitLocalVarDefine is called when exiting the localVarDefine production.
	ExitLocalVarDefine(c *LocalVarDefineContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitFnDefine is called when exiting the fnDefine production.
	ExitFnDefine(c *FnDefineContext)

	// ExitAttrList is called when exiting the attrList production.
	ExitAttrList(c *AttrListContext)

	// ExitAttr is called when exiting the attr production.
	ExitAttr(c *AttrContext)

	// ExitTypeDefine is called when exiting the typeDefine production.
	ExitTypeDefine(c *TypeDefineContext)

	// ExitTmethodList is called when exiting the tmethodList production.
	ExitTmethodList(c *TmethodListContext)

	// ExitTmethod is called when exiting the tmethod production.
	ExitTmethod(c *TmethodContext)

	// ExitTraitDefine is called when exiting the traitDefine production.
	ExitTraitDefine(c *TraitDefineContext)

	// ExitNotEq is called when exiting the NotEq production.
	ExitNotEq(c *NotEqContext)

	// ExitSubExpr is called when exiting the SubExpr production.
	ExitSubExpr(c *SubExprContext)

	// ExitNum is called when exiting the Num production.
	ExitNum(c *NumContext)

	// ExitMulOrDiv is called when exiting the MulOrDiv production.
	ExitMulOrDiv(c *MulOrDivContext)

	// ExitCmp is called when exiting the Cmp production.
	ExitCmp(c *CmpContext)

	// ExitEq is called when exiting the Eq production.
	ExitEq(c *EqContext)

	// ExitAndOrOr is called when exiting the AndOrOr production.
	ExitAndOrOr(c *AndOrOrContext)

	// ExitThreeOpCmp is called when exiting the ThreeOpCmp production.
	ExitThreeOpCmp(c *ThreeOpCmpContext)

	// ExitStr is called when exiting the Str production.
	ExitStr(c *StrContext)

	// ExitAddOrSub is called when exiting the AddOrSub production.
	ExitAddOrSub(c *AddOrSubContext)

	// ExitStatExpr is called when exiting the StatExpr production.
	ExitStatExpr(c *StatExprContext)

	// ExitPow is called when exiting the Pow production.
	ExitPow(c *PowContext)

	// ExitId is called when exiting the Id production.
	ExitId(c *IdContext)
}
