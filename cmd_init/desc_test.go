package cmd_init

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGenDesc(t *testing.T) {
	GenDesc()
}

func GenDesc() error {
	wd, _ := os.Getwd()
	descFile := "app/desc.ap"
	res := filepath.Join(wd, descFile)
	fmt.Printf("wd, descFile: %v, %v \n", wd, descFile)
	fmt.Printf("res: %v \n", res)

	return nil
}
