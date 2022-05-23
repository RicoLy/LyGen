package user

import (
	"{{.Mata}}/cmd/internal/constant"
	"{{.Mata}}/cmd/internal/logic/user"
	"{{.Mata}}/cmd/internal/response"
	"{{.Mata}}/cmd/internal/tools/validate"
	"{{.Mata}}/cmd/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// {{.Name}}Handler {{.Comment}}
func {{.Name}}Handler(c *fiber.Ctx) (err error) {
    {{if ne .Request "CommReq"}}
    req := new(types.{{.Request}})
    {{if len .PathParams}}{{range .PathParams}}req.{{.|Title}} = c.Params("{{.}}")
    {{end}}{{else if ne .MethodType "post"}}
    if err = c.QueryParser(req); err != nil {
        return response.Response(
            c, nil,
            errors.Wrap(err, constant.ErrParserBody),
        )
    }{{else}}
    if err = c.BodyParser(req); err != nil {
        return response.Response(
            c, nil,
            errors.Wrap(err, constant.ErrParserBody),
        )
    }
    {{end}}
    if err = validate.Validator.Struct(req); err != nil {
        return response.Response(
            c, nil,
            errors.Wrap(err, constant.ErrParam),
        )
    }
    {{end}}
	{{if ne .Response "CommRsp"}}rsp, {{end}}err = {{.Group}}.{{.Name}}Logic(c{{if ne .Request "CommReq"}}, req{{end}})
	{{if ne .Response "CommRsp"}}if err != nil {
        return response.Response(c, nil, err)
    }
    return response.Response(c, rsp, nil){{else}}
    return response.Response(c, nil, err)
    {{end}}
}