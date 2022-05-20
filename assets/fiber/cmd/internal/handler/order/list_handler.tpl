package order

import (
	"FiberDemo/cmd/internal/constant"
	"FiberDemo/cmd/internal/logic/order"
	"FiberDemo/cmd/internal/response"
	"FiberDemo/cmd/internal/tools/validate"
	"FiberDemo/cmd/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func ListHandler(c *fiber.Ctx) (err error) {
	req := new(types.ListOrderReq)
	if err = c.QueryParser(req); err != nil {
		return response.Response(
			c, nil,
			errors.Wrap(err, constant.ErrParserBody),
		)
	}
	if err = validate.Validator.Struct(req); err != nil {
		return response.Response(
			c, nil,
			errors.Wrap(err, constant.ErrParam),
		)
	}

	rsp, err := order.ListLogic(c, req)
	if err != nil {
		return response.Response(c, nil, err)
	}
	return response.Response(c, rsp, nil)
}
