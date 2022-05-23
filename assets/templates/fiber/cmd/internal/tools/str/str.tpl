package str

import "strings"

func FindTopStr(str string, sep string) string {
	index := strings.Index(str, sep)
	if index == -1 {
		return str
	} else {
		return str[:index]
	}
}
