package types

{{range .}}
// {{.Name}} {{.Comment}}
type {{.Name}} struct {
    {{range .ElementInfos}}{{.Name}}   {{.Type}}  `{{.Tags}}`
    {{end}}}
{{end}}


