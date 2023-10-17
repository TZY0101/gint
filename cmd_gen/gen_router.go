package cmd_gen

import (
	"gint/collect"
)

func genRouter(projDir string, m collect.DescData) error {

	if err := genFile(config{
		projDir:   projDir,
		subDir:    routerDir,
		fileName:  "router.go",
		isCover:   false,
		tmplName:  "router.template",
		embedTmpl: routerTmpl,
		data:      nil,
	}); err != nil {
		return err
	}

	return nil
}
