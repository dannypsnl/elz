
// Generated from Elz.g4 by ANTLR 4.7

#pragma once


#include "antlr4-runtime.h"
#include "ElzListener.h"


/**
 * This class provides an empty implementation of ElzListener,
 * which can be extended to create a listener which only needs to handle a subset
 * of the available methods.
 */
class  ElzBaseListener : public ElzListener {
public:

  virtual void enterProg(ElzParser::ProgContext * /*ctx*/) override { }
  virtual void exitProg(ElzParser::ProgContext * /*ctx*/) override { }

  virtual void enterTopStatList(ElzParser::TopStatListContext * /*ctx*/) override { }
  virtual void exitTopStatList(ElzParser::TopStatListContext * /*ctx*/) override { }

  virtual void enterTopStat(ElzParser::TopStatContext * /*ctx*/) override { }
  virtual void exitTopStat(ElzParser::TopStatContext * /*ctx*/) override { }

  virtual void enterImportStat(ElzParser::ImportStatContext * /*ctx*/) override { }
  virtual void exitImportStat(ElzParser::ImportStatContext * /*ctx*/) override { }

  virtual void enterStatList(ElzParser::StatListContext * /*ctx*/) override { }
  virtual void exitStatList(ElzParser::StatListContext * /*ctx*/) override { }

  virtual void enterStat(ElzParser::StatContext * /*ctx*/) override { }
  virtual void exitStat(ElzParser::StatContext * /*ctx*/) override { }

  virtual void enterReturnStat(ElzParser::ReturnStatContext * /*ctx*/) override { }
  virtual void exitReturnStat(ElzParser::ReturnStatContext * /*ctx*/) override { }

  virtual void enterLoopStat(ElzParser::LoopStatContext * /*ctx*/) override { }
  virtual void exitLoopStat(ElzParser::LoopStatContext * /*ctx*/) override { }

  virtual void enterExprStat(ElzParser::ExprStatContext * /*ctx*/) override { }
  virtual void exitExprStat(ElzParser::ExprStatContext * /*ctx*/) override { }

  virtual void enterMatchRule(ElzParser::MatchRuleContext * /*ctx*/) override { }
  virtual void exitMatchRule(ElzParser::MatchRuleContext * /*ctx*/) override { }

  virtual void enterAssign(ElzParser::AssignContext * /*ctx*/) override { }
  virtual void exitAssign(ElzParser::AssignContext * /*ctx*/) override { }

  virtual void enterExprList(ElzParser::ExprListContext * /*ctx*/) override { }
  virtual void exitExprList(ElzParser::ExprListContext * /*ctx*/) override { }

  virtual void enterFnCall(ElzParser::FnCallContext * /*ctx*/) override { }
  virtual void exitFnCall(ElzParser::FnCallContext * /*ctx*/) override { }

  virtual void enterTypePass(ElzParser::TypePassContext * /*ctx*/) override { }
  virtual void exitTypePass(ElzParser::TypePassContext * /*ctx*/) override { }

  virtual void enterTypeList(ElzParser::TypeListContext * /*ctx*/) override { }
  virtual void exitTypeList(ElzParser::TypeListContext * /*ctx*/) override { }

  virtual void enterMethodList(ElzParser::MethodListContext * /*ctx*/) override { }
  virtual void exitMethodList(ElzParser::MethodListContext * /*ctx*/) override { }

  virtual void enterMethod(ElzParser::MethodContext * /*ctx*/) override { }
  virtual void exitMethod(ElzParser::MethodContext * /*ctx*/) override { }

  virtual void enterImplBlock(ElzParser::ImplBlockContext * /*ctx*/) override { }
  virtual void exitImplBlock(ElzParser::ImplBlockContext * /*ctx*/) override { }

  virtual void enterExportor(ElzParser::ExportorContext * /*ctx*/) override { }
  virtual void exitExportor(ElzParser::ExportorContext * /*ctx*/) override { }

  virtual void enterDefine(ElzParser::DefineContext * /*ctx*/) override { }
  virtual void exitDefine(ElzParser::DefineContext * /*ctx*/) override { }

  virtual void enterVarDefine(ElzParser::VarDefineContext * /*ctx*/) override { }
  virtual void exitVarDefine(ElzParser::VarDefineContext * /*ctx*/) override { }

  virtual void enterParamList(ElzParser::ParamListContext * /*ctx*/) override { }
  virtual void exitParamList(ElzParser::ParamListContext * /*ctx*/) override { }

  virtual void enterParam(ElzParser::ParamContext * /*ctx*/) override { }
  virtual void exitParam(ElzParser::ParamContext * /*ctx*/) override { }

  virtual void enterFnDefine(ElzParser::FnDefineContext * /*ctx*/) override { }
  virtual void exitFnDefine(ElzParser::FnDefineContext * /*ctx*/) override { }

  virtual void enterAttrList(ElzParser::AttrListContext * /*ctx*/) override { }
  virtual void exitAttrList(ElzParser::AttrListContext * /*ctx*/) override { }

  virtual void enterAttr(ElzParser::AttrContext * /*ctx*/) override { }
  virtual void exitAttr(ElzParser::AttrContext * /*ctx*/) override { }

  virtual void enterTypeDefine(ElzParser::TypeDefineContext * /*ctx*/) override { }
  virtual void exitTypeDefine(ElzParser::TypeDefineContext * /*ctx*/) override { }

  virtual void enterTmethodList(ElzParser::TmethodListContext * /*ctx*/) override { }
  virtual void exitTmethodList(ElzParser::TmethodListContext * /*ctx*/) override { }

  virtual void enterTmethod(ElzParser::TmethodContext * /*ctx*/) override { }
  virtual void exitTmethod(ElzParser::TmethodContext * /*ctx*/) override { }

  virtual void enterTraitDefine(ElzParser::TraitDefineContext * /*ctx*/) override { }
  virtual void exitTraitDefine(ElzParser::TraitDefineContext * /*ctx*/) override { }

  virtual void enterExpr(ElzParser::ExprContext * /*ctx*/) override { }
  virtual void exitExpr(ElzParser::ExprContext * /*ctx*/) override { }

  virtual void enterFactor(ElzParser::FactorContext * /*ctx*/) override { }
  virtual void exitFactor(ElzParser::FactorContext * /*ctx*/) override { }


  virtual void enterEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void exitEveryRule(antlr4::ParserRuleContext * /*ctx*/) override { }
  virtual void visitTerminal(antlr4::tree::TerminalNode * /*node*/) override { }
  virtual void visitErrorNode(antlr4::tree::ErrorNode * /*node*/) override { }

};

