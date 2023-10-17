package cmd_gen

import (
	"gint/collect"
)

func genConfig(projDir string, desc collect.DescData) error {

	if err := genFile(config{
		projDir:   projDir,
		subDir:    configDir,
		fileName:  "config.go",
		isCover:   false,
		tmplName:  "dto.template",
		embedTmpl: configTmpl,
		data:      nil,
	}); err != nil {
		return err
	}

	return nil
}
