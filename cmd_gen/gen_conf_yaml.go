package cmd_gen

import (
	"github.com/TZY0101/gint/collect"
	"github.com/TZY0101/gint/util/pathx"
)

func genConfYaml(projDir string, desc collect.DescData) error {
	modName, err := pathx.GetModName(projDir)
	if err != nil {
		return err
	}

	if err := genFile(config{
		projDir:   projDir,
		subDir:    configDir,
		fileName:  "config.yaml",
		isCover:   false,
		tmplName:  "dto.template",
		embedTmpl: confYamlTmpl,
		data: map[string]string{
			"ModName": modName,
		},
	}); err != nil {
		return err
	}

	return nil
}
