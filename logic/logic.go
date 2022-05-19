package logic

import (
	"LyGen/generater"
	"LyGen/service"
	"LyGen/tools"
	"LyGen/types"
	"fmt"
)

type Logic struct {
}

// CreateMarkdown 生成mysql markdown文档
func (l *Logic) CreateMarkdown() (err error) {
	data := new(types.MarkDownData)
	// 将表结构写入文件
	tables, err := service.DB.FindDbTables()
	if err != nil {
		return
	}
	i := 1
	for _, table := range tables {
		fmt.Println("Doing table:" + table.Name)
		data.TableList = append(data.TableList, &types.TableList{
			Index:          i,
			UpperTableName: tools.ToUpper(table.Name),
			TableName:      table.Name,
			Comment:        table.Comment,
		})
		// 查询表结构信息
		desc := new(types.MarkDownDataChild)
		desc.List, err = service.DB.GetTableDesc(table.Name)
		if err != nil {
			fmt.Println("markdown>>", err)
			continue
		}
		desc.Index = i
		desc.TableName = table.Name
		desc.Comment = table.Comment
		data.DescList = append(data.DescList, desc)
		i++
	}
	// 生成所有表的文件
	err = generater.Gen.GenerateMarkdown(data)
	return
}