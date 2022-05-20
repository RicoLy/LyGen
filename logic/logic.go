package logic

import (
	"LyGen/assets"
	"LyGen/constant"
	"LyGen/generater"
	"LyGen/service"
	"LyGen/tools"
	"LyGen/types"
	"bytes"
	"fmt"
	"log"
	"text/template"
)

var Lg = new(Logic)

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

func (l *Logic) CreateEntity(formatList []string) (err error) {
	// 表结构文件路径
	path := tools.CreateDirs(constant.CustomDir+constant.GoDirEntity) + constant.DS + constant.GoFileEntity

	// 将表结构写入文件
	tables, err := service.DB.FindDbTables()
	if err != nil {
		return
	}
	// 将表结构写入文件
	for idx, table := range tables {
		idx++
		// 查询表结构信息
		tableDesc, err := service.DB.GetTableDesc(table.Name)
		if err != nil {
			fmt.Println("CreateEntityErr>>", table.Name)
			continue
		}
		req := new(types.EntityReq)
		req.Index = idx
		req.EntityPath = path
		req.TableName = table.Name
		req.TableComment = table.Comment
		req.TableDesc = tableDesc
		req.FormatList = formatList
		req.EntityPkg = constant.PkgEntity
		// 生成基础信息
		err = service.DB.GenerateDBEntity(req)
		if err != nil {
			fmt.Println("GenerateDBEntity>>", table.Name)
			continue
		}
	}
	return
}

func (l *Logic) CreateCURD(formatList []string) (err error) {
	tableNameList := make([]*types.TableList, 0)
	// 表结构文件路径
	// 将表结构写入文件
	tables, err := service.DB.FindDbTables()
	for _, table := range tables {
		tableNameList = append(tableNameList, &types.TableList{
			UpperTableName: constant.TablePrefix + tools.ToUpper(table.Name),
			TableName:      table.Name,
			Comment:        table.Comment,
		})
		// 查询表结构信息
		tableDesc, err := service.DB.GetTableDesc(table.Name)
		if err != nil {
			return err
		}
		req := new(types.EntityReq)
		req.TableName = table.Name
		req.TableComment = table.Comment
		req.TableDesc = tableDesc
		req.EntityPath = l.GetEntityDir() + constant.GoFileEntity
		req.FormatList = formatList
		req.Pkg = constant.PkgDbModels
		req.EntityPkg = constant.PkgEntity
		// 生成基础信息
		err = service.DB.GenerateDBEntity(req)
		if err != nil {
			return err
		}
		// 生成增,删,改,查文件
		err = service.DB.GenerateCURDFile(table.Name, table.Comment, tableDesc)
		if err != nil {
			return err
		}
	}

	// 生成所有表的文件
	if err = l.GenerateTableList(tableNameList); err != nil {
		return
	}
	//生成init文件
	if err = l.GenerateInit(); err != nil {
		return
	}
	//生成error文件
	if err = l.GenerateError(); err != nil {
		return
	}
	fmt.Println("`CURD` files created finish!")
	return nil
}

// GenerateError 生成error 文件
func (l *Logic) GenerateError() (err error) {
	file := l.GetMysqlDir() + constant.GoFileError
	// 判断package是否加载过
	checkStr := "package " + constant.PkgDbModels
	if tools.CheckFileContainsChar(file, checkStr) == false {
		if _, err = tools.WriteFile(file, checkStr+"\n"); err != nil {
			return err
		}
	}
	// 判断是否已经生成过此文件
	checkStr = "Stack"
	if tools.CheckFileContainsChar(file, checkStr) {
		log.Println(file + "It already exists. Please delete it and regenerate it")
		return
	}
	tplByte, err := assets.Asset(constant.TplError)
	if err != nil {
		return
	}
	tpl, err := template.New("error").Parse(string(tplByte))
	if err != nil {
		return
	}
	// analysis execute template
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, nil)
	if err != nil {
		return
	}
	// append write to file
	err = tools.WriteAppendFile(file, content.String())
	if err != nil {
		return
	}
	return
}

// 生成init
func (l *Logic) GenerateInit() (err error) {
	file := l.GetMysqlDir() + constant.GoFileInit
	// 判断package是否加载过
	checkStr := "package " + constant.PkgDbModels
	if tools.CheckFileContainsChar(file, checkStr) == false {
		if _, err = tools.WriteFile(file, checkStr+"\n"); err != nil {
			return err
		}
	}
	checkStr = "DBConfig"
	if tools.CheckFileContainsChar(file, checkStr) {
		log.Println(file + "It already exists. Please delete it and regenerate it")
		return
	}
	tplByte, err := assets.Asset(constant.TplInit)
	if err != nil {
		return
	}
	tpl, err := template.New("init").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})

	if err = tpl.Execute(content, nil); err != nil {
		return
	}
	// 表信息写入文件
	return tools.WriteAppendFile(file, content.String())
}

// 生成表列表
func (l *Logic) GenerateTableList(list []*types.TableList) (err error) {
	// 写入表名
	file := l.GetConfigDir() + constant.GoFileTableList
	// 判断package是否加载过
	checkStr := "package " + constant.PkgTable
	if tools.CheckFileContainsChar(file, checkStr) == false {
		if _, err = tools.WriteFile(file, checkStr+"\n"); err != nil {
			return
		}
	}
	checkStr = "const"
	if tools.CheckFileContainsChar(file, checkStr) {
		log.Println(file + "It already exists. Please delete it and regenerate it")
		return
	}
	tplByte, err := assets.Asset(constant.TplTables)
	if err != nil {
		return
	}
	tpl, err := template.New("table_list").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	if err = tpl.Execute(content, list); err != nil {
		return
	}
	// 表信息写入文件
	if err = tools.WriteAppendFile(file, content.String()); err != nil {
		return
	}
	return
}

func (l *Logic) GetEntityDir() string {
	return tools.CreateDirs(constant.CustomDir + constant.GoDirModels + constant.DS + constant.
		GoDirEntity + constant.DS)
}

// 创建和获取MYSQL目录
func (l *Logic) GetConfigDir() string {
	return tools.CreateDirs(constant.CustomDir + constant.GoDirModels + constant.DS + constant.GoDirConfig + constant.DS)
}

// GetMysqlDir 创建和获取MYSQL目录
func (l *Logic) GetMysqlDir() string {
	return tools.CreateDirs(constant.CustomDir + constant.GoDirModels + constant.DS)
}