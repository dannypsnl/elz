
// Generated from Elz.g4 by ANTLR 4.7


#include "ElzListener.h"

#include "ElzParser.h"


using namespace antlrcpp;
using namespace antlr4;

ElzParser::ElzParser(TokenStream *input) : Parser(input) {
  _interpreter = new atn::ParserATNSimulator(this, _atn, _decisionToDFA, _sharedContextCache);
}

ElzParser::~ElzParser() {
  delete _interpreter;
}

std::string ElzParser::getGrammarFileName() const {
  return "Elz.g4";
}

const std::vector<std::string>& ElzParser::getRuleNames() const {
  return _ruleNames;
}

dfa::Vocabulary& ElzParser::getVocabulary() const {
  return _vocabulary;
}


//----------------- ProgContext ------------------------------------------------------------------

ElzParser::ProgContext::ProgContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

ElzParser::TopStatListContext* ElzParser::ProgContext::topStatList() {
  return getRuleContext<ElzParser::TopStatListContext>(0);
}


size_t ElzParser::ProgContext::getRuleIndex() const {
  return ElzParser::RuleProg;
}

void ElzParser::ProgContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterProg(this);
}

void ElzParser::ProgContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitProg(this);
}

ElzParser::ProgContext* ElzParser::prog() {
  ProgContext *_localctx = _tracker.createInstance<ProgContext>(_ctx, getState());
  enterRule(_localctx, 0, ElzParser::RuleProg);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(65);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << ElzParser::T__0)
      | (1ULL << ElzParser::T__12)
      | (1ULL << ElzParser::T__15)
      | (1ULL << ElzParser::T__17)
      | (1ULL << ElzParser::T__18)
      | (1ULL << ElzParser::T__19))) != 0)) {
      setState(64);
      topStatList();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TopStatListContext ------------------------------------------------------------------

ElzParser::TopStatListContext::TopStatListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::TopStatContext *> ElzParser::TopStatListContext::topStat() {
  return getRuleContexts<ElzParser::TopStatContext>();
}

ElzParser::TopStatContext* ElzParser::TopStatListContext::topStat(size_t i) {
  return getRuleContext<ElzParser::TopStatContext>(i);
}


size_t ElzParser::TopStatListContext::getRuleIndex() const {
  return ElzParser::RuleTopStatList;
}

void ElzParser::TopStatListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTopStatList(this);
}

void ElzParser::TopStatListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTopStatList(this);
}

ElzParser::TopStatListContext* ElzParser::topStatList() {
  TopStatListContext *_localctx = _tracker.createInstance<TopStatListContext>(_ctx, getState());
  enterRule(_localctx, 2, ElzParser::RuleTopStatList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(68); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(67);
      topStat();
      setState(70); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << ElzParser::T__0)
      | (1ULL << ElzParser::T__12)
      | (1ULL << ElzParser::T__15)
      | (1ULL << ElzParser::T__17)
      | (1ULL << ElzParser::T__18)
      | (1ULL << ElzParser::T__19))) != 0));
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TopStatContext ------------------------------------------------------------------

ElzParser::TopStatContext::TopStatContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

ElzParser::FnDefineContext* ElzParser::TopStatContext::fnDefine() {
  return getRuleContext<ElzParser::FnDefineContext>(0);
}

ElzParser::VarDefineContext* ElzParser::TopStatContext::varDefine() {
  return getRuleContext<ElzParser::VarDefineContext>(0);
}

ElzParser::TypeDefineContext* ElzParser::TopStatContext::typeDefine() {
  return getRuleContext<ElzParser::TypeDefineContext>(0);
}

ElzParser::ImplBlockContext* ElzParser::TopStatContext::implBlock() {
  return getRuleContext<ElzParser::ImplBlockContext>(0);
}

ElzParser::TraitDefineContext* ElzParser::TopStatContext::traitDefine() {
  return getRuleContext<ElzParser::TraitDefineContext>(0);
}

ElzParser::ImportStatContext* ElzParser::TopStatContext::importStat() {
  return getRuleContext<ElzParser::ImportStatContext>(0);
}


size_t ElzParser::TopStatContext::getRuleIndex() const {
  return ElzParser::RuleTopStat;
}

void ElzParser::TopStatContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTopStat(this);
}

void ElzParser::TopStatContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTopStat(this);
}

ElzParser::TopStatContext* ElzParser::topStat() {
  TopStatContext *_localctx = _tracker.createInstance<TopStatContext>(_ctx, getState());
  enterRule(_localctx, 4, ElzParser::RuleTopStat);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(78);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case ElzParser::T__17: {
        enterOuterAlt(_localctx, 1);
        setState(72);
        fnDefine();
        break;
      }

      case ElzParser::T__15: {
        enterOuterAlt(_localctx, 2);
        setState(73);
        varDefine();
        break;
      }

      case ElzParser::T__18: {
        enterOuterAlt(_localctx, 3);
        setState(74);
        typeDefine();
        break;
      }

      case ElzParser::T__12: {
        enterOuterAlt(_localctx, 4);
        setState(75);
        implBlock();
        break;
      }

      case ElzParser::T__19: {
        enterOuterAlt(_localctx, 5);
        setState(76);
        traitDefine();
        break;
      }

      case ElzParser::T__0: {
        enterOuterAlt(_localctx, 6);
        setState(77);
        importStat();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ImportStatContext ------------------------------------------------------------------

ElzParser::ImportStatContext::ImportStatContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::ImportStatContext::ID() {
  return getToken(ElzParser::ID, 0);
}


size_t ElzParser::ImportStatContext::getRuleIndex() const {
  return ElzParser::RuleImportStat;
}

void ElzParser::ImportStatContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterImportStat(this);
}

void ElzParser::ImportStatContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitImportStat(this);
}

ElzParser::ImportStatContext* ElzParser::importStat() {
  ImportStatContext *_localctx = _tracker.createInstance<ImportStatContext>(_ctx, getState());
  enterRule(_localctx, 6, ElzParser::RuleImportStat);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(80);
    match(ElzParser::T__0);
    setState(81);
    match(ElzParser::ID);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- StatListContext ------------------------------------------------------------------

ElzParser::StatListContext::StatListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::StatContext *> ElzParser::StatListContext::stat() {
  return getRuleContexts<ElzParser::StatContext>();
}

ElzParser::StatContext* ElzParser::StatListContext::stat(size_t i) {
  return getRuleContext<ElzParser::StatContext>(i);
}


size_t ElzParser::StatListContext::getRuleIndex() const {
  return ElzParser::RuleStatList;
}

void ElzParser::StatListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterStatList(this);
}

void ElzParser::StatListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitStatList(this);
}

ElzParser::StatListContext* ElzParser::statList() {
  StatListContext *_localctx = _tracker.createInstance<StatListContext>(_ctx, getState());
  enterRule(_localctx, 8, ElzParser::RuleStatList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(84); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(83);
      stat();
      setState(86); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << ElzParser::T__1)
      | (1ULL << ElzParser::T__2)
      | (1ULL << ElzParser::T__5)
      | (1ULL << ElzParser::T__15)
      | (1ULL << ElzParser::ID))) != 0));
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- StatContext ------------------------------------------------------------------

ElzParser::StatContext::StatContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

ElzParser::VarDefineContext* ElzParser::StatContext::varDefine() {
  return getRuleContext<ElzParser::VarDefineContext>(0);
}

ElzParser::LoopStatContext* ElzParser::StatContext::loopStat() {
  return getRuleContext<ElzParser::LoopStatContext>(0);
}

ElzParser::ReturnStatContext* ElzParser::StatContext::returnStat() {
  return getRuleContext<ElzParser::ReturnStatContext>(0);
}

ElzParser::AssignContext* ElzParser::StatContext::assign() {
  return getRuleContext<ElzParser::AssignContext>(0);
}

ElzParser::ExprStatContext* ElzParser::StatContext::exprStat() {
  return getRuleContext<ElzParser::ExprStatContext>(0);
}


size_t ElzParser::StatContext::getRuleIndex() const {
  return ElzParser::RuleStat;
}

void ElzParser::StatContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterStat(this);
}

void ElzParser::StatContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitStat(this);
}

ElzParser::StatContext* ElzParser::stat() {
  StatContext *_localctx = _tracker.createInstance<StatContext>(_ctx, getState());
  enterRule(_localctx, 10, ElzParser::RuleStat);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(93);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 4, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(88);
      varDefine();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(89);
      loopStat();
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(90);
      returnStat();
      break;
    }

    case 4: {
      enterOuterAlt(_localctx, 4);
      setState(91);
      assign();
      break;
    }

    case 5: {
      enterOuterAlt(_localctx, 5);
      setState(92);
      exprStat();
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ReturnStatContext ------------------------------------------------------------------

ElzParser::ReturnStatContext::ReturnStatContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

ElzParser::ExprContext* ElzParser::ReturnStatContext::expr() {
  return getRuleContext<ElzParser::ExprContext>(0);
}


size_t ElzParser::ReturnStatContext::getRuleIndex() const {
  return ElzParser::RuleReturnStat;
}

void ElzParser::ReturnStatContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterReturnStat(this);
}

void ElzParser::ReturnStatContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitReturnStat(this);
}

ElzParser::ReturnStatContext* ElzParser::returnStat() {
  ReturnStatContext *_localctx = _tracker.createInstance<ReturnStatContext>(_ctx, getState());
  enterRule(_localctx, 12, ElzParser::RuleReturnStat);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(95);
    match(ElzParser::T__1);
    setState(96);
    expr(0);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- LoopStatContext ------------------------------------------------------------------

ElzParser::LoopStatContext::LoopStatContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

ElzParser::StatListContext* ElzParser::LoopStatContext::statList() {
  return getRuleContext<ElzParser::StatListContext>(0);
}


size_t ElzParser::LoopStatContext::getRuleIndex() const {
  return ElzParser::RuleLoopStat;
}

void ElzParser::LoopStatContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterLoopStat(this);
}

void ElzParser::LoopStatContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitLoopStat(this);
}

ElzParser::LoopStatContext* ElzParser::loopStat() {
  LoopStatContext *_localctx = _tracker.createInstance<LoopStatContext>(_ctx, getState());
  enterRule(_localctx, 14, ElzParser::RuleLoopStat);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(98);
    match(ElzParser::T__2);
    setState(99);
    match(ElzParser::T__3);
    setState(101);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << ElzParser::T__1)
      | (1ULL << ElzParser::T__2)
      | (1ULL << ElzParser::T__5)
      | (1ULL << ElzParser::T__15)
      | (1ULL << ElzParser::ID))) != 0)) {
      setState(100);
      statList();
    }
    setState(103);
    match(ElzParser::T__4);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExprStatContext ------------------------------------------------------------------

ElzParser::ExprStatContext::ExprStatContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

ElzParser::MatchRuleContext* ElzParser::ExprStatContext::matchRule() {
  return getRuleContext<ElzParser::MatchRuleContext>(0);
}

ElzParser::FnCallContext* ElzParser::ExprStatContext::fnCall() {
  return getRuleContext<ElzParser::FnCallContext>(0);
}


size_t ElzParser::ExprStatContext::getRuleIndex() const {
  return ElzParser::RuleExprStat;
}

void ElzParser::ExprStatContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExprStat(this);
}

void ElzParser::ExprStatContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExprStat(this);
}

ElzParser::ExprStatContext* ElzParser::exprStat() {
  ExprStatContext *_localctx = _tracker.createInstance<ExprStatContext>(_ctx, getState());
  enterRule(_localctx, 16, ElzParser::RuleExprStat);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(107);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case ElzParser::T__5: {
        enterOuterAlt(_localctx, 1);
        setState(105);
        matchRule();
        break;
      }

      case ElzParser::ID: {
        enterOuterAlt(_localctx, 2);
        setState(106);
        fnCall();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- MatchRuleContext ------------------------------------------------------------------

ElzParser::MatchRuleContext::MatchRuleContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::ExprContext *> ElzParser::MatchRuleContext::expr() {
  return getRuleContexts<ElzParser::ExprContext>();
}

ElzParser::ExprContext* ElzParser::MatchRuleContext::expr(size_t i) {
  return getRuleContext<ElzParser::ExprContext>(i);
}

std::vector<ElzParser::StatContext *> ElzParser::MatchRuleContext::stat() {
  return getRuleContexts<ElzParser::StatContext>();
}

ElzParser::StatContext* ElzParser::MatchRuleContext::stat(size_t i) {
  return getRuleContext<ElzParser::StatContext>(i);
}


size_t ElzParser::MatchRuleContext::getRuleIndex() const {
  return ElzParser::RuleMatchRule;
}

void ElzParser::MatchRuleContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterMatchRule(this);
}

void ElzParser::MatchRuleContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitMatchRule(this);
}

ElzParser::MatchRuleContext* ElzParser::matchRule() {
  MatchRuleContext *_localctx = _tracker.createInstance<MatchRuleContext>(_ctx, getState());
  enterRule(_localctx, 18, ElzParser::RuleMatchRule);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(109);
    match(ElzParser::T__5);
    setState(110);
    expr(0);
    setState(111);
    match(ElzParser::T__3);
    setState(112);
    expr(0);
    setState(113);
    match(ElzParser::T__6);
    setState(114);
    stat();
    setState(122);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 7, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        setState(115);
        match(ElzParser::T__7);
        setState(116);
        expr(0);
        setState(117);
        match(ElzParser::T__6);
        setState(118);
        stat(); 
      }
      setState(124);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 7, _ctx);
    }
    setState(126);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__7) {
      setState(125);
      match(ElzParser::T__7);
    }
    setState(128);
    match(ElzParser::T__4);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- AssignContext ------------------------------------------------------------------

ElzParser::AssignContext::AssignContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::AssignContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::ExprContext* ElzParser::AssignContext::expr() {
  return getRuleContext<ElzParser::ExprContext>(0);
}


size_t ElzParser::AssignContext::getRuleIndex() const {
  return ElzParser::RuleAssign;
}

void ElzParser::AssignContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAssign(this);
}

void ElzParser::AssignContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAssign(this);
}

ElzParser::AssignContext* ElzParser::assign() {
  AssignContext *_localctx = _tracker.createInstance<AssignContext>(_ctx, getState());
  enterRule(_localctx, 20, ElzParser::RuleAssign);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(130);
    match(ElzParser::ID);
    setState(131);
    match(ElzParser::T__8);
    setState(132);
    expr(0);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExprListContext ------------------------------------------------------------------

ElzParser::ExprListContext::ExprListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::ExprContext *> ElzParser::ExprListContext::expr() {
  return getRuleContexts<ElzParser::ExprContext>();
}

ElzParser::ExprContext* ElzParser::ExprListContext::expr(size_t i) {
  return getRuleContext<ElzParser::ExprContext>(i);
}


size_t ElzParser::ExprListContext::getRuleIndex() const {
  return ElzParser::RuleExprList;
}

void ElzParser::ExprListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExprList(this);
}

void ElzParser::ExprListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExprList(this);
}

ElzParser::ExprListContext* ElzParser::exprList() {
  ExprListContext *_localctx = _tracker.createInstance<ExprListContext>(_ctx, getState());
  enterRule(_localctx, 22, ElzParser::RuleExprList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(134);
    expr(0);
    setState(139);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while (_la == ElzParser::T__7) {
      setState(135);
      match(ElzParser::T__7);
      setState(136);
      expr(0);
      setState(141);
      _errHandler->sync(this);
      _la = _input->LA(1);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FnCallContext ------------------------------------------------------------------

ElzParser::FnCallContext::FnCallContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::FnCallContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::ExprListContext* ElzParser::FnCallContext::exprList() {
  return getRuleContext<ElzParser::ExprListContext>(0);
}


size_t ElzParser::FnCallContext::getRuleIndex() const {
  return ElzParser::RuleFnCall;
}

void ElzParser::FnCallContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFnCall(this);
}

void ElzParser::FnCallContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFnCall(this);
}

ElzParser::FnCallContext* ElzParser::fnCall() {
  FnCallContext *_localctx = _tracker.createInstance<FnCallContext>(_ctx, getState());
  enterRule(_localctx, 24, ElzParser::RuleFnCall);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(142);
    match(ElzParser::ID);
    setState(143);
    match(ElzParser::T__9);
    setState(145);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << ElzParser::T__5)
      | (1ULL << ElzParser::T__9)
      | (1ULL << ElzParser::ID)
      | (1ULL << ElzParser::NUM)
      | (1ULL << ElzParser::STRING))) != 0)) {
      setState(144);
      exprList();
    }
    setState(147);
    match(ElzParser::T__10);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TypePassContext ------------------------------------------------------------------

ElzParser::TypePassContext::TypePassContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::TypePassContext::ID() {
  return getToken(ElzParser::ID, 0);
}


size_t ElzParser::TypePassContext::getRuleIndex() const {
  return ElzParser::RuleTypePass;
}

void ElzParser::TypePassContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTypePass(this);
}

void ElzParser::TypePassContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTypePass(this);
}

ElzParser::TypePassContext* ElzParser::typePass() {
  TypePassContext *_localctx = _tracker.createInstance<TypePassContext>(_ctx, getState());
  enterRule(_localctx, 26, ElzParser::RuleTypePass);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(149);
    match(ElzParser::ID);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TypeListContext ------------------------------------------------------------------

ElzParser::TypeListContext::TypeListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::TypePassContext *> ElzParser::TypeListContext::typePass() {
  return getRuleContexts<ElzParser::TypePassContext>();
}

ElzParser::TypePassContext* ElzParser::TypeListContext::typePass(size_t i) {
  return getRuleContext<ElzParser::TypePassContext>(i);
}


size_t ElzParser::TypeListContext::getRuleIndex() const {
  return ElzParser::RuleTypeList;
}

void ElzParser::TypeListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTypeList(this);
}

void ElzParser::TypeListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTypeList(this);
}

ElzParser::TypeListContext* ElzParser::typeList() {
  TypeListContext *_localctx = _tracker.createInstance<TypeListContext>(_ctx, getState());
  enterRule(_localctx, 28, ElzParser::RuleTypeList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(151);
    typePass();
    setState(156);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while (_la == ElzParser::T__7) {
      setState(152);
      match(ElzParser::T__7);
      setState(153);
      typePass();
      setState(158);
      _errHandler->sync(this);
      _la = _input->LA(1);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- MethodListContext ------------------------------------------------------------------

ElzParser::MethodListContext::MethodListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::MethodContext *> ElzParser::MethodListContext::method() {
  return getRuleContexts<ElzParser::MethodContext>();
}

ElzParser::MethodContext* ElzParser::MethodListContext::method(size_t i) {
  return getRuleContext<ElzParser::MethodContext>(i);
}


size_t ElzParser::MethodListContext::getRuleIndex() const {
  return ElzParser::RuleMethodList;
}

void ElzParser::MethodListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterMethodList(this);
}

void ElzParser::MethodListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitMethodList(this);
}

ElzParser::MethodListContext* ElzParser::methodList() {
  MethodListContext *_localctx = _tracker.createInstance<MethodListContext>(_ctx, getState());
  enterRule(_localctx, 30, ElzParser::RuleMethodList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(160); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(159);
      method();
      setState(162); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while (_la == ElzParser::T__14

    || _la == ElzParser::ID);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- MethodContext ------------------------------------------------------------------

ElzParser::MethodContext::MethodContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::MethodContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::ExportorContext* ElzParser::MethodContext::exportor() {
  return getRuleContext<ElzParser::ExportorContext>(0);
}

ElzParser::ParamListContext* ElzParser::MethodContext::paramList() {
  return getRuleContext<ElzParser::ParamListContext>(0);
}

ElzParser::TypePassContext* ElzParser::MethodContext::typePass() {
  return getRuleContext<ElzParser::TypePassContext>(0);
}

ElzParser::StatListContext* ElzParser::MethodContext::statList() {
  return getRuleContext<ElzParser::StatListContext>(0);
}


size_t ElzParser::MethodContext::getRuleIndex() const {
  return ElzParser::RuleMethod;
}

void ElzParser::MethodContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterMethod(this);
}

void ElzParser::MethodContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitMethod(this);
}

ElzParser::MethodContext* ElzParser::method() {
  MethodContext *_localctx = _tracker.createInstance<MethodContext>(_ctx, getState());
  enterRule(_localctx, 32, ElzParser::RuleMethod);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(165);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__14) {
      setState(164);
      exportor();
    }
    setState(167);
    match(ElzParser::ID);
    setState(168);
    match(ElzParser::T__9);
    setState(170);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::ID) {
      setState(169);
      paramList();
    }
    setState(172);
    match(ElzParser::T__10);
    setState(175);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__11) {
      setState(173);
      match(ElzParser::T__11);
      setState(174);
      typePass();
    }
    setState(177);
    match(ElzParser::T__3);
    setState(179);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << ElzParser::T__1)
      | (1ULL << ElzParser::T__2)
      | (1ULL << ElzParser::T__5)
      | (1ULL << ElzParser::T__15)
      | (1ULL << ElzParser::ID))) != 0)) {
      setState(178);
      statList();
    }
    setState(181);
    match(ElzParser::T__4);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ImplBlockContext ------------------------------------------------------------------

ElzParser::ImplBlockContext::ImplBlockContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::ImplBlockContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::TypeListContext* ElzParser::ImplBlockContext::typeList() {
  return getRuleContext<ElzParser::TypeListContext>(0);
}

ElzParser::MethodListContext* ElzParser::ImplBlockContext::methodList() {
  return getRuleContext<ElzParser::MethodListContext>(0);
}


size_t ElzParser::ImplBlockContext::getRuleIndex() const {
  return ElzParser::RuleImplBlock;
}

void ElzParser::ImplBlockContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterImplBlock(this);
}

void ElzParser::ImplBlockContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitImplBlock(this);
}

ElzParser::ImplBlockContext* ElzParser::implBlock() {
  ImplBlockContext *_localctx = _tracker.createInstance<ImplBlockContext>(_ctx, getState());
  enterRule(_localctx, 34, ElzParser::RuleImplBlock);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(183);
    match(ElzParser::T__12);
    setState(184);
    match(ElzParser::ID);
    setState(187);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__13) {
      setState(185);
      match(ElzParser::T__13);
      setState(186);
      typeList();
    }
    setState(189);
    match(ElzParser::T__3);
    setState(191);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__14

    || _la == ElzParser::ID) {
      setState(190);
      methodList();
    }
    setState(193);
    match(ElzParser::T__4);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExportorContext ------------------------------------------------------------------

ElzParser::ExportorContext::ExportorContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}


size_t ElzParser::ExportorContext::getRuleIndex() const {
  return ElzParser::RuleExportor;
}

void ElzParser::ExportorContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExportor(this);
}

void ElzParser::ExportorContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExportor(this);
}

ElzParser::ExportorContext* ElzParser::exportor() {
  ExportorContext *_localctx = _tracker.createInstance<ExportorContext>(_ctx, getState());
  enterRule(_localctx, 36, ElzParser::RuleExportor);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(195);
    match(ElzParser::T__14);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- DefineContext ------------------------------------------------------------------

ElzParser::DefineContext::DefineContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::DefineContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::ExprContext* ElzParser::DefineContext::expr() {
  return getRuleContext<ElzParser::ExprContext>(0);
}

ElzParser::ExportorContext* ElzParser::DefineContext::exportor() {
  return getRuleContext<ElzParser::ExportorContext>(0);
}

ElzParser::TypePassContext* ElzParser::DefineContext::typePass() {
  return getRuleContext<ElzParser::TypePassContext>(0);
}


size_t ElzParser::DefineContext::getRuleIndex() const {
  return ElzParser::RuleDefine;
}

void ElzParser::DefineContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterDefine(this);
}

void ElzParser::DefineContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitDefine(this);
}

ElzParser::DefineContext* ElzParser::define() {
  DefineContext *_localctx = _tracker.createInstance<DefineContext>(_ctx, getState());
  enterRule(_localctx, 38, ElzParser::RuleDefine);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(198);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__14) {
      setState(197);
      exportor();
    }
    setState(200);
    match(ElzParser::ID);
    setState(203);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__13) {
      setState(201);
      match(ElzParser::T__13);
      setState(202);
      typePass();
    }
    setState(205);
    match(ElzParser::T__8);
    setState(206);
    expr(0);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- VarDefineContext ------------------------------------------------------------------

ElzParser::VarDefineContext::VarDefineContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::DefineContext *> ElzParser::VarDefineContext::define() {
  return getRuleContexts<ElzParser::DefineContext>();
}

ElzParser::DefineContext* ElzParser::VarDefineContext::define(size_t i) {
  return getRuleContext<ElzParser::DefineContext>(i);
}


size_t ElzParser::VarDefineContext::getRuleIndex() const {
  return ElzParser::RuleVarDefine;
}

void ElzParser::VarDefineContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterVarDefine(this);
}

void ElzParser::VarDefineContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitVarDefine(this);
}

ElzParser::VarDefineContext* ElzParser::varDefine() {
  VarDefineContext *_localctx = _tracker.createInstance<VarDefineContext>(_ctx, getState());
  enterRule(_localctx, 40, ElzParser::RuleVarDefine);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(208);
    match(ElzParser::T__15);
    setState(210);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__16) {
      setState(209);
      dynamic_cast<VarDefineContext *>(_localctx)->mut = match(ElzParser::T__16);
    }
    setState(212);
    define();
    setState(217);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 22, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        setState(213);
        match(ElzParser::T__7);
        setState(214);
        define(); 
      }
      setState(219);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 22, _ctx);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ParamListContext ------------------------------------------------------------------

ElzParser::ParamListContext::ParamListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::ParamContext *> ElzParser::ParamListContext::param() {
  return getRuleContexts<ElzParser::ParamContext>();
}

ElzParser::ParamContext* ElzParser::ParamListContext::param(size_t i) {
  return getRuleContext<ElzParser::ParamContext>(i);
}


size_t ElzParser::ParamListContext::getRuleIndex() const {
  return ElzParser::RuleParamList;
}

void ElzParser::ParamListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterParamList(this);
}

void ElzParser::ParamListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitParamList(this);
}

ElzParser::ParamListContext* ElzParser::paramList() {
  ParamListContext *_localctx = _tracker.createInstance<ParamListContext>(_ctx, getState());
  enterRule(_localctx, 42, ElzParser::RuleParamList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(220);
    param();
    setState(225);
    _errHandler->sync(this);
    _la = _input->LA(1);
    while (_la == ElzParser::T__7) {
      setState(221);
      match(ElzParser::T__7);
      setState(222);
      param();
      setState(227);
      _errHandler->sync(this);
      _la = _input->LA(1);
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ParamContext ------------------------------------------------------------------

ElzParser::ParamContext::ParamContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::ParamContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::TypePassContext* ElzParser::ParamContext::typePass() {
  return getRuleContext<ElzParser::TypePassContext>(0);
}


size_t ElzParser::ParamContext::getRuleIndex() const {
  return ElzParser::RuleParam;
}

void ElzParser::ParamContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterParam(this);
}

void ElzParser::ParamContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitParam(this);
}

ElzParser::ParamContext* ElzParser::param() {
  ParamContext *_localctx = _tracker.createInstance<ParamContext>(_ctx, getState());
  enterRule(_localctx, 44, ElzParser::RuleParam);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(228);
    match(ElzParser::ID);
    setState(229);
    match(ElzParser::T__13);
    setState(230);
    typePass();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- FnDefineContext ------------------------------------------------------------------

ElzParser::FnDefineContext::FnDefineContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::FnDefineContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::ExportorContext* ElzParser::FnDefineContext::exportor() {
  return getRuleContext<ElzParser::ExportorContext>(0);
}

ElzParser::ParamListContext* ElzParser::FnDefineContext::paramList() {
  return getRuleContext<ElzParser::ParamListContext>(0);
}

ElzParser::TypePassContext* ElzParser::FnDefineContext::typePass() {
  return getRuleContext<ElzParser::TypePassContext>(0);
}

ElzParser::StatListContext* ElzParser::FnDefineContext::statList() {
  return getRuleContext<ElzParser::StatListContext>(0);
}


size_t ElzParser::FnDefineContext::getRuleIndex() const {
  return ElzParser::RuleFnDefine;
}

void ElzParser::FnDefineContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFnDefine(this);
}

void ElzParser::FnDefineContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFnDefine(this);
}

ElzParser::FnDefineContext* ElzParser::fnDefine() {
  FnDefineContext *_localctx = _tracker.createInstance<FnDefineContext>(_ctx, getState());
  enterRule(_localctx, 46, ElzParser::RuleFnDefine);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(232);
    match(ElzParser::T__17);
    setState(234);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__14) {
      setState(233);
      exportor();
    }
    setState(236);
    match(ElzParser::ID);
    setState(237);
    match(ElzParser::T__9);
    setState(239);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::ID) {
      setState(238);
      paramList();
    }
    setState(241);
    match(ElzParser::T__10);
    setState(244);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__11) {
      setState(242);
      match(ElzParser::T__11);
      setState(243);
      typePass();
    }
    setState(246);
    match(ElzParser::T__3);
    setState(248);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if ((((_la & ~ 0x3fULL) == 0) &&
      ((1ULL << _la) & ((1ULL << ElzParser::T__1)
      | (1ULL << ElzParser::T__2)
      | (1ULL << ElzParser::T__5)
      | (1ULL << ElzParser::T__15)
      | (1ULL << ElzParser::ID))) != 0)) {
      setState(247);
      statList();
    }
    setState(250);
    match(ElzParser::T__4);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- AttrListContext ------------------------------------------------------------------

ElzParser::AttrListContext::AttrListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::AttrContext *> ElzParser::AttrListContext::attr() {
  return getRuleContexts<ElzParser::AttrContext>();
}

ElzParser::AttrContext* ElzParser::AttrListContext::attr(size_t i) {
  return getRuleContext<ElzParser::AttrContext>(i);
}


size_t ElzParser::AttrListContext::getRuleIndex() const {
  return ElzParser::RuleAttrList;
}

void ElzParser::AttrListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAttrList(this);
}

void ElzParser::AttrListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAttrList(this);
}

ElzParser::AttrListContext* ElzParser::attrList() {
  AttrListContext *_localctx = _tracker.createInstance<AttrListContext>(_ctx, getState());
  enterRule(_localctx, 48, ElzParser::RuleAttrList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(253); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(252);
      attr();
      setState(255); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while (_la == ElzParser::ID);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- AttrContext ------------------------------------------------------------------

ElzParser::AttrContext::AttrContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::AttrContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::TypePassContext* ElzParser::AttrContext::typePass() {
  return getRuleContext<ElzParser::TypePassContext>(0);
}


size_t ElzParser::AttrContext::getRuleIndex() const {
  return ElzParser::RuleAttr;
}

void ElzParser::AttrContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterAttr(this);
}

void ElzParser::AttrContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitAttr(this);
}

ElzParser::AttrContext* ElzParser::attr() {
  AttrContext *_localctx = _tracker.createInstance<AttrContext>(_ctx, getState());
  enterRule(_localctx, 50, ElzParser::RuleAttr);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(257);
    match(ElzParser::ID);
    setState(258);
    match(ElzParser::T__13);
    setState(259);
    typePass();
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TypeDefineContext ------------------------------------------------------------------

ElzParser::TypeDefineContext::TypeDefineContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::TypeDefineContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::AttrListContext* ElzParser::TypeDefineContext::attrList() {
  return getRuleContext<ElzParser::AttrListContext>(0);
}

ElzParser::ExportorContext* ElzParser::TypeDefineContext::exportor() {
  return getRuleContext<ElzParser::ExportorContext>(0);
}


size_t ElzParser::TypeDefineContext::getRuleIndex() const {
  return ElzParser::RuleTypeDefine;
}

void ElzParser::TypeDefineContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTypeDefine(this);
}

void ElzParser::TypeDefineContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTypeDefine(this);
}

ElzParser::TypeDefineContext* ElzParser::typeDefine() {
  TypeDefineContext *_localctx = _tracker.createInstance<TypeDefineContext>(_ctx, getState());
  enterRule(_localctx, 52, ElzParser::RuleTypeDefine);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(261);
    match(ElzParser::T__18);
    setState(263);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__14) {
      setState(262);
      exportor();
    }
    setState(265);
    match(ElzParser::ID);
    setState(266);
    match(ElzParser::T__9);
    setState(267);
    attrList();
    setState(268);
    match(ElzParser::T__10);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TmethodListContext ------------------------------------------------------------------

ElzParser::TmethodListContext::TmethodListContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::TmethodContext *> ElzParser::TmethodListContext::tmethod() {
  return getRuleContexts<ElzParser::TmethodContext>();
}

ElzParser::TmethodContext* ElzParser::TmethodListContext::tmethod(size_t i) {
  return getRuleContext<ElzParser::TmethodContext>(i);
}


size_t ElzParser::TmethodListContext::getRuleIndex() const {
  return ElzParser::RuleTmethodList;
}

void ElzParser::TmethodListContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTmethodList(this);
}

void ElzParser::TmethodListContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTmethodList(this);
}

ElzParser::TmethodListContext* ElzParser::tmethodList() {
  TmethodListContext *_localctx = _tracker.createInstance<TmethodListContext>(_ctx, getState());
  enterRule(_localctx, 54, ElzParser::RuleTmethodList);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(271); 
    _errHandler->sync(this);
    _la = _input->LA(1);
    do {
      setState(270);
      tmethod();
      setState(273); 
      _errHandler->sync(this);
      _la = _input->LA(1);
    } while (_la == ElzParser::T__14

    || _la == ElzParser::ID);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TmethodContext ------------------------------------------------------------------

ElzParser::TmethodContext::TmethodContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

tree::TerminalNode* ElzParser::TmethodContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::ExportorContext* ElzParser::TmethodContext::exportor() {
  return getRuleContext<ElzParser::ExportorContext>(0);
}

ElzParser::TypeListContext* ElzParser::TmethodContext::typeList() {
  return getRuleContext<ElzParser::TypeListContext>(0);
}

ElzParser::TypePassContext* ElzParser::TmethodContext::typePass() {
  return getRuleContext<ElzParser::TypePassContext>(0);
}


size_t ElzParser::TmethodContext::getRuleIndex() const {
  return ElzParser::RuleTmethod;
}

void ElzParser::TmethodContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTmethod(this);
}

void ElzParser::TmethodContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTmethod(this);
}

ElzParser::TmethodContext* ElzParser::tmethod() {
  TmethodContext *_localctx = _tracker.createInstance<TmethodContext>(_ctx, getState());
  enterRule(_localctx, 56, ElzParser::RuleTmethod);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(276);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__14) {
      setState(275);
      exportor();
    }
    setState(278);
    match(ElzParser::ID);
    setState(279);
    match(ElzParser::T__9);
    setState(281);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::ID) {
      setState(280);
      typeList();
    }
    setState(283);
    match(ElzParser::T__10);
    setState(286);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__11) {
      setState(284);
      match(ElzParser::T__11);
      setState(285);
      typePass();
    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- TraitDefineContext ------------------------------------------------------------------

ElzParser::TraitDefineContext::TraitDefineContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

ElzParser::ExportorContext* ElzParser::TraitDefineContext::exportor() {
  return getRuleContext<ElzParser::ExportorContext>(0);
}

tree::TerminalNode* ElzParser::TraitDefineContext::ID() {
  return getToken(ElzParser::ID, 0);
}

ElzParser::TmethodListContext* ElzParser::TraitDefineContext::tmethodList() {
  return getRuleContext<ElzParser::TmethodListContext>(0);
}


size_t ElzParser::TraitDefineContext::getRuleIndex() const {
  return ElzParser::RuleTraitDefine;
}

void ElzParser::TraitDefineContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterTraitDefine(this);
}

void ElzParser::TraitDefineContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitTraitDefine(this);
}

ElzParser::TraitDefineContext* ElzParser::traitDefine() {
  TraitDefineContext *_localctx = _tracker.createInstance<TraitDefineContext>(_ctx, getState());
  enterRule(_localctx, 58, ElzParser::RuleTraitDefine);
  size_t _la = 0;

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    enterOuterAlt(_localctx, 1);
    setState(288);
    match(ElzParser::T__19);
    setState(289);
    exportor();
    setState(290);
    match(ElzParser::ID);
    setState(291);
    match(ElzParser::T__3);
    setState(293);
    _errHandler->sync(this);

    _la = _input->LA(1);
    if (_la == ElzParser::T__14

    || _la == ElzParser::ID) {
      setState(292);
      tmethodList();
    }
    setState(295);
    match(ElzParser::T__4);
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

//----------------- ExprContext ------------------------------------------------------------------

ElzParser::ExprContext::ExprContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

std::vector<ElzParser::ExprContext *> ElzParser::ExprContext::expr() {
  return getRuleContexts<ElzParser::ExprContext>();
}

ElzParser::ExprContext* ElzParser::ExprContext::expr(size_t i) {
  return getRuleContext<ElzParser::ExprContext>(i);
}

ElzParser::FactorContext* ElzParser::ExprContext::factor() {
  return getRuleContext<ElzParser::FactorContext>(0);
}


size_t ElzParser::ExprContext::getRuleIndex() const {
  return ElzParser::RuleExpr;
}

void ElzParser::ExprContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterExpr(this);
}

void ElzParser::ExprContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitExpr(this);
}


ElzParser::ExprContext* ElzParser::expr() {
   return expr(0);
}

ElzParser::ExprContext* ElzParser::expr(int precedence) {
  ParserRuleContext *parentContext = _ctx;
  size_t parentState = getState();
  ElzParser::ExprContext *_localctx = _tracker.createInstance<ExprContext>(_ctx, parentState);
  ElzParser::ExprContext *previousContext = _localctx;
  size_t startState = 60;
  enterRecursionRule(_localctx, 60, ElzParser::RuleExpr, precedence);

    size_t _la = 0;

  auto onExit = finally([=] {
    unrollRecursionContexts(parentContext);
  });
  try {
    size_t alt;
    enterOuterAlt(_localctx, 1);
    setState(303);
    _errHandler->sync(this);
    switch (_input->LA(1)) {
      case ElzParser::T__9: {
        setState(298);
        match(ElzParser::T__9);
        setState(299);
        expr(0);
        setState(300);
        match(ElzParser::T__10);
        break;
      }

      case ElzParser::T__5:
      case ElzParser::ID:
      case ElzParser::NUM:
      case ElzParser::STRING: {
        setState(302);
        factor();
        break;
      }

    default:
      throw NoViableAltException(this);
    }
    _ctx->stop = _input->LT(-1);
    setState(316);
    _errHandler->sync(this);
    alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 37, _ctx);
    while (alt != 2 && alt != atn::ATN::INVALID_ALT_NUMBER) {
      if (alt == 1) {
        if (!_parseListeners.empty())
          triggerExitRuleEvent();
        previousContext = _localctx;
        setState(314);
        _errHandler->sync(this);
        switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 36, _ctx)) {
        case 1: {
          _localctx = _tracker.createInstance<ExprContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpr);
          setState(305);

          if (!(precpred(_ctx, 4))) throw FailedPredicateException(this, "precpred(_ctx, 4)");
          setState(306);
          dynamic_cast<ExprContext *>(_localctx)->op = match(ElzParser::T__20);
          setState(307);
          expr(5);
          break;
        }

        case 2: {
          _localctx = _tracker.createInstance<ExprContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpr);
          setState(308);

          if (!(precpred(_ctx, 3))) throw FailedPredicateException(this, "precpred(_ctx, 3)");
          setState(309);
          dynamic_cast<ExprContext *>(_localctx)->op = _input->LT(1);
          _la = _input->LA(1);
          if (!(_la == ElzParser::T__21

          || _la == ElzParser::T__22)) {
            dynamic_cast<ExprContext *>(_localctx)->op = _errHandler->recoverInline(this);
          }
          else {
            _errHandler->reportMatch(this);
            consume();
          }
          setState(310);
          expr(4);
          break;
        }

        case 3: {
          _localctx = _tracker.createInstance<ExprContext>(parentContext, parentState);
          pushNewRecursionContext(_localctx, startState, RuleExpr);
          setState(311);

          if (!(precpred(_ctx, 2))) throw FailedPredicateException(this, "precpred(_ctx, 2)");
          setState(312);
          dynamic_cast<ExprContext *>(_localctx)->op = _input->LT(1);
          _la = _input->LA(1);
          if (!(_la == ElzParser::T__14

          || _la == ElzParser::T__23)) {
            dynamic_cast<ExprContext *>(_localctx)->op = _errHandler->recoverInline(this);
          }
          else {
            _errHandler->reportMatch(this);
            consume();
          }
          setState(313);
          expr(3);
          break;
        }

        } 
      }
      setState(318);
      _errHandler->sync(this);
      alt = getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 37, _ctx);
    }
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }
  return _localctx;
}

//----------------- FactorContext ------------------------------------------------------------------

ElzParser::FactorContext::FactorContext(ParserRuleContext *parent, size_t invokingState)
  : ParserRuleContext(parent, invokingState) {
}

ElzParser::ExprStatContext* ElzParser::FactorContext::exprStat() {
  return getRuleContext<ElzParser::ExprStatContext>(0);
}

tree::TerminalNode* ElzParser::FactorContext::NUM() {
  return getToken(ElzParser::NUM, 0);
}

tree::TerminalNode* ElzParser::FactorContext::ID() {
  return getToken(ElzParser::ID, 0);
}

tree::TerminalNode* ElzParser::FactorContext::STRING() {
  return getToken(ElzParser::STRING, 0);
}


size_t ElzParser::FactorContext::getRuleIndex() const {
  return ElzParser::RuleFactor;
}

void ElzParser::FactorContext::enterRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->enterFactor(this);
}

void ElzParser::FactorContext::exitRule(tree::ParseTreeListener *listener) {
  auto parserListener = dynamic_cast<ElzListener *>(listener);
  if (parserListener != nullptr)
    parserListener->exitFactor(this);
}

ElzParser::FactorContext* ElzParser::factor() {
  FactorContext *_localctx = _tracker.createInstance<FactorContext>(_ctx, getState());
  enterRule(_localctx, 62, ElzParser::RuleFactor);

  auto onExit = finally([=] {
    exitRule();
  });
  try {
    setState(323);
    _errHandler->sync(this);
    switch (getInterpreter<atn::ParserATNSimulator>()->adaptivePredict(_input, 38, _ctx)) {
    case 1: {
      enterOuterAlt(_localctx, 1);
      setState(319);
      exprStat();
      break;
    }

    case 2: {
      enterOuterAlt(_localctx, 2);
      setState(320);
      match(ElzParser::NUM);
      break;
    }

    case 3: {
      enterOuterAlt(_localctx, 3);
      setState(321);
      match(ElzParser::ID);
      break;
    }

    case 4: {
      enterOuterAlt(_localctx, 4);
      setState(322);
      match(ElzParser::STRING);
      break;
    }

    }
   
  }
  catch (RecognitionException &e) {
    _errHandler->reportError(this, e);
    _localctx->exception = std::current_exception();
    _errHandler->recover(this, _localctx->exception);
  }

  return _localctx;
}

bool ElzParser::sempred(RuleContext *context, size_t ruleIndex, size_t predicateIndex) {
  switch (ruleIndex) {
    case 30: return exprSempred(dynamic_cast<ExprContext *>(context), predicateIndex);

  default:
    break;
  }
  return true;
}

bool ElzParser::exprSempred(ExprContext *_localctx, size_t predicateIndex) {
  switch (predicateIndex) {
    case 0: return precpred(_ctx, 4);
    case 1: return precpred(_ctx, 3);
    case 2: return precpred(_ctx, 2);

  default:
    break;
  }
  return true;
}

// Static vars and initialization.
std::vector<dfa::DFA> ElzParser::_decisionToDFA;
atn::PredictionContextCache ElzParser::_sharedContextCache;

// We own the ATN which in turn owns the ATN states.
atn::ATN ElzParser::_atn;
std::vector<uint16_t> ElzParser::_serializedATN;

std::vector<std::string> ElzParser::_ruleNames = {
  "prog", "topStatList", "topStat", "importStat", "statList", "stat", "returnStat", 
  "loopStat", "exprStat", "matchRule", "assign", "exprList", "fnCall", "typePass", 
  "typeList", "methodList", "method", "implBlock", "exportor", "define", 
  "varDefine", "paramList", "param", "fnDefine", "attrList", "attr", "typeDefine", 
  "tmethodList", "tmethod", "traitDefine", "expr", "factor"
};

std::vector<std::string> ElzParser::_literalNames = {
  "", "'import'", "'return'", "'loop'", "'{'", "'}'", "'match'", "'=>'", 
  "','", "'='", "'('", "')'", "'->'", "'impl'", "':'", "'+'", "'let'", "'mut'", 
  "'fn'", "'type'", "'trait'", "'^'", "'*'", "'/'", "'-'"
};

std::vector<std::string> ElzParser::_symbolicNames = {
  "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
  "", "", "", "", "", "", "", "WS", "COMMENT", "ID", "NUM", "STRING"
};

dfa::Vocabulary ElzParser::_vocabulary(_literalNames, _symbolicNames);

std::vector<std::string> ElzParser::_tokenNames;

ElzParser::Initializer::Initializer() {
	for (size_t i = 0; i < _symbolicNames.size(); ++i) {
		std::string name = _vocabulary.getLiteralName(i);
		if (name.empty()) {
			name = _vocabulary.getSymbolicName(i);
		}

		if (name.empty()) {
			_tokenNames.push_back("<INVALID>");
		} else {
      _tokenNames.push_back(name);
    }
	}

  _serializedATN = {
    0x3, 0x608b, 0xa72a, 0x8133, 0xb9ed, 0x417c, 0x3be7, 0x7786, 0x5964, 
    0x3, 0x1f, 0x148, 0x4, 0x2, 0x9, 0x2, 0x4, 0x3, 0x9, 0x3, 0x4, 0x4, 
    0x9, 0x4, 0x4, 0x5, 0x9, 0x5, 0x4, 0x6, 0x9, 0x6, 0x4, 0x7, 0x9, 0x7, 
    0x4, 0x8, 0x9, 0x8, 0x4, 0x9, 0x9, 0x9, 0x4, 0xa, 0x9, 0xa, 0x4, 0xb, 
    0x9, 0xb, 0x4, 0xc, 0x9, 0xc, 0x4, 0xd, 0x9, 0xd, 0x4, 0xe, 0x9, 0xe, 
    0x4, 0xf, 0x9, 0xf, 0x4, 0x10, 0x9, 0x10, 0x4, 0x11, 0x9, 0x11, 0x4, 
    0x12, 0x9, 0x12, 0x4, 0x13, 0x9, 0x13, 0x4, 0x14, 0x9, 0x14, 0x4, 0x15, 
    0x9, 0x15, 0x4, 0x16, 0x9, 0x16, 0x4, 0x17, 0x9, 0x17, 0x4, 0x18, 0x9, 
    0x18, 0x4, 0x19, 0x9, 0x19, 0x4, 0x1a, 0x9, 0x1a, 0x4, 0x1b, 0x9, 0x1b, 
    0x4, 0x1c, 0x9, 0x1c, 0x4, 0x1d, 0x9, 0x1d, 0x4, 0x1e, 0x9, 0x1e, 0x4, 
    0x1f, 0x9, 0x1f, 0x4, 0x20, 0x9, 0x20, 0x4, 0x21, 0x9, 0x21, 0x3, 0x2, 
    0x5, 0x2, 0x44, 0xa, 0x2, 0x3, 0x3, 0x6, 0x3, 0x47, 0xa, 0x3, 0xd, 0x3, 
    0xe, 0x3, 0x48, 0x3, 0x4, 0x3, 0x4, 0x3, 0x4, 0x3, 0x4, 0x3, 0x4, 0x3, 
    0x4, 0x5, 0x4, 0x51, 0xa, 0x4, 0x3, 0x5, 0x3, 0x5, 0x3, 0x5, 0x3, 0x6, 
    0x6, 0x6, 0x57, 0xa, 0x6, 0xd, 0x6, 0xe, 0x6, 0x58, 0x3, 0x7, 0x3, 0x7, 
    0x3, 0x7, 0x3, 0x7, 0x3, 0x7, 0x5, 0x7, 0x60, 0xa, 0x7, 0x3, 0x8, 0x3, 
    0x8, 0x3, 0x8, 0x3, 0x9, 0x3, 0x9, 0x3, 0x9, 0x5, 0x9, 0x68, 0xa, 0x9, 
    0x3, 0x9, 0x3, 0x9, 0x3, 0xa, 0x3, 0xa, 0x5, 0xa, 0x6e, 0xa, 0xa, 0x3, 
    0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 
    0xb, 0x3, 0xb, 0x3, 0xb, 0x3, 0xb, 0x7, 0xb, 0x7b, 0xa, 0xb, 0xc, 0xb, 
    0xe, 0xb, 0x7e, 0xb, 0xb, 0x3, 0xb, 0x5, 0xb, 0x81, 0xa, 0xb, 0x3, 0xb, 
    0x3, 0xb, 0x3, 0xc, 0x3, 0xc, 0x3, 0xc, 0x3, 0xc, 0x3, 0xd, 0x3, 0xd, 
    0x3, 0xd, 0x7, 0xd, 0x8c, 0xa, 0xd, 0xc, 0xd, 0xe, 0xd, 0x8f, 0xb, 0xd, 
    0x3, 0xe, 0x3, 0xe, 0x3, 0xe, 0x5, 0xe, 0x94, 0xa, 0xe, 0x3, 0xe, 0x3, 
    0xe, 0x3, 0xf, 0x3, 0xf, 0x3, 0x10, 0x3, 0x10, 0x3, 0x10, 0x7, 0x10, 
    0x9d, 0xa, 0x10, 0xc, 0x10, 0xe, 0x10, 0xa0, 0xb, 0x10, 0x3, 0x11, 0x6, 
    0x11, 0xa3, 0xa, 0x11, 0xd, 0x11, 0xe, 0x11, 0xa4, 0x3, 0x12, 0x5, 0x12, 
    0xa8, 0xa, 0x12, 0x3, 0x12, 0x3, 0x12, 0x3, 0x12, 0x5, 0x12, 0xad, 0xa, 
    0x12, 0x3, 0x12, 0x3, 0x12, 0x3, 0x12, 0x5, 0x12, 0xb2, 0xa, 0x12, 0x3, 
    0x12, 0x3, 0x12, 0x5, 0x12, 0xb6, 0xa, 0x12, 0x3, 0x12, 0x3, 0x12, 0x3, 
    0x13, 0x3, 0x13, 0x3, 0x13, 0x3, 0x13, 0x5, 0x13, 0xbe, 0xa, 0x13, 0x3, 
    0x13, 0x3, 0x13, 0x5, 0x13, 0xc2, 0xa, 0x13, 0x3, 0x13, 0x3, 0x13, 0x3, 
    0x14, 0x3, 0x14, 0x3, 0x15, 0x5, 0x15, 0xc9, 0xa, 0x15, 0x3, 0x15, 0x3, 
    0x15, 0x3, 0x15, 0x5, 0x15, 0xce, 0xa, 0x15, 0x3, 0x15, 0x3, 0x15, 0x3, 
    0x15, 0x3, 0x16, 0x3, 0x16, 0x5, 0x16, 0xd5, 0xa, 0x16, 0x3, 0x16, 0x3, 
    0x16, 0x3, 0x16, 0x7, 0x16, 0xda, 0xa, 0x16, 0xc, 0x16, 0xe, 0x16, 0xdd, 
    0xb, 0x16, 0x3, 0x17, 0x3, 0x17, 0x3, 0x17, 0x7, 0x17, 0xe2, 0xa, 0x17, 
    0xc, 0x17, 0xe, 0x17, 0xe5, 0xb, 0x17, 0x3, 0x18, 0x3, 0x18, 0x3, 0x18, 
    0x3, 0x18, 0x3, 0x19, 0x3, 0x19, 0x5, 0x19, 0xed, 0xa, 0x19, 0x3, 0x19, 
    0x3, 0x19, 0x3, 0x19, 0x5, 0x19, 0xf2, 0xa, 0x19, 0x3, 0x19, 0x3, 0x19, 
    0x3, 0x19, 0x5, 0x19, 0xf7, 0xa, 0x19, 0x3, 0x19, 0x3, 0x19, 0x5, 0x19, 
    0xfb, 0xa, 0x19, 0x3, 0x19, 0x3, 0x19, 0x3, 0x1a, 0x6, 0x1a, 0x100, 
    0xa, 0x1a, 0xd, 0x1a, 0xe, 0x1a, 0x101, 0x3, 0x1b, 0x3, 0x1b, 0x3, 0x1b, 
    0x3, 0x1b, 0x3, 0x1c, 0x3, 0x1c, 0x5, 0x1c, 0x10a, 0xa, 0x1c, 0x3, 0x1c, 
    0x3, 0x1c, 0x3, 0x1c, 0x3, 0x1c, 0x3, 0x1c, 0x3, 0x1d, 0x6, 0x1d, 0x112, 
    0xa, 0x1d, 0xd, 0x1d, 0xe, 0x1d, 0x113, 0x3, 0x1e, 0x5, 0x1e, 0x117, 
    0xa, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 0x5, 0x1e, 0x11c, 0xa, 0x1e, 
    0x3, 0x1e, 0x3, 0x1e, 0x3, 0x1e, 0x5, 0x1e, 0x121, 0xa, 0x1e, 0x3, 0x1f, 
    0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x3, 0x1f, 0x5, 0x1f, 0x128, 0xa, 0x1f, 
    0x3, 0x1f, 0x3, 0x1f, 0x3, 0x20, 0x3, 0x20, 0x3, 0x20, 0x3, 0x20, 0x3, 
    0x20, 0x3, 0x20, 0x5, 0x20, 0x132, 0xa, 0x20, 0x3, 0x20, 0x3, 0x20, 
    0x3, 0x20, 0x3, 0x20, 0x3, 0x20, 0x3, 0x20, 0x3, 0x20, 0x3, 0x20, 0x3, 
    0x20, 0x7, 0x20, 0x13d, 0xa, 0x20, 0xc, 0x20, 0xe, 0x20, 0x140, 0xb, 
    0x20, 0x3, 0x21, 0x3, 0x21, 0x3, 0x21, 0x3, 0x21, 0x5, 0x21, 0x146, 
    0xa, 0x21, 0x3, 0x21, 0x2, 0x3, 0x3e, 0x22, 0x2, 0x4, 0x6, 0x8, 0xa, 
    0xc, 0xe, 0x10, 0x12, 0x14, 0x16, 0x18, 0x1a, 0x1c, 0x1e, 0x20, 0x22, 
    0x24, 0x26, 0x28, 0x2a, 0x2c, 0x2e, 0x30, 0x32, 0x34, 0x36, 0x38, 0x3a, 
    0x3c, 0x3e, 0x40, 0x2, 0x4, 0x3, 0x2, 0x18, 0x19, 0x4, 0x2, 0x11, 0x11, 
    0x1a, 0x1a, 0x2, 0x158, 0x2, 0x43, 0x3, 0x2, 0x2, 0x2, 0x4, 0x46, 0x3, 
    0x2, 0x2, 0x2, 0x6, 0x50, 0x3, 0x2, 0x2, 0x2, 0x8, 0x52, 0x3, 0x2, 0x2, 
    0x2, 0xa, 0x56, 0x3, 0x2, 0x2, 0x2, 0xc, 0x5f, 0x3, 0x2, 0x2, 0x2, 0xe, 
    0x61, 0x3, 0x2, 0x2, 0x2, 0x10, 0x64, 0x3, 0x2, 0x2, 0x2, 0x12, 0x6d, 
    0x3, 0x2, 0x2, 0x2, 0x14, 0x6f, 0x3, 0x2, 0x2, 0x2, 0x16, 0x84, 0x3, 
    0x2, 0x2, 0x2, 0x18, 0x88, 0x3, 0x2, 0x2, 0x2, 0x1a, 0x90, 0x3, 0x2, 
    0x2, 0x2, 0x1c, 0x97, 0x3, 0x2, 0x2, 0x2, 0x1e, 0x99, 0x3, 0x2, 0x2, 
    0x2, 0x20, 0xa2, 0x3, 0x2, 0x2, 0x2, 0x22, 0xa7, 0x3, 0x2, 0x2, 0x2, 
    0x24, 0xb9, 0x3, 0x2, 0x2, 0x2, 0x26, 0xc5, 0x3, 0x2, 0x2, 0x2, 0x28, 
    0xc8, 0x3, 0x2, 0x2, 0x2, 0x2a, 0xd2, 0x3, 0x2, 0x2, 0x2, 0x2c, 0xde, 
    0x3, 0x2, 0x2, 0x2, 0x2e, 0xe6, 0x3, 0x2, 0x2, 0x2, 0x30, 0xea, 0x3, 
    0x2, 0x2, 0x2, 0x32, 0xff, 0x3, 0x2, 0x2, 0x2, 0x34, 0x103, 0x3, 0x2, 
    0x2, 0x2, 0x36, 0x107, 0x3, 0x2, 0x2, 0x2, 0x38, 0x111, 0x3, 0x2, 0x2, 
    0x2, 0x3a, 0x116, 0x3, 0x2, 0x2, 0x2, 0x3c, 0x122, 0x3, 0x2, 0x2, 0x2, 
    0x3e, 0x131, 0x3, 0x2, 0x2, 0x2, 0x40, 0x145, 0x3, 0x2, 0x2, 0x2, 0x42, 
    0x44, 0x5, 0x4, 0x3, 0x2, 0x43, 0x42, 0x3, 0x2, 0x2, 0x2, 0x43, 0x44, 
    0x3, 0x2, 0x2, 0x2, 0x44, 0x3, 0x3, 0x2, 0x2, 0x2, 0x45, 0x47, 0x5, 
    0x6, 0x4, 0x2, 0x46, 0x45, 0x3, 0x2, 0x2, 0x2, 0x47, 0x48, 0x3, 0x2, 
    0x2, 0x2, 0x48, 0x46, 0x3, 0x2, 0x2, 0x2, 0x48, 0x49, 0x3, 0x2, 0x2, 
    0x2, 0x49, 0x5, 0x3, 0x2, 0x2, 0x2, 0x4a, 0x51, 0x5, 0x30, 0x19, 0x2, 
    0x4b, 0x51, 0x5, 0x2a, 0x16, 0x2, 0x4c, 0x51, 0x5, 0x36, 0x1c, 0x2, 
    0x4d, 0x51, 0x5, 0x24, 0x13, 0x2, 0x4e, 0x51, 0x5, 0x3c, 0x1f, 0x2, 
    0x4f, 0x51, 0x5, 0x8, 0x5, 0x2, 0x50, 0x4a, 0x3, 0x2, 0x2, 0x2, 0x50, 
    0x4b, 0x3, 0x2, 0x2, 0x2, 0x50, 0x4c, 0x3, 0x2, 0x2, 0x2, 0x50, 0x4d, 
    0x3, 0x2, 0x2, 0x2, 0x50, 0x4e, 0x3, 0x2, 0x2, 0x2, 0x50, 0x4f, 0x3, 
    0x2, 0x2, 0x2, 0x51, 0x7, 0x3, 0x2, 0x2, 0x2, 0x52, 0x53, 0x7, 0x3, 
    0x2, 0x2, 0x53, 0x54, 0x7, 0x1d, 0x2, 0x2, 0x54, 0x9, 0x3, 0x2, 0x2, 
    0x2, 0x55, 0x57, 0x5, 0xc, 0x7, 0x2, 0x56, 0x55, 0x3, 0x2, 0x2, 0x2, 
    0x57, 0x58, 0x3, 0x2, 0x2, 0x2, 0x58, 0x56, 0x3, 0x2, 0x2, 0x2, 0x58, 
    0x59, 0x3, 0x2, 0x2, 0x2, 0x59, 0xb, 0x3, 0x2, 0x2, 0x2, 0x5a, 0x60, 
    0x5, 0x2a, 0x16, 0x2, 0x5b, 0x60, 0x5, 0x10, 0x9, 0x2, 0x5c, 0x60, 0x5, 
    0xe, 0x8, 0x2, 0x5d, 0x60, 0x5, 0x16, 0xc, 0x2, 0x5e, 0x60, 0x5, 0x12, 
    0xa, 0x2, 0x5f, 0x5a, 0x3, 0x2, 0x2, 0x2, 0x5f, 0x5b, 0x3, 0x2, 0x2, 
    0x2, 0x5f, 0x5c, 0x3, 0x2, 0x2, 0x2, 0x5f, 0x5d, 0x3, 0x2, 0x2, 0x2, 
    0x5f, 0x5e, 0x3, 0x2, 0x2, 0x2, 0x60, 0xd, 0x3, 0x2, 0x2, 0x2, 0x61, 
    0x62, 0x7, 0x4, 0x2, 0x2, 0x62, 0x63, 0x5, 0x3e, 0x20, 0x2, 0x63, 0xf, 
    0x3, 0x2, 0x2, 0x2, 0x64, 0x65, 0x7, 0x5, 0x2, 0x2, 0x65, 0x67, 0x7, 
    0x6, 0x2, 0x2, 0x66, 0x68, 0x5, 0xa, 0x6, 0x2, 0x67, 0x66, 0x3, 0x2, 
    0x2, 0x2, 0x67, 0x68, 0x3, 0x2, 0x2, 0x2, 0x68, 0x69, 0x3, 0x2, 0x2, 
    0x2, 0x69, 0x6a, 0x7, 0x7, 0x2, 0x2, 0x6a, 0x11, 0x3, 0x2, 0x2, 0x2, 
    0x6b, 0x6e, 0x5, 0x14, 0xb, 0x2, 0x6c, 0x6e, 0x5, 0x1a, 0xe, 0x2, 0x6d, 
    0x6b, 0x3, 0x2, 0x2, 0x2, 0x6d, 0x6c, 0x3, 0x2, 0x2, 0x2, 0x6e, 0x13, 
    0x3, 0x2, 0x2, 0x2, 0x6f, 0x70, 0x7, 0x8, 0x2, 0x2, 0x70, 0x71, 0x5, 
    0x3e, 0x20, 0x2, 0x71, 0x72, 0x7, 0x6, 0x2, 0x2, 0x72, 0x73, 0x5, 0x3e, 
    0x20, 0x2, 0x73, 0x74, 0x7, 0x9, 0x2, 0x2, 0x74, 0x7c, 0x5, 0xc, 0x7, 
    0x2, 0x75, 0x76, 0x7, 0xa, 0x2, 0x2, 0x76, 0x77, 0x5, 0x3e, 0x20, 0x2, 
    0x77, 0x78, 0x7, 0x9, 0x2, 0x2, 0x78, 0x79, 0x5, 0xc, 0x7, 0x2, 0x79, 
    0x7b, 0x3, 0x2, 0x2, 0x2, 0x7a, 0x75, 0x3, 0x2, 0x2, 0x2, 0x7b, 0x7e, 
    0x3, 0x2, 0x2, 0x2, 0x7c, 0x7a, 0x3, 0x2, 0x2, 0x2, 0x7c, 0x7d, 0x3, 
    0x2, 0x2, 0x2, 0x7d, 0x80, 0x3, 0x2, 0x2, 0x2, 0x7e, 0x7c, 0x3, 0x2, 
    0x2, 0x2, 0x7f, 0x81, 0x7, 0xa, 0x2, 0x2, 0x80, 0x7f, 0x3, 0x2, 0x2, 
    0x2, 0x80, 0x81, 0x3, 0x2, 0x2, 0x2, 0x81, 0x82, 0x3, 0x2, 0x2, 0x2, 
    0x82, 0x83, 0x7, 0x7, 0x2, 0x2, 0x83, 0x15, 0x3, 0x2, 0x2, 0x2, 0x84, 
    0x85, 0x7, 0x1d, 0x2, 0x2, 0x85, 0x86, 0x7, 0xb, 0x2, 0x2, 0x86, 0x87, 
    0x5, 0x3e, 0x20, 0x2, 0x87, 0x17, 0x3, 0x2, 0x2, 0x2, 0x88, 0x8d, 0x5, 
    0x3e, 0x20, 0x2, 0x89, 0x8a, 0x7, 0xa, 0x2, 0x2, 0x8a, 0x8c, 0x5, 0x3e, 
    0x20, 0x2, 0x8b, 0x89, 0x3, 0x2, 0x2, 0x2, 0x8c, 0x8f, 0x3, 0x2, 0x2, 
    0x2, 0x8d, 0x8b, 0x3, 0x2, 0x2, 0x2, 0x8d, 0x8e, 0x3, 0x2, 0x2, 0x2, 
    0x8e, 0x19, 0x3, 0x2, 0x2, 0x2, 0x8f, 0x8d, 0x3, 0x2, 0x2, 0x2, 0x90, 
    0x91, 0x7, 0x1d, 0x2, 0x2, 0x91, 0x93, 0x7, 0xc, 0x2, 0x2, 0x92, 0x94, 
    0x5, 0x18, 0xd, 0x2, 0x93, 0x92, 0x3, 0x2, 0x2, 0x2, 0x93, 0x94, 0x3, 
    0x2, 0x2, 0x2, 0x94, 0x95, 0x3, 0x2, 0x2, 0x2, 0x95, 0x96, 0x7, 0xd, 
    0x2, 0x2, 0x96, 0x1b, 0x3, 0x2, 0x2, 0x2, 0x97, 0x98, 0x7, 0x1d, 0x2, 
    0x2, 0x98, 0x1d, 0x3, 0x2, 0x2, 0x2, 0x99, 0x9e, 0x5, 0x1c, 0xf, 0x2, 
    0x9a, 0x9b, 0x7, 0xa, 0x2, 0x2, 0x9b, 0x9d, 0x5, 0x1c, 0xf, 0x2, 0x9c, 
    0x9a, 0x3, 0x2, 0x2, 0x2, 0x9d, 0xa0, 0x3, 0x2, 0x2, 0x2, 0x9e, 0x9c, 
    0x3, 0x2, 0x2, 0x2, 0x9e, 0x9f, 0x3, 0x2, 0x2, 0x2, 0x9f, 0x1f, 0x3, 
    0x2, 0x2, 0x2, 0xa0, 0x9e, 0x3, 0x2, 0x2, 0x2, 0xa1, 0xa3, 0x5, 0x22, 
    0x12, 0x2, 0xa2, 0xa1, 0x3, 0x2, 0x2, 0x2, 0xa3, 0xa4, 0x3, 0x2, 0x2, 
    0x2, 0xa4, 0xa2, 0x3, 0x2, 0x2, 0x2, 0xa4, 0xa5, 0x3, 0x2, 0x2, 0x2, 
    0xa5, 0x21, 0x3, 0x2, 0x2, 0x2, 0xa6, 0xa8, 0x5, 0x26, 0x14, 0x2, 0xa7, 
    0xa6, 0x3, 0x2, 0x2, 0x2, 0xa7, 0xa8, 0x3, 0x2, 0x2, 0x2, 0xa8, 0xa9, 
    0x3, 0x2, 0x2, 0x2, 0xa9, 0xaa, 0x7, 0x1d, 0x2, 0x2, 0xaa, 0xac, 0x7, 
    0xc, 0x2, 0x2, 0xab, 0xad, 0x5, 0x2c, 0x17, 0x2, 0xac, 0xab, 0x3, 0x2, 
    0x2, 0x2, 0xac, 0xad, 0x3, 0x2, 0x2, 0x2, 0xad, 0xae, 0x3, 0x2, 0x2, 
    0x2, 0xae, 0xb1, 0x7, 0xd, 0x2, 0x2, 0xaf, 0xb0, 0x7, 0xe, 0x2, 0x2, 
    0xb0, 0xb2, 0x5, 0x1c, 0xf, 0x2, 0xb1, 0xaf, 0x3, 0x2, 0x2, 0x2, 0xb1, 
    0xb2, 0x3, 0x2, 0x2, 0x2, 0xb2, 0xb3, 0x3, 0x2, 0x2, 0x2, 0xb3, 0xb5, 
    0x7, 0x6, 0x2, 0x2, 0xb4, 0xb6, 0x5, 0xa, 0x6, 0x2, 0xb5, 0xb4, 0x3, 
    0x2, 0x2, 0x2, 0xb5, 0xb6, 0x3, 0x2, 0x2, 0x2, 0xb6, 0xb7, 0x3, 0x2, 
    0x2, 0x2, 0xb7, 0xb8, 0x7, 0x7, 0x2, 0x2, 0xb8, 0x23, 0x3, 0x2, 0x2, 
    0x2, 0xb9, 0xba, 0x7, 0xf, 0x2, 0x2, 0xba, 0xbd, 0x7, 0x1d, 0x2, 0x2, 
    0xbb, 0xbc, 0x7, 0x10, 0x2, 0x2, 0xbc, 0xbe, 0x5, 0x1e, 0x10, 0x2, 0xbd, 
    0xbb, 0x3, 0x2, 0x2, 0x2, 0xbd, 0xbe, 0x3, 0x2, 0x2, 0x2, 0xbe, 0xbf, 
    0x3, 0x2, 0x2, 0x2, 0xbf, 0xc1, 0x7, 0x6, 0x2, 0x2, 0xc0, 0xc2, 0x5, 
    0x20, 0x11, 0x2, 0xc1, 0xc0, 0x3, 0x2, 0x2, 0x2, 0xc1, 0xc2, 0x3, 0x2, 
    0x2, 0x2, 0xc2, 0xc3, 0x3, 0x2, 0x2, 0x2, 0xc3, 0xc4, 0x7, 0x7, 0x2, 
    0x2, 0xc4, 0x25, 0x3, 0x2, 0x2, 0x2, 0xc5, 0xc6, 0x7, 0x11, 0x2, 0x2, 
    0xc6, 0x27, 0x3, 0x2, 0x2, 0x2, 0xc7, 0xc9, 0x5, 0x26, 0x14, 0x2, 0xc8, 
    0xc7, 0x3, 0x2, 0x2, 0x2, 0xc8, 0xc9, 0x3, 0x2, 0x2, 0x2, 0xc9, 0xca, 
    0x3, 0x2, 0x2, 0x2, 0xca, 0xcd, 0x7, 0x1d, 0x2, 0x2, 0xcb, 0xcc, 0x7, 
    0x10, 0x2, 0x2, 0xcc, 0xce, 0x5, 0x1c, 0xf, 0x2, 0xcd, 0xcb, 0x3, 0x2, 
    0x2, 0x2, 0xcd, 0xce, 0x3, 0x2, 0x2, 0x2, 0xce, 0xcf, 0x3, 0x2, 0x2, 
    0x2, 0xcf, 0xd0, 0x7, 0xb, 0x2, 0x2, 0xd0, 0xd1, 0x5, 0x3e, 0x20, 0x2, 
    0xd1, 0x29, 0x3, 0x2, 0x2, 0x2, 0xd2, 0xd4, 0x7, 0x12, 0x2, 0x2, 0xd3, 
    0xd5, 0x7, 0x13, 0x2, 0x2, 0xd4, 0xd3, 0x3, 0x2, 0x2, 0x2, 0xd4, 0xd5, 
    0x3, 0x2, 0x2, 0x2, 0xd5, 0xd6, 0x3, 0x2, 0x2, 0x2, 0xd6, 0xdb, 0x5, 
    0x28, 0x15, 0x2, 0xd7, 0xd8, 0x7, 0xa, 0x2, 0x2, 0xd8, 0xda, 0x5, 0x28, 
    0x15, 0x2, 0xd9, 0xd7, 0x3, 0x2, 0x2, 0x2, 0xda, 0xdd, 0x3, 0x2, 0x2, 
    0x2, 0xdb, 0xd9, 0x3, 0x2, 0x2, 0x2, 0xdb, 0xdc, 0x3, 0x2, 0x2, 0x2, 
    0xdc, 0x2b, 0x3, 0x2, 0x2, 0x2, 0xdd, 0xdb, 0x3, 0x2, 0x2, 0x2, 0xde, 
    0xe3, 0x5, 0x2e, 0x18, 0x2, 0xdf, 0xe0, 0x7, 0xa, 0x2, 0x2, 0xe0, 0xe2, 
    0x5, 0x2e, 0x18, 0x2, 0xe1, 0xdf, 0x3, 0x2, 0x2, 0x2, 0xe2, 0xe5, 0x3, 
    0x2, 0x2, 0x2, 0xe3, 0xe1, 0x3, 0x2, 0x2, 0x2, 0xe3, 0xe4, 0x3, 0x2, 
    0x2, 0x2, 0xe4, 0x2d, 0x3, 0x2, 0x2, 0x2, 0xe5, 0xe3, 0x3, 0x2, 0x2, 
    0x2, 0xe6, 0xe7, 0x7, 0x1d, 0x2, 0x2, 0xe7, 0xe8, 0x7, 0x10, 0x2, 0x2, 
    0xe8, 0xe9, 0x5, 0x1c, 0xf, 0x2, 0xe9, 0x2f, 0x3, 0x2, 0x2, 0x2, 0xea, 
    0xec, 0x7, 0x14, 0x2, 0x2, 0xeb, 0xed, 0x5, 0x26, 0x14, 0x2, 0xec, 0xeb, 
    0x3, 0x2, 0x2, 0x2, 0xec, 0xed, 0x3, 0x2, 0x2, 0x2, 0xed, 0xee, 0x3, 
    0x2, 0x2, 0x2, 0xee, 0xef, 0x7, 0x1d, 0x2, 0x2, 0xef, 0xf1, 0x7, 0xc, 
    0x2, 0x2, 0xf0, 0xf2, 0x5, 0x2c, 0x17, 0x2, 0xf1, 0xf0, 0x3, 0x2, 0x2, 
    0x2, 0xf1, 0xf2, 0x3, 0x2, 0x2, 0x2, 0xf2, 0xf3, 0x3, 0x2, 0x2, 0x2, 
    0xf3, 0xf6, 0x7, 0xd, 0x2, 0x2, 0xf4, 0xf5, 0x7, 0xe, 0x2, 0x2, 0xf5, 
    0xf7, 0x5, 0x1c, 0xf, 0x2, 0xf6, 0xf4, 0x3, 0x2, 0x2, 0x2, 0xf6, 0xf7, 
    0x3, 0x2, 0x2, 0x2, 0xf7, 0xf8, 0x3, 0x2, 0x2, 0x2, 0xf8, 0xfa, 0x7, 
    0x6, 0x2, 0x2, 0xf9, 0xfb, 0x5, 0xa, 0x6, 0x2, 0xfa, 0xf9, 0x3, 0x2, 
    0x2, 0x2, 0xfa, 0xfb, 0x3, 0x2, 0x2, 0x2, 0xfb, 0xfc, 0x3, 0x2, 0x2, 
    0x2, 0xfc, 0xfd, 0x7, 0x7, 0x2, 0x2, 0xfd, 0x31, 0x3, 0x2, 0x2, 0x2, 
    0xfe, 0x100, 0x5, 0x34, 0x1b, 0x2, 0xff, 0xfe, 0x3, 0x2, 0x2, 0x2, 0x100, 
    0x101, 0x3, 0x2, 0x2, 0x2, 0x101, 0xff, 0x3, 0x2, 0x2, 0x2, 0x101, 0x102, 
    0x3, 0x2, 0x2, 0x2, 0x102, 0x33, 0x3, 0x2, 0x2, 0x2, 0x103, 0x104, 0x7, 
    0x1d, 0x2, 0x2, 0x104, 0x105, 0x7, 0x10, 0x2, 0x2, 0x105, 0x106, 0x5, 
    0x1c, 0xf, 0x2, 0x106, 0x35, 0x3, 0x2, 0x2, 0x2, 0x107, 0x109, 0x7, 
    0x15, 0x2, 0x2, 0x108, 0x10a, 0x5, 0x26, 0x14, 0x2, 0x109, 0x108, 0x3, 
    0x2, 0x2, 0x2, 0x109, 0x10a, 0x3, 0x2, 0x2, 0x2, 0x10a, 0x10b, 0x3, 
    0x2, 0x2, 0x2, 0x10b, 0x10c, 0x7, 0x1d, 0x2, 0x2, 0x10c, 0x10d, 0x7, 
    0xc, 0x2, 0x2, 0x10d, 0x10e, 0x5, 0x32, 0x1a, 0x2, 0x10e, 0x10f, 0x7, 
    0xd, 0x2, 0x2, 0x10f, 0x37, 0x3, 0x2, 0x2, 0x2, 0x110, 0x112, 0x5, 0x3a, 
    0x1e, 0x2, 0x111, 0x110, 0x3, 0x2, 0x2, 0x2, 0x112, 0x113, 0x3, 0x2, 
    0x2, 0x2, 0x113, 0x111, 0x3, 0x2, 0x2, 0x2, 0x113, 0x114, 0x3, 0x2, 
    0x2, 0x2, 0x114, 0x39, 0x3, 0x2, 0x2, 0x2, 0x115, 0x117, 0x5, 0x26, 
    0x14, 0x2, 0x116, 0x115, 0x3, 0x2, 0x2, 0x2, 0x116, 0x117, 0x3, 0x2, 
    0x2, 0x2, 0x117, 0x118, 0x3, 0x2, 0x2, 0x2, 0x118, 0x119, 0x7, 0x1d, 
    0x2, 0x2, 0x119, 0x11b, 0x7, 0xc, 0x2, 0x2, 0x11a, 0x11c, 0x5, 0x1e, 
    0x10, 0x2, 0x11b, 0x11a, 0x3, 0x2, 0x2, 0x2, 0x11b, 0x11c, 0x3, 0x2, 
    0x2, 0x2, 0x11c, 0x11d, 0x3, 0x2, 0x2, 0x2, 0x11d, 0x120, 0x7, 0xd, 
    0x2, 0x2, 0x11e, 0x11f, 0x7, 0xe, 0x2, 0x2, 0x11f, 0x121, 0x5, 0x1c, 
    0xf, 0x2, 0x120, 0x11e, 0x3, 0x2, 0x2, 0x2, 0x120, 0x121, 0x3, 0x2, 
    0x2, 0x2, 0x121, 0x3b, 0x3, 0x2, 0x2, 0x2, 0x122, 0x123, 0x7, 0x16, 
    0x2, 0x2, 0x123, 0x124, 0x5, 0x26, 0x14, 0x2, 0x124, 0x125, 0x7, 0x1d, 
    0x2, 0x2, 0x125, 0x127, 0x7, 0x6, 0x2, 0x2, 0x126, 0x128, 0x5, 0x38, 
    0x1d, 0x2, 0x127, 0x126, 0x3, 0x2, 0x2, 0x2, 0x127, 0x128, 0x3, 0x2, 
    0x2, 0x2, 0x128, 0x129, 0x3, 0x2, 0x2, 0x2, 0x129, 0x12a, 0x7, 0x7, 
    0x2, 0x2, 0x12a, 0x3d, 0x3, 0x2, 0x2, 0x2, 0x12b, 0x12c, 0x8, 0x20, 
    0x1, 0x2, 0x12c, 0x12d, 0x7, 0xc, 0x2, 0x2, 0x12d, 0x12e, 0x5, 0x3e, 
    0x20, 0x2, 0x12e, 0x12f, 0x7, 0xd, 0x2, 0x2, 0x12f, 0x132, 0x3, 0x2, 
    0x2, 0x2, 0x130, 0x132, 0x5, 0x40, 0x21, 0x2, 0x131, 0x12b, 0x3, 0x2, 
    0x2, 0x2, 0x131, 0x130, 0x3, 0x2, 0x2, 0x2, 0x132, 0x13e, 0x3, 0x2, 
    0x2, 0x2, 0x133, 0x134, 0xc, 0x6, 0x2, 0x2, 0x134, 0x135, 0x7, 0x17, 
    0x2, 0x2, 0x135, 0x13d, 0x5, 0x3e, 0x20, 0x7, 0x136, 0x137, 0xc, 0x5, 
    0x2, 0x2, 0x137, 0x138, 0x9, 0x2, 0x2, 0x2, 0x138, 0x13d, 0x5, 0x3e, 
    0x20, 0x6, 0x139, 0x13a, 0xc, 0x4, 0x2, 0x2, 0x13a, 0x13b, 0x9, 0x3, 
    0x2, 0x2, 0x13b, 0x13d, 0x5, 0x3e, 0x20, 0x5, 0x13c, 0x133, 0x3, 0x2, 
    0x2, 0x2, 0x13c, 0x136, 0x3, 0x2, 0x2, 0x2, 0x13c, 0x139, 0x3, 0x2, 
    0x2, 0x2, 0x13d, 0x140, 0x3, 0x2, 0x2, 0x2, 0x13e, 0x13c, 0x3, 0x2, 
    0x2, 0x2, 0x13e, 0x13f, 0x3, 0x2, 0x2, 0x2, 0x13f, 0x3f, 0x3, 0x2, 0x2, 
    0x2, 0x140, 0x13e, 0x3, 0x2, 0x2, 0x2, 0x141, 0x146, 0x5, 0x12, 0xa, 
    0x2, 0x142, 0x146, 0x7, 0x1e, 0x2, 0x2, 0x143, 0x146, 0x7, 0x1d, 0x2, 
    0x2, 0x144, 0x146, 0x7, 0x1f, 0x2, 0x2, 0x145, 0x141, 0x3, 0x2, 0x2, 
    0x2, 0x145, 0x142, 0x3, 0x2, 0x2, 0x2, 0x145, 0x143, 0x3, 0x2, 0x2, 
    0x2, 0x145, 0x144, 0x3, 0x2, 0x2, 0x2, 0x146, 0x41, 0x3, 0x2, 0x2, 0x2, 
    0x29, 0x43, 0x48, 0x50, 0x58, 0x5f, 0x67, 0x6d, 0x7c, 0x80, 0x8d, 0x93, 
    0x9e, 0xa4, 0xa7, 0xac, 0xb1, 0xb5, 0xbd, 0xc1, 0xc8, 0xcd, 0xd4, 0xdb, 
    0xe3, 0xec, 0xf1, 0xf6, 0xfa, 0x101, 0x109, 0x113, 0x116, 0x11b, 0x120, 
    0x127, 0x131, 0x13c, 0x13e, 0x145, 
  };

  atn::ATNDeserializer deserializer;
  _atn = deserializer.deserialize(_serializedATN);

  size_t count = _atn.getNumberOfDecisions();
  _decisionToDFA.reserve(count);
  for (size_t i = 0; i < count; i++) { 
    _decisionToDFA.emplace_back(_atn.getDecisionState(i), i);
  }
}

ElzParser::Initializer ElzParser::_init;
