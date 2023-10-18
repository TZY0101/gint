package cmd_gen

import (
	"github.com/TZY0101/gint/collect"
)

func genDto(projDir string, desc collect.DescData) error {
	if err := genFile(config{
		projDir:   projDir,
		subDir:    dtoDir,
		fileName:  "dto.go",
		isCover:   true,
		tmplName:  "dto.template",
		embedTmpl: dtoTmpl,
		data:      desc.Dtos,
	}); err != nil {
		return err
	}

	return nil
}
