package middleware

import (
	"{{.Mata}}/internal/tools/logger"
	"github.com/gofiber/fiber/v2"
)

func {{.Name}}Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger.GetLogger().Info("start {{.Name}}Middleware")

		err := c.Next()

		logger.GetLogger().Info("End {{.Name}}Middleware")
		return err
	}
}