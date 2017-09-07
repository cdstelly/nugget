// Generated from C:/Users/drew/GoglandProjects/nugget\Nugget2.g4 by ANTLR 4.7.

package parser // Nugget2

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 124,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 7, 2, 34, 10, 2, 12,
	2, 14, 2, 37, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 44, 10, 3, 3,
	4, 3, 4, 3, 4, 5, 4, 49, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 54, 10, 5, 3, 5,
	6, 5, 57, 10, 5, 13, 5, 14, 5, 58, 3, 5, 3, 5, 5, 5, 63, 10, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 71, 10, 6, 12, 6, 14, 6, 74, 11, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 81, 10, 6, 12, 6, 14, 6, 84, 11, 6, 5,
	6, 86, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 5, 13, 109, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 115, 10,
	14, 12, 14, 14, 14, 118, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 2,
	2, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2, 4, 3, 2,
	10, 12, 3, 2, 14, 22, 2, 126, 2, 35, 3, 2, 2, 2, 4, 43, 3, 2, 2, 2, 6,
	45, 3, 2, 2, 2, 8, 50, 3, 2, 2, 2, 10, 85, 3, 2, 2, 2, 12, 87, 3, 2, 2,
	2, 14, 92, 3, 2, 2, 2, 16, 94, 3, 2, 2, 2, 18, 96, 3, 2, 2, 2, 20, 99,
	3, 2, 2, 2, 22, 101, 3, 2, 2, 2, 24, 108, 3, 2, 2, 2, 26, 110, 3, 2, 2,
	2, 28, 119, 3, 2, 2, 2, 30, 34, 5, 4, 3, 2, 31, 34, 5, 12, 7, 2, 32, 34,
	5, 16, 9, 2, 33, 30, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 32, 3, 2, 2, 2,
	34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 38, 3,
	2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 39, 7, 2, 2, 3, 39, 3, 3, 2, 2, 2, 40,
	44, 5, 6, 4, 2, 41, 44, 5, 8, 5, 2, 42, 44, 5, 10, 6, 2, 43, 40, 3, 2,
	2, 2, 43, 41, 3, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44, 5, 3, 2, 2, 2, 45, 46,
	7, 29, 2, 2, 46, 48, 5, 20, 11, 2, 47, 49, 7, 27, 2, 2, 48, 47, 3, 2, 2,
	2, 48, 49, 3, 2, 2, 2, 49, 7, 3, 2, 2, 2, 50, 51, 7, 29, 2, 2, 51, 56,
	7, 3, 2, 2, 52, 54, 7, 4, 2, 2, 53, 52, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2,
	54, 55, 3, 2, 2, 2, 55, 57, 5, 20, 11, 2, 56, 53, 3, 2, 2, 2, 57, 58, 3,
	2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60,
	62, 7, 5, 2, 2, 61, 63, 7, 27, 2, 2, 62, 61, 3, 2, 2, 2, 62, 63, 3, 2,
	2, 2, 63, 9, 3, 2, 2, 2, 64, 65, 7, 29, 2, 2, 65, 66, 7, 6, 2, 2, 66, 67,
	7, 30, 2, 2, 67, 72, 5, 18, 10, 2, 68, 69, 7, 7, 2, 2, 69, 71, 5, 22, 12,
	2, 70, 68, 3, 2, 2, 2, 71, 74, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72, 73,
	3, 2, 2, 2, 73, 86, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 75, 76, 7, 29, 2, 2,
	76, 77, 7, 6, 2, 2, 77, 82, 7, 29, 2, 2, 78, 79, 7, 7, 2, 2, 79, 81, 5,
	22, 12, 2, 80, 78, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2,
	82, 83, 3, 2, 2, 2, 83, 86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 64, 3,
	2, 2, 2, 85, 75, 3, 2, 2, 2, 86, 11, 3, 2, 2, 2, 87, 88, 5, 14, 8, 2, 88,
	89, 7, 8, 2, 2, 89, 90, 7, 29, 2, 2, 90, 91, 7, 9, 2, 2, 91, 13, 3, 2,
	2, 2, 92, 93, 9, 2, 2, 2, 93, 15, 3, 2, 2, 2, 94, 95, 7, 29, 2, 2, 95,
	17, 3, 2, 2, 2, 96, 97, 7, 13, 2, 2, 97, 98, 5, 20, 11, 2, 98, 19, 3, 2,
	2, 2, 99, 100, 9, 3, 2, 2, 100, 21, 3, 2, 2, 2, 101, 102, 5, 24, 13, 2,
	102, 23, 3, 2, 2, 2, 103, 109, 5, 26, 14, 2, 104, 109, 7, 23, 2, 2, 105,
	109, 7, 15, 2, 2, 106, 109, 7, 16, 2, 2, 107, 109, 7, 24, 2, 2, 108, 103,
	3, 2, 2, 2, 108, 104, 3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 108, 106, 3, 2,
	2, 2, 108, 107, 3, 2, 2, 2, 109, 25, 3, 2, 2, 2, 110, 111, 7, 25, 2, 2,
	111, 116, 5, 28, 15, 2, 112, 113, 7, 4, 2, 2, 113, 115, 5, 28, 15, 2, 114,
	112, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 117,
	3, 2, 2, 2, 117, 27, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 120, 7, 29,
	2, 2, 120, 121, 7, 26, 2, 2, 121, 122, 7, 30, 2, 2, 122, 29, 3, 2, 2, 2,
	14, 33, 35, 43, 48, 53, 58, 62, 72, 82, 85, 108, 116,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'tuple['", "','", "']'", "'='", "'|'", "'('", "')'", "'type'", "'print'",
	"'size'", "'as'", "'string'", "'sha1'", "'md5'", "'ntfs'", "'file'", "'packet'",
	"'pcap'", "'exifinfo'", "'datetime'", "'extract'", "'NUGGETGENERATORPLACEHOLDER'",
	"'filter'", "", "'[]'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "COMPOP", "LISTOP", "INT", "ID", "STRING", "WS",
	"NL", "LINE_COMMENT",
}

var ruleNames = []string{
	"prog", "define_assign", "define", "define_tuple", "assign", "operation_on_singleton",
	"singleton_op", "singleton_var", "asType", "nugget_type", "nugget_action",
	"action_word", "filter", "filter_term",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Nugget2Parser struct {
	*antlr.BaseParser
}

func NewNugget2Parser(input antlr.TokenStream) *Nugget2Parser {
	this := new(Nugget2Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Nugget2.g4"

	return this
}

// Nugget2Parser tokens.
const (
	Nugget2ParserEOF          = antlr.TokenEOF
	Nugget2ParserT__0         = 1
	Nugget2ParserT__1         = 2
	Nugget2ParserT__2         = 3
	Nugget2ParserT__3         = 4
	Nugget2ParserT__4         = 5
	Nugget2ParserT__5         = 6
	Nugget2ParserT__6         = 7
	Nugget2ParserT__7         = 8
	Nugget2ParserT__8         = 9
	Nugget2ParserT__9         = 10
	Nugget2ParserT__10        = 11
	Nugget2ParserT__11        = 12
	Nugget2ParserT__12        = 13
	Nugget2ParserT__13        = 14
	Nugget2ParserT__14        = 15
	Nugget2ParserT__15        = 16
	Nugget2ParserT__16        = 17
	Nugget2ParserT__17        = 18
	Nugget2ParserT__18        = 19
	Nugget2ParserT__19        = 20
	Nugget2ParserT__20        = 21
	Nugget2ParserT__21        = 22
	Nugget2ParserT__22        = 23
	Nugget2ParserCOMPOP       = 24
	Nugget2ParserLISTOP       = 25
	Nugget2ParserINT          = 26
	Nugget2ParserID           = 27
	Nugget2ParserSTRING       = 28
	Nugget2ParserWS           = 29
	Nugget2ParserNL           = 30
	Nugget2ParserLINE_COMMENT = 31
)

// Nugget2Parser rules.
const (
	Nugget2ParserRULE_prog                   = 0
	Nugget2ParserRULE_define_assign          = 1
	Nugget2ParserRULE_define                 = 2
	Nugget2ParserRULE_define_tuple           = 3
	Nugget2ParserRULE_assign                 = 4
	Nugget2ParserRULE_operation_on_singleton = 5
	Nugget2ParserRULE_singleton_op           = 6
	Nugget2ParserRULE_singleton_var          = 7
	Nugget2ParserRULE_asType                 = 8
	Nugget2ParserRULE_nugget_type            = 9
	Nugget2ParserRULE_nugget_action          = 10
	Nugget2ParserRULE_action_word            = 11
	Nugget2ParserRULE_filter                 = 12
	Nugget2ParserRULE_filter_term            = 13
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
	p.RuleIndex = Nugget2ParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(Nugget2ParserEOF, 0)
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
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *Nugget2Parser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Nugget2ParserRULE_prog)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Nugget2ParserT__7)|(1<<Nugget2ParserT__8)|(1<<Nugget2ParserT__9)|(1<<Nugget2ParserID))) != 0 {
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(28)
				p.Define_assign()
			}

		case 2:
			{
				p.SetState(29)
				p.Operation_on_singleton()
			}

		case 3:
			{
				p.SetState(30)
				p.Singleton_var()
			}

		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
		p.Match(Nugget2ParserEOF)
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
	p.RuleIndex = Nugget2ParserRULE_define_assign
	return p
}

func (*Define_assignContext) IsDefine_assignContext() {}

func NewDefine_assignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Define_assignContext {
	var p = new(Define_assignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_define_assign

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
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterDefine_assign(s)
	}
}

func (s *Define_assignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitDefine_assign(s)
	}
}

func (p *Nugget2Parser) Define_assign() (localctx IDefine_assignContext) {
	localctx = NewDefine_assignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Nugget2ParserRULE_define_assign)

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

	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Define()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Define_tuple()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
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
	p.RuleIndex = Nugget2ParserRULE_define
	return p
}

func (*DefineContext) IsDefineContext() {}

func NewDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefineContext {
	var p = new(DefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_define

	return p
}

func (s *DefineContext) GetParser() antlr.Parser { return s.parser }

func (s *DefineContext) ID() antlr.TerminalNode {
	return s.GetToken(Nugget2ParserID, 0)
}

func (s *DefineContext) Nugget_type() INugget_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INugget_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INugget_typeContext)
}

func (s *DefineContext) LISTOP() antlr.TerminalNode {
	return s.GetToken(Nugget2ParserLISTOP, 0)
}

func (s *DefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterDefine(s)
	}
}

func (s *DefineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitDefine(s)
	}
}

func (p *Nugget2Parser) Define() (localctx IDefineContext) {
	localctx = NewDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Nugget2ParserRULE_define)
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
		p.SetState(43)
		p.Match(Nugget2ParserID)
	}
	{
		p.SetState(44)
		p.Nugget_type()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Nugget2ParserLISTOP {
		{
			p.SetState(45)
			p.Match(Nugget2ParserLISTOP)
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
	p.RuleIndex = Nugget2ParserRULE_define_tuple
	return p
}

func (*Define_tupleContext) IsDefine_tupleContext() {}

func NewDefine_tupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Define_tupleContext {
	var p = new(Define_tupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_define_tuple

	return p
}

func (s *Define_tupleContext) GetParser() antlr.Parser { return s.parser }

func (s *Define_tupleContext) ID() antlr.TerminalNode {
	return s.GetToken(Nugget2ParserID, 0)
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
	return s.GetToken(Nugget2ParserLISTOP, 0)
}

func (s *Define_tupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Define_tupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Define_tupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterDefine_tuple(s)
	}
}

func (s *Define_tupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitDefine_tuple(s)
	}
}

func (p *Nugget2Parser) Define_tuple() (localctx IDefine_tupleContext) {
	localctx = NewDefine_tupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Nugget2ParserRULE_define_tuple)
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
		p.SetState(48)
		p.Match(Nugget2ParserID)
	}
	{
		p.SetState(49)
		p.Match(Nugget2ParserT__0)
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Nugget2ParserT__1)|(1<<Nugget2ParserT__11)|(1<<Nugget2ParserT__12)|(1<<Nugget2ParserT__13)|(1<<Nugget2ParserT__14)|(1<<Nugget2ParserT__15)|(1<<Nugget2ParserT__16)|(1<<Nugget2ParserT__17)|(1<<Nugget2ParserT__18)|(1<<Nugget2ParserT__19))) != 0) {
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == Nugget2ParserT__1 {
			{
				p.SetState(50)
				p.Match(Nugget2ParserT__1)
			}

		}
		{
			p.SetState(53)
			p.Nugget_type()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(Nugget2ParserT__2)
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == Nugget2ParserLISTOP {
		{
			p.SetState(59)
			p.Match(Nugget2ParserLISTOP)
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
	p.RuleIndex = Nugget2ParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(Nugget2ParserID)
}

func (s *AssignContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(Nugget2ParserID, i)
}

func (s *AssignContext) STRING() antlr.TerminalNode {
	return s.GetToken(Nugget2ParserSTRING, 0)
}

func (s *AssignContext) AsType() IAsTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsTypeContext)
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
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *Nugget2Parser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Nugget2ParserRULE_assign)
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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(Nugget2ParserID)
		}
		{
			p.SetState(63)
			p.Match(Nugget2ParserT__3)
		}
		{
			p.SetState(64)
			p.Match(Nugget2ParserSTRING)
		}
		{
			p.SetState(65)
			p.AsType()
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Nugget2ParserT__4 {
			{
				p.SetState(66)
				p.Match(Nugget2ParserT__4)
			}
			{
				p.SetState(67)
				p.Nugget_action()
			}

			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Match(Nugget2ParserID)
		}
		{
			p.SetState(74)
			p.Match(Nugget2ParserT__3)
		}
		{
			p.SetState(75)
			p.Match(Nugget2ParserID)
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == Nugget2ParserT__4 {
			{
				p.SetState(76)
				p.Match(Nugget2ParserT__4)
			}
			{
				p.SetState(77)
				p.Nugget_action()
			}

			p.SetState(82)
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
	p.RuleIndex = Nugget2ParserRULE_operation_on_singleton
	return p
}

func (*Operation_on_singletonContext) IsOperation_on_singletonContext() {}

func NewOperation_on_singletonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operation_on_singletonContext {
	var p = new(Operation_on_singletonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_operation_on_singleton

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
	return s.GetToken(Nugget2ParserID, 0)
}

func (s *Operation_on_singletonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operation_on_singletonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operation_on_singletonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterOperation_on_singleton(s)
	}
}

func (s *Operation_on_singletonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitOperation_on_singleton(s)
	}
}

func (p *Nugget2Parser) Operation_on_singleton() (localctx IOperation_on_singletonContext) {
	localctx = NewOperation_on_singletonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Nugget2ParserRULE_operation_on_singleton)

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
		p.SetState(85)
		p.Singleton_op()
	}
	{
		p.SetState(86)
		p.Match(Nugget2ParserT__5)
	}
	{
		p.SetState(87)
		p.Match(Nugget2ParserID)
	}
	{
		p.SetState(88)
		p.Match(Nugget2ParserT__6)
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
	p.RuleIndex = Nugget2ParserRULE_singleton_op
	return p
}

func (*Singleton_opContext) IsSingleton_opContext() {}

func NewSingleton_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Singleton_opContext {
	var p = new(Singleton_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_singleton_op

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
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterSingleton_op(s)
	}
}

func (s *Singleton_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitSingleton_op(s)
	}
}

func (p *Nugget2Parser) Singleton_op() (localctx ISingleton_opContext) {
	localctx = NewSingleton_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Nugget2ParserRULE_singleton_op)
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
	p.SetState(90)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Nugget2ParserT__7)|(1<<Nugget2ParserT__8)|(1<<Nugget2ParserT__9))) != 0) {
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
	p.RuleIndex = Nugget2ParserRULE_singleton_var
	return p
}

func (*Singleton_varContext) IsSingleton_varContext() {}

func NewSingleton_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Singleton_varContext {
	var p = new(Singleton_varContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_singleton_var

	return p
}

func (s *Singleton_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Singleton_varContext) ID() antlr.TerminalNode {
	return s.GetToken(Nugget2ParserID, 0)
}

func (s *Singleton_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Singleton_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Singleton_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterSingleton_var(s)
	}
}

func (s *Singleton_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitSingleton_var(s)
	}
}

func (p *Nugget2Parser) Singleton_var() (localctx ISingleton_varContext) {
	localctx = NewSingleton_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Nugget2ParserRULE_singleton_var)

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
		p.Match(Nugget2ParserID)
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
	p.RuleIndex = Nugget2ParserRULE_asType
	return p
}

func (*AsTypeContext) IsAsTypeContext() {}

func NewAsTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsTypeContext {
	var p = new(AsTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_asType

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

func (s *AsTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterAsType(s)
	}
}

func (s *AsTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitAsType(s)
	}
}

func (p *Nugget2Parser) AsType() (localctx IAsTypeContext) {
	localctx = NewAsTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Nugget2ParserRULE_asType)

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
		p.SetState(94)
		p.Match(Nugget2ParserT__10)
	}
	{
		p.SetState(95)
		p.Nugget_type()
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
	p.RuleIndex = Nugget2ParserRULE_nugget_type
	return p
}

func (*Nugget_typeContext) IsNugget_typeContext() {}

func NewNugget_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nugget_typeContext {
	var p = new(Nugget_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_nugget_type

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
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterNugget_type(s)
	}
}

func (s *Nugget_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitNugget_type(s)
	}
}

func (p *Nugget2Parser) Nugget_type() (localctx INugget_typeContext) {
	localctx = NewNugget_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Nugget2ParserRULE_nugget_type)
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

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Nugget2ParserT__11)|(1<<Nugget2ParserT__12)|(1<<Nugget2ParserT__13)|(1<<Nugget2ParserT__14)|(1<<Nugget2ParserT__15)|(1<<Nugget2ParserT__16)|(1<<Nugget2ParserT__17)|(1<<Nugget2ParserT__18)|(1<<Nugget2ParserT__19))) != 0) {
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
	p.RuleIndex = Nugget2ParserRULE_nugget_action
	return p
}

func (*Nugget_actionContext) IsNugget_actionContext() {}

func NewNugget_actionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nugget_actionContext {
	var p = new(Nugget_actionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_nugget_action

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
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterNugget_action(s)
	}
}

func (s *Nugget_actionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitNugget_action(s)
	}
}

func (p *Nugget2Parser) Nugget_action() (localctx INugget_actionContext) {
	localctx = NewNugget_actionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Nugget2ParserRULE_nugget_action)

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
	p.RuleIndex = Nugget2ParserRULE_action_word
	return p
}

func (*Action_wordContext) IsAction_wordContext() {}

func NewAction_wordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Action_wordContext {
	var p = new(Action_wordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_action_word

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

func (s *Action_wordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Action_wordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Action_wordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterAction_word(s)
	}
}

func (s *Action_wordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitAction_word(s)
	}
}

func (p *Nugget2Parser) Action_word() (localctx IAction_wordContext) {
	localctx = NewAction_wordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Nugget2ParserRULE_action_word)

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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Nugget2ParserT__22:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Filter()
		}

	case Nugget2ParserT__20:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(Nugget2ParserT__20)
		}

	case Nugget2ParserT__12:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.Match(Nugget2ParserT__12)
		}

	case Nugget2ParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(104)
			p.Match(Nugget2ParserT__13)
		}

	case Nugget2ParserT__21:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(105)
			p.Match(Nugget2ParserT__21)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = Nugget2ParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_filter

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
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *Nugget2Parser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Nugget2ParserRULE_filter)
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
		p.SetState(108)
		p.Match(Nugget2ParserT__22)
	}
	{
		p.SetState(109)
		p.Filter_term()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Nugget2ParserT__1 {
		{
			p.SetState(110)
			p.Match(Nugget2ParserT__1)
		}
		{
			p.SetState(111)
			p.Filter_term()
		}

		p.SetState(116)
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
	p.RuleIndex = Nugget2ParserRULE_filter_term
	return p
}

func (*Filter_termContext) IsFilter_termContext() {}

func NewFilter_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_termContext {
	var p = new(Filter_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Nugget2ParserRULE_filter_term

	return p
}

func (s *Filter_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_termContext) ID() antlr.TerminalNode {
	return s.GetToken(Nugget2ParserID, 0)
}

func (s *Filter_termContext) COMPOP() antlr.TerminalNode {
	return s.GetToken(Nugget2ParserCOMPOP, 0)
}

func (s *Filter_termContext) STRING() antlr.TerminalNode {
	return s.GetToken(Nugget2ParserSTRING, 0)
}

func (s *Filter_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.EnterFilter_term(s)
	}
}

func (s *Filter_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Nugget2Listener); ok {
		listenerT.ExitFilter_term(s)
	}
}

func (p *Nugget2Parser) Filter_term() (localctx IFilter_termContext) {
	localctx = NewFilter_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Nugget2ParserRULE_filter_term)

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
		p.Match(Nugget2ParserID)
	}
	{
		p.SetState(118)
		p.Match(Nugget2ParserCOMPOP)
	}
	{
		p.SetState(119)
		p.Match(Nugget2ParserSTRING)
	}

	return localctx
}
