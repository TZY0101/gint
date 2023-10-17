package cmd_gen

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {
	absolutePath, err := filepath.Abs("test.d")
	if err != nil {
		fmt.Println("无法转换为绝对路径:", err)
		return
	}

	fmt.Println("绝对路径:", absolutePath)
}

func TestGenRouter(t *testing.T) {

}

func TestGenDtoByCollect(t *testing.T) {
	//dtos, _ := collect.CollectData("test.desc")
	//dir, _ := pathx.GetProjDir()
	//if err := genFile(config{
	//	projDir:   dir,
	//	subDir:    dtoDir,
	//	fileName:  "dto.go",
	//	isCover:   true,
	//	tmplName:  "dto.template",
	//	embedTmpl: dtoTmpl,
	//	data:      dtos,
	//}); err != nil {
	//	fmt.Printf("cmd_gen err: %v \n", err)
	//}
}

func TestGenApiByCollect(t *testing.T) {
	//desc, _ := collect.CollectData("./","test.desc")
	//apiModule := desc.ApiModule
	//dataSet := make([]ApiTmpl, 0)
	//for _, api := range apiModule.Apis {
	//	a := ApiTmpl{
	//		PackageName: apiModule.Module,
	//		ApiNameToTitle:     api.ApiNameToTitle,
	//		HandlerName: api.ApiNameToTitle,
	//		LogicName:   api.ApiNameToTitle,
	//		Req:         api.Req,
	//		Resp:        api.Resp,
	//	}
	//	dataSet = append(dataSet, a)
	//}
	//
	//for _, data := range dataSet {
	//	if err := genFile(config{
	//		//FileName: data.ApiNameToTitle + ".go",
	//		//SubDir:   ApiSubDir,
	//		//DirName:  data.PackageName,
	//		//IsCover:  true,
	//		//TmplPath: TmplDir,
	//		//TmplName: ApiTmplName,
	//		data: data,
	//	}); err != nil {
	//		fmt.Printf("cmd_gen err: %v \n", err)
	//	}
	//	if err := genFile(config{
	//		//FileName: data.ApiNameToTitle + ".go",
	//		//SubDir:   LogicSubDir,
	//		//DirName:  data.PackageName,
	//		//IsCover:  true,
	//		//TmplPath: TmplDir,
	//		//TmplName: LogicTmplName,
	//		data: data,
	//	}); err != nil {
	//		fmt.Printf("cmd_gen err: %v \n", err)
	//		return
	//	}
	//}
	//
	////apis := make([]*Route, len(apiModule.Apis))
	//var apis []Route
	//if err := copier.Copy(&apis, apiModule.Apis); err != nil {
	//	fmt.Printf("copy err: %v \n", err)
	//	return
	//}
	//
	//routesData := RoutesTmplData{
	//	Imports:            []string{path.Join(importApiPrefix, strings.ToLower(apiModule.Module))},
	//	ModuleToUpperCamel: strings.Title(apiModule.Module),
	//	ModuleToLowerCamel: apiModule.Module,
	//	Apis:               apis,
	//}
	//if err := genFile(config{
	//	//FileName: apiModule.Module + "_routes.go",
	//	//SubDir:   RouterSubDir,
	//	//DirName:  "",
	//	//IsCover:  true,
	//	//TmplPath: TmplDir,
	//	//TmplName: RoutesTmplName,
	//	data: routesData,
	//}); err != nil {
	//	fmt.Printf("cmd_gen err: %v \n", err)
	//}
}

func TestInitProj(t *testing.T) {

	//if err :=genFile(config{
	//	FileName: "go.mod",
	//	SubDir:
	//})
}
