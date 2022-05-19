package tools

import (
	"errors"
	"os"
	"strings"
)

func GenerateDir(path string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("directory is null")
	}
	last := path[len(path)-1:]
	if !strings.EqualFold(last, string(os.PathSeparator)) {
		path = path + string(os.PathSeparator)
	}
	if !IsDir(path) {
		if CreateDir(path) {
			return path, nil
		}
		return "", errors.New(path + "Failed to create or insufficient permissions")
	}
	return path, nil
}
