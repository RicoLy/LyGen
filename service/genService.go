package service

import (
	"LyGen/asset"
	"LyGen/constant"
	"LyGen/tools"
	"LyGen/types"
	"bytes"
	"fmt"
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

	tplByte, err := asset.Asset(temp)
	if err != nil {
		return
	}
	tpl, err := template.New("tpl").Parse(string(tplByte))
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
