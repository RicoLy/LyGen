package tools

import (
	"LyGen/constant"
	"fmt"
	"strings"
	"testing"
)

func TestFindLastStr(t *testing.T) {
	custom := "F:\\code\\goProject\\LyGen"
	projectName := FindLastStr(custom, constant.DS)
	fmt.Println("projectName", projectName)
	for _, tpl := range constant.TplFixedList {
		pos := strings.ReplaceAll(tpl, constant.TplFiberPrefix, custom)
		path, fileName := SeparateByLastStr(pos, "/")
		ok := CreateDir(path)
		fmt.Println("ok:", ok)
		err := WriteAppendFile(pos, "hello")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("pos", fileName)
	}

}