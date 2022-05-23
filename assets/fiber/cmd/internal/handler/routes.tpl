package handler

import (
	"{{.Name}}/cmd/internal/handler/order"
	"{{.Name}}/cmd/internal/handler/user"
	"{{.Name}}/cmd/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func RegisterHandlers(app *fiber.App) {
    {{range .Services}}{{$name := .Name}}
    // {{$name}}Group {{.Comment}}
    {{$name}}Group := app.Group("{{.Group}}")
    {
        {{range .Methods}}
        {{$name}}Group.{{.MethodType}}("{{.Path}}",{{range .MiddleWares}} middleware.{{.|Title}}Middleware(),{{end}} {{.Group}}.CaptchaHandler){{end}}
    }{{end}}
}
