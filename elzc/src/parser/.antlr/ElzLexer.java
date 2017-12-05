// Generated from /home/routedan/workspace/go/src/github.com/elz-lang/elz/elzc/src/parser/Elz.g4 by ANTLR 4.7
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ElzLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		WS=25, COMMENT=26, ID=27, NUM=28, STRING=29;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "WS", "COMMENT", 
		"ID", "StartLetter", "Letter", "NUM", "StartDigit", "Digit", "STRING"
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


	public ElzLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Elz.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\37\u00cc\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b"+
		"\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\16\3\16"+
		"\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\23"+
		"\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26"+
		"\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\6\32\u0099\n\32\r\32\16\32\u009a"+
		"\3\32\3\32\3\33\3\33\3\33\3\33\7\33\u00a3\n\33\f\33\16\33\u00a6\13\33"+
		"\3\33\3\33\3\33\3\33\3\34\3\34\7\34\u00ae\n\34\f\34\16\34\u00b1\13\34"+
		"\3\35\3\35\3\36\3\36\5\36\u00b7\n\36\3\37\3\37\7\37\u00bb\n\37\f\37\16"+
		"\37\u00be\13\37\3 \3 \3!\3!\3\"\3\"\7\"\u00c6\n\"\f\"\16\"\u00c9\13\""+
		"\3\"\3\"\4\u00a4\u00c7\2#\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25"+
		"\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32"+
		"\63\33\65\34\67\359\2;\2=\36?\2A\2C\37\3\2\6\5\2\13\f\17\17\"\"\5\2C\\"+
		"aac|\3\2\62;\4\2\60\60\62;\2\u00cd\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2"+
		"\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3"+
		"\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2"+
		"\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2"+
		"\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2"+
		"\2\2\2\67\3\2\2\2\2=\3\2\2\2\2C\3\2\2\2\3E\3\2\2\2\5L\3\2\2\2\7S\3\2\2"+
		"\2\tX\3\2\2\2\13Z\3\2\2\2\r\\\3\2\2\2\17b\3\2\2\2\21e\3\2\2\2\23g\3\2"+
		"\2\2\25i\3\2\2\2\27k\3\2\2\2\31m\3\2\2\2\33p\3\2\2\2\35u\3\2\2\2\37w\3"+
		"\2\2\2!y\3\2\2\2#}\3\2\2\2%\u0081\3\2\2\2\'\u0084\3\2\2\2)\u0089\3\2\2"+
		"\2+\u008f\3\2\2\2-\u0091\3\2\2\2/\u0093\3\2\2\2\61\u0095\3\2\2\2\63\u0098"+
		"\3\2\2\2\65\u009e\3\2\2\2\67\u00ab\3\2\2\29\u00b2\3\2\2\2;\u00b6\3\2\2"+
		"\2=\u00b8\3\2\2\2?\u00bf\3\2\2\2A\u00c1\3\2\2\2C\u00c3\3\2\2\2EF\7k\2"+
		"\2FG\7o\2\2GH\7r\2\2HI\7q\2\2IJ\7t\2\2JK\7v\2\2K\4\3\2\2\2LM\7t\2\2MN"+
		"\7g\2\2NO\7v\2\2OP\7w\2\2PQ\7t\2\2QR\7p\2\2R\6\3\2\2\2ST\7n\2\2TU\7q\2"+
		"\2UV\7q\2\2VW\7r\2\2W\b\3\2\2\2XY\7}\2\2Y\n\3\2\2\2Z[\7\177\2\2[\f\3\2"+
		"\2\2\\]\7o\2\2]^\7c\2\2^_\7v\2\2_`\7e\2\2`a\7j\2\2a\16\3\2\2\2bc\7?\2"+
		"\2cd\7@\2\2d\20\3\2\2\2ef\7.\2\2f\22\3\2\2\2gh\7?\2\2h\24\3\2\2\2ij\7"+
		"*\2\2j\26\3\2\2\2kl\7+\2\2l\30\3\2\2\2mn\7/\2\2no\7@\2\2o\32\3\2\2\2p"+
		"q\7k\2\2qr\7o\2\2rs\7r\2\2st\7n\2\2t\34\3\2\2\2uv\7<\2\2v\36\3\2\2\2w"+
		"x\7-\2\2x \3\2\2\2yz\7n\2\2z{\7g\2\2{|\7v\2\2|\"\3\2\2\2}~\7o\2\2~\177"+
		"\7w\2\2\177\u0080\7v\2\2\u0080$\3\2\2\2\u0081\u0082\7h\2\2\u0082\u0083"+
		"\7p\2\2\u0083&\3\2\2\2\u0084\u0085\7v\2\2\u0085\u0086\7{\2\2\u0086\u0087"+
		"\7r\2\2\u0087\u0088\7g\2\2\u0088(\3\2\2\2\u0089\u008a\7v\2\2\u008a\u008b"+
		"\7t\2\2\u008b\u008c\7c\2\2\u008c\u008d\7k\2\2\u008d\u008e\7v\2\2\u008e"+
		"*\3\2\2\2\u008f\u0090\7`\2\2\u0090,\3\2\2\2\u0091\u0092\7,\2\2\u0092."+
		"\3\2\2\2\u0093\u0094\7\61\2\2\u0094\60\3\2\2\2\u0095\u0096\7/\2\2\u0096"+
		"\62\3\2\2\2\u0097\u0099\t\2\2\2\u0098\u0097\3\2\2\2\u0099\u009a\3\2\2"+
		"\2\u009a\u0098\3\2\2\2\u009a\u009b\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u009d"+
		"\b\32\2\2\u009d\64\3\2\2\2\u009e\u009f\7\61\2\2\u009f\u00a0\7\61\2\2\u00a0"+
		"\u00a4\3\2\2\2\u00a1\u00a3\13\2\2\2\u00a2\u00a1\3\2\2\2\u00a3\u00a6\3"+
		"\2\2\2\u00a4\u00a5\3\2\2\2\u00a4\u00a2\3\2\2\2\u00a5\u00a7\3\2\2\2\u00a6"+
		"\u00a4\3\2\2\2\u00a7\u00a8\7\f\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00aa\b\33"+
		"\2\2\u00aa\66\3\2\2\2\u00ab\u00af\59\35\2\u00ac\u00ae\5;\36\2\u00ad\u00ac"+
		"\3\2\2\2\u00ae\u00b1\3\2\2\2\u00af\u00ad\3\2\2\2\u00af\u00b0\3\2\2\2\u00b0"+
		"8\3\2\2\2\u00b1\u00af\3\2\2\2\u00b2\u00b3\t\3\2\2\u00b3:\3\2\2\2\u00b4"+
		"\u00b7\t\4\2\2\u00b5\u00b7\59\35\2\u00b6\u00b4\3\2\2\2\u00b6\u00b5\3\2"+
		"\2\2\u00b7<\3\2\2\2\u00b8\u00bc\5? \2\u00b9\u00bb\5A!\2\u00ba\u00b9\3"+
		"\2\2\2\u00bb\u00be\3\2\2\2\u00bc\u00ba\3\2\2\2\u00bc\u00bd\3\2\2\2\u00bd"+
		">\3\2\2\2\u00be\u00bc\3\2\2\2\u00bf\u00c0\t\5\2\2\u00c0@\3\2\2\2\u00c1"+
		"\u00c2\t\4\2\2\u00c2B\3\2\2\2\u00c3\u00c7\7$\2\2\u00c4\u00c6\13\2\2\2"+
		"\u00c5\u00c4\3\2\2\2\u00c6\u00c9\3\2\2\2\u00c7\u00c8\3\2\2\2\u00c7\u00c5"+
		"\3\2\2\2\u00c8\u00ca\3\2\2\2\u00c9\u00c7\3\2\2\2\u00ca\u00cb\7$\2\2\u00cb"+
		"D\3\2\2\2\t\2\u009a\u00a4\u00af\u00b6\u00bc\u00c7\3\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}