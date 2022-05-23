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

func AddHandler(c *fiber.Ctx) (err error) {
	req := new(types.AddUserReq)
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

	err = user.AddLogic(c, req)
	return response.Response(c, nil, err)
}