package middleware

import (
	"{{.}}/cmd/internal/constant"
	"{{.}}/cmd/internal/tools/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

// WrapCtxMiddleware : 请求信息记录中间件
func WrapCtxMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		startTime := time.Now()
		traceId := uuid.NewString()
		c.Locals(constant.CTXTraceId, traceId)
		c.Locals(constant.CTXStartTime, startTime)
		err := c.Next()
		endTime := time.Now()
		logger.GetLogger().Info("success:",
			zap.Any("traceId:", traceId),
			zap.Any("startTime:", startTime.Format("2006-01-02 15:04:05")),
			zap.Any("cost:", endTime.Sub(startTime).Microseconds()),
		)
		return err
	}
}