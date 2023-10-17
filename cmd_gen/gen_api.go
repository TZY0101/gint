package cmd_gen

import (
	"gint/collect"
	"gint/util/pathx"
	"path"
	"path/filepath"
	"strings"
)

func genApi(projDir string, desc collect.DescData) error {
	modName, err := pathx.GetModName(projDir)
	if err != nil {
		return err
	}

	apis := make([]ApiTmpl, 0)
	for _, api := range desc.ApiModule.Apis {
		a := ApiTmpl{
			PackageName: desc.ApiModule.Module,
			Imports: []string{
				path.Join(modName, logicDir, desc.ApiModule.Module),
				path.Join(modName, ginxDir),
				path.Join(modName, dtoDir),
			},
			ApiName:     api.ApiName,
			HandlerName: strings.Title(api.ApiName),
			LogicName:   strings.Title(api.ApiName),
			Req:         api.Req,
			Resp:        api.Resp,
		}
		apis = append(apis, a)
	}

	for _, api := range apis {
		if err := genFile(config{
			projDir:   projDir,
			subDir:    filepath.Join(apiDir, desc.ApiModule.Module),
			fileName:  strings.ToLower(api.ApiName) + ".go",
			isCover:   true,
			tmplName:  "api.template",
			embedTmpl: apiTmpl,
			data:      api,
		}); err != nil {
			return err
		}
	}

	return nil
}
