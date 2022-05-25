package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GlobalConfig = &Config{
	Viper: viper.New(),
}

type Config struct {
	Viper  *viper.Viper
	Name string
    Host string
    Port string
	Logger struct {
		Name  string
		IsDev bool
		Level string
	}
	Captcha struct {
		KeyLong   int
		ImgWidth  int
		ImgHeight int
	}
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr string
	}
	JWT struct {
		SigningKey  string // jwt签名
		ExpiresTime int64  // 过期时间
		BufferTime  int64  // 缓冲时间
	}
	Cypher struct {
		AesKey string // aes 秘钥
	}
}

func NewConfig() *Config {
	return &Config{
		Viper: viper.New(),
	}
}

func (c *Config) InitConfig(configPaths ...string) (err error) {
	c.Viper.SetConfigName("config")
	for _, configPath := range configPaths {
		c.Viper.AddConfigPath(configPath)
	}
	c.Viper.SetConfigType("yaml")
	if err = c.Viper.ReadInConfig(); err != nil {
		return
	}
	c.Viper.WatchConfig()

	c.Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := c.Viper.Unmarshal(&c); err != nil {
			fmt.Println(err)
		}
	})
	return c.Viper.Unmarshal(&c)
}
