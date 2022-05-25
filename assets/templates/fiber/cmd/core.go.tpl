package main

import (
	"{{.}}/cmd/global"
	"{{.}}/cmd/internal/config"
	"{{.}}/cmd/internal/handler"
	"{{.}}/cmd/internal/middleware"
	"{{.}}/cmd/internal/tools/logger"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var configPath = flag.String("f", "etc/", "the config Path")

func main() {
	flag.Parse()
	if err := config.GlobalConfig.InitConfig(*configPath); err != nil {
		panic(err)
	}
	if err := global.Init(config.GlobalConfig); err != nil {
		panic(err)
	}
	if err := logger.InitLoggerServer(
		config.GlobalConfig.Logger.Name,
		config.GlobalConfig.Logger.IsDev,
		config.GlobalConfig.Logger.Level,
	); err != nil {
		panic(err)
	}
	
	app := fiber.New()
	app.Use(middleware.WrapCtxMiddleware())
	handler.RegisterHandlers(app)

	_ = app.Listen(fmt.Sprintf("%s:%s", config.GlobalConfig.Host, config.GlobalConfig.Port))
}
