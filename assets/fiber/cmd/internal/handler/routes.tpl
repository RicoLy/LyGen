package handler

import (
	"{{.ProjectName}}/cmd/internal/handler/order"
	"{{.ProjectName}}/cmd/internal/handler/user"
	"{{.ProjectName}}/cmd/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandlers(app *fiber.App) {
	userGroup := app.Group("/user")
	{
		userGroup.Get("/captcha", middleware.CasBinMiddleware(), user.CaptchaHandler)
		userGroup.Post("/login", user.LoginHandler)
	}
	{
		userGroup.Use(middleware.JwtMiddleware())
		userGroup.Get("/detail/:id", user.DetailHandler)
		userGroup.Post("/add", user.AddHandler)
	}

	orderGroup := app.Group("/order")
	{
		userGroup.Use(middleware.JwtMiddleware())
		orderGroup.Get("/list", order.ListHandler)
	}
}
