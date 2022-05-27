package {{.Group}}

import (
	"{{.Mata}}/internal/types"
	"github.com/gofiber/fiber/v2"
)

// {{.Name}}Logic {{.Comment}}
func {{.Name}}Logic(c *fiber.Ctx{{if ne .Request "CommReq"}}, req *types.{{.Request}}{{end}}) ({{if ne .Response "CommRsp"}}rsp *types.{{.Response}}, {{end}}err error) {
	// todo write logic

	return
}