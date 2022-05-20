package tools

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

//添加注释 //
func AddToComment(s string, suf string) string {
	if strings.EqualFold(s, "") {
		return ""
	}
	return "// " + s + suf
}

// FormatField 拼接tag
func FormatField(field string, formats []string) string {
	if len(formats) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	for key := range formats {
		buf.WriteString(fmt.Sprintf(`%s:"%s" `, formats[key], field))
	}
	return "`" + strings.TrimRight(buf.String(), " ") + "`"
}

// Capitalize 字符首字母大写,将字符串转换成驼峰格式
func Capitalize(s string) string {
	var upperStr string
	chars := strings.Split(s, "_")
	for _, val := range chars {
		vv := []rune(val) // 后文有介绍
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
					vv[i] -= 32 // string的码表相差32位
					upperStr += string(vv[i])
				}
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

// ToUpper 将分隔_拆掉,全大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// CheckCharDoSpecialArr 检查字符串,去掉特殊字符
func CheckCharDoSpecialArr(s string, char byte, reg string) []string {
	s = CheckCharDoSpecial(s, char, reg)
	return strings.Split(s, string(char))
}

// CheckCharDoSpecial 检查字符串,去掉特殊字符
func CheckCharDoSpecial(s string, char byte, regs string) string {
	reg := regexp.MustCompile(regs)
	var result []string
	if arr := reg.FindAllString(s, -1); len(arr) > 0 {
		buf := bytes.Buffer{}
		for key, val := range arr {
			if val != string(char) {
				buf.WriteString(val)
			}
			if val == string(char) && buf.Len() > 0 {
				result = append(result, buf.String())
				buf.Reset()
			}
			//处理最后一批数据
			if buf.Len() > 0 && key == len(arr)-1 {
				result = append(result, buf.String())
			}
		}
	}
	return strings.Join(result, string(char))
}

// CheckFileContainsChar 检查某字符是否存在文件里
func CheckFileContainsChar(filename, s string) bool {
	data := ReadFile(filename)
	if len(data) > 0 {
		return strings.LastIndex(data, s) > 0
	}
	return false
}

//读取文件内容
func ReadFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

// WriteFile 写文件
func WriteFile(filename string, data string) (count int, err error) {
	var f *os.File
	if IsDirOrFileExist(filename) == false {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	}
	defer f.Close()
	count, err = io.WriteString(f, data)
	if err != nil {
		return
	}
	return
}

// AddQuote 添加``符号
func AddQuote(str string) string {
	return "`" + str + "`"
}