package middleware

import (
	"{{.}}/cmd/internal/constant"
	jwt2 "{{.ProjectName}}/cmd/internal/tools/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"time"
)

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get(constant.HeaderJwtAuthKey)
		if token == "" {
			return errors.New(constant.ErrNoToken)
		}
		customClaims, err := jwt2.Jwt.ParseToken(token)
		if err != nil {
			return errors.New(constant.ErrTokenNotValidYet)
		}
		if time.Now().Unix() > customClaims.ExpiresAt {
			return errors.New(constant.ErrTokenExpired)
		}

		err = c.Next()

		// 达到缓冲时间自动更新token
		if time.Now().Unix()+customClaims.BufferTime > customClaims.ExpiresAt {
			token, err = jwt2.Jwt.CreateToken(customClaims)
			if err != nil {
				return errors.Wrap(err, constant.ErrCreateTokenErr)
			}
		}
		c.Set(constant.HeaderJwtAuthKey, token)
		return err
	}
}
