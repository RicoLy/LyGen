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

//
func DetailHandler(c *fiber.Ctx) (err error) {
	req := new(types.DetailReq)
	req.Id = c.Params("id")
	if err = validate.Validator.Struct(req); err != nil {
		return response.Response(
			c, nil,
			errors.Wrap(err, constant.ErrParam),
		)
	}

	rsp, err := user.DetailLogic(c, req)
	if err != nil {
		return response.Response(c, nil, err)
	}
	return response.Response(c, rsp, nil)
}