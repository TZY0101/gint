package test_proj

import (
	"fmt"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	_, err := os.ReadFile("api.desc")
	if err != nil {
		fmt.Printf("err: %v \n", err)
	}
}

func TestPath(t *testing.T) {

}

func TestGoInitMode(t *testing.T) {

}
