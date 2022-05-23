package {{.PkgName}}

import (
	"LyAdmin/response"
    "go.opentelemetry.io/otel/trace"
    "net/http"
    
    {{if .After1_1_10}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
    {{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        {{if .HasRequest}}var req types.{{.RequestType}}
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }

        {{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        var traceId string
        spanCtx := trace.SpanContextFromContext(r.Context())
        if spanCtx.HasTraceID() {
            traceId = spanCtx.TraceID().String()
        }
        {{if .HasResp}}response.Response(w, traceId, resp, err){{else}}response.Response(w, traceId, nil, err){{end}}
            
    }
}
