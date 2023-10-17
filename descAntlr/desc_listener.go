// Code generated from E:/Project/GoProject-Larning/antlr\gint_demo.g4 by ANTLR 4.12.0. DO NOT EDIT.

package desc // gint_demo
import (
	"github.com/antlr4-go/antlr"
)

// descListener is a complete listener for a parse tree produced by descParser.
type descListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterDto is called when entering the dto production.
	EnterDto(c *DtoContext)

	// EnterDtoName is called when entering the dtoName production.
	EnterDtoName(c *DtoNameContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterApis is called when entering the apis production.
	EnterApis(c *ApisContext)

	// EnterApiDesc is called when entering the apiDesc production.
	EnterApiDesc(c *ApiDescContext)

	// EnterReq is called when entering the req production.
	EnterReq(c *ReqContext)

	// EnterResp is called when entering the resp production.
	EnterResp(c *RespContext)

	// EnterApiModuleName is called when entering the apiModuleName production.
	EnterApiModuleName(c *ApiModuleNameContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterPathComponent is called when entering the pathComponent production.
	EnterPathComponent(c *PathComponentContext)

	// EnterApiName is called when entering the apiName production.
	EnterApiName(c *ApiNameContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitDto is called when exiting the dto production.
	ExitDto(c *DtoContext)

	// ExitDtoName is called when exiting the dtoName production.
	ExitDtoName(c *DtoNameContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitApis is called when exiting the apis production.
	ExitApis(c *ApisContext)

	// ExitApiDesc is called when exiting the apiDesc production.
	ExitApiDesc(c *ApiDescContext)

	// ExitReq is called when exiting the req production.
	ExitReq(c *ReqContext)

	// ExitResp is called when exiting the resp production.
	ExitResp(c *RespContext)

	// ExitApiModuleName is called when exiting the apiModuleName production.
	ExitApiModuleName(c *ApiModuleNameContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitPathComponent is called when exiting the pathComponent production.
	ExitPathComponent(c *PathComponentContext)

	// ExitApiName is called when exiting the apiName production.
	ExitApiName(c *ApiNameContext)
}
