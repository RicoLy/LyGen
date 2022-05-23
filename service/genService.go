package service

import (
	"LyGen/assets"
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
	tplByte, err := assets.Asset(constant.TplMarkdown)
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	tpl, err := template.New("markdown").Parse(string(tplByte))
	if err != nil {
		return err
	}
	err = tpl.Execute(content, data)
	if err != nil {
		return
	}
	// 表信息写入文件
	_, err = tools.WriteFileAppend(file, content.String())
	if err != nil {
		return
	}
	fmt.Printf("Generate success:%s\n", file)
	return
}
