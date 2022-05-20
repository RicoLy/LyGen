package main

import (
	"FiberDemo/cmd/global"
	"FiberDemo/cmd/internal/config"
	"FiberDemo/cmd/internal/handler"
	"FiberDemo/cmd/internal/middleware"
	"FiberDemo/cmd/internal/tools/logger"
	"flag"
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

	_ = app.Listen(":3000")
}
