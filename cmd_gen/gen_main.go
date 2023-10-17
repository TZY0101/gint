package cmd_gen

import (
	"gint/collect"
	"gint/util/pathx"
)

func genMain(projDir string, desc collect.DescData) error {
	modName, err := pathx.GetModName(projDir)
	if err != nil {
		return err
	}

	if err := genFile(config{
		projDir:   projDir,
		subDir:    cmdDir,
		fileName:  "main.go",
		isCover:   false,
		tmplName:  "main.template",
		embedTmpl: mainTmpl,
		data: map[string]string{
			"ModName": modName,
		},
	}); err != nil {
		return err
	}

	return nil
}
