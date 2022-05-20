package service

import (
	"LyGen/assets"
	"LyGen/constant"
	"LyGen/db"
	"LyGen/tools"
	"LyGen/types"
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
	"text/template"
)

var DB = new(DbSrv)

type DbSrv struct {
	DBName string
	Once sync.Once
}

// FindDbTables 向数据库里读取所有的表
func (s *DbSrv) FindDbTables() (NameAndComment []*types.TableNameAndComment, err error) {
	result, err := db.MysqlRepo.Find("SELECT `TABLE_NAME` AS 'table_name', `TABLE_COMMENT` AS 'table_comment' FROM "+
		"information_schema.tables WHERE `TABLE_SCHEMA` = ?", s.DBName)
	if err != nil {
		return
	}
	//获取表名 与 表注释
	NameAndComment = make([]*types.TableNameAndComment, 0)
	//获取库里所有的表名
	for idx, info := range result {
		idx++
		NameAndComment = append(NameAndComment, &types.TableNameAndComment{
			Index:   idx,
			Name:    info["table_name"].(string),
			Comment: info["table_comment"].(string),
		})
	}
	//排序, 采用升序
	sort.Slice(NameAndComment, func(i, j int) bool {
		return strings.ToLower(NameAndComment[i].Name) < strings.ToLower(NameAndComment[j].Name)
	})

	return
}

// GetTableDesc 获取表结构详情
func (s *DbSrv) GetTableDesc(tableName string) (reply []*types.TableDesc, err error) {
	result, err := db.MysqlRepo.Find("select `COLUMN_NAME` AS column_name,`DATA_TYPE` AS data_type, `COLUMN_KEY` AS column_key, "+
		"`IS_NULLABLE` AS is_nullable, `COLUMN_DEFAULT` AS column_default,`COLUMN_TYPE` AS column_type, `COLUMN_COMMENT` "+
		"AS column_comment from information_schema.columns where table_name = ? and table_schema = ? ORDER BY ordinal_position",
		tableName, s.DBName)
	if err != nil {
		return
	}
	reply = make([]*types.TableDesc, 0)
	i := 0
	for _, row := range result {
		var keyBool bool
		if strings.ToUpper(row["column_key"].(string)) == "PRI" {
			keyBool = true
		}
		oriType := row["data_type"].(string)
		var columnDefault string
		val, ok := row["column_default"].(string)
		if ok {
			columnDefault = val
		}
		reply = append(reply, &types.TableDesc{
			Index:            i,
			ColumnName:       row["column_name"].(string),
			GoColumnName:     tools.Capitalize(row["column_name"].(string)),
			OriMysqlType:     oriType,
			UpperMysqlType:   strings.ToUpper(oriType),
			GolangType:       constant.MysqlTypeToGoType[oriType],
			MysqlNullType:    constant.MysqlTypeToGoNullType[oriType],
			ColumnComment:    row["column_comment"].(string),
			IsNull:           row["is_nullable"].(string),
			DefaultValue:     columnDefault,
			ColumnTypeNumber: row["column_type"].(string),
			PrimaryKey:       keyBool,
		})
		i++
	}
	return
}

func (s *DbSrv) GenerateDBEntity(req *types.EntityReq) (err error) {
	var str string
	str = fmt.Sprintf(`// 判断package是否加载过
								package %s
								import (
									"database/sql"
									"github.com/go-sql-driver/mysql"
									"time"
								)
								`, req.EntityPkg)
	// 判断import是否加载过
	check := "github.com/go-sql-driver/mysql"
	if tools.CheckFileContainsChar(req.EntityPath, check) == false {
		_, err := tools.WriteFile(req.EntityPath, str)
		return err
	}
	// 声明表结构变量
	TableData := new(types.TableInfo)
	TableData.Table = tools.Capitalize(req.TableName)
	TableData.TableName = req.TableName
	TableData.NullTable = TableData.Table + constant.DbNullPrefix
	TableData.TableComment = tools.AddToComment(req.TableComment, "")
	TableData.TableCommentNull = tools.AddToComment(req.TableComment, " Null Entity")
	// 判断表结构是否加载过
	if tools.CheckFileContainsChar(req.EntityPath, "type "+TableData.Table+" struct") == true {
		log.Println(req.EntityPath + "It already exists. Please delete it and regenerate it")
		return
	}
	// 加载模板文件
	tplByte, err := assets.Asset(constant.TplEntity)
	if err != nil {
		return
	}
	tpl, err := template.New("entity").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 装载表字段信息
	for _, val := range req.TableDesc {
		TableData.Fields = append(TableData.Fields, &types.FieldsInfo{
			Name:         tools.Capitalize(val.ColumnName),
			Type:         val.GolangType,
			NullType:     val.MysqlNullType,
			DbOriField:   val.ColumnName,
			FormatFields: tools.FormatField(val.ColumnName, req.FormatList),
			Remark:       tools.AddToComment(val.ColumnComment, ""),
		})
	}
	content := bytes.NewBuffer([]byte{})
	if err = tpl.Execute(content, TableData); err != nil {
		return err
	}
	// 表信息写入文件
	con := strings.Replace(content.String(), "&#34;", `"`, -1)
	err = tools.WriteAppendFile(req.EntityPath, con)
	return
}

func (s *DbSrv) GenerateCURDFile(tableName, tableComment string, tableDesc []*types.TableDesc) (err error) {
	var (
		allFields       = make([]string, 0)
		insertFields    = make([]string, 0)
		InsertInfo      = make([]*types.SqlFieldInfo, 0)
		fieldsList      = make([]*types.SqlFieldInfo, 0)
		nullFieldList   = make([]*types.NullSqlFieldInfo, 0)
		updateList      = make([]string, 0)
		updateListField = make([]string, 0)
		PrimaryKey      = ""
		primaryType     = ""
	)
	// 存放第个字段
	var secondField string
	for _, item := range tableDesc {
		allFields = append(allFields, tools.AddQuote(item.ColumnName))
		if item.PrimaryKey == false && item.ColumnName != "updated_at" && item.ColumnName != "created_at" {
			insertFields = append(insertFields, tools.AddQuote(item.ColumnName))
			InsertInfo = append(InsertInfo, &types.SqlFieldInfo{
				HumpName: tools.Capitalize(item.ColumnName),
				Comment:  item.ColumnComment,
			})
			if item.ColumnName == "identify" {
				updateList = append(updateList, tools.AddQuote(item.ColumnName)+"="+item.ColumnName+"+1")
			} else {
				updateList = append(updateList, tools.AddQuote(item.ColumnName)+"=?")
				if item.PrimaryKey == false {
					updateListField = append(updateListField, "value."+tools.Capitalize(item.ColumnName))
				}
			}
		}
		if item.PrimaryKey {
			PrimaryKey = item.ColumnName
			primaryType = item.GolangType
		} else {
			// 除了主键外的任意一个字段即可。
			if secondField == "" {
				secondField = item.ColumnName
			}
		}
		fieldsList = append(fieldsList, &types.SqlFieldInfo{
			HumpName: tools.Capitalize(item.ColumnName),
			Comment:  item.ColumnComment,
		})
		nullFieldList = append(nullFieldList, &types.NullSqlFieldInfo{
			HumpName:     tools.Capitalize(item.ColumnName),
			OriFieldType: item.OriMysqlType,
			GoType:       constant.MysqlTypeToGoType[item.OriMysqlType],
			Comment:      item.ColumnComment,
		})
	}
	// 主键ID,用于更新
	if PrimaryKey != "" {
		updateListField = append(updateListField, "value."+tools.Capitalize(PrimaryKey))
	}
	// 拼出SQL所需要结构数据
	InsertMark := strings.Repeat("?,", len(insertFields))
	if len(InsertMark) > 0 {
		InsertMark = InsertMark[:len(InsertMark)-1]
	}
	sqlInfo := &types.SqlInfo{
		TableName:           tableName,
		PrimaryKey:          tools.AddQuote(PrimaryKey),
		PrimaryType:         primaryType,
		StructTableName:     tools.Capitalize(tableName),
		NullStructTableName: tools.Capitalize(tableName) + constant.DbNullPrefix,
		PkgEntity:           constant.PkgEntity + ".",
		PkgTable:            constant.PkgTable + ".",
		UpperTableName:      constant.TablePrefix + tools.ToUpper(tableName),
		AllFieldList:        strings.Join(allFields, ","),
		InsertFieldList:     strings.Join(insertFields, ","),
		InsertMark:          InsertMark,
		UpdateFieldList:     strings.Join(updateList, ","),
		UpdateListField:     updateListField,
		FieldsInfo:          fieldsList,
		NullFieldsInfo:      nullFieldList,
		InsertInfo:          InsertInfo,
		SecondField:         tools.AddQuote(secondField),
	}
	if err = s.GenerateSQL(sqlInfo, tableComment); err != nil {
		return
	}
	// 添加一个实例
	s.Once.Do(func() {
		s.GenerateExample(sqlInfo.StructTableName)
	})

	return
}

// 生成一个实例文件
func (s *DbSrv) GenerateExample(name string) {
	// 写入表名
	file := tools.CreateDirs(constant.CustomDir + constant.GoDirModels + constant.DS) + constant.GoFileExample

	// 解析模板
	tplByte, err := assets.Asset(constant.TplExample)
	if err != nil {
		return
	}
	tpl, err := template.New("example").Parse(string(tplByte))
	if err != nil {
		return
	}
	type st struct {
		Name string
	}
	ss := st{
		Name: name,
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, ss)
	if err != nil {
		return
	}
	// 表信息写入文件
	_, err = tools.WriteFile(file, content.String())
	if err != nil {
		return
	}
	return
}

// GenerateSQL 生成SQL文件
func (s *DbSrv) GenerateSQL(info *types.SqlInfo, tableComment string) (err error) {
	// 写入表名
	goFile := tools.CreateDirs(constant.CustomDir + constant.GoDirModels + constant.DS) + info.TableName + ".go"
	str := fmt.Sprintf(`
								// %s
								package %s
								import(
									"database/sql"
									"strings"
									_ "github.com/go-sql-driver/mysql"
								)
								`, tableComment, constant.PkgDbModels)
	// 判断package是否加载过
	if tools.CheckFileContainsChar(goFile, "database/sql") == false {

		if _, err = tools.WriteFile(goFile, str); err != nil {
			return
		}
	}

	// 解析模板
	tplByte, err := assets.Asset(constant.TplCurd)
	if err != nil {
		return
	}
	tpl, err := template.New("CURD").Parse(string(tplByte))
	if err != nil {
		return
	}
	// 解析
	content := bytes.NewBuffer([]byte{})
	err = tpl.Execute(content, info)
	if err != nil {
		return
	}
	// 表信息写入文件
	if tools.CheckFileContainsChar(goFile, info.StructTableName) == false {
		err = tools.WriteAppendFile(goFile, content.String())
		if err != nil {
			return
		}
	}
	return
}
