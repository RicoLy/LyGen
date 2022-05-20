package user

import (
	"FiberDemo/cmd/internal/logic/user"
	"FiberDemo/cmd/internal/response"
	"github.com/gofiber/fiber/v2"
)


// 获取验证码
func CaptchaHandler(c *fiber.Ctx) (err error) {

	rsp, err := user.CaptchaLogic(c)
	if err != nil {
		return response.Response(c, nil, err)
	}

	return response.Response(c, rsp, nil)
}