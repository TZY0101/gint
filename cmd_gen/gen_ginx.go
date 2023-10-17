package cmd_gen

import (
	"gint/collect"
	"gint/util/pathx"
	"path"
)

func genGinx(projDir string, desc collect.DescData) error {
	modName, err := pathx.GetModName(projDir)
	if err != nil {
		return err
	}

	if err := genFile(config{
		projDir:   projDir,
		subDir:    ginxDir,
		fileName:  "ginx.go",
		isCover:   false,
		tmplName:  "ginx.template",
		embedTmpl: ginxTmpl,
		data:      map[string]interface{}{
			"Imports": []string{path.Join(modName, errorxDir)},
		},
	}); err != nil {
		return err
	}

	return nil
}
