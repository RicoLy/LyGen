package handler

import (
    {{range .Services}}
    "{{.Meta}}/cmd/internal/handler/{{.Group}}"{{end}}
	"{{.Name}}/cmd/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandlers(app *fiber.App) {
    {{range .Services}}{{$name := .Name}}
    // {{$name}}Group {{.Comment}}
    {{$name|LeftLower}}Group := app.Group("{{.Prefix}}")
    {
        {{range .Methods}}
        {{$name|LeftLower}}Group.{{.MethodType|LeftUpper}}("{{.Path}}",{{range .MiddleWares}} middleware.{{.|LeftUpper}}Middleware(), {{end}} {{.Group}}.{{.Name}}Handler){{end}}
    }{{end}}
}
