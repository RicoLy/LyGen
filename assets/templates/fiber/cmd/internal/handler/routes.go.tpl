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
    {{$name}}Group := app.Group("{{.Prefix}}")
    {
        {{range .Methods}}
        {{$name}}Group.{{.MethodType|Title}}("{{.Path}}",{{range .MiddleWares}} middleware.{{.|Title}}Middleware(), {{end}} {{.Group}}.{{.Name}}Handler){{end}}
    }{{end}}
}
