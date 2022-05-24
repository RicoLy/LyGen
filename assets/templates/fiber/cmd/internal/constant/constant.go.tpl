package constant

const (
	CTXTraceId       = "traceId"
	CTXStartTime     = "startTime"
	HeaderJwtAuthKey = "Authorization"
)

const (
	ErrParserBody       = "请求参数解析错误"
	ErrParam            = "请求参数有误"
	ErrCreateTokenErr   = "生成Token失败"
	ErrNoToken          = "权限校验失败;没有token信息"
	ErrTokenExpired     = "Token过期"
	ErrTokenNotValidYet = "Token无效"
	ErrUserOrPassword   = "密码或账号错误"
	ErrRecordNotFound   = "record not found"
	ErrDbException      = "数据库开了小差"
)
