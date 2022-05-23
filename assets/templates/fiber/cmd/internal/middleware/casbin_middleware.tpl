package middleware

import (
	"{{.ProjectName}}/cmd/internal/tools/logger"
	"github.com/gofiber/fiber/v2"
)

func CasBinMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger.GetLogger().Info("start CasBinMiddleware")

		err := c.Next()

		logger.GetLogger().Info("End CasBinMiddleware")
		return err
	}
}