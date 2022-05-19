package constant

import "database/sql"

const (
	ProjectName = "LyGen"
	Author      = "RicoLy"
	AuthorEmail = "rico_ly@163.com"
	Version     = "0.0.1"
	Copyright   = "2022.5.20"
)

var (
	MysqlDataSource = ""
	CustomDir       = ""
)

var (
	MysqlConn *sql.DB
)