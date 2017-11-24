// Generated from Elz.g4 by ANTLR 4.7.

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

	// EnterTypePass is called when entering the typePass production.
	EnterTypePass(c *TypePassContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterMethodList is called when entering the methodList production.
	EnterMethodList(c *MethodListContext)

	// EnterMethod is called when entering the method production.
	EnterMethod(c *MethodContext)

	// EnterImplBlock is called when entering the implBlock production.
	EnterImplBlock(c *ImplBlockContext)

	// EnterExportor is called when entering the exportor production.
	EnterExportor(c *ExportorContext)

	// EnterDefine is called when entering the define production.
	EnterDefine(c *DefineContext)

	// EnterVarDefine is called when entering the varDefine production.
	EnterVarDefine(c *VarDefineContext)

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

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitTopStatList is called when exiting the topStatList production.
	ExitTopStatList(c *TopStatListContext)

	// ExitTopStat is called when exiting the topStat production.
	ExitTopStat(c *TopStatContext)

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

	// ExitTypePass is called when exiting the typePass production.
	ExitTypePass(c *TypePassContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitMethodList is called when exiting the methodList production.
	ExitMethodList(c *MethodListContext)

	// ExitMethod is called when exiting the method production.
	ExitMethod(c *MethodContext)

	// ExitImplBlock is called when exiting the implBlock production.
	ExitImplBlock(c *ImplBlockContext)

	// ExitExportor is called when exiting the exportor production.
	ExitExportor(c *ExportorContext)

	// ExitDefine is called when exiting the define production.
	ExitDefine(c *DefineContext)

	// ExitVarDefine is called when exiting the varDefine production.
	ExitVarDefine(c *VarDefineContext)

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

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)
}
