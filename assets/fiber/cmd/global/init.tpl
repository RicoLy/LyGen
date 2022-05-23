package global

import "{{.}}/cmd/internal/config"

func Init(c *config.Config) (err error) {
	// InitRedis(c)
	if err = InitMysqlDb(c); err != nil {
		return err
	}
	return
}
