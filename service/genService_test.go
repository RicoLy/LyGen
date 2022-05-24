package service

import (
	"fmt"
	"testing"
)

func TestGenerateFixedFiles(t *testing.T) {
	err := Gen.GenerateFixedFiles()
	if err != nil {
		fmt.Println(err)
	}
}
