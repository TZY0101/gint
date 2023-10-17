package cmd_gen

import (
	"gint/collect"
	"gint/util/pathx"
	"github.com/jinzhu/copier"
	"path"
	"strings"
)

// todo: The next update will instead generate multiple groups of routes
func genRoutes(projDir string, desc collect.DescData) error {
	apis := make([]Route, len(desc.ApiModule.Apis))
	if err := copier.Copy(&apis, &desc.ApiModule.Apis); err != nil {
		return err
	}

	modName, err := pathx.GetModName(projDir)
	if err != nil {
		return err
	}

	importApiPrefix := path.Join(modName, apiDir)
	module := desc.ApiModule.Module
	routesData := RoutesTmplData{
		Imports:            []string{path.Join(importApiPrefix, module)},
		ModuleToUpperCamel: strings.Title(module),
		ModuleToLowerCamel: module,
		Apis:               apis,
	}
	if err := genFile(config{
		projDir:   projDir,
		subDir:    routerDir,
		fileName:  module + "_routes.go",
		isCover:   true,
		tmplName:  "routes.template",
		embedTmpl: routesTmpl,
		data:      routesData,
	}); err != nil {
		return err
	}

	return nil
}

// Method toUpper method of descAntlr.Api
func (r *Route) Method(method string) {
	r.MethodToUpper = strings.ToUpper(method)
}

func (r *Route) ApiName(apiName string) {
	r.ApiNameToTitle = strings.Title(apiName)
}
