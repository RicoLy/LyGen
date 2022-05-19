package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//连接数据库
func InitMysqlDB(dataSource string) (*sql.DB, error) {

	connection, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	if err = connection.Ping(); err != nil {
		return nil, err
	}
	connection.SetMaxIdleConns(5)
	connection.SetMaxOpenConns(5)
	return connection, nil
}