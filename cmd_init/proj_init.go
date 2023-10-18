package cmd_init

import (
	_ "embed"
	"fmt"
	"github.com/TZY0101/gint/cmd_gen"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed template/desc.template
var descTmpl string

func InitProjDir(projName string) error {
	// 是否已初始化
	abs, err := filepath.Abs(projName) // Abs 会做一个clean？
	if err != nil {
		return err
	}

	if _, err := os.Stat(abs); err == nil {
		fmt.Printf("project %s already exists, do not initialize repeatedly \n", projName)
		return nil
	}

	// 创建项目文件，为什么用MkdirAll？projName可能包含多层
	if err := os.MkdirAll(abs, os.ModePerm); err != nil {
		return err
	}

	// gen go.mod
	if err := cmd_gen.GenMod(abs, filepath.Base(abs)); err != nil {
		return err
	}

	// gen desc file
	fileName := filepath.Base(abs) + ".desc"
	descDir := filepath.Join(abs, "desc")
	if err := os.Mkdir(descDir, os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(filepath.Join(descDir, fileName))
	if err != nil {
		return err
	}
	t := template.Must(template.New("desc").Parse(descTmpl))
	if err := t.Execute(file, nil); err != nil {
		return err
	}

	// gen proj
	err = cmd_gen.GenProj(abs, filepath.Join(descDir, fileName))

	return err
}
