package response

import (
	"{{.}}/internal/constant"
	"{{.}}/internal/tools/logger"
	"{{.}}/internal/tools/str"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Body struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	TraceId string      `json:"traceId"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c *fiber.Ctx, rsp interface{}, err error) error {
	var body Body
	body.TraceId = c.Locals(constant.CTXTraceId).(string)
	if err != nil {
		logger.GetLogger().Error("error:",
			zap.Any("traceId:", body.TraceId),
			zap.Any("err:", err),
		)
		body.Code = -1
		body.Msg = str.FindTopStr(err.Error(), ":")
	} else {
		logger.GetLogger().Info("success:",
			zap.Any("traceId:", body.TraceId),
			zap.Any("rsp:", rsp),
		)
		body.Msg = "OK"
		body.Data = rsp
	}
	return c.JSON(body)
}