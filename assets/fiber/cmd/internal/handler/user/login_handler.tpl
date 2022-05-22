package user

import (
	"{{.ProjectName}}/cmd/internal/constant"
	"{{.ProjectName}}/cmd/internal/logic/user"
	"{{.ProjectName}}/cmd/internal/response"
	"{{.ProjectName}}/cmd/internal/tools/validate"
	"{{.ProjectName}}/cmd/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// 登录
func LoginHandler(c *fiber.Ctx) (err error) {
	req := new(types.LoginReq)
	if err = c.BodyParser(req); err != nil {
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

	rsp, err := user.LoginLogic(c, req)
	if err != nil {
		return response.Response(c, nil, err)
	}

	return response.Response(c, rsp, nil)
}
