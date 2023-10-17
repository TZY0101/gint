package cmd_gen

import (
	"gint/collect"
)

func genErrorx(projDir string, desc collect.DescData) error {
	if err := genFile(config{
		projDir:   projDir,
		subDir:    errorxDir,
		fileName:  "base.go",
		isCover:   false,
		tmplName:  "base.template",
		embedTmpl: errorxBaseTmpl,
		data:      nil,
	}); err != nil {
		return err
	}

	if err := genFile(config{
		projDir:   projDir,
		subDir:    errorxDir,
		fileName:  "code.go",
		isCover:   false,
		tmplName:  "code.template",
		embedTmpl: errorxCodeTmpl,
		data:      nil,
	}); err != nil {
		return err
	}

	if err := genFile(config{
		projDir:   projDir,
		subDir:    errorxDir,
		fileName:  "msg.go",
		isCover:   false,
		tmplName:  "msg.template",
		embedTmpl: errorxMsgTmpl,
		data:      nil,
	}); err != nil {
		return err
	}

	return nil
}
