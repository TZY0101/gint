// Code generated from E:/Project/GoProject-Larning/antlr\gint_demo.g4 by ANTLR 4.12.0. DO NOT EDIT.

package desc // gint_demo
import (
	"github.com/antlr4-go/antlr"
	"strings"
)

// BasedescListener is a complete listener for a parse tree produced by descParser.
type BasedescListener struct{
	Dtos []Dto
	ApiModule
}

type Dto struct {
    DtoName string
	Fields  []Field
}

type Field struct {
    Name string
	Type string
}

type ApiModule struct {
    Module string
	Prefix string
	Apis []Api
}

type Api struct {
    Method string
	Path string
	ApiName string
	Req string
	Resp string
}

var _ descListener = &BasedescListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedescListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedescListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedescListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedescListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BasedescListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BasedescListener) ExitFile(ctx *FileContext) {}

// EnterDto is called when production dto is entered.
func (s *BasedescListener) EnterDto(ctx *DtoContext) {
	d := Dto{
		DtoName: ctx.DtoName().GetText(),
	}
	fileds := make([]Field, 0)
	for _, c := range ctx.AllFieldDeclaration() {
		fileds = append(fileds, Field{c.FieldName().GetText(), c.FieldType().GetText()})
	}
	d.Fields = fileds
	s.Dtos = append(s.Dtos, d)
}

// ExitDto is called when production dto is exited.
func (s *BasedescListener) ExitDto(ctx *DtoContext) {
}

// EnterDtoName is called when production dtoName is entered.
func (s *BasedescListener) EnterDtoName(ctx *DtoNameContext) {}

// ExitDtoName is called when production dtoName is exited.
func (s *BasedescListener) ExitDtoName(ctx *DtoNameContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BasedescListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BasedescListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BasedescListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BasedescListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BasedescListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BasedescListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterApis is called when production apis is entered.
func (s *BasedescListener) EnterApis(ctx *ApisContext) {
	m := ApiModule{
		Module: strings.ToLower(ctx.ApiModuleName().GetText()),
	}

	apis := make([]Api, 0)
	for _, c := range ctx.AllApiDesc() {
		api := Api{
			Method: c.METHODS().GetText(),
			Path: c.Path().GetText(),
			ApiName: c.ApiName().GetText(),
			Req: c.Req().GetText(),
			Resp: c.Resp().GetText(),
		}
		apis = append(apis, api)
	}
	m.Apis = apis

	s.ApiModule = m
}

func (s *BasedescListener) EnterReq(c *ReqContext) {

}

func (s *BasedescListener) EnterResp(c *RespContext) {

}

func (s *BasedescListener) ExitReq(c *ReqContext) {

}

func (s *BasedescListener) ExitResp(c *RespContext) {

}

// ExitApis is called when production apis is exited.
func (s *BasedescListener) ExitApis(ctx *ApisContext) {}

// EnterApiDesc is called when production apiDesc is entered.
func (s *BasedescListener) EnterApiDesc(ctx *ApiDescContext) {}

// ExitApiDesc is called when production apiDesc is exited.
func (s *BasedescListener) ExitApiDesc(ctx *ApiDescContext) {}

// EnterApiModuleName is called when production apiModuleName is entered.
func (s *BasedescListener) EnterApiModuleName(ctx *ApiModuleNameContext) {}

// ExitApiModuleName is called when production apiModuleName is exited.
func (s *BasedescListener) ExitApiModuleName(ctx *ApiModuleNameContext) {}

// EnterPath is called when production path is entered.
func (s *BasedescListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BasedescListener) ExitPath(ctx *PathContext) {}

// EnterPathComponent is called when production pathComponent is entered.
func (s *BasedescListener) EnterPathComponent(ctx *PathComponentContext) {}

// ExitPathComponent is called when production pathComponent is exited.
func (s *BasedescListener) ExitPathComponent(ctx *PathComponentContext) {}

// EnterApiName is called when production apiName is entered.
func (s *BasedescListener) EnterApiName(ctx *ApiNameContext) {}

// ExitApiName is called when production apiName is exited.
func (s *BasedescListener) ExitApiName(ctx *ApiNameContext) {}
