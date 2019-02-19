// Generated from /home/cyclops/code/nugget/Nugget.g4 by ANTLR 4.7.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 48, 163,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 7, 2, 42, 10, 2, 12, 2, 14, 2, 45,
	11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 52, 10, 3, 3, 4, 3, 4, 3, 4,
	5, 4, 57, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 62, 10, 5, 3, 5, 6, 5, 65, 10,
	5, 13, 5, 14, 5, 66, 3, 5, 3, 5, 5, 5, 71, 10, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 7, 6, 78, 10, 6, 12, 6, 14, 6, 81, 11, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 7, 6, 88, 10, 6, 12, 6, 14, 6, 91, 11, 6, 5, 6, 93, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 7, 7, 99, 10, 7, 12, 7, 14, 7, 102, 11, 7, 3, 7, 5,
	7, 105, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 134, 10, 14,
	3, 15, 3, 15, 3, 15, 5, 15, 139, 10, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 154,
	10, 18, 12, 18, 14, 18, 157, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	2, 2, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 2, 4, 3, 2, 10, 15, 3, 2, 16, 29, 2, 170, 2, 43, 3, 2, 2, 2, 4, 51,
	3, 2, 2, 2, 6, 53, 3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 92, 3, 2, 2, 2, 12,
	94, 3, 2, 2, 2, 14, 106, 3, 2, 2, 2, 16, 109, 3, 2, 2, 2, 18, 111, 3, 2,
	2, 2, 20, 113, 3, 2, 2, 2, 22, 115, 3, 2, 2, 2, 24, 117, 3, 2, 2, 2, 26,
	133, 3, 2, 2, 2, 28, 135, 3, 2, 2, 2, 30, 140, 3, 2, 2, 2, 32, 143, 3,
	2, 2, 2, 34, 149, 3, 2, 2, 2, 36, 158, 3, 2, 2, 2, 38, 42, 5, 4, 3, 2,
	39, 42, 5, 12, 7, 2, 40, 42, 5, 20, 11, 2, 41, 38, 3, 2, 2, 2, 41, 39,
	3, 2, 2, 2, 41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2,
	43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 47, 7,
	2, 2, 3, 47, 3, 3, 2, 2, 2, 48, 52, 5, 6, 4, 2, 49, 52, 5, 8, 5, 2, 50,
	52, 5, 10, 6, 2, 51, 48, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2,
	2, 2, 52, 5, 3, 2, 2, 2, 53, 54, 7, 44, 2, 2, 54, 56, 5, 22, 12, 2, 55,
	57, 7, 42, 2, 2, 56, 55, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 7, 3, 2, 2,
	2, 58, 59, 7, 44, 2, 2, 59, 64, 7, 3, 2, 2, 60, 62, 7, 4, 2, 2, 61, 60,
	3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 5, 22, 12,
	2, 64, 61, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67,
	3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 7, 5, 2, 2, 69, 71, 7, 42, 2, 2,
	70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 9, 3, 2, 2, 2, 72, 73, 7, 44,
	2, 2, 73, 74, 7, 6, 2, 2, 74, 79, 7, 45, 2, 2, 75, 76, 7, 7, 2, 2, 76,
	78, 5, 24, 13, 2, 77, 75, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2,
	2, 2, 79, 80, 3, 2, 2, 2, 80, 93, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 83,
	7, 44, 2, 2, 83, 84, 7, 6, 2, 2, 84, 89, 7, 44, 2, 2, 85, 86, 7, 7, 2,
	2, 86, 88, 5, 24, 13, 2, 87, 85, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87,
	3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2,
	92, 72, 3, 2, 2, 2, 92, 82, 3, 2, 2, 2, 93, 11, 3, 2, 2, 2, 94, 95, 5,
	18, 10, 2, 95, 100, 7, 44, 2, 2, 96, 97, 7, 4, 2, 2, 97, 99, 7, 44, 2,
	2, 98, 96, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101,
	3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 5, 14,
	8, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 13, 3, 2, 2, 2,
	106, 107, 7, 8, 2, 2, 107, 108, 5, 16, 9, 2, 108, 15, 3, 2, 2, 2, 109,
	110, 7, 9, 2, 2, 110, 17, 3, 2, 2, 2, 111, 112, 9, 2, 2, 2, 112, 19, 3,
	2, 2, 2, 113, 114, 7, 44, 2, 2, 114, 21, 3, 2, 2, 2, 115, 116, 9, 3, 2,
	2, 116, 23, 3, 2, 2, 2, 117, 118, 5, 26, 14, 2, 118, 25, 3, 2, 2, 2, 119,
	134, 5, 34, 18, 2, 120, 121, 7, 30, 2, 2, 121, 134, 5, 28, 15, 2, 122,
	123, 7, 31, 2, 2, 123, 134, 5, 30, 16, 2, 124, 134, 7, 17, 2, 2, 125, 134,
	7, 18, 2, 2, 126, 134, 7, 32, 2, 2, 127, 134, 7, 33, 2, 2, 128, 134, 7,
	34, 2, 2, 129, 130, 7, 35, 2, 2, 130, 134, 7, 44, 2, 2, 131, 134, 7, 36,
	2, 2, 132, 134, 7, 37, 2, 2, 133, 119, 3, 2, 2, 2, 133, 120, 3, 2, 2, 2,
	133, 122, 3, 2, 2, 2, 133, 124, 3, 2, 2, 2, 133, 125, 3, 2, 2, 2, 133,
	126, 3, 2, 2, 2, 133, 127, 3, 2, 2, 2, 133, 128, 3, 2, 2, 2, 133, 129,
	3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 132, 3, 2, 2, 2, 134, 27, 3, 2,
	2, 2, 135, 136, 7, 8, 2, 2, 136, 138, 5, 22, 12, 2, 137, 139, 5, 32, 17,
	2, 138, 137, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 29, 3, 2, 2, 2, 140,
	141, 7, 38, 2, 2, 141, 142, 7, 44, 2, 2, 142, 31, 3, 2, 2, 2, 143, 144,
	7, 39, 2, 2, 144, 145, 7, 43, 2, 2, 145, 146, 7, 4, 2, 2, 146, 147, 7,
	43, 2, 2, 147, 148, 7, 5, 2, 2, 148, 33, 3, 2, 2, 2, 149, 150, 7, 40, 2,
	2, 150, 155, 5, 36, 19, 2, 151, 152, 7, 4, 2, 2, 152, 154, 5, 36, 19, 2,
	153, 151, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155,
	156, 3, 2, 2, 2, 156, 35, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 158, 159, 7,
	44, 2, 2, 159, 160, 7, 41, 2, 2, 160, 161, 7, 45, 2, 2, 161, 37, 3, 2,
	2, 2, 17, 41, 43, 51, 56, 61, 66, 70, 79, 89, 92, 100, 104, 133, 138, 155,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'tuple['", "','", "']'", "'='", "'|'", "'as'", "'json'", "'type'",
	"'print'", "'size'", "'typex'", "'printx'", "'raw'", "'string'", "'sha1'",
	"'md5'", "'ntfs'", "'file'", "'packet'", "'pcap'", "'exifinfo'", "'datetime'",
	"'memory'", "'http'", "'listof-md5'", "'listof-sha1'", "'listof-sha256'",
	"'extract'", "'sort'", "'sha256'", "'getGetRequests'", "'diskinfo'", "'union'",
	"'pslist'", "'%%%'", "'by'", "'['", "'filter'", "", "'[]'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "COMPOP", "LISTOP", "INT", "ID", "STRING", "WS", "NL", "LINE_COMMENT",
}

var ruleNames = []string{
	"prog", "define_assign", "define", "define_tuple", "assign", "operation_on_singleton",
	"output_as", "output_type", "singleton_op", "singleton_var", "nugget_type",
	"nugget_action", "action_word", "asType", "byField", "byteOffsetSize",
	"filter", "filter_term",
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
	NuggetParserRULE_output_as              = 6
	NuggetParserRULE_output_type            = 7
	NuggetParserRULE_singleton_op           = 8
	NuggetParserRULE_singleton_var          = 9
	NuggetParserRULE_nugget_type            = 10
	NuggetParserRULE_nugget_action          = 11
	NuggetParserRULE_action_word            = 12
	NuggetParserRULE_asType                 = 13
	NuggetParserRULE_byField                = 14
	NuggetParserRULE_byteOffsetSize         = 15
	NuggetParserRULE_filter                 = 16
	NuggetParserRULE_filter_term            = 17
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserT__7)|(1<<NuggetParserT__8)|(1<<NuggetParserT__9)|(1<<NuggetParserT__10)|(1<<NuggetParserT__11)|(1<<NuggetParserT__12))) != 0) || _la == NuggetParserID {
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(36)
				p.Define_assign()
			}

		case 2:
			{
				p.SetState(37)
				p.Operation_on_singleton()
			}

		case 3:
			{
				p.SetState(38)
				p.Singleton_var()
			}

		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)
			p.Define()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)
			p.Define_tuple()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(48)
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
		p.SetState(51)
		p.Match(NuggetParserID)
	}
	{
		p.SetState(52)
		p.Nugget_type()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NuggetParserLISTOP {
		{
			p.SetState(53)
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
		p.SetState(56)
		p.Match(NuggetParserID)
	}
	{
		p.SetState(57)
		p.Match(NuggetParserT__0)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserT__1)|(1<<NuggetParserT__13)|(1<<NuggetParserT__14)|(1<<NuggetParserT__15)|(1<<NuggetParserT__16)|(1<<NuggetParserT__17)|(1<<NuggetParserT__18)|(1<<NuggetParserT__19)|(1<<NuggetParserT__20)|(1<<NuggetParserT__21)|(1<<NuggetParserT__22)|(1<<NuggetParserT__23)|(1<<NuggetParserT__24)|(1<<NuggetParserT__25)|(1<<NuggetParserT__26))) != 0) {
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NuggetParserT__1 {
			{
				p.SetState(58)
				p.Match(NuggetParserT__1)
			}

		}
		{
			p.SetState(61)
			p.Nugget_type()
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(66)
		p.Match(NuggetParserT__2)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NuggetParserLISTOP {
		{
			p.SetState(67)
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

	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Match(NuggetParserID)
		}
		{
			p.SetState(71)
			p.Match(NuggetParserT__3)
		}
		{
			p.SetState(72)
			p.Match(NuggetParserSTRING)
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__4 {
			{
				p.SetState(73)
				p.Match(NuggetParserT__4)
			}
			{
				p.SetState(74)
				p.Nugget_action()
			}

			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(NuggetParserID)
		}
		{
			p.SetState(81)
			p.Match(NuggetParserT__3)
		}
		{
			p.SetState(82)
			p.Match(NuggetParserID)
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__4 {
			{
				p.SetState(83)
				p.Match(NuggetParserT__4)
			}
			{
				p.SetState(84)
				p.Nugget_action()
			}

			p.SetState(89)
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

func (s *Operation_on_singletonContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(NuggetParserID)
}

func (s *Operation_on_singletonContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(NuggetParserID, i)
}

func (s *Operation_on_singletonContext) Output_as() IOutput_asContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutput_asContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOutput_asContext)
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
		p.SetState(92)
		p.Singleton_op()
	}
	{
		p.SetState(93)
		p.Match(NuggetParserID)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NuggetParserT__1 {
		{
			p.SetState(94)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(95)
			p.Match(NuggetParserID)
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NuggetParserT__5 {
		{
			p.SetState(101)
			p.Output_as()
		}

	}

	return localctx
}

// IOutput_asContext is an interface to support dynamic dispatch.
type IOutput_asContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutput_asContext differentiates from other interfaces.
	IsOutput_asContext()
}

type Output_asContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutput_asContext() *Output_asContext {
	var p = new(Output_asContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_output_as
	return p
}

func (*Output_asContext) IsOutput_asContext() {}

func NewOutput_asContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Output_asContext {
	var p = new(Output_asContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_output_as

	return p
}

func (s *Output_asContext) GetParser() antlr.Parser { return s.parser }

func (s *Output_asContext) Output_type() IOutput_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutput_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOutput_typeContext)
}

func (s *Output_asContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Output_asContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Output_asContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterOutput_as(s)
	}
}

func (s *Output_asContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitOutput_as(s)
	}
}

func (p *NuggetParser) Output_as() (localctx IOutput_asContext) {
	localctx = NewOutput_asContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NuggetParserRULE_output_as)

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
		p.SetState(104)
		p.Match(NuggetParserT__5)
	}
	{
		p.SetState(105)
		p.Output_type()
	}

	return localctx
}

// IOutput_typeContext is an interface to support dynamic dispatch.
type IOutput_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutput_typeContext differentiates from other interfaces.
	IsOutput_typeContext()
}

type Output_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutput_typeContext() *Output_typeContext {
	var p = new(Output_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_output_type
	return p
}

func (*Output_typeContext) IsOutput_typeContext() {}

func NewOutput_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Output_typeContext {
	var p = new(Output_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_output_type

	return p
}

func (s *Output_typeContext) GetParser() antlr.Parser { return s.parser }
func (s *Output_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Output_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Output_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterOutput_type(s)
	}
}

func (s *Output_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitOutput_type(s)
	}
}

func (p *NuggetParser) Output_type() (localctx IOutput_typeContext) {
	localctx = NewOutput_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NuggetParserRULE_output_type)

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
		p.SetState(107)
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
	p.EnterRule(localctx, 16, NuggetParserRULE_singleton_op)
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
	p.SetState(109)
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
	p.EnterRule(localctx, 18, NuggetParserRULE_singleton_var)

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
		p.SetState(111)
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
	p.EnterRule(localctx, 20, NuggetParserRULE_nugget_type)
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
	p.SetState(113)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserT__13)|(1<<NuggetParserT__14)|(1<<NuggetParserT__15)|(1<<NuggetParserT__16)|(1<<NuggetParserT__17)|(1<<NuggetParserT__18)|(1<<NuggetParserT__19)|(1<<NuggetParserT__20)|(1<<NuggetParserT__21)|(1<<NuggetParserT__22)|(1<<NuggetParserT__23)|(1<<NuggetParserT__24)|(1<<NuggetParserT__25)|(1<<NuggetParserT__26))) != 0) {
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
	p.EnterRule(localctx, 22, NuggetParserRULE_nugget_action)

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
		p.SetState(115)
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
	p.EnterRule(localctx, 24, NuggetParserRULE_action_word)

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

	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserT__37:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Filter()
		}

	case NuggetParserT__27:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(NuggetParserT__27)
		}
		{
			p.SetState(119)
			p.AsType()
		}

	case NuggetParserT__28:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.Match(NuggetParserT__28)
		}
		{
			p.SetState(121)
			p.ByField()
		}

	case NuggetParserT__14:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.Match(NuggetParserT__14)
		}

	case NuggetParserT__15:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(123)
			p.Match(NuggetParserT__15)
		}

	case NuggetParserT__29:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(124)
			p.Match(NuggetParserT__29)
		}

	case NuggetParserT__30:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(125)
			p.Match(NuggetParserT__30)
		}

	case NuggetParserT__31:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(126)
			p.Match(NuggetParserT__31)
		}

	case NuggetParserT__32:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(127)
			p.Match(NuggetParserT__32)
		}
		{
			p.SetState(128)
			p.Match(NuggetParserID)
		}

	case NuggetParserT__33:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(129)
			p.Match(NuggetParserT__33)
		}

	case NuggetParserT__34:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(130)
			p.Match(NuggetParserT__34)
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
	p.EnterRule(localctx, 26, NuggetParserRULE_asType)
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
		p.SetState(133)
		p.Match(NuggetParserT__5)
	}
	{
		p.SetState(134)
		p.Nugget_type()
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NuggetParserT__36 {
		{
			p.SetState(135)
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
	p.EnterRule(localctx, 28, NuggetParserRULE_byField)

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
		p.SetState(138)
		p.Match(NuggetParserT__35)
	}
	{
		p.SetState(139)
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
	p.EnterRule(localctx, 30, NuggetParserRULE_byteOffsetSize)

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
		p.SetState(141)
		p.Match(NuggetParserT__36)
	}
	{
		p.SetState(142)
		p.Match(NuggetParserINT)
	}
	{
		p.SetState(143)
		p.Match(NuggetParserT__1)
	}
	{
		p.SetState(144)
		p.Match(NuggetParserINT)
	}
	{
		p.SetState(145)
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
	p.EnterRule(localctx, 32, NuggetParserRULE_filter)
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
		p.SetState(147)
		p.Match(NuggetParserT__37)
	}
	{
		p.SetState(148)
		p.Filter_term()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NuggetParserT__1 {
		{
			p.SetState(149)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(150)
			p.Filter_term()
		}

		p.SetState(155)
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
	p.EnterRule(localctx, 34, NuggetParserRULE_filter_term)

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
		p.SetState(156)
		p.Match(NuggetParserID)
	}
	{
		p.SetState(157)
		p.Match(NuggetParserCOMPOP)
	}
	{
		p.SetState(158)
		p.Match(NuggetParserSTRING)
	}

	return localctx
}
