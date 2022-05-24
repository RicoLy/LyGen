package global

import (
	"{{.}}/cmd/internal/config"
	"{{.}}/cmd/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)


func InitMysqlDb(c *config.Config) error {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Mysql.DataSource, // DSN data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return err
	}

	err = migration(db)
	if err != nil {
		return err
	}
	MysqlDb = db
	return nil
}

//数据迁移
func migration(db *gorm.DB) error {
	//传入要迁移的模型指针
	return db.AutoMigrate(
		&model.User{},
	)

}
