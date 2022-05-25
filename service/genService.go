package service

import (
	"LyGen/asset"
	"LyGen/constant"
	"LyGen/tools"
	"LyGen/types"
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"
)

type Generator struct {
}

var Gen = new(Generator)

// GenerateMarkdown 生成表列表
func (g *Generator) GenerateMarkdown(data *types.MarkDownData) (err error) {
	// 写入markdown
	file := constant.CustomDir + fmt.Sprintf("markdown%s.md", time.Now().Format("2006-01-02_150405"))

	return g.GenerateFiles(constant.TplMarkdown, data, file)
}

// 根据模板生成文件
func (g *Generator) GenerateFiles(temp string, data interface{}, path string) (err error) {
	if tools.IsDirOrFileExist(path) {
		return
	}
	tplByte, err := asset.Asset(temp)
	if err != nil {
		return
	}
	tpl, err := template.New("tpl").Funcs(map[string]interface{}{
		"LeftUpper": func(s string) string {
			if len(s) > 0{
				return strings.ToUpper(string(s[0])) + s[1:]
			}
			return s
		},
		"LeftLower": func(s string) string {
			if len(s) > 0 {
				return strings.ToLower(string(s[0])) + s[1:]
			}
			return s
		},
	}).Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, data)
	if err != nil {
		return
	}
	// 表信息写入文件
	_, err = tools.WriteFile(path, content.String())
	if err != nil {
		return
	}
	return
}

// 生成fiber固定文件模板
func (g Generator) GenerateFixedFiles() (err error) {
	fmt.Println("GenerateFixedFiles|constant.Project", constant.Project)
	for _, tpl := range constant.TplFixedList {
		pos := strings.ReplaceAll(tpl, constant.TplFiberPrefix, constant.CustomDir)
		dir, _ := tools.SeparateByLastStr(pos, "/")
		if tpl != constant.TplFiberGoMod {
			_ = tools.CreateDir(dir)
		}
		if err = g.GenerateFiles(tpl, constant.Project, tools.FindTopStr(pos, ".tpl")); err != nil {
			return
		}
	}
	return
}
