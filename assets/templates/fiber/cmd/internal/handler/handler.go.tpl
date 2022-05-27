package {{.Group}}

import (
	{{if ne .Request "CommReq"}}"{{.Mata}}/internal/constant"
    "{{.Mata}}/internal/tools/validate"
    "{{.Mata}}/internal/types"
    "github.com/pkg/errors"{{end}}
	"{{.Mata}}/internal/logic/{{.Group}}"
	"{{.Mata}}/internal/response"
	"github.com/gofiber/fiber/v2"
)

// {{.Name}}Handler {{.Comment}}
func {{.Name}}Handler(c *fiber.Ctx) (err error) {
    {{if ne .Request "CommReq"}}
    req := new(types.{{.Request}})
    {{if len .PathParams}}{{range .PathParams}}req.{{.|LeftUpper}} = c.Params("{{.}}")
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