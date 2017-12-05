// Generated from /home/routedan/workspace/go/src/github.com/elz-lang/elz/elzc/src/parser/Elz.g4 by ANTLR 4.7
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ElzParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		WS=25, COMMENT=26, ID=27, NUM=28, STRING=29;
	public static final int
		RULE_prog = 0, RULE_topStatList = 1, RULE_topStat = 2, RULE_importStat = 3, 
		RULE_statList = 4, RULE_stat = 5, RULE_returnStat = 6, RULE_loopStat = 7, 
		RULE_exprStat = 8, RULE_matchRule = 9, RULE_assign = 10, RULE_exprList = 11, 
		RULE_fnCall = 12, RULE_typePass = 13, RULE_typeList = 14, RULE_methodList = 15, 
		RULE_method = 16, RULE_implBlock = 17, RULE_exportor = 18, RULE_define = 19, 
		RULE_varDefine = 20, RULE_paramList = 21, RULE_param = 22, RULE_fnDefine = 23, 
		RULE_attrList = 24, RULE_attr = 25, RULE_typeDefine = 26, RULE_tmethodList = 27, 
		RULE_tmethod = 28, RULE_traitDefine = 29, RULE_expr = 30, RULE_factor = 31;
	public static final String[] ruleNames = {
		"prog", "topStatList", "topStat", "importStat", "statList", "stat", "returnStat", 
		"loopStat", "exprStat", "matchRule", "assign", "exprList", "fnCall", "typePass", 
		"typeList", "methodList", "method", "implBlock", "exportor", "define", 
		"varDefine", "paramList", "param", "fnDefine", "attrList", "attr", "typeDefine", 
		"tmethodList", "tmethod", "traitDefine", "expr", "factor"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'import'", "'return'", "'loop'", "'{'", "'}'", "'match'", "'=>'", 
		"','", "'='", "'('", "')'", "'->'", "'impl'", "':'", "'+'", "'let'", "'mut'", 
		"'fn'", "'type'", "'trait'", "'^'", "'*'", "'/'", "'-'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, null, null, null, null, null, null, null, null, null, null, null, 
		null, "WS", "COMMENT", "ID", "NUM", "STRING"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Elz.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ElzParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class ProgContext extends ParserRuleContext {
		public TopStatListContext topStatList() {
			return getRuleContext(TopStatListContext.class,0);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(65);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__12) | (1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__19))) != 0)) {
				{
				setState(64);
				topStatList();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TopStatListContext extends ParserRuleContext {
		public List<TopStatContext> topStat() {
			return getRuleContexts(TopStatContext.class);
		}
		public TopStatContext topStat(int i) {
			return getRuleContext(TopStatContext.class,i);
		}
		public TopStatListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_topStatList; }
	}

	public final TopStatListContext topStatList() throws RecognitionException {
		TopStatListContext _localctx = new TopStatListContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_topStatList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(68); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(67);
				topStat();
				}
				}
				setState(70); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__12) | (1L << T__14) | (1L << T__15) | (1L << T__17) | (1L << T__18) | (1L << T__19))) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TopStatContext extends ParserRuleContext {
		public FnDefineContext fnDefine() {
			return getRuleContext(FnDefineContext.class,0);
		}
		public VarDefineContext varDefine() {
			return getRuleContext(VarDefineContext.class,0);
		}
		public TypeDefineContext typeDefine() {
			return getRuleContext(TypeDefineContext.class,0);
		}
		public ImplBlockContext implBlock() {
			return getRuleContext(ImplBlockContext.class,0);
		}
		public TraitDefineContext traitDefine() {
			return getRuleContext(TraitDefineContext.class,0);
		}
		public ImportStatContext importStat() {
			return getRuleContext(ImportStatContext.class,0);
		}
		public TopStatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_topStat; }
	}

	public final TopStatContext topStat() throws RecognitionException {
		TopStatContext _localctx = new TopStatContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_topStat);
		try {
			setState(78);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__14:
			case T__17:
				enterOuterAlt(_localctx, 1);
				{
				setState(72);
				fnDefine();
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 2);
				{
				setState(73);
				varDefine();
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 3);
				{
				setState(74);
				typeDefine();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 4);
				{
				setState(75);
				implBlock();
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 5);
				{
				setState(76);
				traitDefine();
				}
				break;
			case T__0:
				enterOuterAlt(_localctx, 6);
				{
				setState(77);
				importStat();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportStatContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public ImportStatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importStat; }
	}

	public final ImportStatContext importStat() throws RecognitionException {
		ImportStatContext _localctx = new ImportStatContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_importStat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			match(T__0);
			setState(81);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatListContext extends ParserRuleContext {
		public List<StatContext> stat() {
			return getRuleContexts(StatContext.class);
		}
		public StatContext stat(int i) {
			return getRuleContext(StatContext.class,i);
		}
		public StatListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statList; }
	}

	public final StatListContext statList() throws RecognitionException {
		StatListContext _localctx = new StatListContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_statList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(84); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(83);
				stat();
				}
				}
				setState(86); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__2) | (1L << T__5) | (1L << T__15) | (1L << ID))) != 0) );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatContext extends ParserRuleContext {
		public VarDefineContext varDefine() {
			return getRuleContext(VarDefineContext.class,0);
		}
		public LoopStatContext loopStat() {
			return getRuleContext(LoopStatContext.class,0);
		}
		public ReturnStatContext returnStat() {
			return getRuleContext(ReturnStatContext.class,0);
		}
		public AssignContext assign() {
			return getRuleContext(AssignContext.class,0);
		}
		public ExprStatContext exprStat() {
			return getRuleContext(ExprStatContext.class,0);
		}
		public StatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stat; }
	}

	public final StatContext stat() throws RecognitionException {
		StatContext _localctx = new StatContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_stat);
		try {
			setState(93);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(88);
				varDefine();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(89);
				loopStat();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(90);
				returnStat();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(91);
				assign();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(92);
				exprStat();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReturnStatContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReturnStatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnStat; }
	}

	public final ReturnStatContext returnStat() throws RecognitionException {
		ReturnStatContext _localctx = new ReturnStatContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_returnStat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			match(T__1);
			setState(96);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LoopStatContext extends ParserRuleContext {
		public StatListContext statList() {
			return getRuleContext(StatListContext.class,0);
		}
		public LoopStatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loopStat; }
	}

	public final LoopStatContext loopStat() throws RecognitionException {
		LoopStatContext _localctx = new LoopStatContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_loopStat);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			match(T__2);
			setState(99);
			match(T__3);
			setState(101);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__2) | (1L << T__5) | (1L << T__15) | (1L << ID))) != 0)) {
				{
				setState(100);
				statList();
				}
			}

			setState(103);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprStatContext extends ParserRuleContext {
		public MatchRuleContext matchRule() {
			return getRuleContext(MatchRuleContext.class,0);
		}
		public FnCallContext fnCall() {
			return getRuleContext(FnCallContext.class,0);
		}
		public ExprStatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exprStat; }
	}

	public final ExprStatContext exprStat() throws RecognitionException {
		ExprStatContext _localctx = new ExprStatContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_exprStat);
		try {
			setState(107);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__5:
				enterOuterAlt(_localctx, 1);
				{
				setState(105);
				matchRule();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(106);
				fnCall();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MatchRuleContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<StatContext> stat() {
			return getRuleContexts(StatContext.class);
		}
		public StatContext stat(int i) {
			return getRuleContext(StatContext.class,i);
		}
		public MatchRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchRule; }
	}

	public final MatchRuleContext matchRule() throws RecognitionException {
		MatchRuleContext _localctx = new MatchRuleContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_matchRule);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			match(T__5);
			setState(110);
			expr(0);
			setState(111);
			match(T__3);
			setState(112);
			expr(0);
			setState(113);
			match(T__6);
			setState(114);
			stat();
			setState(122);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(115);
					match(T__7);
					setState(116);
					expr(0);
					setState(117);
					match(T__6);
					setState(118);
					stat();
					}
					} 
				}
				setState(124);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			setState(126);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(125);
				match(T__7);
				}
			}

			setState(128);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign; }
	}

	public final AssignContext assign() throws RecognitionException {
		AssignContext _localctx = new AssignContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_assign);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			match(ID);
			setState(131);
			match(T__8);
			setState(132);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprListContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ExprListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exprList; }
	}

	public final ExprListContext exprList() throws RecognitionException {
		ExprListContext _localctx = new ExprListContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_exprList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			expr(0);
			setState(139);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(135);
				match(T__7);
				setState(136);
				expr(0);
				}
				}
				setState(141);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FnCallContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public ExprListContext exprList() {
			return getRuleContext(ExprListContext.class,0);
		}
		public FnCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fnCall; }
	}

	public final FnCallContext fnCall() throws RecognitionException {
		FnCallContext _localctx = new FnCallContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_fnCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(142);
			match(ID);
			setState(143);
			match(T__9);
			setState(145);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__5) | (1L << T__9) | (1L << ID) | (1L << NUM) | (1L << STRING))) != 0)) {
				{
				setState(144);
				exprList();
				}
			}

			setState(147);
			match(T__10);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypePassContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public TypePassContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typePass; }
	}

	public final TypePassContext typePass() throws RecognitionException {
		TypePassContext _localctx = new TypePassContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_typePass);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeListContext extends ParserRuleContext {
		public List<TypePassContext> typePass() {
			return getRuleContexts(TypePassContext.class);
		}
		public TypePassContext typePass(int i) {
			return getRuleContext(TypePassContext.class,i);
		}
		public TypeListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeList; }
	}

	public final TypeListContext typeList() throws RecognitionException {
		TypeListContext _localctx = new TypeListContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_typeList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			typePass();
			setState(156);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(152);
				match(T__7);
				setState(153);
				typePass();
				}
				}
				setState(158);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MethodListContext extends ParserRuleContext {
		public List<MethodContext> method() {
			return getRuleContexts(MethodContext.class);
		}
		public MethodContext method(int i) {
			return getRuleContext(MethodContext.class,i);
		}
		public MethodListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_methodList; }
	}

	public final MethodListContext methodList() throws RecognitionException {
		MethodListContext _localctx = new MethodListContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_methodList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(160); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(159);
				method();
				}
				}
				setState(162); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__14 || _la==ID );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MethodContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public ExportorContext exportor() {
			return getRuleContext(ExportorContext.class,0);
		}
		public ParamListContext paramList() {
			return getRuleContext(ParamListContext.class,0);
		}
		public TypePassContext typePass() {
			return getRuleContext(TypePassContext.class,0);
		}
		public StatListContext statList() {
			return getRuleContext(StatListContext.class,0);
		}
		public MethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_method; }
	}

	public final MethodContext method() throws RecognitionException {
		MethodContext _localctx = new MethodContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_method);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14) {
				{
				setState(164);
				exportor();
				}
			}

			setState(167);
			match(ID);
			setState(168);
			match(T__9);
			setState(170);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(169);
				paramList();
				}
			}

			setState(172);
			match(T__10);
			setState(175);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(173);
				match(T__11);
				setState(174);
				typePass();
				}
			}

			setState(177);
			match(T__3);
			setState(179);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__2) | (1L << T__5) | (1L << T__15) | (1L << ID))) != 0)) {
				{
				setState(178);
				statList();
				}
			}

			setState(181);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImplBlockContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public TypeListContext typeList() {
			return getRuleContext(TypeListContext.class,0);
		}
		public MethodListContext methodList() {
			return getRuleContext(MethodListContext.class,0);
		}
		public ImplBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_implBlock; }
	}

	public final ImplBlockContext implBlock() throws RecognitionException {
		ImplBlockContext _localctx = new ImplBlockContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_implBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			match(T__12);
			setState(184);
			match(ID);
			setState(187);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__13) {
				{
				setState(185);
				match(T__13);
				setState(186);
				typeList();
				}
			}

			setState(189);
			match(T__3);
			setState(191);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14 || _la==ID) {
				{
				setState(190);
				methodList();
				}
			}

			setState(193);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExportorContext extends ParserRuleContext {
		public ExportorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exportor; }
	}

	public final ExportorContext exportor() throws RecognitionException {
		ExportorContext _localctx = new ExportorContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_exportor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(195);
			match(T__14);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DefineContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ExportorContext exportor() {
			return getRuleContext(ExportorContext.class,0);
		}
		public TypePassContext typePass() {
			return getRuleContext(TypePassContext.class,0);
		}
		public DefineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_define; }
	}

	public final DefineContext define() throws RecognitionException {
		DefineContext _localctx = new DefineContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_define);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(198);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14) {
				{
				setState(197);
				exportor();
				}
			}

			setState(200);
			match(ID);
			setState(203);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__13) {
				{
				setState(201);
				match(T__13);
				setState(202);
				typePass();
				}
			}

			setState(205);
			match(T__8);
			setState(206);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VarDefineContext extends ParserRuleContext {
		public Token mut;
		public List<DefineContext> define() {
			return getRuleContexts(DefineContext.class);
		}
		public DefineContext define(int i) {
			return getRuleContext(DefineContext.class,i);
		}
		public VarDefineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDefine; }
	}

	public final VarDefineContext varDefine() throws RecognitionException {
		VarDefineContext _localctx = new VarDefineContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_varDefine);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			match(T__15);
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__16) {
				{
				setState(209);
				((VarDefineContext)_localctx).mut = match(T__16);
				}
			}

			setState(212);
			define();
			setState(217);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(213);
					match(T__7);
					setState(214);
					define();
					}
					} 
				}
				setState(219);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamListContext extends ParserRuleContext {
		public List<ParamContext> param() {
			return getRuleContexts(ParamContext.class);
		}
		public ParamContext param(int i) {
			return getRuleContext(ParamContext.class,i);
		}
		public ParamListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramList; }
	}

	public final ParamListContext paramList() throws RecognitionException {
		ParamListContext _localctx = new ParamListContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_paramList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(220);
			param();
			setState(225);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__7) {
				{
				{
				setState(221);
				match(T__7);
				setState(222);
				param();
				}
				}
				setState(227);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParamContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public TypePassContext typePass() {
			return getRuleContext(TypePassContext.class,0);
		}
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			match(ID);
			setState(229);
			match(T__13);
			setState(230);
			typePass();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FnDefineContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public ExportorContext exportor() {
			return getRuleContext(ExportorContext.class,0);
		}
		public ParamListContext paramList() {
			return getRuleContext(ParamListContext.class,0);
		}
		public TypePassContext typePass() {
			return getRuleContext(TypePassContext.class,0);
		}
		public StatListContext statList() {
			return getRuleContext(StatListContext.class,0);
		}
		public FnDefineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fnDefine; }
	}

	public final FnDefineContext fnDefine() throws RecognitionException {
		FnDefineContext _localctx = new FnDefineContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_fnDefine);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(233);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14) {
				{
				setState(232);
				exportor();
				}
			}

			setState(235);
			match(T__17);
			setState(236);
			match(ID);
			setState(237);
			match(T__9);
			setState(239);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(238);
				paramList();
				}
			}

			setState(241);
			match(T__10);
			setState(244);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(242);
				match(T__11);
				setState(243);
				typePass();
				}
			}

			setState(246);
			match(T__3);
			setState(248);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__1) | (1L << T__2) | (1L << T__5) | (1L << T__15) | (1L << ID))) != 0)) {
				{
				setState(247);
				statList();
				}
			}

			setState(250);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AttrListContext extends ParserRuleContext {
		public List<AttrContext> attr() {
			return getRuleContexts(AttrContext.class);
		}
		public AttrContext attr(int i) {
			return getRuleContext(AttrContext.class,i);
		}
		public AttrListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attrList; }
	}

	public final AttrListContext attrList() throws RecognitionException {
		AttrListContext _localctx = new AttrListContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_attrList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(253); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(252);
				attr();
				}
				}
				setState(255); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AttrContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public TypePassContext typePass() {
			return getRuleContext(TypePassContext.class,0);
		}
		public AttrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attr; }
	}

	public final AttrContext attr() throws RecognitionException {
		AttrContext _localctx = new AttrContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_attr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(257);
			match(ID);
			setState(258);
			match(T__13);
			setState(259);
			typePass();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeDefineContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public AttrListContext attrList() {
			return getRuleContext(AttrListContext.class,0);
		}
		public ExportorContext exportor() {
			return getRuleContext(ExportorContext.class,0);
		}
		public TypeDefineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDefine; }
	}

	public final TypeDefineContext typeDefine() throws RecognitionException {
		TypeDefineContext _localctx = new TypeDefineContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_typeDefine);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(261);
			match(T__18);
			setState(263);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14) {
				{
				setState(262);
				exportor();
				}
			}

			setState(265);
			match(ID);
			setState(266);
			match(T__9);
			setState(267);
			attrList();
			setState(268);
			match(T__10);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TmethodListContext extends ParserRuleContext {
		public List<TmethodContext> tmethod() {
			return getRuleContexts(TmethodContext.class);
		}
		public TmethodContext tmethod(int i) {
			return getRuleContext(TmethodContext.class,i);
		}
		public TmethodListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tmethodList; }
	}

	public final TmethodListContext tmethodList() throws RecognitionException {
		TmethodListContext _localctx = new TmethodListContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_tmethodList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(271); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(270);
				tmethod();
				}
				}
				setState(273); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__14 || _la==ID );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TmethodContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public ExportorContext exportor() {
			return getRuleContext(ExportorContext.class,0);
		}
		public TypeListContext typeList() {
			return getRuleContext(TypeListContext.class,0);
		}
		public TypePassContext typePass() {
			return getRuleContext(TypePassContext.class,0);
		}
		public TmethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tmethod; }
	}

	public final TmethodContext tmethod() throws RecognitionException {
		TmethodContext _localctx = new TmethodContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_tmethod);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14) {
				{
				setState(275);
				exportor();
				}
			}

			setState(278);
			match(ID);
			setState(279);
			match(T__9);
			setState(281);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(280);
				typeList();
				}
			}

			setState(283);
			match(T__10);
			setState(286);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(284);
				match(T__11);
				setState(285);
				typePass();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TraitDefineContext extends ParserRuleContext {
		public ExportorContext exportor() {
			return getRuleContext(ExportorContext.class,0);
		}
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public TmethodListContext tmethodList() {
			return getRuleContext(TmethodListContext.class,0);
		}
		public TraitDefineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_traitDefine; }
	}

	public final TraitDefineContext traitDefine() throws RecognitionException {
		TraitDefineContext _localctx = new TraitDefineContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_traitDefine);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(288);
			match(T__19);
			setState(289);
			exportor();
			setState(290);
			match(ID);
			setState(291);
			match(T__3);
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14 || _la==ID) {
				{
				setState(292);
				tmethodList();
				}
			}

			setState(295);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprContext extends ParserRuleContext {
		public Token op;
		public FactorContext factor() {
			return getRuleContext(FactorContext.class,0);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 60;
		enterRecursionRule(_localctx, 60, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(298);
			factor();
			}
			_ctx.stop = _input.LT(-1);
			setState(311);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(309);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(300);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(301);
						((ExprContext)_localctx).op = match(T__20);
						setState(302);
						expr(5);
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(303);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(304);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__21 || _la==T__22) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(305);
						expr(4);
						}
						break;
					case 3:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(306);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(307);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__14 || _la==T__23) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(308);
						expr(3);
						}
						break;
					}
					} 
				}
				setState(313);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class FactorContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ExprStatContext exprStat() {
			return getRuleContext(ExprStatContext.class,0);
		}
		public TerminalNode NUM() { return getToken(ElzParser.NUM, 0); }
		public TerminalNode ID() { return getToken(ElzParser.ID, 0); }
		public TerminalNode STRING() { return getToken(ElzParser.STRING, 0); }
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_factor);
		try {
			setState(322);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(314);
				match(T__9);
				setState(315);
				expr(0);
				setState(316);
				match(T__10);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(318);
				exprStat();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(319);
				match(NUM);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(320);
				match(ID);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(321);
				match(STRING);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 30:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 4);
		case 1:
			return precpred(_ctx, 3);
		case 2:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\37\u0147\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\3\2\5\2D\n\2\3\3\6\3G\n\3\r\3\16\3H\3\4\3\4\3\4\3\4\3\4\3\4\5\4Q\n"+
		"\4\3\5\3\5\3\5\3\6\6\6W\n\6\r\6\16\6X\3\7\3\7\3\7\3\7\3\7\5\7`\n\7\3\b"+
		"\3\b\3\b\3\t\3\t\3\t\5\th\n\t\3\t\3\t\3\n\3\n\5\nn\n\n\3\13\3\13\3\13"+
		"\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\7\13{\n\13\f\13\16\13~\13\13"+
		"\3\13\5\13\u0081\n\13\3\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r\7\r\u008c"+
		"\n\r\f\r\16\r\u008f\13\r\3\16\3\16\3\16\5\16\u0094\n\16\3\16\3\16\3\17"+
		"\3\17\3\20\3\20\3\20\7\20\u009d\n\20\f\20\16\20\u00a0\13\20\3\21\6\21"+
		"\u00a3\n\21\r\21\16\21\u00a4\3\22\5\22\u00a8\n\22\3\22\3\22\3\22\5\22"+
		"\u00ad\n\22\3\22\3\22\3\22\5\22\u00b2\n\22\3\22\3\22\5\22\u00b6\n\22\3"+
		"\22\3\22\3\23\3\23\3\23\3\23\5\23\u00be\n\23\3\23\3\23\5\23\u00c2\n\23"+
		"\3\23\3\23\3\24\3\24\3\25\5\25\u00c9\n\25\3\25\3\25\3\25\5\25\u00ce\n"+
		"\25\3\25\3\25\3\25\3\26\3\26\5\26\u00d5\n\26\3\26\3\26\3\26\7\26\u00da"+
		"\n\26\f\26\16\26\u00dd\13\26\3\27\3\27\3\27\7\27\u00e2\n\27\f\27\16\27"+
		"\u00e5\13\27\3\30\3\30\3\30\3\30\3\31\5\31\u00ec\n\31\3\31\3\31\3\31\3"+
		"\31\5\31\u00f2\n\31\3\31\3\31\3\31\5\31\u00f7\n\31\3\31\3\31\5\31\u00fb"+
		"\n\31\3\31\3\31\3\32\6\32\u0100\n\32\r\32\16\32\u0101\3\33\3\33\3\33\3"+
		"\33\3\34\3\34\5\34\u010a\n\34\3\34\3\34\3\34\3\34\3\34\3\35\6\35\u0112"+
		"\n\35\r\35\16\35\u0113\3\36\5\36\u0117\n\36\3\36\3\36\3\36\5\36\u011c"+
		"\n\36\3\36\3\36\3\36\5\36\u0121\n\36\3\37\3\37\3\37\3\37\3\37\5\37\u0128"+
		"\n\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \7 \u0138\n \f \16"+
		" \u013b\13 \3!\3!\3!\3!\3!\3!\3!\3!\5!\u0145\n!\3!\2\3>\"\2\4\6\b\n\f"+
		"\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@\2\4\3\2\30\31\4"+
		"\2\21\21\32\32\2\u0157\2C\3\2\2\2\4F\3\2\2\2\6P\3\2\2\2\bR\3\2\2\2\nV"+
		"\3\2\2\2\f_\3\2\2\2\16a\3\2\2\2\20d\3\2\2\2\22m\3\2\2\2\24o\3\2\2\2\26"+
		"\u0084\3\2\2\2\30\u0088\3\2\2\2\32\u0090\3\2\2\2\34\u0097\3\2\2\2\36\u0099"+
		"\3\2\2\2 \u00a2\3\2\2\2\"\u00a7\3\2\2\2$\u00b9\3\2\2\2&\u00c5\3\2\2\2"+
		"(\u00c8\3\2\2\2*\u00d2\3\2\2\2,\u00de\3\2\2\2.\u00e6\3\2\2\2\60\u00eb"+
		"\3\2\2\2\62\u00ff\3\2\2\2\64\u0103\3\2\2\2\66\u0107\3\2\2\28\u0111\3\2"+
		"\2\2:\u0116\3\2\2\2<\u0122\3\2\2\2>\u012b\3\2\2\2@\u0144\3\2\2\2BD\5\4"+
		"\3\2CB\3\2\2\2CD\3\2\2\2D\3\3\2\2\2EG\5\6\4\2FE\3\2\2\2GH\3\2\2\2HF\3"+
		"\2\2\2HI\3\2\2\2I\5\3\2\2\2JQ\5\60\31\2KQ\5*\26\2LQ\5\66\34\2MQ\5$\23"+
		"\2NQ\5<\37\2OQ\5\b\5\2PJ\3\2\2\2PK\3\2\2\2PL\3\2\2\2PM\3\2\2\2PN\3\2\2"+
		"\2PO\3\2\2\2Q\7\3\2\2\2RS\7\3\2\2ST\7\35\2\2T\t\3\2\2\2UW\5\f\7\2VU\3"+
		"\2\2\2WX\3\2\2\2XV\3\2\2\2XY\3\2\2\2Y\13\3\2\2\2Z`\5*\26\2[`\5\20\t\2"+
		"\\`\5\16\b\2]`\5\26\f\2^`\5\22\n\2_Z\3\2\2\2_[\3\2\2\2_\\\3\2\2\2_]\3"+
		"\2\2\2_^\3\2\2\2`\r\3\2\2\2ab\7\4\2\2bc\5> \2c\17\3\2\2\2de\7\5\2\2eg"+
		"\7\6\2\2fh\5\n\6\2gf\3\2\2\2gh\3\2\2\2hi\3\2\2\2ij\7\7\2\2j\21\3\2\2\2"+
		"kn\5\24\13\2ln\5\32\16\2mk\3\2\2\2ml\3\2\2\2n\23\3\2\2\2op\7\b\2\2pq\5"+
		"> \2qr\7\6\2\2rs\5> \2st\7\t\2\2t|\5\f\7\2uv\7\n\2\2vw\5> \2wx\7\t\2\2"+
		"xy\5\f\7\2y{\3\2\2\2zu\3\2\2\2{~\3\2\2\2|z\3\2\2\2|}\3\2\2\2}\u0080\3"+
		"\2\2\2~|\3\2\2\2\177\u0081\7\n\2\2\u0080\177\3\2\2\2\u0080\u0081\3\2\2"+
		"\2\u0081\u0082\3\2\2\2\u0082\u0083\7\7\2\2\u0083\25\3\2\2\2\u0084\u0085"+
		"\7\35\2\2\u0085\u0086\7\13\2\2\u0086\u0087\5> \2\u0087\27\3\2\2\2\u0088"+
		"\u008d\5> \2\u0089\u008a\7\n\2\2\u008a\u008c\5> \2\u008b\u0089\3\2\2\2"+
		"\u008c\u008f\3\2\2\2\u008d\u008b\3\2\2\2\u008d\u008e\3\2\2\2\u008e\31"+
		"\3\2\2\2\u008f\u008d\3\2\2\2\u0090\u0091\7\35\2\2\u0091\u0093\7\f\2\2"+
		"\u0092\u0094\5\30\r\2\u0093\u0092\3\2\2\2\u0093\u0094\3\2\2\2\u0094\u0095"+
		"\3\2\2\2\u0095\u0096\7\r\2\2\u0096\33\3\2\2\2\u0097\u0098\7\35\2\2\u0098"+
		"\35\3\2\2\2\u0099\u009e\5\34\17\2\u009a\u009b\7\n\2\2\u009b\u009d\5\34"+
		"\17\2\u009c\u009a\3\2\2\2\u009d\u00a0\3\2\2\2\u009e\u009c\3\2\2\2\u009e"+
		"\u009f\3\2\2\2\u009f\37\3\2\2\2\u00a0\u009e\3\2\2\2\u00a1\u00a3\5\"\22"+
		"\2\u00a2\u00a1\3\2\2\2\u00a3\u00a4\3\2\2\2\u00a4\u00a2\3\2\2\2\u00a4\u00a5"+
		"\3\2\2\2\u00a5!\3\2\2\2\u00a6\u00a8\5&\24\2\u00a7\u00a6\3\2\2\2\u00a7"+
		"\u00a8\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00aa\7\35\2\2\u00aa\u00ac\7"+
		"\f\2\2\u00ab\u00ad\5,\27\2\u00ac\u00ab\3\2\2\2\u00ac\u00ad\3\2\2\2\u00ad"+
		"\u00ae\3\2\2\2\u00ae\u00b1\7\r\2\2\u00af\u00b0\7\16\2\2\u00b0\u00b2\5"+
		"\34\17\2\u00b1\u00af\3\2\2\2\u00b1\u00b2\3\2\2\2\u00b2\u00b3\3\2\2\2\u00b3"+
		"\u00b5\7\6\2\2\u00b4\u00b6\5\n\6\2\u00b5\u00b4\3\2\2\2\u00b5\u00b6\3\2"+
		"\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b8\7\7\2\2\u00b8#\3\2\2\2\u00b9\u00ba"+
		"\7\17\2\2\u00ba\u00bd\7\35\2\2\u00bb\u00bc\7\20\2\2\u00bc\u00be\5\36\20"+
		"\2\u00bd\u00bb\3\2\2\2\u00bd\u00be\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00c1"+
		"\7\6\2\2\u00c0\u00c2\5 \21\2\u00c1\u00c0\3\2\2\2\u00c1\u00c2\3\2\2\2\u00c2"+
		"\u00c3\3\2\2\2\u00c3\u00c4\7\7\2\2\u00c4%\3\2\2\2\u00c5\u00c6\7\21\2\2"+
		"\u00c6\'\3\2\2\2\u00c7\u00c9\5&\24\2\u00c8\u00c7\3\2\2\2\u00c8\u00c9\3"+
		"\2\2\2\u00c9\u00ca\3\2\2\2\u00ca\u00cd\7\35\2\2\u00cb\u00cc\7\20\2\2\u00cc"+
		"\u00ce\5\34\17\2\u00cd\u00cb\3\2\2\2\u00cd\u00ce\3\2\2\2\u00ce\u00cf\3"+
		"\2\2\2\u00cf\u00d0\7\13\2\2\u00d0\u00d1\5> \2\u00d1)\3\2\2\2\u00d2\u00d4"+
		"\7\22\2\2\u00d3\u00d5\7\23\2\2\u00d4\u00d3\3\2\2\2\u00d4\u00d5\3\2\2\2"+
		"\u00d5\u00d6\3\2\2\2\u00d6\u00db\5(\25\2\u00d7\u00d8\7\n\2\2\u00d8\u00da"+
		"\5(\25\2\u00d9\u00d7\3\2\2\2\u00da\u00dd\3\2\2\2\u00db\u00d9\3\2\2\2\u00db"+
		"\u00dc\3\2\2\2\u00dc+\3\2\2\2\u00dd\u00db\3\2\2\2\u00de\u00e3\5.\30\2"+
		"\u00df\u00e0\7\n\2\2\u00e0\u00e2\5.\30\2\u00e1\u00df\3\2\2\2\u00e2\u00e5"+
		"\3\2\2\2\u00e3\u00e1\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4-\3\2\2\2\u00e5"+
		"\u00e3\3\2\2\2\u00e6\u00e7\7\35\2\2\u00e7\u00e8\7\20\2\2\u00e8\u00e9\5"+
		"\34\17\2\u00e9/\3\2\2\2\u00ea\u00ec\5&\24\2\u00eb\u00ea\3\2\2\2\u00eb"+
		"\u00ec\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed\u00ee\7\24\2\2\u00ee\u00ef\7"+
		"\35\2\2\u00ef\u00f1\7\f\2\2\u00f0\u00f2\5,\27\2\u00f1\u00f0\3\2\2\2\u00f1"+
		"\u00f2\3\2\2\2\u00f2\u00f3\3\2\2\2\u00f3\u00f6\7\r\2\2\u00f4\u00f5\7\16"+
		"\2\2\u00f5\u00f7\5\34\17\2\u00f6\u00f4\3\2\2\2\u00f6\u00f7\3\2\2\2\u00f7"+
		"\u00f8\3\2\2\2\u00f8\u00fa\7\6\2\2\u00f9\u00fb\5\n\6\2\u00fa\u00f9\3\2"+
		"\2\2\u00fa\u00fb\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc\u00fd\7\7\2\2\u00fd"+
		"\61\3\2\2\2\u00fe\u0100\5\64\33\2\u00ff\u00fe\3\2\2\2\u0100\u0101\3\2"+
		"\2\2\u0101\u00ff\3\2\2\2\u0101\u0102\3\2\2\2\u0102\63\3\2\2\2\u0103\u0104"+
		"\7\35\2\2\u0104\u0105\7\20\2\2\u0105\u0106\5\34\17\2\u0106\65\3\2\2\2"+
		"\u0107\u0109\7\25\2\2\u0108\u010a\5&\24\2\u0109\u0108\3\2\2\2\u0109\u010a"+
		"\3\2\2\2\u010a\u010b\3\2\2\2\u010b\u010c\7\35\2\2\u010c\u010d\7\f\2\2"+
		"\u010d\u010e\5\62\32\2\u010e\u010f\7\r\2\2\u010f\67\3\2\2\2\u0110\u0112"+
		"\5:\36\2\u0111\u0110\3\2\2\2\u0112\u0113\3\2\2\2\u0113\u0111\3\2\2\2\u0113"+
		"\u0114\3\2\2\2\u01149\3\2\2\2\u0115\u0117\5&\24\2\u0116\u0115\3\2\2\2"+
		"\u0116\u0117\3\2\2\2\u0117\u0118\3\2\2\2\u0118\u0119\7\35\2\2\u0119\u011b"+
		"\7\f\2\2\u011a\u011c\5\36\20\2\u011b\u011a\3\2\2\2\u011b\u011c\3\2\2\2"+
		"\u011c\u011d\3\2\2\2\u011d\u0120\7\r\2\2\u011e\u011f\7\16\2\2\u011f\u0121"+
		"\5\34\17\2\u0120\u011e\3\2\2\2\u0120\u0121\3\2\2\2\u0121;\3\2\2\2\u0122"+
		"\u0123\7\26\2\2\u0123\u0124\5&\24\2\u0124\u0125\7\35\2\2\u0125\u0127\7"+
		"\6\2\2\u0126\u0128\58\35\2\u0127\u0126\3\2\2\2\u0127\u0128\3\2\2\2\u0128"+
		"\u0129\3\2\2\2\u0129\u012a\7\7\2\2\u012a=\3\2\2\2\u012b\u012c\b \1\2\u012c"+
		"\u012d\5@!\2\u012d\u0139\3\2\2\2\u012e\u012f\f\6\2\2\u012f\u0130\7\27"+
		"\2\2\u0130\u0138\5> \7\u0131\u0132\f\5\2\2\u0132\u0133\t\2\2\2\u0133\u0138"+
		"\5> \6\u0134\u0135\f\4\2\2\u0135\u0136\t\3\2\2\u0136\u0138\5> \5\u0137"+
		"\u012e\3\2\2\2\u0137\u0131\3\2\2\2\u0137\u0134\3\2\2\2\u0138\u013b\3\2"+
		"\2\2\u0139\u0137\3\2\2\2\u0139\u013a\3\2\2\2\u013a?\3\2\2\2\u013b\u0139"+
		"\3\2\2\2\u013c\u013d\7\f\2\2\u013d\u013e\5> \2\u013e\u013f\7\r\2\2\u013f"+
		"\u0145\3\2\2\2\u0140\u0145\5\22\n\2\u0141\u0145\7\36\2\2\u0142\u0145\7"+
		"\35\2\2\u0143\u0145\7\37\2\2\u0144\u013c\3\2\2\2\u0144\u0140\3\2\2\2\u0144"+
		"\u0141\3\2\2\2\u0144\u0142\3\2\2\2\u0144\u0143\3\2\2\2\u0145A\3\2\2\2"+
		"(CHPX_gm|\u0080\u008d\u0093\u009e\u00a4\u00a7\u00ac\u00b1\u00b5\u00bd"+
		"\u00c1\u00c8\u00cd\u00d4\u00db\u00e3\u00eb\u00f1\u00f6\u00fa\u0101\u0109"+
		"\u0113\u0116\u011b\u0120\u0127\u0137\u0139\u0144";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}