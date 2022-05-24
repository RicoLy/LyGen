package {{.Group}}

import (
	{{if ne .Request "CommReq"}}"{{.Mata}}/cmd/internal/constant"
    "{{.Mata}}/cmd/internal/tools/validate"
    "{{.Mata}}/cmd/internal/types"
    "github.com/pkg/errors"{{end}}
	"{{.Mata}}/cmd/internal/logic/{{.Group}}"
	"{{.Mata}}/cmd/internal/response"
	"github.com/gofiber/fiber/v2"
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
	{{if ne .Response "CommRsp"}}rsp, err := {{else}}err ={{end}}{{.Group}}.{{.Name}}Logic(c{{if ne .Request "CommReq"}}, req{{end}})
	{{if ne .Response "CommRsp"}}if err != nil {
        return response.Response(c, nil, err)
    }
    return response.Response(c, rsp, nil){{else}}
    return response.Response(c, nil, err)
    {{end}}
}