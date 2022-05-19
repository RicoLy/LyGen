package service

import (
	"LyGen/constant"
	"LyGen/db"
	"LyGen/tools"
	"LyGen/types"
	"sort"
	"strings"
)

var DB = new(DbSrv)

type DbSrv struct {
	DBName string
}

// FindDbTables 向数据库里读取所有的表
func (s *DbSrv)FindDbTables() (NameAndComment []*types.TableNameAndComment, err error) {
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

