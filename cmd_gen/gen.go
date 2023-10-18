package cmd_gen

import (
	"gint/collect"
	"gint/util/logx"
	"gint/util/pathx"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (
	VarStringDesc string
)

func GenProjByDesc(_ *cobra.Command, _ []string) error {
	descFile := VarStringDesc
	projDir, err := pathx.GetProjDir(descFile)
	if err != nil {
		return err
	}

	wd, _ := os.Getwd()
	descFile = filepath.Join(wd, descFile)
	err = GenProj(projDir, descFile)

	return nil
}

// GenProj read data from descFile
// 初始化时，projDir是abs("dir")，其他命令调用是pathx.GetProjDir
// descFile 是描述文件绝对路径
func GenProj(projDir, descFile string) error {
	// collect data
	desc, err := collect.CollectData(descFile)
	if err != nil {
		return err
	}

	// cmd_gen files
	logx.MustNil(genMain(projDir, desc))
	logx.MustNil(genConfYaml(projDir, desc))
	logx.MustNil(genConfig(projDir, desc))
	logx.MustNil(genRouter(projDir, desc))
	logx.MustNil(genRoutes(projDir, desc))
	logx.MustNil(genDto(projDir, desc))
	logx.MustNil(genApi(projDir, desc))
	logx.MustNil(genLogic(projDir, desc))
	logx.MustNil(genGinx(projDir, desc))
	logx.MustNil(genErrorx(projDir, desc))

	return nil
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstLower 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
