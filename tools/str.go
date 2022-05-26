package tools

import (
	"bytes"
	"fmt"
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

// AddQuote 添加``符号
func AddQuote(str string) string {
	return "`" + str + "`"
}

// FindTopStr
func FindTopStr(str string, sep string) string {
	index := strings.Index(str, sep)
	if index == -1 {
		return str
	} else {
		return str[:index]
	}
}

// FindLastStr
func FindLastStr(str string, sep string) string {
	index := strings.LastIndex(str, sep)
	if index == -1 {
		return str
	} else {
		return str[index+1:]
	}
}

// SeparateByLastStr
func SeparateByLastStr(str string, sep string) (prefix, suffix string) {
	index := strings.LastIndex(str, sep)
	if index == -1 {
		return str, ""
	} else {
		return str[:index], str[index+1:]
	}
}

// LeftLower
func LeftLower(s string) string {
	if len(s) > 0 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	return s
}