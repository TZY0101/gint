// Code generated from E:/Project/GoProject-Larning/antlr\gint_demo.g4 by ANTLR 4.12.0. DO NOT EDIT.

package desc // gint_demo
import (
	"fmt"
	"strconv"
	"sync"

)

import "github.com/antlr4-go/antlr"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type descParser struct {
	*antlr.BaseParser
}

var descParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func descParserInit() {
	staticData := &descParserStaticData
	staticData.literalNames = []string{
		"", "'{'", "'}'", "'('", "')'", "'/'", "':'", "", "", "", "", "", "'byte'",
		"'rune'", "'string'", "'bool'", "", "", "'APIS'", "", "'DTO'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "DATATYPE", "INT", "UINT", "FLOAT", "COMPLEX",
		"BYTE", "RUNE", "STRING", "BOOL", "SLICE", "MAP", "APIS", "METHODS",
		"DTO", "IDENTIFIER", "WS", "NEWLINE",
	}
	staticData.ruleNames = []string{
		"file", "dto", "dtoName", "fieldDeclaration", "fieldName", "fieldType",
		"apis", "apiDesc", "req", "resp", "apiModuleName", "path", "pathComponent",
		"apiName",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 126, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 5, 0, 30, 8, 0, 10,
		0, 12, 0, 33, 9, 0, 1, 0, 3, 0, 36, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		42, 8, 1, 10, 1, 12, 1, 45, 9, 1, 1, 1, 4, 1, 48, 8, 1, 11, 1, 12, 1, 49,
		1, 1, 1, 1, 5, 1, 54, 8, 1, 10, 1, 12, 1, 57, 9, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 5, 3, 64, 8, 3, 10, 3, 12, 3, 67, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 77, 8, 6, 10, 6, 12, 6, 80, 9, 6, 1, 6, 4,
		6, 83, 8, 6, 11, 6, 12, 6, 84, 1, 6, 1, 6, 5, 6, 89, 8, 6, 10, 6, 12, 6,
		92, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 102, 8,
		7, 10, 7, 12, 7, 105, 9, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 4, 11, 115, 8, 11, 11, 11, 12, 11, 116, 1, 12, 1, 12, 1, 12, 3,
		12, 122, 8, 12, 1, 13, 1, 13, 1, 13, 0, 0, 14, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 0, 0, 123, 0, 31, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0,
		4, 58, 1, 0, 0, 0, 6, 60, 1, 0, 0, 0, 8, 68, 1, 0, 0, 0, 10, 70, 1, 0,
		0, 0, 12, 72, 1, 0, 0, 0, 14, 93, 1, 0, 0, 0, 16, 106, 1, 0, 0, 0, 18,
		108, 1, 0, 0, 0, 20, 110, 1, 0, 0, 0, 22, 114, 1, 0, 0, 0, 24, 121, 1,
		0, 0, 0, 26, 123, 1, 0, 0, 0, 28, 30, 3, 2, 1, 0, 29, 28, 1, 0, 0, 0, 30,
		33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 35, 1, 0, 0,
		0, 33, 31, 1, 0, 0, 0, 34, 36, 3, 12, 6, 0, 35, 34, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 1, 1, 0, 0, 0, 37, 38, 5, 20, 0, 0, 38, 39, 3, 4, 2, 0,
		39, 43, 5, 1, 0, 0, 40, 42, 5, 23, 0, 0, 41, 40, 1, 0, 0, 0, 42, 45, 1,
		0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45,
		43, 1, 0, 0, 0, 46, 48, 3, 6, 3, 0, 47, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0,
		0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 55,
		5, 2, 0, 0, 52, 54, 5, 23, 0, 0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0,
		55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 3, 1, 0, 0, 0, 57, 55, 1, 0,
		0, 0, 58, 59, 5, 21, 0, 0, 59, 5, 1, 0, 0, 0, 60, 61, 3, 8, 4, 0, 61, 65,
		3, 10, 5, 0, 62, 64, 5, 23, 0, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0,
		0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 7, 1, 0, 0, 0, 67, 65, 1,
		0, 0, 0, 68, 69, 5, 21, 0, 0, 69, 9, 1, 0, 0, 0, 70, 71, 5, 7, 0, 0, 71,
		11, 1, 0, 0, 0, 72, 73, 5, 18, 0, 0, 73, 74, 3, 20, 10, 0, 74, 78, 5, 1,
		0, 0, 75, 77, 5, 23, 0, 0, 76, 75, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78,
		76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 78, 1, 0, 0,
		0, 81, 83, 3, 14, 7, 0, 82, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 82,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 90, 5, 2, 0, 0,
		87, 89, 5, 23, 0, 0, 88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 13, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93,
		94, 5, 19, 0, 0, 94, 95, 3, 22, 11, 0, 95, 96, 3, 26, 13, 0, 96, 97, 5,
		3, 0, 0, 97, 98, 3, 16, 8, 0, 98, 99, 5, 4, 0, 0, 99, 103, 3, 18, 9, 0,
		100, 102, 5, 23, 0, 0, 101, 100, 1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103,
		101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 15, 1, 0, 0, 0, 105, 103, 1,
		0, 0, 0, 106, 107, 5, 21, 0, 0, 107, 17, 1, 0, 0, 0, 108, 109, 5, 21, 0,
		0, 109, 19, 1, 0, 0, 0, 110, 111, 5, 21, 0, 0, 111, 21, 1, 0, 0, 0, 112,
		113, 5, 5, 0, 0, 113, 115, 3, 24, 12, 0, 114, 112, 1, 0, 0, 0, 115, 116,
		1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 23, 1, 0,
		0, 0, 118, 122, 5, 21, 0, 0, 119, 120, 5, 6, 0, 0, 120, 122, 5, 21, 0,
		0, 121, 118, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 25, 1, 0, 0, 0, 123,
		124, 5, 21, 0, 0, 124, 27, 1, 0, 0, 0, 12, 31, 35, 43, 49, 55, 65, 78,
		84, 90, 103, 116, 121,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// descParserInit initializes any static state used to implement descParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewdescParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DescParserInit() {
	staticData := &descParserStaticData
	staticData.once.Do(descParserInit)
}

// NewdescParser produces a new parser instance for the optional input antlr.TokenStream.
func NewdescParser(input antlr.TokenStream) *descParser {
	DescParserInit()
	this := new(descParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &descParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "gint_demo.g4"

	return this
}

// descParser tokens.
const (
	descParserEOF        = antlr.TokenEOF
	descParserT__0       = 1
	descParserT__1       = 2
	descParserT__2       = 3
	descParserT__3       = 4
	descParserT__4       = 5
	descParserT__5       = 6
	descParserDATATYPE   = 7
	descParserINT        = 8
	descParserUINT       = 9
	descParserFLOAT      = 10
	descParserCOMPLEX    = 11
	descParserBYTE       = 12
	descParserRUNE       = 13
	descParserSTRING     = 14
	descParserBOOL       = 15
	descParserSLICE      = 16
	descParserMAP        = 17
	descParserAPIS       = 18
	descParserMETHODS    = 19
	descParserDTO        = 20
	descParserIDENTIFIER = 21
	descParserWS         = 22
	descParserNEWLINE    = 23
)

// descParser rules.
const (
	descParserRULE_file             = 0
	descParserRULE_dto              = 1
	descParserRULE_dtoName          = 2
	descParserRULE_fieldDeclaration = 3
	descParserRULE_fieldName        = 4
	descParserRULE_fieldType        = 5
	descParserRULE_apis             = 6
	descParserRULE_apiDesc          = 7
	descParserRULE_req              = 8
	descParserRULE_resp             = 9
	descParserRULE_apiModuleName    = 10
	descParserRULE_path             = 11
	descParserRULE_pathComponent    = 12
	descParserRULE_apiName          = 13
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDto() []IDtoContext
	Dto(i int) IDtoContext
	Apis() IApisContext

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) AllDto() []IDtoContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDtoContext); ok {
			len++
		}
	}

	tst := make([]IDtoContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDtoContext); ok {
			tst[i] = t.(IDtoContext)
			i++
		}
	}

	return tst
}

func (s *FileContext) Dto(i int) IDtoContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDtoContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDtoContext)
}

func (s *FileContext) Apis() IApisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApisContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *descParser) File() (localctx IFileContext) {
	this := p
	_ = this

	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, descParserRULE_file)
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
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == descParserDTO {
		{
			p.SetState(28)
			p.Dto()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == descParserAPIS {
		{
			p.SetState(34)
			p.Apis()
		}

	}

	return localctx
}

// IDtoContext is an interface to support dynamic dispatch.
type IDtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DTO() antlr.TerminalNode
	DtoName() IDtoNameContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllFieldDeclaration() []IFieldDeclarationContext
	FieldDeclaration(i int) IFieldDeclarationContext

	// IsDtoContext differentiates from other interfaces.
	IsDtoContext()
}

type DtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDtoContext() *DtoContext {
	var p = new(DtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_dto
	return p
}

func (*DtoContext) IsDtoContext() {}

func NewDtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DtoContext {
	var p = new(DtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_dto

	return p
}

func (s *DtoContext) GetParser() antlr.Parser { return s.parser }

func (s *DtoContext) DTO() antlr.TerminalNode {
	return s.GetToken(descParserDTO, 0)
}

func (s *DtoContext) DtoName() IDtoNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDtoNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDtoNameContext)
}

func (s *DtoContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(descParserNEWLINE)
}

func (s *DtoContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(descParserNEWLINE, i)
}

func (s *DtoContext) AllFieldDeclaration() []IFieldDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IFieldDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDeclarationContext); ok {
			tst[i] = t.(IFieldDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *DtoContext) FieldDeclaration(i int) IFieldDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDeclarationContext)
}

func (s *DtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DtoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterDto(s)
	}
}

func (s *DtoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitDto(s)
	}
}

func (p *descParser) Dto() (localctx IDtoContext) {
	this := p
	_ = this

	localctx = NewDtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, descParserRULE_dto)
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
		p.SetState(37)
		p.Match(descParserDTO)
	}
	{
		p.SetState(38)
		p.DtoName()
	}
	{
		p.SetState(39)
		p.Match(descParserT__0)
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == descParserNEWLINE {
		{
			p.SetState(40)
			p.Match(descParserNEWLINE)
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == descParserIDENTIFIER {
		{
			p.SetState(46)
			p.FieldDeclaration()
		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(51)
		p.Match(descParserT__1)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == descParserNEWLINE {
		{
			p.SetState(52)
			p.Match(descParserNEWLINE)
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDtoNameContext is an interface to support dynamic dispatch.
type IDtoNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsDtoNameContext differentiates from other interfaces.
	IsDtoNameContext()
}

type DtoNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDtoNameContext() *DtoNameContext {
	var p = new(DtoNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_dtoName
	return p
}

func (*DtoNameContext) IsDtoNameContext() {}

func NewDtoNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DtoNameContext {
	var p = new(DtoNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_dtoName

	return p
}

func (s *DtoNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DtoNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(descParserIDENTIFIER, 0)
}

func (s *DtoNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DtoNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DtoNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterDtoName(s)
	}
}

func (s *DtoNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitDtoName(s)
	}
}

func (p *descParser) DtoName() (localctx IDtoNameContext) {
	this := p
	_ = this

	localctx = NewDtoNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, descParserRULE_dtoName)

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
		p.SetState(58)
		p.Match(descParserIDENTIFIER)
	}

	return localctx
}

// IFieldDeclarationContext is an interface to support dynamic dispatch.
type IFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldName() IFieldNameContext
	FieldType() IFieldTypeContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsFieldDeclarationContext differentiates from other interfaces.
	IsFieldDeclarationContext()
}

type FieldDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclarationContext() *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_fieldDeclaration
	return p
}

func (*FieldDeclarationContext) IsFieldDeclarationContext() {}

func NewFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_fieldDeclaration

	return p
}

func (s *FieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclarationContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldDeclarationContext) FieldType() IFieldTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldDeclarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(descParserNEWLINE)
}

func (s *FieldDeclarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(descParserNEWLINE, i)
}

func (s *FieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitFieldDeclaration(s)
	}
}

func (p *descParser) FieldDeclaration() (localctx IFieldDeclarationContext) {
	this := p
	_ = this

	localctx = NewFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, descParserRULE_fieldDeclaration)
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
		p.SetState(60)
		p.FieldName()
	}
	{
		p.SetState(61)
		p.FieldType()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == descParserNEWLINE {
		{
			p.SetState(62)
			p.Match(descParserNEWLINE)
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(descParserIDENTIFIER, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *descParser) FieldName() (localctx IFieldNameContext) {
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, descParserRULE_fieldName)

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
		p.SetState(68)
		p.Match(descParserIDENTIFIER)
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATATYPE() antlr.TerminalNode

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) DATATYPE() antlr.TerminalNode {
	return s.GetToken(descParserDATATYPE, 0)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *descParser) FieldType() (localctx IFieldTypeContext) {
	this := p
	_ = this

	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, descParserRULE_fieldType)

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
		p.SetState(70)
		p.Match(descParserDATATYPE)
	}

	return localctx
}

// IApisContext is an interface to support dynamic dispatch.
type IApisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	APIS() antlr.TerminalNode
	ApiModuleName() IApiModuleNameContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllApiDesc() []IApiDescContext
	ApiDesc(i int) IApiDescContext

	// IsApisContext differentiates from other interfaces.
	IsApisContext()
}

type ApisContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApisContext() *ApisContext {
	var p = new(ApisContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_apis
	return p
}

func (*ApisContext) IsApisContext() {}

func NewApisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApisContext {
	var p = new(ApisContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_apis

	return p
}

func (s *ApisContext) GetParser() antlr.Parser { return s.parser }

func (s *ApisContext) APIS() antlr.TerminalNode {
	return s.GetToken(descParserAPIS, 0)
}

func (s *ApisContext) ApiModuleName() IApiModuleNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApiModuleNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApiModuleNameContext)
}

func (s *ApisContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(descParserNEWLINE)
}

func (s *ApisContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(descParserNEWLINE, i)
}

func (s *ApisContext) AllApiDesc() []IApiDescContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IApiDescContext); ok {
			len++
		}
	}

	tst := make([]IApiDescContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IApiDescContext); ok {
			tst[i] = t.(IApiDescContext)
			i++
		}
	}

	return tst
}

func (s *ApisContext) ApiDesc(i int) IApiDescContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApiDescContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApiDescContext)
}

func (s *ApisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterApis(s)
	}
}

func (s *ApisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitApis(s)
	}
}

func (p *descParser) Apis() (localctx IApisContext) {
	this := p
	_ = this

	localctx = NewApisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, descParserRULE_apis)
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
		p.SetState(72)
		p.Match(descParserAPIS)
	}
	{
		p.SetState(73)
		p.ApiModuleName()
	}
	{
		p.SetState(74)
		p.Match(descParserT__0)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == descParserNEWLINE {
		{
			p.SetState(75)
			p.Match(descParserNEWLINE)
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == descParserMETHODS {
		{
			p.SetState(81)
			p.ApiDesc()
		}

		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(descParserT__1)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == descParserNEWLINE {
		{
			p.SetState(87)
			p.Match(descParserNEWLINE)
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IApiDescContext is an interface to support dynamic dispatch.
type IApiDescContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	METHODS() antlr.TerminalNode
	Path() IPathContext
	ApiName() IApiNameContext
	Req() IReqContext
	Resp() IRespContext
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsApiDescContext differentiates from other interfaces.
	IsApiDescContext()
}

type ApiDescContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiDescContext() *ApiDescContext {
	var p = new(ApiDescContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_apiDesc
	return p
}

func (*ApiDescContext) IsApiDescContext() {}

func NewApiDescContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiDescContext {
	var p = new(ApiDescContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_apiDesc

	return p
}

func (s *ApiDescContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiDescContext) METHODS() antlr.TerminalNode {
	return s.GetToken(descParserMETHODS, 0)
}

func (s *ApiDescContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *ApiDescContext) ApiName() IApiNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApiNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApiNameContext)
}

func (s *ApiDescContext) Req() IReqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReqContext)
}

func (s *ApiDescContext) Resp() IRespContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRespContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRespContext)
}

func (s *ApiDescContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(descParserNEWLINE)
}

func (s *ApiDescContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(descParserNEWLINE, i)
}

func (s *ApiDescContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiDescContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiDescContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterApiDesc(s)
	}
}

func (s *ApiDescContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitApiDesc(s)
	}
}

func (p *descParser) ApiDesc() (localctx IApiDescContext) {
	this := p
	_ = this

	localctx = NewApiDescContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, descParserRULE_apiDesc)
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
		p.SetState(93)
		p.Match(descParserMETHODS)
	}
	{
		p.SetState(94)
		p.Path()
	}
	{
		p.SetState(95)
		p.ApiName()
	}
	{
		p.SetState(96)
		p.Match(descParserT__2)
	}
	{
		p.SetState(97)
		p.Req()
	}
	{
		p.SetState(98)
		p.Match(descParserT__3)
	}
	{
		p.SetState(99)
		p.Resp()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == descParserNEWLINE {
		{
			p.SetState(100)
			p.Match(descParserNEWLINE)
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IReqContext is an interface to support dynamic dispatch.
type IReqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsReqContext differentiates from other interfaces.
	IsReqContext()
}

type ReqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReqContext() *ReqContext {
	var p = new(ReqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_req
	return p
}

func (*ReqContext) IsReqContext() {}

func NewReqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReqContext {
	var p = new(ReqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_req

	return p
}

func (s *ReqContext) GetParser() antlr.Parser { return s.parser }

func (s *ReqContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(descParserIDENTIFIER, 0)
}

func (s *ReqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterReq(s)
	}
}

func (s *ReqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitReq(s)
	}
}

func (p *descParser) Req() (localctx IReqContext) {
	this := p
	_ = this

	localctx = NewReqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, descParserRULE_req)

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
		p.SetState(106)
		p.Match(descParserIDENTIFIER)
	}

	return localctx
}

// IRespContext is an interface to support dynamic dispatch.
type IRespContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsRespContext differentiates from other interfaces.
	IsRespContext()
}

type RespContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRespContext() *RespContext {
	var p = new(RespContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_resp
	return p
}

func (*RespContext) IsRespContext() {}

func NewRespContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RespContext {
	var p = new(RespContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_resp

	return p
}

func (s *RespContext) GetParser() antlr.Parser { return s.parser }

func (s *RespContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(descParserIDENTIFIER, 0)
}

func (s *RespContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RespContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RespContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterResp(s)
	}
}

func (s *RespContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitResp(s)
	}
}

func (p *descParser) Resp() (localctx IRespContext) {
	this := p
	_ = this

	localctx = NewRespContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, descParserRULE_resp)

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
		p.Match(descParserIDENTIFIER)
	}

	return localctx
}

// IApiModuleNameContext is an interface to support dynamic dispatch.
type IApiModuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsApiModuleNameContext differentiates from other interfaces.
	IsApiModuleNameContext()
}

type ApiModuleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiModuleNameContext() *ApiModuleNameContext {
	var p = new(ApiModuleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_apiModuleName
	return p
}

func (*ApiModuleNameContext) IsApiModuleNameContext() {}

func NewApiModuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiModuleNameContext {
	var p = new(ApiModuleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_apiModuleName

	return p
}

func (s *ApiModuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiModuleNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(descParserIDENTIFIER, 0)
}

func (s *ApiModuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiModuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiModuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterApiModuleName(s)
	}
}

func (s *ApiModuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitApiModuleName(s)
	}
}

func (p *descParser) ApiModuleName() (localctx IApiModuleNameContext) {
	this := p
	_ = this

	localctx = NewApiModuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, descParserRULE_apiModuleName)

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
		p.SetState(110)
		p.Match(descParserIDENTIFIER)
	}

	return localctx
}

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPathComponent() []IPathComponentContext
	PathComponent(i int) IPathComponentContext

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) AllPathComponent() []IPathComponentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPathComponentContext); ok {
			len++
		}
	}

	tst := make([]IPathComponentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPathComponentContext); ok {
			tst[i] = t.(IPathComponentContext)
			i++
		}
	}

	return tst
}

func (s *PathContext) PathComponent(i int) IPathComponentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathComponentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathComponentContext)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitPath(s)
	}
}

func (p *descParser) Path() (localctx IPathContext) {
	this := p
	_ = this

	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, descParserRULE_path)
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
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == descParserT__4 {
		{
			p.SetState(112)
			p.Match(descParserT__4)
		}
		{
			p.SetState(113)
			p.PathComponent()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPathComponentContext is an interface to support dynamic dispatch.
type IPathComponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsPathComponentContext differentiates from other interfaces.
	IsPathComponentContext()
}

type PathComponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathComponentContext() *PathComponentContext {
	var p = new(PathComponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_pathComponent
	return p
}

func (*PathComponentContext) IsPathComponentContext() {}

func NewPathComponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathComponentContext {
	var p = new(PathComponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_pathComponent

	return p
}

func (s *PathComponentContext) GetParser() antlr.Parser { return s.parser }

func (s *PathComponentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(descParserIDENTIFIER, 0)
}

func (s *PathComponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathComponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathComponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterPathComponent(s)
	}
}

func (s *PathComponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitPathComponent(s)
	}
}

func (p *descParser) PathComponent() (localctx IPathComponentContext) {
	this := p
	_ = this

	localctx = NewPathComponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, descParserRULE_pathComponent)

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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case descParserIDENTIFIER:
		{
			p.SetState(118)
			p.Match(descParserIDENTIFIER)
		}

	case descParserT__5:
		{
			p.SetState(119)
			p.Match(descParserT__5)
		}
		{
			p.SetState(120)
			p.Match(descParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IApiNameContext is an interface to support dynamic dispatch.
type IApiNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsApiNameContext differentiates from other interfaces.
	IsApiNameContext()
}

type ApiNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApiNameContext() *ApiNameContext {
	var p = new(ApiNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = descParserRULE_apiName
	return p
}

func (*ApiNameContext) IsApiNameContext() {}

func NewApiNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApiNameContext {
	var p = new(ApiNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = descParserRULE_apiName

	return p
}

func (s *ApiNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ApiNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(descParserIDENTIFIER, 0)
}

func (s *ApiNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApiNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApiNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.EnterApiName(s)
	}
}

func (s *ApiNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(descListener); ok {
		listenerT.ExitApiName(s)
	}
}

func (p *descParser) ApiName() (localctx IApiNameContext) {
	this := p
	_ = this

	localctx = NewApiNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, descParserRULE_apiName)

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
		p.SetState(123)
		p.Match(descParserIDENTIFIER)
	}

	return localctx
}
