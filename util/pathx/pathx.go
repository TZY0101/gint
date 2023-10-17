package pathx

import (
	"golang.org/x/mod/modfile"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// GetProjDir params dir can be absolute path or relative path
func GetProjDir(dir string) (string, error) {
	// 获取工作路径 绝对路径
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	for {
		//	判断当前目录是否为项目根目录，即是否含有go.mod
		if _, err := os.Stat(filepath.Join(abs, "go.mod")); err == nil {
			return abs, nil
		}

		// 到达根目录，停止遍历
		if abs == filepath.Dir(abs) {
			break
		}

		abs = filepath.Dir(abs)
	}

	return "", nil
}

func GetModName(dir string) (string, error) {
	root, err := GetProjDir(dir)
	if err != nil {
		return "", err
	}

	// 读取 go.mod 文件内容

	content, err := os.ReadFile(path.Join(root, "go.mod"))
	if err != nil {
		return "", err
	}

	f, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
		return "", err
	}

	return f.Module.Mod.String(), nil
}

func GetCurPath() string {
	_, filename, _, _ := runtime.Caller(1)
	path := path.Dir(filename)
	return path
}

func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

func IsExist(dir string) bool {
	if len(dir) == 0 {
		return false
	}

	_, err := os.Stat(dir)

	return os.IsExist(err)
}
