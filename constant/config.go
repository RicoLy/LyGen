package constant

import (
	"database/sql"
	"os"
)

const (
	ProjectName = "LyGen"
	Author      = "RicoLy"
	AuthorEmail = "rico_ly@163.com"
	Version     = "0.0.1"
	Copyright   = "2022.5.20"
)

const (
	DS              = string(os.PathSeparator) // 通用/
	DbNullPrefix    = "Null"                   // 处理数据为空时结构的前缀定义
	TablePrefix     = "TABLE_"                 // 表前缀
	DefaultSavePath = "output"                 // 默认生成目录名称
)

// generate file name
const (
	GoDirModels     = "db_models"       // model dir
	GoDirConfig     = "config"          // config dir
	GoDirEntity     = "entity"          // entity dir
	GoFileEntity    = "db_entity.go"    // entity table file
	GoFileTableList = "table_list.go"   // table file
	GoFileInit      = "init.go"         // init file
	GoFileError     = "e.go"            // error file
	GoFileExample   = "example_test.go" // example file
)

const (
	PkgDbModels = "mysql"  // db_models package name
	PkgEntity   = "entity" // entity package name
	PkgTable    = "config" // table package name
)

const (
	Unknown = iota
	Darwin
	Window
	Linux
)

var (
	MysqlDataSource = ""
	CustomDir       = ""
	Project         = ""
)

var (
	MysqlConn *sql.DB
)
