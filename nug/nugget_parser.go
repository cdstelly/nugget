// Generated from C:/Users/drew/GoglandProjects/nugget\Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// import "../NTypes"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 147,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 7, 2, 38, 10, 2, 12, 2, 14, 2, 41, 11, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 5, 3, 48, 10, 3, 3, 4, 3, 4, 3, 4, 5, 4, 53, 10, 4, 3, 5, 3, 5,
	3, 5, 5, 5, 58, 10, 5, 3, 5, 6, 5, 61, 10, 5, 13, 5, 14, 5, 62, 3, 5, 3,
	5, 5, 5, 67, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 74, 10, 6, 12,
	6, 14, 6, 77, 11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 84, 10, 6, 12,
	6, 14, 6, 87, 11, 6, 5, 6, 89, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	5, 12, 118, 10, 12, 3, 13, 3, 13, 3, 13, 5, 13, 123, 10, 13, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 7, 16, 138, 10, 16, 12, 16, 14, 16, 141, 11, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 2, 4, 3, 2, 10, 15, 3, 2, 16, 28, 2, 154, 2, 39, 3, 2,
	2, 2, 4, 47, 3, 2, 2, 2, 6, 49, 3, 2, 2, 2, 8, 54, 3, 2, 2, 2, 10, 88,
	3, 2, 2, 2, 12, 90, 3, 2, 2, 2, 14, 95, 3, 2, 2, 2, 16, 97, 3, 2, 2, 2,
	18, 99, 3, 2, 2, 2, 20, 101, 3, 2, 2, 2, 22, 117, 3, 2, 2, 2, 24, 119,
	3, 2, 2, 2, 26, 124, 3, 2, 2, 2, 28, 127, 3, 2, 2, 2, 30, 133, 3, 2, 2,
	2, 32, 142, 3, 2, 2, 2, 34, 38, 5, 4, 3, 2, 35, 38, 5, 12, 7, 2, 36, 38,
	5, 16, 9, 2, 37, 34, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 36, 3, 2, 2, 2,
	38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 42, 3,
	2, 2, 2, 41, 39, 3, 2, 2, 2, 42, 43, 7, 2, 2, 3, 43, 3, 3, 2, 2, 2, 44,
	48, 5, 6, 4, 2, 45, 48, 5, 8, 5, 2, 46, 48, 5, 10, 6, 2, 47, 44, 3, 2,
	2, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 5, 3, 2, 2, 2, 49, 50,
	7, 44, 2, 2, 50, 52, 5, 18, 10, 2, 51, 53, 7, 42, 2, 2, 52, 51, 3, 2, 2,
	2, 52, 53, 3, 2, 2, 2, 53, 7, 3, 2, 2, 2, 54, 55, 7, 44, 2, 2, 55, 60,
	7, 3, 2, 2, 56, 58, 7, 4, 2, 2, 57, 56, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2,
	58, 59, 3, 2, 2, 2, 59, 61, 5, 18, 10, 2, 60, 57, 3, 2, 2, 2, 61, 62, 3,
	2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64,
	66, 7, 5, 2, 2, 65, 67, 7, 42, 2, 2, 66, 65, 3, 2, 2, 2, 66, 67, 3, 2,
	2, 2, 67, 9, 3, 2, 2, 2, 68, 69, 7, 44, 2, 2, 69, 70, 7, 6, 2, 2, 70, 75,
	7, 45, 2, 2, 71, 72, 7, 7, 2, 2, 72, 74, 5, 20, 11, 2, 73, 71, 3, 2, 2,
	2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 89,
	3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 79, 7, 44, 2, 2, 79, 80, 7, 6, 2, 2,
	80, 85, 7, 44, 2, 2, 81, 82, 7, 7, 2, 2, 82, 84, 5, 20, 11, 2, 83, 81,
	3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2,
	86, 89, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 68, 3, 2, 2, 2, 88, 78, 3,
	2, 2, 2, 89, 11, 3, 2, 2, 2, 90, 91, 5, 14, 8, 2, 91, 92, 7, 8, 2, 2, 92,
	93, 7, 44, 2, 2, 93, 94, 7, 9, 2, 2, 94, 13, 3, 2, 2, 2, 95, 96, 9, 2,
	2, 2, 96, 15, 3, 2, 2, 2, 97, 98, 7, 44, 2, 2, 98, 17, 3, 2, 2, 2, 99,
	100, 9, 3, 2, 2, 100, 19, 3, 2, 2, 2, 101, 102, 5, 22, 12, 2, 102, 21,
	3, 2, 2, 2, 103, 118, 5, 30, 16, 2, 104, 105, 7, 29, 2, 2, 105, 118, 5,
	24, 13, 2, 106, 107, 7, 30, 2, 2, 107, 118, 5, 26, 14, 2, 108, 118, 7,
	17, 2, 2, 109, 118, 7, 18, 2, 2, 110, 118, 7, 31, 2, 2, 111, 118, 7, 32,
	2, 2, 112, 118, 7, 33, 2, 2, 113, 114, 7, 34, 2, 2, 114, 118, 7, 44, 2,
	2, 115, 118, 7, 35, 2, 2, 116, 118, 7, 36, 2, 2, 117, 103, 3, 2, 2, 2,
	117, 104, 3, 2, 2, 2, 117, 106, 3, 2, 2, 2, 117, 108, 3, 2, 2, 2, 117,
	109, 3, 2, 2, 2, 117, 110, 3, 2, 2, 2, 117, 111, 3, 2, 2, 2, 117, 112,
	3, 2, 2, 2, 117, 113, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 116, 3, 2,
	2, 2, 118, 23, 3, 2, 2, 2, 119, 120, 7, 37, 2, 2, 120, 122, 5, 18, 10,
	2, 121, 123, 5, 28, 15, 2, 122, 121, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2,
	123, 25, 3, 2, 2, 2, 124, 125, 7, 38, 2, 2, 125, 126, 7, 44, 2, 2, 126,
	27, 3, 2, 2, 2, 127, 128, 7, 39, 2, 2, 128, 129, 7, 43, 2, 2, 129, 130,
	7, 4, 2, 2, 130, 131, 7, 43, 2, 2, 131, 132, 7, 5, 2, 2, 132, 29, 3, 2,
	2, 2, 133, 134, 7, 40, 2, 2, 134, 139, 5, 32, 17, 2, 135, 136, 7, 4, 2,
	2, 136, 138, 5, 32, 17, 2, 137, 135, 3, 2, 2, 2, 138, 141, 3, 2, 2, 2,
	139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 31, 3, 2, 2, 2, 141, 139,
	3, 2, 2, 2, 142, 143, 7, 44, 2, 2, 143, 144, 7, 41, 2, 2, 144, 145, 7,
	45, 2, 2, 145, 33, 3, 2, 2, 2, 15, 37, 39, 47, 52, 57, 62, 66, 75, 85,
	88, 117, 122, 139,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'tuple['", "','", "']'", "'='", "'|'", "'('", "')'", "'type'", "'print'",
	"'size'", "'typex'", "'printx'", "'raw'", "'string'", "'sha1'", "'md5'",
	"'ntfs'", "'file'", "'packet'", "'pcap'", "'exifinfo'", "'datetime'", "'memory'",
	"'listof-md5'", "'listof-sha1'", "'listof-sha256'", "'extract'", "'sort'",
	"'sha256'", "'getGetRequests'", "'diskinfo'", "'union'", "'pslist'", "'%%%'",
	"'as'", "'by'", "'['", "'filter'", "", "'[]'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "COMPOP", "LISTOP", "INT", "ID", "STRING", "WS", "NL", "LINE_COMMENT",
}

var ruleNames = []string{
	"prog", "define_assign", "define", "define_tuple", "assign", "operation_on_singleton",
	"singleton_op", "singleton_var", "nugget_type", "nugget_action", "action_word",
	"asType", "byField", "byteOffsetSize", "filter", "filter_term",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type NuggetParser struct {
	*antlr.BaseParser
}

func NewNuggetParser(input antlr.TokenStream) *NuggetParser {
	this := new(NuggetParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Nugget.g4"

	return this
}

// NuggetParser tokens.
const (
	NuggetParserEOF          = antlr.TokenEOF
	NuggetParserT__0         = 1
	NuggetParserT__1         = 2
	NuggetParserT__2         = 3
	NuggetParserT__3         = 4
	NuggetParserT__4         = 5
	NuggetParserT__5         = 6
	NuggetParserT__6         = 7
	NuggetParserT__7         = 8
	NuggetParserT__8         = 9
	NuggetParserT__9         = 10
	NuggetParserT__10        = 11
	NuggetParserT__11        = 12
	NuggetParserT__12        = 13
	NuggetParserT__13        = 14
	NuggetParserT__14        = 15
	NuggetParserT__15        = 16
	NuggetParserT__16        = 17
	NuggetParserT__17        = 18
	NuggetParserT__18        = 19
	NuggetParserT__19        = 20
	NuggetParserT__20        = 21
	NuggetParserT__21        = 22
	NuggetParserT__22        = 23
	NuggetParserT__23        = 24
	NuggetParserT__24        = 25
	NuggetParserT__25        = 26
	NuggetParserT__26        = 27
	NuggetParserT__27        = 28
	NuggetParserT__28        = 29
	NuggetParserT__29        = 30
	NuggetParserT__30        = 31
	NuggetParserT__31        = 32
	NuggetParserT__32        = 33
	NuggetParserT__33        = 34
	NuggetParserT__34        = 35
	NuggetParserT__35        = 36
	NuggetParserT__36        = 37
	NuggetParserT__37        = 38
	NuggetParserCOMPOP       = 39
	NuggetParserLISTOP       = 40
	NuggetParserINT          = 41
	NuggetParserID           = 42
	NuggetParserSTRING       = 43
	NuggetParserWS           = 44
	NuggetParserNL           = 45
	NuggetParserLINE_COMMENT = 46
)

// NuggetParser rules.
const (
	NuggetParserRULE_prog                   = 0
	NuggetParserRULE_define_assign          = 1
	NuggetParserRULE_define                 = 2
	NuggetParserRULE_define_tuple           = 3
	NuggetParserRULE_assign                 = 4
	NuggetParserRULE_operation_on_singleton = 5
	NuggetParserRULE_singleton_op           = 6
	NuggetParserRULE_singleton_var          = 7
	NuggetParserRULE_nugget_type            = 8
	NuggetParserRULE_nugget_action          = 9
	NuggetParserRULE_action_word            = 10
	NuggetParserRULE_asType                 = 11
	NuggetParserRULE_byField                = 12
	NuggetParserRULE_byteOffsetSize         = 13
	NuggetParserRULE_filter                 = 14
	NuggetParserRULE_filter_term            = 15
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(NuggetParserEOF, 0)
}

func (s *ProgContext) AllDefine_assign() []IDefine_assignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefine_assignContext)(nil)).Elem())
	var tst = make([]IDefine_assignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefine_assignContext)
		}
	}

	return tst
}

func (s *ProgContext) Define_assign(i int) IDefine_assignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefine_assignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefine_assignContext)
}

func (s *ProgContext) AllOperation_on_singleton() []IOperation_on_singletonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperation_on_singletonContext)(nil)).Elem())
	var tst = make([]IOperation_on_singletonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperation_on_singletonContext)
		}
	}

	return tst
}

func (s *ProgContext) Operation_on_singleton(i int) IOperation_on_singletonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperation_on_singletonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperation_on_singletonContext)
}

func (s *ProgContext) AllSingleton_var() []ISingleton_varContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleton_varContext)(nil)).Elem())
	var tst = make([]ISingleton_varContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleton_varContext)
		}
	}

	return tst
}

func (s *ProgContext) Singleton_var(i int) ISingleton_varContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleton_varContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleton_varContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *NuggetParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NuggetParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserT__7)|(1<<NuggetParserT__8)|(1<<NuggetParserT__9)|(1<<NuggetParserT__10)|(1<<NuggetParserT__11)|(1<<NuggetParserT__12))) != 0) || _la == NuggetParserID {
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(32)
				p.Define_assign()
			}

		case 2:
			{
				p.SetState(33)
				p.Operation_on_singleton()
			}

		case 3:
			{
				p.SetState(34)
				p.Singleton_var()
			}

		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(NuggetParserEOF)
	}

	return localctx
}

// IDefine_assignContext is an interface to support dynamic dispatch.
type IDefine_assignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefine_assignContext differentiates from other interfaces.
	IsDefine_assignContext()
}

type Define_assignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefine_assignContext() *Define_assignContext {
	var p = new(Define_assignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_define_assign
	return p
}

func (*Define_assignContext) IsDefine_assignContext() {}

func NewDefine_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Define_assignContext {
	var p = new(Define_assignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_define_assign

	return p
}

func (s *Define_assignContext) GetParser() antlr.Parser { return s.parser }

func (s *Define_assignContext) Define() IDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefineContext)
}

func (s *Define_assignContext) Define_tuple() IDefine_tupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefine_tupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefine_tupleContext)
}

func (s *Define_assignContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *Define_assignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Define_assignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Define_assignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterDefine_assign(s)
	}
}

func (s *Define_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitDefine_assign(s)
	}
}

func (p *NuggetParser) Define_assign() (localctx IDefine_assignContext) {
	localctx = NewDefine_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NuggetParserRULE_define_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Define()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Define_tuple()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Assign()
		}

	}

	return localctx
}

// IDefineContext is an interface to support dynamic dispatch.
type IDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefineContext differentiates from other interfaces.
	IsDefineContext()
}

type DefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefineContext() *DefineContext {
	var p = new(DefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_define
	return p
}

func (*DefineContext) IsDefineContext() {}

func NewDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineContext {
	var p = new(DefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_define

	return p
}

func (s *DefineContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineContext) ID() antlr.TerminalNode {
	return s.GetToken(NuggetParserID, 0)
}

func (s *DefineContext) Nugget_type() INugget_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INugget_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INugget_typeContext)
}

func (s *DefineContext) LISTOP() antlr.TerminalNode {
	return s.GetToken(NuggetParserLISTOP, 0)
}

func (s *DefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterDefine(s)
	}
}

func (s *DefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitDefine(s)
	}
}

func (p *NuggetParser) Define() (localctx IDefineContext) {
	localctx = NewDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NuggetParserRULE_define)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(NuggetParserID)
	}
	{
		p.SetState(48)
		p.Nugget_type()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NuggetParserLISTOP {
		{
			p.SetState(49)
			p.Match(NuggetParserLISTOP)
		}

	}

	return localctx
}

// IDefine_tupleContext is an interface to support dynamic dispatch.
type IDefine_tupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefine_tupleContext differentiates from other interfaces.
	IsDefine_tupleContext()
}

type Define_tupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefine_tupleContext() *Define_tupleContext {
	var p = new(Define_tupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_define_tuple
	return p
}

func (*Define_tupleContext) IsDefine_tupleContext() {}

func NewDefine_tupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Define_tupleContext {
	var p = new(Define_tupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_define_tuple

	return p
}

func (s *Define_tupleContext) GetParser() antlr.Parser { return s.parser }

func (s *Define_tupleContext) ID() antlr.TerminalNode {
	return s.GetToken(NuggetParserID, 0)
}

func (s *Define_tupleContext) AllNugget_type() []INugget_typeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INugget_typeContext)(nil)).Elem())
	var tst = make([]INugget_typeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INugget_typeContext)
		}
	}

	return tst
}

func (s *Define_tupleContext) Nugget_type(i int) INugget_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INugget_typeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INugget_typeContext)
}

func (s *Define_tupleContext) LISTOP() antlr.TerminalNode {
	return s.GetToken(NuggetParserLISTOP, 0)
}

func (s *Define_tupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Define_tupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Define_tupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterDefine_tuple(s)
	}
}

func (s *Define_tupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitDefine_tuple(s)
	}
}

func (p *NuggetParser) Define_tuple() (localctx IDefine_tupleContext) {
	localctx = NewDefine_tupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NuggetParserRULE_define_tuple)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(NuggetParserID)
	}
	{
		p.SetState(53)
		p.Match(NuggetParserT__0)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserT__1)|(1<<NuggetParserT__13)|(1<<NuggetParserT__14)|(1<<NuggetParserT__15)|(1<<NuggetParserT__16)|(1<<NuggetParserT__17)|(1<<NuggetParserT__18)|(1<<NuggetParserT__19)|(1<<NuggetParserT__20)|(1<<NuggetParserT__21)|(1<<NuggetParserT__22)|(1<<NuggetParserT__23)|(1<<NuggetParserT__24)|(1<<NuggetParserT__25))) != 0) {
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NuggetParserT__1 {
			{
				p.SetState(54)
				p.Match(NuggetParserT__1)
			}

		}
		{
			p.SetState(57)
			p.Nugget_type()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(NuggetParserT__2)
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NuggetParserLISTOP {
		{
			p.SetState(63)
			p.Match(NuggetParserLISTOP)
		}

	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(NuggetParserID)
}

func (s *AssignContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(NuggetParserID, i)
}

func (s *AssignContext) STRING() antlr.TerminalNode {
	return s.GetToken(NuggetParserSTRING, 0)
}

func (s *AssignContext) AllNugget_action() []INugget_actionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INugget_actionContext)(nil)).Elem())
	var tst = make([]INugget_actionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INugget_actionContext)
		}
	}

	return tst
}

func (s *AssignContext) Nugget_action(i int) INugget_actionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INugget_actionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INugget_actionContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *NuggetParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NuggetParserRULE_assign)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Match(NuggetParserID)
		}
		{
			p.SetState(67)
			p.Match(NuggetParserT__3)
		}
		{
			p.SetState(68)
			p.Match(NuggetParserSTRING)
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__4 {
			{
				p.SetState(69)
				p.Match(NuggetParserT__4)
			}
			{
				p.SetState(70)
				p.Nugget_action()
			}

			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Match(NuggetParserID)
		}
		{
			p.SetState(77)
			p.Match(NuggetParserT__3)
		}
		{
			p.SetState(78)
			p.Match(NuggetParserID)
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__4 {
			{
				p.SetState(79)
				p.Match(NuggetParserT__4)
			}
			{
				p.SetState(80)
				p.Nugget_action()
			}

			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IOperation_on_singletonContext is an interface to support dynamic dispatch.
type IOperation_on_singletonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperation_on_singletonContext differentiates from other interfaces.
	IsOperation_on_singletonContext()
}

type Operation_on_singletonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperation_on_singletonContext() *Operation_on_singletonContext {
	var p = new(Operation_on_singletonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_operation_on_singleton
	return p
}

func (*Operation_on_singletonContext) IsOperation_on_singletonContext() {}

func NewOperation_on_singletonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operation_on_singletonContext {
	var p = new(Operation_on_singletonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_operation_on_singleton

	return p
}

func (s *Operation_on_singletonContext) GetParser() antlr.Parser { return s.parser }

func (s *Operation_on_singletonContext) Singleton_op() ISingleton_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleton_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleton_opContext)
}

func (s *Operation_on_singletonContext) ID() antlr.TerminalNode {
	return s.GetToken(NuggetParserID, 0)
}

func (s *Operation_on_singletonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operation_on_singletonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operation_on_singletonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterOperation_on_singleton(s)
	}
}

func (s *Operation_on_singletonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitOperation_on_singleton(s)
	}
}

func (p *NuggetParser) Operation_on_singleton() (localctx IOperation_on_singletonContext) {
	localctx = NewOperation_on_singletonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NuggetParserRULE_operation_on_singleton)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Singleton_op()
	}
	{
		p.SetState(89)
		p.Match(NuggetParserT__5)
	}
	{
		p.SetState(90)
		p.Match(NuggetParserID)
	}
	{
		p.SetState(91)
		p.Match(NuggetParserT__6)
	}

	return localctx
}

// ISingleton_opContext is an interface to support dynamic dispatch.
type ISingleton_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleton_opContext differentiates from other interfaces.
	IsSingleton_opContext()
}

type Singleton_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleton_opContext() *Singleton_opContext {
	var p = new(Singleton_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_singleton_op
	return p
}

func (*Singleton_opContext) IsSingleton_opContext() {}

func NewSingleton_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Singleton_opContext {
	var p = new(Singleton_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_singleton_op

	return p
}

func (s *Singleton_opContext) GetParser() antlr.Parser { return s.parser }
func (s *Singleton_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Singleton_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Singleton_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterSingleton_op(s)
	}
}

func (s *Singleton_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitSingleton_op(s)
	}
}

func (p *NuggetParser) Singleton_op() (localctx ISingleton_opContext) {
	localctx = NewSingleton_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NuggetParserRULE_singleton_op)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserT__7)|(1<<NuggetParserT__8)|(1<<NuggetParserT__9)|(1<<NuggetParserT__10)|(1<<NuggetParserT__11)|(1<<NuggetParserT__12))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ISingleton_varContext is an interface to support dynamic dispatch.
type ISingleton_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleton_varContext differentiates from other interfaces.
	IsSingleton_varContext()
}

type Singleton_varContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleton_varContext() *Singleton_varContext {
	var p = new(Singleton_varContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_singleton_var
	return p
}

func (*Singleton_varContext) IsSingleton_varContext() {}

func NewSingleton_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Singleton_varContext {
	var p = new(Singleton_varContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_singleton_var

	return p
}

func (s *Singleton_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Singleton_varContext) ID() antlr.TerminalNode {
	return s.GetToken(NuggetParserID, 0)
}

func (s *Singleton_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Singleton_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Singleton_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterSingleton_var(s)
	}
}

func (s *Singleton_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitSingleton_var(s)
	}
}

func (p *NuggetParser) Singleton_var() (localctx ISingleton_varContext) {
	localctx = NewSingleton_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NuggetParserRULE_singleton_var)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(NuggetParserID)
	}

	return localctx
}

// INugget_typeContext is an interface to support dynamic dispatch.
type INugget_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNugget_typeContext differentiates from other interfaces.
	IsNugget_typeContext()
}

type Nugget_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNugget_typeContext() *Nugget_typeContext {
	var p = new(Nugget_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_nugget_type
	return p
}

func (*Nugget_typeContext) IsNugget_typeContext() {}

func NewNugget_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nugget_typeContext {
	var p = new(Nugget_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_nugget_type

	return p
}

func (s *Nugget_typeContext) GetParser() antlr.Parser { return s.parser }
func (s *Nugget_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nugget_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nugget_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterNugget_type(s)
	}
}

func (s *Nugget_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitNugget_type(s)
	}
}

func (p *NuggetParser) Nugget_type() (localctx INugget_typeContext) {
	localctx = NewNugget_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NuggetParserRULE_nugget_type)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserT__13)|(1<<NuggetParserT__14)|(1<<NuggetParserT__15)|(1<<NuggetParserT__16)|(1<<NuggetParserT__17)|(1<<NuggetParserT__18)|(1<<NuggetParserT__19)|(1<<NuggetParserT__20)|(1<<NuggetParserT__21)|(1<<NuggetParserT__22)|(1<<NuggetParserT__23)|(1<<NuggetParserT__24)|(1<<NuggetParserT__25))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// INugget_actionContext is an interface to support dynamic dispatch.
type INugget_actionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNugget_actionContext differentiates from other interfaces.
	IsNugget_actionContext()
}

type Nugget_actionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNugget_actionContext() *Nugget_actionContext {
	var p = new(Nugget_actionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_nugget_action
	return p
}

func (*Nugget_actionContext) IsNugget_actionContext() {}

func NewNugget_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nugget_actionContext {
	var p = new(Nugget_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_nugget_action

	return p
}

func (s *Nugget_actionContext) GetParser() antlr.Parser { return s.parser }

func (s *Nugget_actionContext) Action_word() IAction_wordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAction_wordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAction_wordContext)
}

func (s *Nugget_actionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nugget_actionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nugget_actionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterNugget_action(s)
	}
}

func (s *Nugget_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitNugget_action(s)
	}
}

func (p *NuggetParser) Nugget_action() (localctx INugget_actionContext) {
	localctx = NewNugget_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NuggetParserRULE_nugget_action)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Action_word()
	}

	return localctx
}

// IAction_wordContext is an interface to support dynamic dispatch.
type IAction_wordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAction_wordContext differentiates from other interfaces.
	IsAction_wordContext()
}

type Action_wordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAction_wordContext() *Action_wordContext {
	var p = new(Action_wordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_action_word
	return p
}

func (*Action_wordContext) IsAction_wordContext() {}

func NewAction_wordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_wordContext {
	var p = new(Action_wordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_action_word

	return p
}

func (s *Action_wordContext) GetParser() antlr.Parser { return s.parser }

func (s *Action_wordContext) Filter() IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *Action_wordContext) AsType() IAsTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsTypeContext)
}

func (s *Action_wordContext) ByField() IByFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IByFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IByFieldContext)
}

func (s *Action_wordContext) ID() antlr.TerminalNode {
	return s.GetToken(NuggetParserID, 0)
}

func (s *Action_wordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_wordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_wordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterAction_word(s)
	}
}

func (s *Action_wordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitAction_word(s)
	}
}

func (p *NuggetParser) Action_word() (localctx IAction_wordContext) {
	localctx = NewAction_wordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NuggetParserRULE_action_word)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserT__37:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Filter()
		}

	case NuggetParserT__26:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(NuggetParserT__26)
		}
		{
			p.SetState(103)
			p.AsType()
		}

	case NuggetParserT__27:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Match(NuggetParserT__27)
		}
		{
			p.SetState(105)
			p.ByField()
		}

	case NuggetParserT__14:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(106)
			p.Match(NuggetParserT__14)
		}

	case NuggetParserT__15:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(107)
			p.Match(NuggetParserT__15)
		}

	case NuggetParserT__28:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(108)
			p.Match(NuggetParserT__28)
		}

	case NuggetParserT__29:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(109)
			p.Match(NuggetParserT__29)
		}

	case NuggetParserT__30:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(110)
			p.Match(NuggetParserT__30)
		}

	case NuggetParserT__31:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(111)
			p.Match(NuggetParserT__31)
		}
		{
			p.SetState(112)
			p.Match(NuggetParserID)
		}

	case NuggetParserT__32:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(113)
			p.Match(NuggetParserT__32)
		}

	case NuggetParserT__33:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(114)
			p.Match(NuggetParserT__33)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAsTypeContext is an interface to support dynamic dispatch.
type IAsTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAsTypeContext differentiates from other interfaces.
	IsAsTypeContext()
}

type AsTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsTypeContext() *AsTypeContext {
	var p = new(AsTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_asType
	return p
}

func (*AsTypeContext) IsAsTypeContext() {}

func NewAsTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsTypeContext {
	var p = new(AsTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_asType

	return p
}

func (s *AsTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AsTypeContext) Nugget_type() INugget_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INugget_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INugget_typeContext)
}

func (s *AsTypeContext) ByteOffsetSize() IByteOffsetSizeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IByteOffsetSizeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IByteOffsetSizeContext)
}

func (s *AsTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterAsType(s)
	}
}

func (s *AsTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitAsType(s)
	}
}

func (p *NuggetParser) AsType() (localctx IAsTypeContext) {
	localctx = NewAsTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NuggetParserRULE_asType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(NuggetParserT__34)
	}
	{
		p.SetState(118)
		p.Nugget_type()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NuggetParserT__36 {
		{
			p.SetState(119)
			p.ByteOffsetSize()
		}

	}

	return localctx
}

// IByFieldContext is an interface to support dynamic dispatch.
type IByFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsByFieldContext differentiates from other interfaces.
	IsByFieldContext()
}

type ByFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyByFieldContext() *ByFieldContext {
	var p = new(ByFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_byField
	return p
}

func (*ByFieldContext) IsByFieldContext() {}

func NewByFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByFieldContext {
	var p = new(ByFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_byField

	return p
}

func (s *ByFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ByFieldContext) ID() antlr.TerminalNode {
	return s.GetToken(NuggetParserID, 0)
}

func (s *ByFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ByFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterByField(s)
	}
}

func (s *ByFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitByField(s)
	}
}

func (p *NuggetParser) ByField() (localctx IByFieldContext) {
	localctx = NewByFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NuggetParserRULE_byField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(NuggetParserT__35)
	}
	{
		p.SetState(123)
		p.Match(NuggetParserID)
	}

	return localctx
}

// IByteOffsetSizeContext is an interface to support dynamic dispatch.
type IByteOffsetSizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsByteOffsetSizeContext differentiates from other interfaces.
	IsByteOffsetSizeContext()
}

type ByteOffsetSizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyByteOffsetSizeContext() *ByteOffsetSizeContext {
	var p = new(ByteOffsetSizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_byteOffsetSize
	return p
}

func (*ByteOffsetSizeContext) IsByteOffsetSizeContext() {}

func NewByteOffsetSizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByteOffsetSizeContext {
	var p = new(ByteOffsetSizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_byteOffsetSize

	return p
}

func (s *ByteOffsetSizeContext) GetParser() antlr.Parser { return s.parser }

func (s *ByteOffsetSizeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(NuggetParserINT)
}

func (s *ByteOffsetSizeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(NuggetParserINT, i)
}

func (s *ByteOffsetSizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByteOffsetSizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ByteOffsetSizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterByteOffsetSize(s)
	}
}

func (s *ByteOffsetSizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitByteOffsetSize(s)
	}
}

func (p *NuggetParser) ByteOffsetSize() (localctx IByteOffsetSizeContext) {
	localctx = NewByteOffsetSizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NuggetParserRULE_byteOffsetSize)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(NuggetParserT__36)
	}
	{
		p.SetState(126)
		p.Match(NuggetParserINT)
	}
	{
		p.SetState(127)
		p.Match(NuggetParserT__1)
	}
	{
		p.SetState(128)
		p.Match(NuggetParserINT)
	}
	{
		p.SetState(129)
		p.Match(NuggetParserT__2)
	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) AllFilter_term() []IFilter_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilter_termContext)(nil)).Elem())
	var tst = make([]IFilter_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilter_termContext)
		}
	}

	return tst
}

func (s *FilterContext) Filter_term(i int) IFilter_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilter_termContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *NuggetParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NuggetParserRULE_filter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(NuggetParserT__37)
	}
	{
		p.SetState(132)
		p.Filter_term()
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NuggetParserT__1 {
		{
			p.SetState(133)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(134)
			p.Filter_term()
		}

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFilter_termContext is an interface to support dynamic dispatch.
type IFilter_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilter_termContext differentiates from other interfaces.
	IsFilter_termContext()
}

type Filter_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_termContext() *Filter_termContext {
	var p = new(Filter_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_filter_term
	return p
}

func (*Filter_termContext) IsFilter_termContext() {}

func NewFilter_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_termContext {
	var p = new(Filter_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_filter_term

	return p
}

func (s *Filter_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_termContext) ID() antlr.TerminalNode {
	return s.GetToken(NuggetParserID, 0)
}

func (s *Filter_termContext) COMPOP() antlr.TerminalNode {
	return s.GetToken(NuggetParserCOMPOP, 0)
}

func (s *Filter_termContext) STRING() antlr.TerminalNode {
	return s.GetToken(NuggetParserSTRING, 0)
}

func (s *Filter_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterFilter_term(s)
	}
}

func (s *Filter_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitFilter_term(s)
	}
}

func (p *NuggetParser) Filter_term() (localctx IFilter_termContext) {
	localctx = NewFilter_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NuggetParserRULE_filter_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(NuggetParserID)
	}
	{
		p.SetState(141)
		p.Match(NuggetParserCOMPOP)
	}
	{
		p.SetState(142)
		p.Match(NuggetParserSTRING)
	}

	return localctx
}
