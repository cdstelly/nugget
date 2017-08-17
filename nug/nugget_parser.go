// Generated from C:/Users/drew/GoglandProjects/nugget\Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 33, 199,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 56, 10, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 67, 10, 3, 12, 3,
	14, 3, 70, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 7, 5, 82, 10, 5, 12, 5, 14, 5, 85, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 95, 10, 5, 12, 5, 14, 5, 98, 11, 5, 3, 5,
	7, 5, 101, 10, 5, 12, 5, 14, 5, 104, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 7, 5, 113, 10, 5, 12, 5, 14, 5, 116, 11, 5, 5, 5, 118, 10,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 127, 10, 6, 12, 6, 14,
	6, 130, 11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 136, 10, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8, 145, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 169, 10, 15, 3,
	16, 3, 16, 3, 16, 7, 16, 174, 10, 16, 12, 16, 14, 16, 177, 11, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 5, 16, 183, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 5, 17, 190, 10, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 2, 3, 4, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 2, 5, 3, 2, 12, 13, 4, 2, 17, 17, 20, 22, 4, 2, 18, 19,
	23, 24, 2, 201, 2, 55, 3, 2, 2, 2, 4, 57, 3, 2, 2, 2, 6, 71, 3, 2, 2, 2,
	8, 117, 3, 2, 2, 2, 10, 135, 3, 2, 2, 2, 12, 137, 3, 2, 2, 2, 14, 144,
	3, 2, 2, 2, 16, 146, 3, 2, 2, 2, 18, 150, 3, 2, 2, 2, 20, 154, 3, 2, 2,
	2, 22, 156, 3, 2, 2, 2, 24, 160, 3, 2, 2, 2, 26, 162, 3, 2, 2, 2, 28, 168,
	3, 2, 2, 2, 30, 182, 3, 2, 2, 2, 32, 189, 3, 2, 2, 2, 34, 191, 3, 2, 2,
	2, 36, 193, 3, 2, 2, 2, 38, 39, 5, 36, 19, 2, 39, 40, 7, 2, 2, 3, 40, 56,
	3, 2, 2, 2, 41, 42, 5, 8, 5, 2, 42, 43, 7, 2, 2, 3, 43, 56, 3, 2, 2, 2,
	44, 56, 5, 4, 3, 2, 45, 46, 5, 34, 18, 2, 46, 47, 7, 2, 2, 3, 47, 56, 3,
	2, 2, 2, 48, 49, 5, 6, 4, 2, 49, 50, 7, 2, 2, 3, 50, 56, 3, 2, 2, 2, 51,
	52, 5, 12, 7, 2, 52, 53, 7, 2, 2, 3, 53, 56, 3, 2, 2, 2, 54, 56, 7, 2,
	2, 3, 55, 38, 3, 2, 2, 2, 55, 41, 3, 2, 2, 2, 55, 44, 3, 2, 2, 2, 55, 45,
	3, 2, 2, 2, 55, 48, 3, 2, 2, 2, 55, 51, 3, 2, 2, 2, 55, 54, 3, 2, 2, 2,
	56, 3, 3, 2, 2, 2, 57, 58, 8, 3, 1, 2, 58, 59, 7, 28, 2, 2, 59, 60, 7,
	3, 2, 2, 60, 61, 5, 28, 15, 2, 61, 62, 7, 23, 2, 2, 62, 63, 5, 24, 13,
	2, 63, 68, 3, 2, 2, 2, 64, 65, 12, 3, 2, 2, 65, 67, 5, 10, 6, 2, 66, 64,
	3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2,
	69, 5, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 72, 7, 28, 2, 2, 72, 73, 7,
	3, 2, 2, 73, 74, 7, 28, 2, 2, 74, 7, 3, 2, 2, 2, 75, 76, 7, 28, 2, 2, 76,
	77, 7, 3, 2, 2, 77, 78, 7, 28, 2, 2, 78, 79, 7, 4, 2, 2, 79, 83, 5, 26,
	14, 2, 80, 82, 5, 10, 6, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83,
	81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 118, 3, 2, 2, 2, 85, 83, 3, 2,
	2, 2, 86, 87, 7, 28, 2, 2, 87, 88, 7, 3, 2, 2, 88, 89, 7, 28, 2, 2, 89,
	90, 7, 4, 2, 2, 90, 91, 5, 26, 14, 2, 91, 96, 5, 14, 8, 2, 92, 93, 7, 5,
	2, 2, 93, 95, 5, 14, 8, 2, 94, 92, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96,
	94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 102, 3, 2, 2, 2, 98, 96, 3, 2,
	2, 2, 99, 101, 5, 10, 6, 2, 100, 99, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2,
	102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 118, 3, 2, 2, 2, 104,
	102, 3, 2, 2, 2, 105, 106, 7, 28, 2, 2, 106, 107, 7, 3, 2, 2, 107, 108,
	7, 28, 2, 2, 108, 109, 7, 4, 2, 2, 109, 110, 5, 26, 14, 2, 110, 114, 7,
	28, 2, 2, 111, 113, 5, 10, 6, 2, 112, 111, 3, 2, 2, 2, 113, 116, 3, 2,
	2, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2,
	116, 114, 3, 2, 2, 2, 117, 75, 3, 2, 2, 2, 117, 86, 3, 2, 2, 2, 117, 105,
	3, 2, 2, 2, 118, 9, 3, 2, 2, 2, 119, 120, 7, 4, 2, 2, 120, 136, 5, 26,
	14, 2, 121, 122, 7, 4, 2, 2, 122, 123, 5, 26, 14, 2, 123, 128, 5, 14, 8,
	2, 124, 125, 7, 5, 2, 2, 125, 127, 5, 14, 8, 2, 126, 124, 3, 2, 2, 2, 127,
	130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 136,
	3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 4, 2, 2, 132, 133, 5, 26,
	14, 2, 133, 134, 7, 28, 2, 2, 134, 136, 3, 2, 2, 2, 135, 119, 3, 2, 2,
	2, 135, 121, 3, 2, 2, 2, 135, 131, 3, 2, 2, 2, 136, 11, 3, 2, 2, 2, 137,
	138, 7, 6, 2, 2, 138, 139, 7, 28, 2, 2, 139, 140, 7, 7, 2, 2, 140, 141,
	7, 28, 2, 2, 141, 13, 3, 2, 2, 2, 142, 145, 5, 16, 9, 2, 143, 145, 5, 18,
	10, 2, 144, 142, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145, 15, 3, 2, 2, 2,
	146, 147, 7, 8, 2, 2, 147, 148, 7, 28, 2, 2, 148, 149, 7, 8, 2, 2, 149,
	17, 3, 2, 2, 2, 150, 151, 5, 20, 11, 2, 151, 152, 7, 29, 2, 2, 152, 153,
	5, 22, 12, 2, 153, 19, 3, 2, 2, 2, 154, 155, 9, 2, 2, 2, 155, 21, 3, 2,
	2, 2, 156, 157, 7, 8, 2, 2, 157, 158, 7, 30, 2, 2, 158, 159, 7, 8, 2, 2,
	159, 23, 3, 2, 2, 2, 160, 161, 9, 3, 2, 2, 161, 25, 3, 2, 2, 2, 162, 163,
	9, 4, 2, 2, 163, 27, 3, 2, 2, 2, 164, 165, 7, 8, 2, 2, 165, 166, 7, 28,
	2, 2, 166, 169, 7, 8, 2, 2, 167, 169, 7, 28, 2, 2, 168, 164, 3, 2, 2, 2,
	168, 167, 3, 2, 2, 2, 169, 29, 3, 2, 2, 2, 170, 175, 7, 28, 2, 2, 171,
	172, 7, 5, 2, 2, 172, 174, 7, 28, 2, 2, 173, 171, 3, 2, 2, 2, 174, 177,
	3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 183, 3, 2,
	2, 2, 177, 175, 3, 2, 2, 2, 178, 179, 7, 9, 2, 2, 179, 180, 5, 30, 16,
	2, 180, 181, 7, 9, 2, 2, 181, 183, 3, 2, 2, 2, 182, 170, 3, 2, 2, 2, 182,
	178, 3, 2, 2, 2, 183, 31, 3, 2, 2, 2, 184, 190, 7, 28, 2, 2, 185, 186,
	7, 9, 2, 2, 186, 187, 5, 32, 17, 2, 187, 188, 7, 9, 2, 2, 188, 190, 3,
	2, 2, 2, 189, 184, 3, 2, 2, 2, 189, 185, 3, 2, 2, 2, 190, 33, 3, 2, 2,
	2, 191, 192, 7, 28, 2, 2, 192, 35, 3, 2, 2, 2, 193, 194, 7, 14, 2, 2, 194,
	195, 7, 10, 2, 2, 195, 196, 7, 27, 2, 2, 196, 197, 7, 11, 2, 2, 197, 37,
	3, 2, 2, 2, 16, 55, 68, 83, 96, 102, 114, 117, 128, 135, 144, 168, 175,
	182, 189,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'|'", "','", "'save'", "'to'", "'\"'", "'''", "'('", "')'",
	"", "", "'sin'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "CTIME", "MTIME", "SIN", "LOAD",
	"FROM", "ALL", "HASH", "SELECT", "FILES", "IMAGES", "DOCUMENTS", "EXTRACT",
	"JOIN", "WS", "COMMENT", "NUMBER", "StrLit", "OPERATION", "DATE", "LINE_COMMENT",
	"ROOT", "ATTR_LIST",
}

var ruleNames = []string{
	"nugget", "initextract", "assign", "execute", "streamTask", "save", "filter",
	"filename", "timefilter", "reference", "date", "subtype", "task", "target",
	"field", "sourceidentifier", "printId", "sin",
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
	NuggetParserCTIME        = 10
	NuggetParserMTIME        = 11
	NuggetParserSIN          = 12
	NuggetParserLOAD         = 13
	NuggetParserFROM         = 14
	NuggetParserALL          = 15
	NuggetParserHASH         = 16
	NuggetParserSELECT       = 17
	NuggetParserFILES        = 18
	NuggetParserIMAGES       = 19
	NuggetParserDOCUMENTS    = 20
	NuggetParserEXTRACT      = 21
	NuggetParserJOIN         = 22
	NuggetParserWS           = 23
	NuggetParserCOMMENT      = 24
	NuggetParserNUMBER       = 25
	NuggetParserStrLit       = 26
	NuggetParserOPERATION    = 27
	NuggetParserDATE         = 28
	NuggetParserLINE_COMMENT = 29
	NuggetParserROOT         = 30
	NuggetParserATTR_LIST    = 31
)

// NuggetParser rules.
const (
	NuggetParserRULE_nugget           = 0
	NuggetParserRULE_initextract      = 1
	NuggetParserRULE_assign           = 2
	NuggetParserRULE_execute          = 3
	NuggetParserRULE_streamTask       = 4
	NuggetParserRULE_save             = 5
	NuggetParserRULE_filter           = 6
	NuggetParserRULE_filename         = 7
	NuggetParserRULE_timefilter       = 8
	NuggetParserRULE_reference        = 9
	NuggetParserRULE_date             = 10
	NuggetParserRULE_subtype          = 11
	NuggetParserRULE_task             = 12
	NuggetParserRULE_target           = 13
	NuggetParserRULE_field            = 14
	NuggetParserRULE_sourceidentifier = 15
	NuggetParserRULE_printId          = 16
	NuggetParserRULE_sin              = 17
)

// INuggetContext is an interface to support dynamic dispatch.
type INuggetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNuggetContext differentiates from other interfaces.
	IsNuggetContext()
}

type NuggetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNuggetContext() *NuggetContext {
	var p = new(NuggetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_nugget
	return p
}

func (*NuggetContext) IsNuggetContext() {}

func NewNuggetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NuggetContext {
	var p = new(NuggetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_nugget

	return p
}

func (s *NuggetContext) GetParser() antlr.Parser { return s.parser }

func (s *NuggetContext) Sin() ISinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISinContext)
}

func (s *NuggetContext) EOF() antlr.TerminalNode {
	return s.GetToken(NuggetParserEOF, 0)
}

func (s *NuggetContext) Execute() IExecuteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecuteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecuteContext)
}

func (s *NuggetContext) Initextract() IInitextractContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitextractContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitextractContext)
}

func (s *NuggetContext) PrintId() IPrintIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintIdContext)
}

func (s *NuggetContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *NuggetContext) Save() ISaveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISaveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISaveContext)
}

func (s *NuggetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NuggetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NuggetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterNugget(s)
	}
}

func (s *NuggetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitNugget(s)
	}
}

func (p *NuggetParser) Nugget() (localctx INuggetContext) {
	localctx = NewNuggetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NuggetParserRULE_nugget)

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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Sin()
		}
		{
			p.SetState(37)
			p.Match(NuggetParserEOF)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Execute()
		}
		{
			p.SetState(40)
			p.Match(NuggetParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.initextract(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(43)
			p.PrintId()
		}
		{
			p.SetState(44)
			p.Match(NuggetParserEOF)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(46)
			p.Assign()
		}
		{
			p.SetState(47)
			p.Match(NuggetParserEOF)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(49)
			p.Save()
		}
		{
			p.SetState(50)
			p.Match(NuggetParserEOF)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(52)
			p.Match(NuggetParserEOF)
		}

	}

	return localctx
}

// IInitextractContext is an interface to support dynamic dispatch.
type IInitextractContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitextractContext differentiates from other interfaces.
	IsInitextractContext()
}

type InitextractContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitextractContext() *InitextractContext {
	var p = new(InitextractContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_initextract
	return p
}

func (*InitextractContext) IsInitextractContext() {}

func NewInitextractContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitextractContext {
	var p = new(InitextractContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_initextract

	return p
}

func (s *InitextractContext) GetParser() antlr.Parser { return s.parser }

func (s *InitextractContext) StrLit() antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, 0)
}

func (s *InitextractContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *InitextractContext) EXTRACT() antlr.TerminalNode {
	return s.GetToken(NuggetParserEXTRACT, 0)
}

func (s *InitextractContext) Subtype() ISubtypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubtypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubtypeContext)
}

func (s *InitextractContext) Initextract() IInitextractContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitextractContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitextractContext)
}

func (s *InitextractContext) StreamTask() IStreamTaskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStreamTaskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStreamTaskContext)
}

func (s *InitextractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitextractContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitextractContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterInitextract(s)
	}
}

func (s *InitextractContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitInitextract(s)
	}
}

func (p *NuggetParser) Initextract() (localctx IInitextractContext) {
	return p.initextract(0)
}

func (p *NuggetParser) initextract(_p int) (localctx IInitextractContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewInitextractContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IInitextractContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, NuggetParserRULE_initextract, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(NuggetParserStrLit)
	}
	{
		p.SetState(57)
		p.Match(NuggetParserT__0)
	}
	{
		p.SetState(58)
		p.Target()
	}
	{
		p.SetState(59)
		p.Match(NuggetParserEXTRACT)
	}
	{
		p.SetState(60)
		p.Subtype()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewInitextractContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, NuggetParserRULE_initextract)
			p.SetState(62)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(63)
				p.StreamTask()
			}

		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
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

func (s *AssignContext) AllStrLit() []antlr.TerminalNode {
	return s.GetTokens(NuggetParserStrLit)
}

func (s *AssignContext) StrLit(i int) antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, i)
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
	p.EnterRule(localctx, 4, NuggetParserRULE_assign)

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
		p.SetState(69)
		p.Match(NuggetParserStrLit)
	}
	{
		p.SetState(70)
		p.Match(NuggetParserT__0)
	}
	{
		p.SetState(71)
		p.Match(NuggetParserStrLit)
	}

	return localctx
}

// IExecuteContext is an interface to support dynamic dispatch.
type IExecuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecuteContext differentiates from other interfaces.
	IsExecuteContext()
}

type ExecuteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecuteContext() *ExecuteContext {
	var p = new(ExecuteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_execute
	return p
}

func (*ExecuteContext) IsExecuteContext() {}

func NewExecuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteContext {
	var p = new(ExecuteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_execute

	return p
}

func (s *ExecuteContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteContext) AllStrLit() []antlr.TerminalNode {
	return s.GetTokens(NuggetParserStrLit)
}

func (s *ExecuteContext) StrLit(i int) antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, i)
}

func (s *ExecuteContext) Task() ITaskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITaskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITaskContext)
}

func (s *ExecuteContext) AllStreamTask() []IStreamTaskContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStreamTaskContext)(nil)).Elem())
	var tst = make([]IStreamTaskContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStreamTaskContext)
		}
	}

	return tst
}

func (s *ExecuteContext) StreamTask(i int) IStreamTaskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStreamTaskContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStreamTaskContext)
}

func (s *ExecuteContext) AllFilter() []IFilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterContext)(nil)).Elem())
	var tst = make([]IFilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterContext)
		}
	}

	return tst
}

func (s *ExecuteContext) Filter(i int) IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *ExecuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecuteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterExecute(s)
	}
}

func (s *ExecuteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitExecute(s)
	}
}

func (p *NuggetParser) Execute() (localctx IExecuteContext) {
	localctx = NewExecuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NuggetParserRULE_execute)
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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Match(NuggetParserStrLit)
		}
		{
			p.SetState(74)
			p.Match(NuggetParserT__0)
		}
		{
			p.SetState(75)
			p.Match(NuggetParserStrLit)
		}
		{
			p.SetState(76)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(77)
			p.Task()
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__1 {
			{
				p.SetState(78)
				p.StreamTask()
			}

			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Match(NuggetParserStrLit)
		}
		{
			p.SetState(85)
			p.Match(NuggetParserT__0)
		}
		{
			p.SetState(86)
			p.Match(NuggetParserStrLit)
		}
		{
			p.SetState(87)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(88)
			p.Task()
		}
		{
			p.SetState(89)
			p.Filter()
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__2 {
			{
				p.SetState(90)
				p.Match(NuggetParserT__2)
			}
			{
				p.SetState(91)
				p.Filter()
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__1 {
			{
				p.SetState(97)
				p.StreamTask()
			}

			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.Match(NuggetParserStrLit)
		}
		{
			p.SetState(104)
			p.Match(NuggetParserT__0)
		}
		{
			p.SetState(105)
			p.Match(NuggetParserStrLit)
		}
		{
			p.SetState(106)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(107)
			p.Task()
		}
		{
			p.SetState(108)
			p.Match(NuggetParserStrLit)
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__1 {
			{
				p.SetState(109)
				p.StreamTask()
			}

			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IStreamTaskContext is an interface to support dynamic dispatch.
type IStreamTaskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreamTaskContext differentiates from other interfaces.
	IsStreamTaskContext()
}

type StreamTaskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreamTaskContext() *StreamTaskContext {
	var p = new(StreamTaskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_streamTask
	return p
}

func (*StreamTaskContext) IsStreamTaskContext() {}

func NewStreamTaskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreamTaskContext {
	var p = new(StreamTaskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_streamTask

	return p
}

func (s *StreamTaskContext) GetParser() antlr.Parser { return s.parser }

func (s *StreamTaskContext) Task() ITaskContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITaskContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITaskContext)
}

func (s *StreamTaskContext) AllFilter() []IFilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterContext)(nil)).Elem())
	var tst = make([]IFilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterContext)
		}
	}

	return tst
}

func (s *StreamTaskContext) Filter(i int) IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFilterContext)
}

func (s *StreamTaskContext) StrLit() antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, 0)
}

func (s *StreamTaskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreamTaskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StreamTaskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterStreamTask(s)
	}
}

func (s *StreamTaskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitStreamTask(s)
	}
}

func (p *NuggetParser) StreamTask() (localctx IStreamTaskContext) {
	localctx = NewStreamTaskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NuggetParserRULE_streamTask)

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

	var _alt int

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(118)
			p.Task()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(120)
			p.Task()
		}
		{
			p.SetState(121)
			p.Filter()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(122)
					p.Match(NuggetParserT__2)
				}
				{
					p.SetState(123)
					p.Filter()
				}

			}
			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.Match(NuggetParserT__1)
		}
		{
			p.SetState(130)
			p.Task()
		}
		{
			p.SetState(131)
			p.Match(NuggetParserStrLit)
		}

	}

	return localctx
}

// ISaveContext is an interface to support dynamic dispatch.
type ISaveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSaveContext differentiates from other interfaces.
	IsSaveContext()
}

type SaveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySaveContext() *SaveContext {
	var p = new(SaveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_save
	return p
}

func (*SaveContext) IsSaveContext() {}

func NewSaveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SaveContext {
	var p = new(SaveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_save

	return p
}

func (s *SaveContext) GetParser() antlr.Parser { return s.parser }

func (s *SaveContext) AllStrLit() []antlr.TerminalNode {
	return s.GetTokens(NuggetParserStrLit)
}

func (s *SaveContext) StrLit(i int) antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, i)
}

func (s *SaveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SaveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SaveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterSave(s)
	}
}

func (s *SaveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitSave(s)
	}
}

func (p *NuggetParser) Save() (localctx ISaveContext) {
	localctx = NewSaveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NuggetParserRULE_save)

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
		p.SetState(135)
		p.Match(NuggetParserT__3)
	}
	{
		p.SetState(136)
		p.Match(NuggetParserStrLit)
	}
	{
		p.SetState(137)
		p.Match(NuggetParserT__4)
	}
	{
		p.SetState(138)
		p.Match(NuggetParserStrLit)
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

func (s *FilterContext) Filename() IFilenameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilenameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilenameContext)
}

func (s *FilterContext) Timefilter() ITimefilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimefilterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimefilterContext)
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
	p.EnterRule(localctx, 12, NuggetParserRULE_filter)

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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Filename()
		}

	case NuggetParserCTIME, NuggetParserMTIME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Timefilter()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFilenameContext is an interface to support dynamic dispatch.
type IFilenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilenameContext differentiates from other interfaces.
	IsFilenameContext()
}

type FilenameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilenameContext() *FilenameContext {
	var p = new(FilenameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_filename
	return p
}

func (*FilenameContext) IsFilenameContext() {}

func NewFilenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilenameContext {
	var p = new(FilenameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_filename

	return p
}

func (s *FilenameContext) GetParser() antlr.Parser { return s.parser }

func (s *FilenameContext) StrLit() antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, 0)
}

func (s *FilenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilenameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterFilename(s)
	}
}

func (s *FilenameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitFilename(s)
	}
}

func (p *NuggetParser) Filename() (localctx IFilenameContext) {
	localctx = NewFilenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NuggetParserRULE_filename)

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
		p.SetState(144)
		p.Match(NuggetParserT__5)
	}
	{
		p.SetState(145)
		p.Match(NuggetParserStrLit)
	}
	{
		p.SetState(146)
		p.Match(NuggetParserT__5)
	}

	return localctx
}

// ITimefilterContext is an interface to support dynamic dispatch.
type ITimefilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimefilterContext differentiates from other interfaces.
	IsTimefilterContext()
}

type TimefilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimefilterContext() *TimefilterContext {
	var p = new(TimefilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_timefilter
	return p
}

func (*TimefilterContext) IsTimefilterContext() {}

func NewTimefilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimefilterContext {
	var p = new(TimefilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_timefilter

	return p
}

func (s *TimefilterContext) GetParser() antlr.Parser { return s.parser }

func (s *TimefilterContext) Reference() IReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *TimefilterContext) OPERATION() antlr.TerminalNode {
	return s.GetToken(NuggetParserOPERATION, 0)
}

func (s *TimefilterContext) Date() IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *TimefilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimefilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimefilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterTimefilter(s)
	}
}

func (s *TimefilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitTimefilter(s)
	}
}

func (p *NuggetParser) Timefilter() (localctx ITimefilterContext) {
	localctx = NewTimefilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NuggetParserRULE_timefilter)

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
		p.SetState(148)
		p.Reference()
	}
	{
		p.SetState(149)
		p.Match(NuggetParserOPERATION)
	}
	{
		p.SetState(150)
		p.Date()
	}

	return localctx
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_reference
	return p
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) CTIME() antlr.TerminalNode {
	return s.GetToken(NuggetParserCTIME, 0)
}

func (s *ReferenceContext) MTIME() antlr.TerminalNode {
	return s.GetToken(NuggetParserMTIME, 0)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitReference(s)
	}
}

func (p *NuggetParser) Reference() (localctx IReferenceContext) {
	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NuggetParserRULE_reference)
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
	p.SetState(152)
	_la = p.GetTokenStream().LA(1)

	if !(_la == NuggetParserCTIME || _la == NuggetParserMTIME) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_date
	return p
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) DATE() antlr.TerminalNode {
	return s.GetToken(NuggetParserDATE, 0)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitDate(s)
	}
}

func (p *NuggetParser) Date() (localctx IDateContext) {
	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NuggetParserRULE_date)

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
		p.SetState(154)
		p.Match(NuggetParserT__5)
	}
	{
		p.SetState(155)
		p.Match(NuggetParserDATE)
	}
	{
		p.SetState(156)
		p.Match(NuggetParserT__5)
	}

	return localctx
}

// ISubtypeContext is an interface to support dynamic dispatch.
type ISubtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubtypeContext differentiates from other interfaces.
	IsSubtypeContext()
}

type SubtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubtypeContext() *SubtypeContext {
	var p = new(SubtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_subtype
	return p
}

func (*SubtypeContext) IsSubtypeContext() {}

func NewSubtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubtypeContext {
	var p = new(SubtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_subtype

	return p
}

func (s *SubtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SubtypeContext) FILES() antlr.TerminalNode {
	return s.GetToken(NuggetParserFILES, 0)
}

func (s *SubtypeContext) IMAGES() antlr.TerminalNode {
	return s.GetToken(NuggetParserIMAGES, 0)
}

func (s *SubtypeContext) DOCUMENTS() antlr.TerminalNode {
	return s.GetToken(NuggetParserDOCUMENTS, 0)
}

func (s *SubtypeContext) ALL() antlr.TerminalNode {
	return s.GetToken(NuggetParserALL, 0)
}

func (s *SubtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterSubtype(s)
	}
}

func (s *SubtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitSubtype(s)
	}
}

func (p *NuggetParser) Subtype() (localctx ISubtypeContext) {
	localctx = NewSubtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NuggetParserRULE_subtype)
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
	p.SetState(158)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserALL)|(1<<NuggetParserFILES)|(1<<NuggetParserIMAGES)|(1<<NuggetParserDOCUMENTS))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ITaskContext is an interface to support dynamic dispatch.
type ITaskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTaskContext differentiates from other interfaces.
	IsTaskContext()
}

type TaskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTaskContext() *TaskContext {
	var p = new(TaskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_task
	return p
}

func (*TaskContext) IsTaskContext() {}

func NewTaskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TaskContext {
	var p = new(TaskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_task

	return p
}

func (s *TaskContext) GetParser() antlr.Parser { return s.parser }

func (s *TaskContext) HASH() antlr.TerminalNode {
	return s.GetToken(NuggetParserHASH, 0)
}

func (s *TaskContext) EXTRACT() antlr.TerminalNode {
	return s.GetToken(NuggetParserEXTRACT, 0)
}

func (s *TaskContext) SELECT() antlr.TerminalNode {
	return s.GetToken(NuggetParserSELECT, 0)
}

func (s *TaskContext) JOIN() antlr.TerminalNode {
	return s.GetToken(NuggetParserJOIN, 0)
}

func (s *TaskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TaskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterTask(s)
	}
}

func (s *TaskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitTask(s)
	}
}

func (p *NuggetParser) Task() (localctx ITaskContext) {
	localctx = NewTaskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NuggetParserRULE_task)
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
	p.SetState(160)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NuggetParserHASH)|(1<<NuggetParserSELECT)|(1<<NuggetParserEXTRACT)|(1<<NuggetParserJOIN))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_target
	return p
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) StrLit() antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, 0)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (p *NuggetParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NuggetParserRULE_target)

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

	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(NuggetParserT__5)
		}
		{
			p.SetState(163)
			p.Match(NuggetParserStrLit)
		}
		{
			p.SetState(164)
			p.Match(NuggetParserT__5)
		}

	case NuggetParserStrLit:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(NuggetParserStrLit)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) AllStrLit() []antlr.TerminalNode {
	return s.GetTokens(NuggetParserStrLit)
}

func (s *FieldContext) StrLit(i int) antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, i)
}

func (s *FieldContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *NuggetParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NuggetParserRULE_field)
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

	p.SetState(180)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserStrLit:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(168)
			p.Match(NuggetParserStrLit)
		}
		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NuggetParserT__2 {
			{
				p.SetState(169)
				p.Match(NuggetParserT__2)
			}
			{
				p.SetState(170)
				p.Match(NuggetParserStrLit)
			}

			p.SetState(175)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case NuggetParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.Match(NuggetParserT__6)
		}
		{
			p.SetState(177)
			p.Field()
		}
		{
			p.SetState(178)
			p.Match(NuggetParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISourceidentifierContext is an interface to support dynamic dispatch.
type ISourceidentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceidentifierContext differentiates from other interfaces.
	IsSourceidentifierContext()
}

type SourceidentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceidentifierContext() *SourceidentifierContext {
	var p = new(SourceidentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_sourceidentifier
	return p
}

func (*SourceidentifierContext) IsSourceidentifierContext() {}

func NewSourceidentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceidentifierContext {
	var p = new(SourceidentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_sourceidentifier

	return p
}

func (s *SourceidentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceidentifierContext) StrLit() antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, 0)
}

func (s *SourceidentifierContext) Sourceidentifier() ISourceidentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceidentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceidentifierContext)
}

func (s *SourceidentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceidentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceidentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterSourceidentifier(s)
	}
}

func (s *SourceidentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitSourceidentifier(s)
	}
}

func (p *NuggetParser) Sourceidentifier() (localctx ISourceidentifierContext) {
	localctx = NewSourceidentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NuggetParserRULE_sourceidentifier)

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

	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NuggetParserStrLit:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(182)
			p.Match(NuggetParserStrLit)
		}

	case NuggetParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Match(NuggetParserT__6)
		}
		{
			p.SetState(184)
			p.Sourceidentifier()
		}
		{
			p.SetState(185)
			p.Match(NuggetParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrintIdContext is an interface to support dynamic dispatch.
type IPrintIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrintIdContext differentiates from other interfaces.
	IsPrintIdContext()
}

type PrintIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintIdContext() *PrintIdContext {
	var p = new(PrintIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_printId
	return p
}

func (*PrintIdContext) IsPrintIdContext() {}

func NewPrintIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintIdContext {
	var p = new(PrintIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_printId

	return p
}

func (s *PrintIdContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintIdContext) StrLit() antlr.TerminalNode {
	return s.GetToken(NuggetParserStrLit, 0)
}

func (s *PrintIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterPrintId(s)
	}
}

func (s *PrintIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitPrintId(s)
	}
}

func (p *NuggetParser) PrintId() (localctx IPrintIdContext) {
	localctx = NewPrintIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NuggetParserRULE_printId)

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
		p.SetState(189)
		p.Match(NuggetParserStrLit)
	}

	return localctx
}

// ISinContext is an interface to support dynamic dispatch.
type ISinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSinContext differentiates from other interfaces.
	IsSinContext()
}

type SinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySinContext() *SinContext {
	var p = new(SinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NuggetParserRULE_sin
	return p
}

func (*SinContext) IsSinContext() {}

func NewSinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SinContext {
	var p = new(SinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NuggetParserRULE_sin

	return p
}

func (s *SinContext) GetParser() antlr.Parser { return s.parser }

func (s *SinContext) SIN() antlr.TerminalNode {
	return s.GetToken(NuggetParserSIN, 0)
}

func (s *SinContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(NuggetParserNUMBER, 0)
}

func (s *SinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.EnterSin(s)
	}
}

func (s *SinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NuggetListener); ok {
		listenerT.ExitSin(s)
	}
}

func (p *NuggetParser) Sin() (localctx ISinContext) {
	localctx = NewSinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NuggetParserRULE_sin)

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
		p.SetState(191)
		p.Match(NuggetParserSIN)
	}
	{
		p.SetState(192)
		p.Match(NuggetParserT__7)
	}
	{
		p.SetState(193)
		p.Match(NuggetParserNUMBER)
	}
	{
		p.SetState(194)
		p.Match(NuggetParserT__8)
	}

	return localctx
}

func (p *NuggetParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *InitextractContext = nil
		if localctx != nil {
			t = localctx.(*InitextractContext)
		}
		return p.Initextract_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NuggetParser) Initextract_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
