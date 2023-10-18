package cmd_gen

import (
	"bytes"
	"errors"
	"github.com/TZY0101/gint/util/pathx"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"text/template"
)

var (
	ErrReCreated = errors.New("projDir already exist")
)

// Config 尝试定义生成代码的配置
type config struct {
	projDir   string // 项目所在目录，即go.mod所在目录
	subDir    string // 子目录
	fileName  string // 文件名
	tmplPath  string
	tmplName  string
	isCover   bool // 是否覆盖
	embedTmpl string
	data      interface{}
}

func GenMod(wd, modName string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.Chdir(wd); err != nil {
		return err
	}
	defer os.Chdir(pwd)

	cmd := exec.Command("go", "mod", "init", modName)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func ModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func genFile(conf config) error {
	fpath := filepath.Join(conf.projDir, conf.subDir, conf.fileName)

	// 是否覆盖原文件
	if !conf.isCover { // 不覆盖
		_, err := os.Stat(fpath)
		if err == nil { // 已存在
			return nil
		}
	}

	// create projDir
	if err := pathx.MkdirIfNotExist(filepath.Join(conf.projDir, conf.subDir)); err != nil {
		return err
	}

	// create file
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 目标文件目录初始化
	if err := os.Mkdir(path.Join(conf.projDir, conf.subDir), os.ModePerm); err != nil && !os.IsExist(err) {
		return err
	}

	t := template.Must(template.New(conf.tmplName).Parse(conf.embedTmpl))

	// 数据写入模板
	var buffer bytes.Buffer
	if err := t.Execute(&buffer, conf.data); err != nil {
		return err
	}

	// 模板写入文件
	if _, err = file.WriteString(buffer.String()); err != nil {
		return err
	}

	return nil
}
