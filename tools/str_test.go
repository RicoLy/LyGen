package tools

import (
	"LyGen/constant"
	"fmt"
	"strings"
	"testing"
)

func TestFindLastStr(t *testing.T) {

	pos := strings.ReplaceAll(constant.TplFiberInHandler, constant.TplFiberPrefix, "F:\\code\\goProject\\LyGen")
	fmt.Println("pos", pos)
	dir, fileName := SeparateByLastStr(pos, "/")
	fmt.Println("dir", dir)
	fmt.Println("fileName", fileName)
	dir = dir + "/" + "user"
	fileName = "add" + "_" + fileName
	pos = dir + "/" + fileName
	fmt.Println("pos", pos)
}

func TestCheckExist(t *testing.T) {
	fmt.Println(IsDirOrFileExist("F:\\code\\goProject\\LyGen/go.mod"))
}

func TestFindLastStr2(t *testing.T) {
	pos := strings.ReplaceAll(constant.TplFiberGoMod, constant.TplFiberPrefix, "F:\\code\\goProject\\LyGen\\Demo/abc")
	dir, _ := SeparateByLastStr(pos, constant.DS)
	fmt.Println("dir", dir)
}