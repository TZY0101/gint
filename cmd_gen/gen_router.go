package cmd_gen

import (
	"github.com/TZY0101/gint/collect"
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
