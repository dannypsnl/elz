
// Generated from Elz.g4 by ANTLR 4.7

#pragma once


#include "antlr4-runtime.h"
#include "ElzParser.h"


/**
 * This interface defines an abstract listener for a parse tree produced by ElzParser.
 */
class  ElzListener : public antlr4::tree::ParseTreeListener {
public:

  virtual void enterProg(ElzParser::ProgContext *ctx) = 0;
  virtual void exitProg(ElzParser::ProgContext *ctx) = 0;

  virtual void enterTopStatList(ElzParser::TopStatListContext *ctx) = 0;
  virtual void exitTopStatList(ElzParser::TopStatListContext *ctx) = 0;

  virtual void enterTopStat(ElzParser::TopStatContext *ctx) = 0;
  virtual void exitTopStat(ElzParser::TopStatContext *ctx) = 0;

  virtual void enterImportStat(ElzParser::ImportStatContext *ctx) = 0;
  virtual void exitImportStat(ElzParser::ImportStatContext *ctx) = 0;

  virtual void enterStatList(ElzParser::StatListContext *ctx) = 0;
  virtual void exitStatList(ElzParser::StatListContext *ctx) = 0;

  virtual void enterStat(ElzParser::StatContext *ctx) = 0;
  virtual void exitStat(ElzParser::StatContext *ctx) = 0;

  virtual void enterReturnStat(ElzParser::ReturnStatContext *ctx) = 0;
  virtual void exitReturnStat(ElzParser::ReturnStatContext *ctx) = 0;

  virtual void enterLoopStat(ElzParser::LoopStatContext *ctx) = 0;
  virtual void exitLoopStat(ElzParser::LoopStatContext *ctx) = 0;

  virtual void enterExprStat(ElzParser::ExprStatContext *ctx) = 0;
  virtual void exitExprStat(ElzParser::ExprStatContext *ctx) = 0;

  virtual void enterMatchRule(ElzParser::MatchRuleContext *ctx) = 0;
  virtual void exitMatchRule(ElzParser::MatchRuleContext *ctx) = 0;

  virtual void enterAssign(ElzParser::AssignContext *ctx) = 0;
  virtual void exitAssign(ElzParser::AssignContext *ctx) = 0;

  virtual void enterExprList(ElzParser::ExprListContext *ctx) = 0;
  virtual void exitExprList(ElzParser::ExprListContext *ctx) = 0;

  virtual void enterFnCall(ElzParser::FnCallContext *ctx) = 0;
  virtual void exitFnCall(ElzParser::FnCallContext *ctx) = 0;

  virtual void enterTypePass(ElzParser::TypePassContext *ctx) = 0;
  virtual void exitTypePass(ElzParser::TypePassContext *ctx) = 0;

  virtual void enterTypeList(ElzParser::TypeListContext *ctx) = 0;
  virtual void exitTypeList(ElzParser::TypeListContext *ctx) = 0;

  virtual void enterMethodList(ElzParser::MethodListContext *ctx) = 0;
  virtual void exitMethodList(ElzParser::MethodListContext *ctx) = 0;

  virtual void enterMethod(ElzParser::MethodContext *ctx) = 0;
  virtual void exitMethod(ElzParser::MethodContext *ctx) = 0;

  virtual void enterImplBlock(ElzParser::ImplBlockContext *ctx) = 0;
  virtual void exitImplBlock(ElzParser::ImplBlockContext *ctx) = 0;

  virtual void enterExportor(ElzParser::ExportorContext *ctx) = 0;
  virtual void exitExportor(ElzParser::ExportorContext *ctx) = 0;

  virtual void enterDefine(ElzParser::DefineContext *ctx) = 0;
  virtual void exitDefine(ElzParser::DefineContext *ctx) = 0;

  virtual void enterVarDefine(ElzParser::VarDefineContext *ctx) = 0;
  virtual void exitVarDefine(ElzParser::VarDefineContext *ctx) = 0;

  virtual void enterParamList(ElzParser::ParamListContext *ctx) = 0;
  virtual void exitParamList(ElzParser::ParamListContext *ctx) = 0;

  virtual void enterParam(ElzParser::ParamContext *ctx) = 0;
  virtual void exitParam(ElzParser::ParamContext *ctx) = 0;

  virtual void enterFnDefine(ElzParser::FnDefineContext *ctx) = 0;
  virtual void exitFnDefine(ElzParser::FnDefineContext *ctx) = 0;

  virtual void enterAttrList(ElzParser::AttrListContext *ctx) = 0;
  virtual void exitAttrList(ElzParser::AttrListContext *ctx) = 0;

  virtual void enterAttr(ElzParser::AttrContext *ctx) = 0;
  virtual void exitAttr(ElzParser::AttrContext *ctx) = 0;

  virtual void enterTypeDefine(ElzParser::TypeDefineContext *ctx) = 0;
  virtual void exitTypeDefine(ElzParser::TypeDefineContext *ctx) = 0;

  virtual void enterTmethodList(ElzParser::TmethodListContext *ctx) = 0;
  virtual void exitTmethodList(ElzParser::TmethodListContext *ctx) = 0;

  virtual void enterTmethod(ElzParser::TmethodContext *ctx) = 0;
  virtual void exitTmethod(ElzParser::TmethodContext *ctx) = 0;

  virtual void enterTraitDefine(ElzParser::TraitDefineContext *ctx) = 0;
  virtual void exitTraitDefine(ElzParser::TraitDefineContext *ctx) = 0;

  virtual void enterExpr(ElzParser::ExprContext *ctx) = 0;
  virtual void exitExpr(ElzParser::ExprContext *ctx) = 0;

  virtual void enterFactor(ElzParser::FactorContext *ctx) = 0;
  virtual void exitFactor(ElzParser::FactorContext *ctx) = 0;


};

