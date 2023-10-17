package pathx

import (
	"fmt"
	"testing"
)

func TestGetModName(t *testing.T) {
	name, err := GetModName()
	if err != nil {
		fmt.Printf("err: %v \n", err)
		return
	}

	fmt.Printf("mod name: %v \n", name)
}

func TestGetRootDir(t *testing.T) {
	root, err := GetProjDir()
	if err != nil {
		return
	}
	fmt.Printf("root: %v \n", root)
}
